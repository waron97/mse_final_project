import os

MAX_QUERY_SIZE = int(os.getenv('MAX_QUERY_SIZE', 100))
BERT_MODEL = os.getenv('BERT_MODEL', 'bert-base-uncased')
LOGS_URL = "http://logs:8080"
LOGS_APP_NAME = os.getenv("LOGS_APP_NAME")
LOGS_API_KEY = os.getenv("LOGS_KEY")
