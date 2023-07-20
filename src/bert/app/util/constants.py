import os

MAX_QUERY_SIZE = int(os.getenv('MAX_QUERY_SIZE', 50))
BERT_MODEL = os.getenv('BERT_MODEL', 'distilbert-base-multilingual-cased')
LOGS_URL = "http://logs:8080"
LOGS_APP_NAME = os.getenv("LOGS_APP_NAME")
LOGS_API_KEY = os.getenv("LOGS_KEY")
