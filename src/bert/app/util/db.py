from pymongo import MongoClient
from typing import Iterator

from app.util.types import PageData

connstring = "mongodb://db:27017"
client = MongoClient(connstring)
db = client["mse"]
Crawl = db["crawl"]


def get_crawl_pages() -> Iterator[PageData]:
    cursor = Crawl.find({})
    return cursor
