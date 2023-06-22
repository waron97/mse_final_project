from flask import Flask
from flask_cors import CORS
from app.api.blueprint import api_blueprint
from app.util.db import get_crawl_pages
from app.util.setup_jobs import setup_jobs

setup_jobs()

app = Flask(__name__)
CORS(app)
app.register_blueprint(api_blueprint)


@app.route('/')
def greet():
    pages = get_crawl_pages()
    page = next(pages)
    return page["date"].date().isoformat()
