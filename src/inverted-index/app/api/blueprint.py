from flask import blueprints
from .query.blueprint import blueprint as query


api_blueprint = blueprints.Blueprint('api', __name__, url_prefix='/api')


api_blueprint.register_blueprint(query)
