package kanban

import (
	"aion-instant-kanban-delivery/mock"
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"context"
	"sync"
)

var (
	once sync.Once
	kanbanClient msclient.MicroserviceClient
)

func InitKanbanClient(ctx context.Context,isDebug bool) error {
	var err error
	once.Do(func () {
		if isDebug {
			kanbanClient = &mock.MockKanbanServer{}
		} else {
			kanbanClient,err = msclient.NewKanbanClient(ctx)
		}
	})
	if err != nil {
		return err
	}
	return nil
}

func KanbanClient() msclient.MicroserviceClient {
	return kanbanClient
}