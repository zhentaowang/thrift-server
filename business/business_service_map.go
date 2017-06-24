package business

import "thrift-server/gen-go/server"

type BusinessServiceMap struct {
	ServiceMap map[string]IBusinessService
}
func (b *BusinessServiceMap) Handle(serviceName string, operation string, paramJSON []byte) (*server.Response, error) {
	businessService := b.ServiceMap[serviceName]
	return businessService.Handle(operation, paramJSON)
}

func (b *BusinessServiceMap) RegisterService(serviceName string, businessService IBusinessService) {
	b.ServiceMap[serviceName] = businessService
}