package thriftserver

import (
	"testing"
	"code.aliyun.com/wyunshare/thrift-server/business"
	"code.aliyun.com/wyunshare/thrift-server/processor"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"encoding/json"
)
type BusinessServiceImpl struct {}
func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {
	resMap := make(map[string]string)
	resMap["operation"] = operation
	resJSON, _ := json.Marshal(resMap)

	res := server.Response{
		ResponeCode: server.RESCODE__200,
		ResponseJSON: resJSON,
	}
	return &res, nil
}

func TestStartServer(t *testing.T) {


	businessServiceMap := &business.BusinessServiceMap{
		ServiceMap: make(map[string]business.IBusinessService),
	}
	businessServiceMap.RegisterService("businessService", &BusinessServiceImpl{})
	wyunServiceImpl := processor.WyunServiceImpl{
		BusinessServiceMap: businessServiceMap,
	}
	go StartServer("localhost", "9092", &wyunServiceImpl)
}

func TestStartSingleServer(t *testing.T) {
	go StartSingleServer("0.0.0.0", "9093", "businessService", &BusinessServiceImpl{})
}