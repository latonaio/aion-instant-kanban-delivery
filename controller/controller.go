package controller

import (
	"aion-instant-kanban-delivery/domain"
	"aion-instant-kanban-delivery/kanban"
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"bitbucket.org/latonaio/aion-core/pkg/log"
	"github.com/labstack/echo"
	"net/http"

)

func SetKanban(c echo.Context) error {
	setter := &domain.KanbanSetter{}
	err := c.Bind(setter)
	if err != nil {
		return err
	}
	kb := kanban.KanbanClient()
	_,err = kb.SetKanban(setter.MimickedMsName,setter.ProcessNum)
	if err != nil {
		return err
	}

	log.Printf("microservice watcher set as %s",setter.MimickedMsName)
	return c.JSON(http.StatusOK,"OK")
}

func WriteKanban(c echo.Context) error {
	writer := &domain.KanbanWriter{}
	err := c.Bind(writer)
	if err != nil {
		return err
	}
	opt := writer.BuildOutputData()
	req,err := msclient.NewOutputData(opt...)
	if err != nil {
		return err
	}
	kb := kanban.KanbanClient()
	err = kb.OutputKanban(req)
	if err != nil {
		return err
	}


	return c.JSON(http.StatusOK,"OK")
}