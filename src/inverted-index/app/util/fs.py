import os

OUT_FOLDER = 'out'


def create_out_folder():
    """Create the output folder if it doesn't exist."""
    if not os.path.exists(OUT_FOLDER):
        os.makedirs(OUT_FOLDER)
