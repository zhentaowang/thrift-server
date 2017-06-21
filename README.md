服务端使用

router.go
```go
import "wyun-thrift/src/gen-go/server"

type BusinessServiceImpl struct {
}

// 通过 BusinessServiceImpl 实现 IBusinessService 接口的 Send 方法，从而实现 IBusinessService 接口
func (msi *BusinessServiceImpl) Handle(operation string, paramJSON []byte) (r *server.Response, err error) {

	paramInput := make(map[string]interface{})
	json.Unmarshal(paramJSON, &paramInput)

	switch request.Operation {
	case <<your operation name>>:
		// TODO 业务逻辑
	}
	...
}
```
main.go
```go
import "wyun-thrift/src/server"

// 创建 BusinessServiceMap, 并注册支持的 service
businessServiceMap := &business.BusinessServiceMap{
	ServiceMap: make(map[string]business.IBusinessService),
}
businessServiceMap.RegisterService("businessService", routers.GetHandler())

wyunServiceImpl := processor.WyunServiceImpl{
	BusinessServiceMap: businessServiceMap,
}
server.StartServer(thrift_server_address, thrift_server_port, &wyunServiceImpl) // thrift server

```
客户端使用
```go

import (
	"wyun-thrift/src/gen-go/server"
	"wyun-thrift/src/client"
	)

	service := client.Service //申明一个请求服务
	service.Host = beego.AppConfig.String("thrift_client_address") // 初始化请求地址
	service.Port = beego.AppConfig.String("thrift_client_port") // 初始化请求端口
	pooledClient, err := service.Pool.Get()  //从连接池获取thrift client
	if err != nil {
		log.Println("Thrift pool get client error", err)
		return
	}
	defer service.Pool.Put(pooledClient, false)

	rawClient, ok := pooledClient.(*server.MyServiceClient)
	if !ok {
		log.Println("convert to raw client failed")
		return
	}

	req := server.NewRequest() //创建request
	req.ServiceName = <<your service name>>
	req.Operation = <<your operation name>>
	paramInput := make(map[string]interface{})
	req.ParamJSON, _ = json.Marshal(paramInput)
	r, err = rawClient.Send(req) //发送request
```