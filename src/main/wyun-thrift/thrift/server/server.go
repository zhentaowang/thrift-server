package server

import (
	"fmt"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"wyun-thrift/thrift/gen-go/server"
)

func StartServer(address string, port string, handler server.MyService) {

	NetworkAddress := address + ":" + port
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddress)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	processor := server.NewMyServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddress)
	server.Serve()
}