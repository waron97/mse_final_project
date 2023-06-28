from flask import Flask
from flask_cors import CORS
from app.api.blueprint import api_blueprint
from app.util.setup_jobs import setup_jobs

setup_jobs()

app = Flask(__name__)
CORS(app)

app.register_blueprint(api_blueprint)
