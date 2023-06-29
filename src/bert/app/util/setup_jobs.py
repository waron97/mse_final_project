from app.util.fs import create_out_folder
from app.util.logger import Logger


def setup_jobs():
    Logger.info("setup_jobs", "Setting up python application")
    create_out_folder()
