# -*- coding: utf-8 -*-
from server.BusinessHandler import LogBusinessHandler
from server.MyServiceHandler import MyServiceHandler
from server.personality import MyService
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer


class Server:
    service_handler = MyServiceHandler()

    def __init__(self, port):
        self.port = port

    def start_single_server(self, business_handler, service_name):
        self.service_handler.register_business_handler(service_name, business_handler)
        self.start_server(self.port)

    # 启动thrift服务器
    def start_server(self, port):
        processor = MyService.Processor(self.service_handler)
        transport = TSocket.TServerSocket('0.0.0.0', port)
        tfactory = TTransport.TFramedTransportFactory()
        pfactory = TBinaryProtocol.TBinaryProtocolFactory()
        server = TServer.TThreadPoolServer(processor, transport, tfactory, pfactory)
        print("Starting thrift server...")
        server.serve()


if __name__ == '__main__':
    server = Server(9090)
    server.start_single_server(LogBusinessHandler(), 'thrift')
