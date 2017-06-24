package thriftserver

import (
	"fmt"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrift-server/gen-go/server"
)

func StartServer(address string, port string, handler server.MyService) {

	processor := server.NewMyServiceProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(address + ":" + port)
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", address + ":" + port)
	server.Serve()
}