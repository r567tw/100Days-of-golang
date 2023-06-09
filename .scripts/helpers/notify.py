import os
import requests

def send(message):
    token = os.getenv("TOKEN")
    notify_url = "https://notify-api.line.me/api/notify"

    requests.post(
        notify_url,
        headers={'Authorization': "Bearer {}".format(token)},
        data={"message": message}
    )