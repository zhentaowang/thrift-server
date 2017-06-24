package processor

import (
	"thrift-server/business"
	"thrift-server/gen-go/server"
)

type WyunServiceImpl struct{
	BusinessServiceMap *business.BusinessServiceMap
}

func (w *WyunServiceImpl) Send(request *server.Request) (r *server.Response, err error) {
	return w.BusinessServiceMap.Handle(request.ServiceName, request.Operation, request.ParamJSON)
}