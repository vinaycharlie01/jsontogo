import re
import socket
import logging
import json

# import jsonschema
from datetime import datetime

class Email:
    def __init__(self, mail="", username="", domain=""):
        self.mail = mail
        self.username = username
        self.domain = domain

    def is_valid_username(self, username):
        username_regex = r'^[a-zA-Z0-9_]{3,20}$'
        return bool(re.match(username_regex, username))

    def is_valid_domain(self, domain):
        try:
            socket.gethostbyname(domain)
            return True
        except socket.error:
            return False

    def is_email_slicer(self):
        result = self.mail.strip().split("@")
        if len(result) == 2:
            if self.is_valid_username(result[0]) and self.is_valid_domain(result[1]):
                logging.info(json.dumps({
                    "level": "info",
                    "msg": "Valid Mail address",
                    "mail": self.mail,
                    "username": result[0],
                    "domain": result[1],
                    "time": datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
                }))
                return Email(username=result[0], domain=result[1])
            else:
                logging.error(json.dumps({
                    "level": "error",
                    "msg": "Invalid mail",
                    "mail": self.mail,
                    "username": result[0],
                    "domain": result[1],
                    "time": datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
                }))
                return None
        else:
            logging.error(json.dumps({
                "level": "error",
                "msg": "Invalid mail",
                "mail": self.mail,
                "time": datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
            }))
            return None


# logging.basicConfig(filename='logfile.log', level=logging.INFO, format='%(asctime)s - %(message)s')

logging.basicConfig(filename='logfile.log', level=logging.INFO, format='%(message)s')

if __name__ == "__main__":
    b = Email()

    try:
        res = input()
        b.mail = res
        result_email = b.is_email_slicer()
        if result_email:
            print(f"Username: {result_email.username}, Domain: {result_email.domain}")
    except Exception as e:
        logging.error(json.dumps({
            "level": "error",
            "msg": f"Error reading input: {e}",
            "time": datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
        }))
