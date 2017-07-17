from thrift.protocol import TBinaryProtocol
from thrift.transport import TSocket
from thrift.transport.TTransport import TFramedTransport

from server.personality import MyService
from server.personality.ttypes import Request


class SimpleClient:
    def __init__(self, host, port):
        self.host = host
        self.port = port

    def send(self, request):
        t_frame_transport = TFramedTransport(TSocket.TSocket(self.host, self.port))
        t_frame_transport.open()
        t_protocol = TBinaryProtocol.TBinaryProtocol(t_frame_transport)
        client = MyService.Client(t_protocol)
        response = client.send(request)
        t_frame_transport.close()
        return response


if __name__ == '__main__':
    client = SimpleClient('127.0.0.1', '9090')
    request = Request()
    request.serviceName = 'thrift'
    request.operation = 'test'
    request.paramJSON = 'abc'
    response = client.send(request)
    print response
