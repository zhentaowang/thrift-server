package thriftserver

import (
	"testing"
	"code.aliyun.com/wyunshare/thrift-server/business"
	"code.aliyun.com/wyunshare/thrift-server/processor"
	"code.aliyun.com/wyunshare/thrift-server/gen-go/server"
	"encoding/json"
	"log"
	"fmt"
)
type BusinessServiceImpl struct {}
func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {
	fmt.Println(operation)
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

	po := GetPool("localhost", "9092")
	pooledClient, err := po.Get() //从连接池获取thrift client
	if err != nil {
		log.Println("Thrift pool get client error", err)
		return
	}

	defer po.Put(pooledClient, false)

	rawClient, ok := pooledClient.(*server.MyServiceClient)
	if !ok {
		log.Println("convert to raw client failed")
		return
	}

	//req := &server.Request{[]byte(dict), "name", "test"}

	req := server.NewRequest() //创建request
	req.ServiceName = "businessService"
	req.Operation = "test"

	paramInput := make(map[string]interface{})

	req.ParamJSON, _ = json.Marshal(paramInput)

	r, err := rawClient.Send(req) //发送request
	t.Log(r.ResponseJSON)
}

func TestStartSingleServer(t *testing.T) {
	go StartSingleServer("0.0.0.0", "9093", "businessService", &BusinessServiceImpl{})
}