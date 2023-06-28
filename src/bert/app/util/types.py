
import datetime
from typing import TypedDict, List

from bson import ObjectId


class Anchor(TypedDict):
    text: str
    href: str


class PageData(TypedDict):
    _id: ObjectId
    url: str
    title: str
    links: List[Anchor]
    bodyText: str
    mainText: str
    description: str
    keywords: List[str]
    rawHtml: str
    date: datetime.datetime
