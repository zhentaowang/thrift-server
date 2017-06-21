package client

import (
	"net"
	"fmt"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"wyun-thrift/thrift/gen-go/server"
)

func SendClient(request *server.Request, host string, port string) (r *server.Response, err error){

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := server.NewMyServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to" + host + ":" + port, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	r, err = client.Send(request)
	return r, err

}
