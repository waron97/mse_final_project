import os

MAX_QUERY_SIZE = int(os.getenv('MAX_QUERY_SIZE', 100))
BERT_MODEL = os.getenv('BERT_MODEL', 'bert-base-uncased')
