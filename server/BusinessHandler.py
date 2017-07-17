import json
import logging

from server.personality.ttypes import *


class IBusinessHandler:
    def handle(self, operation, params):
        pass


class LogBusinessHandler(IBusinessHandler):
    def handle(self, operation, params):
        logging.info(operation)
        logging.info(params)
        return Response(200, json.dumps({
            'success': True
        }))
