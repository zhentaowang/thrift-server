package server

import (
	"testing"
	"wyun-thrift/src/business"
	"wyun-thrift/src/processor"
	"wyun-thrift/src/gen-go/server"
)
type BusinessServiceImpl struct {}
func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {
	return nil, nil
}

func TestStartServer(t *testing.T) {


	businessServiceMap := &business.BusinessServiceMap{}
	businessServiceMap.RegisterService("businessService", &BusinessServiceImpl{})
	wyunServiceImpl := processor.WyunServiceImpl{
		BusinessServiceMap: businessServiceMap,
	}
	StartServer("localhost", "9092", &wyunServiceImpl)
}