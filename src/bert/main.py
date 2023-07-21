from transformers import DistilBertTokenizer, DistilBertModel
import torch
import string

from flask import Flask, request, jsonify
from flask_cors import CORS
from app.api.blueprint import api_blueprint
# from app.util.setup_jobs import setup_jobs
from app.util.constants import MAX_QUERY_SIZE, BERT_MODEL

# Instantiate the tokenizer and model in the global scope
tokenizer = DistilBertTokenizer.from_pretrained(BERT_MODEL, pad_token="[MASK]")
model = DistilBertModel.from_pretrained(BERT_MODEL)

# setup_jobs()

app = Flask(__name__)
CORS(app)

app.register_blueprint(api_blueprint)


@app.route("/encode", methods=["POST"])
def process_text():
    """
    Input: {text, type}
    "type" is either "document" or "query"
    Output: {embeddings: float[][]}
    """
    text = request.json["text"]
    try:
        if request.json["type"] == 'query':
            result = get_bert_embedding_query(text)
        elif request.json["type"] == 'document':
            result = get_bert_embedding_document(text)
        return jsonify({"embeddings": result})
    except Exception as e:
        print(request.json)
        print(e)
        raise e


@app.route("/health", methods=["GET"])
def health():
    return "OK"


def get_bert_embedding_query(text):
    encoded_input = tokenizer.encode_plus(
        text,
        add_special_tokens=True,
        return_tensors='pt',
        max_length=MAX_QUERY_SIZE,
        padding='max_length',
        truncation=True,
        verbose=True,
    )

    with torch.no_grad():
        model_output = model(**encoded_input)

    embeddings = model_output.last_hidden_state.squeeze().tolist()

    return embeddings


def get_bert_embedding_document(text, max_input_size=512):
    words = text.split()
    chunks = []
    current_chunk = ""
    for word in words:
        if len(current_chunk) + len(word) + 1 <= max_input_size:
            if current_chunk:
                current_chunk += " " + word
            else:
                current_chunk = word
        else:
            chunks.append(current_chunk)
            current_chunk = word

    if current_chunk:
        chunks.append(current_chunk)

    embeddings_list = []
    punctuation_positions_list = []
    # Initiate the size of the last chunk (basing on tokens) as 0.
    last_chunk_tokens_length = 0
    for chunk in chunks:
        encoded_input = tokenizer.encode_plus(
            chunk,
            add_special_tokens=True,
            return_tensors='pt',
            padding=True,
            truncation=True,
        )

        with torch.no_grad():
            model_output = model(**encoded_input)

        embeddings = model_output.last_hidden_state.squeeze().tolist()
        embeddings_list.append(embeddings)

        # Get the positions of punctuations from each chunk with considering special tokens
        punctuation_positions = __get_punctuation_positions(chunk, tokenizer)

        punctuation_positions = [
            i + last_chunk_tokens_length for i in punctuation_positions]
        punctuation_positions_list.extend(punctuation_positions)
        # Update variable "last_chunk_tokens_length" with the size of the current chunk (basing on tokens) for
        # computing punctuation positions in the next chunk.
        # two special tokens for each chunk in BERT
        last_chunk_tokens_length = len(tokenizer.tokenize(chunk)) + 2

    concatenated_embeddings = [
        emb for chunk in embeddings_list for emb in chunk]
    # Remove punctuation embeddings
    __remove_elements_by_index(
        concatenated_embeddings, punctuation_positions_list)

    return concatenated_embeddings


def __get_punctuation_positions(chunk, tokenizer):
    punctuation_positions = []
    tokens = tokenizer.tokenize(chunk)
    for i, token in enumerate(tokens):
        if token in string.punctuation:
            punctuation_positions.append(i + 1)  # the 1 is for [cls]
    return punctuation_positions


def __remove_elements_by_index(embeddings, positions):
    # Sort to not influence the actual indexes
    positions = sorted(positions, reverse=True)
    for pos in positions:
        embeddings.pop(pos)
    return embeddings


if __name__ == "__main__":
    app.run(port=8888, host="0.0.0.0")

    # ### Check when "type" is "document"
    # BERT_MODEL = "bert-base-uncased"
    # MAX_QUERY_SIZE = 128
    #
    # text = "Tomorrow is Thursday."
    # embeddings = get_bert_embedding_query(text, BERT_MODEL, MAX_QUERY_SIZE)
    #
    # print("BERT embeddings:", embeddings)

    # ### Check when "type" is "document"
    # BERT_MODEL = "bert-base-uncased"
    # MAX_INPUT_SIZE = 512
    #
    # text = "A book."
    # embeddings = get_bert_embedding_document(text, BERT_MODEL, MAX_INPUT_SIZE)
    #
    # print(len(embeddings))
    # # print(embeddings)
