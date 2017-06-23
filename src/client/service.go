package client

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
	"wyun-thrift/src/pool"
	"wyun-thrift/src/gen-go/server"
	"wyun-thrift/src/conf"
)

type service struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Pool *pool.Pool  `json:"-"`
}

var Service = &service{}

func init() {
	fmt.Println("init  hahahahah")
	configByte, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Fatal(err)
	}

	conf.TConfig = conf.T{}
	err = yaml.Unmarshal(configByte, &conf.TConfig)
	if nil != err {
		log.Panic("thrift load conf error: ", err)
	}
	// client
    Service.Pool = &pool.Pool{
		Dial: func() (interface{}, error) {

			transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
			protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

			transport, err := thrift.NewTSocket(net.JoinHostPort(Service.Host, Service.Port))
			if err != nil {
				fmt.Fprintln(os.Stderr, "error resolving address:", err)
				os.Exit(1)
			}

			useTransport := transportFactory.GetTransport(transport)
			client := server.NewMyServiceClientFactory(useTransport, protocolFactory)
			if err := transport.Open(); err != nil {
				fmt.Fprintln(os.Stderr, "Error opening socket to"+Service.Host+":"+Service.Port, " ", err)
				os.Exit(1)
			}
			return client, nil
		},
		Close: func(v interface{}) error {
			v.(*server.MyServiceClient).Transport.Close()
			return nil
		},
		MaxActive:   conf.TConfig.MaxConns,
		MaxIdle:     conf.TConfig.MaxIdle,
		IdleTimeout: time.Duration(conf.TConfig.MaxIdleConnDuration),
	}

	return
}
