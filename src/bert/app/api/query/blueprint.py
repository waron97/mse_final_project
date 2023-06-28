from flask.blueprints import Blueprint

blueprint = Blueprint('api', __name__, url_prefix='/query')


@blueprint.route('/')
def get():
    return 'Hello World'
