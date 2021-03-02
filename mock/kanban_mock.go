package mock

import (
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"bitbucket.org/latonaio/aion-core/pkg/log"
	"bitbucket.org/latonaio/aion-core/proto/kanbanpb"
)

type MockKanbanServer struct {}

func (mk *MockKanbanServer) GetOneKanban(serviceName string, processNumber int) (*msclient.WrapKanban, error) {
	log.Printf("serviceName %s, processNumber %d",serviceName,processNumber)
	return &msclient.WrapKanban{},nil
}

func (mk *MockKanbanServer) GetKanbanCh(serviceName string, processNumber int) (chan *msclient.WrapKanban, error) {
	log.Printf("serviceName %s, processNumber %d",serviceName,processNumber)
	c := make(chan *msclient.WrapKanban)
	return c,nil
}

func (mk *MockKanbanServer) SetKanban(serviceName string, processNumber int) (*msclient.WrapKanban, error) {
	log.Printf("serviceName %s, processNumber %d",serviceName,processNumber)
	return &msclient.WrapKanban{},nil
}
func (mk *MockKanbanServer) OutputKanban(data *kanbanpb.OutputRequest) error {
	log.Printf("OutputRequest %+v",data)
	return nil
}
func (mk *MockKanbanServer) Close() error {
	log.Print("close")
	return nil
}
func (mk *MockKanbanServer) GetProcessNumber() int {
	return 0
}
