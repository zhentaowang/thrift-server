package business

import "wyun-thrift/src/gen-go/server"

type IBusinessService interface {
	Handle(operation string, paramJSON []byte) (r *server.Response, err error)
}
