import time
import requests

from app.util.constants import LOGS_API_KEY, LOGS_APP_NAME, LOGS_URL


class Logger:
    @staticmethod
    def _send_log(level, location, message, data):
        body = {
            "level": level,
            "location": location,
            "message": message,
            "data": data,
            "appId": LOGS_APP_NAME
        }

        while not Logger.is_alive():
            time.sleep(1)

        try:
            response = requests.post(
                LOGS_URL + "/logs",
                headers={"Content-Type": "application/json",
                         "Authorization": f"apiKey {LOGS_API_KEY}"},
                json=body
            )

            print(response.status_code, response.text)

        except Exception as e:
            print("Failed to send log to logs service", e)

    @staticmethod
    def debug(location: str, message: str, data=None):
        Logger._send_log("debug", location, message, data)

    @staticmethod
    def info(location: str, message: str, data=None):
        Logger._send_log("info", location, message, data)

    @staticmethod
    def warning(location: str, message: str, data=None):
        Logger._send_log("warning", location, message, data)

    @staticmethod
    def error(location: str, message: str, data=None):
        Logger._send_log("error", location, message, data)

    @staticmethod
    def critical(location: str, message: str, data=None):
        Logger._send_log("critical", location, message, data)

    @staticmethod
    def is_alive() -> bool:
        try:
            requests.get(LOGS_URL)
            return True
        except requests.exceptions.ConnectionError:
            return False
