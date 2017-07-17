from server.personality import MyService
from server.personality.ttypes import *


class MyServiceHandler(MyService.Iface):
    business_handlers = {}

    def send(self, request):
        service_name = request.serviceName
        operation = request.operation
        param_json = request.paramJSON
        if service_name not in self.business_handlers:
            return Response(404)
        business_handler = self.business_handlers[service_name]
        if business_handler is not None:
            return business_handler.handle(operation, param_json)

    def register_business_handler(self, name, business_handler):
        self.business_handlers[name] = business_handler
