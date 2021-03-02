package main

import (
	"aion-instant-kanban-delivery/controller"
	"aion-instant-kanban-delivery/kanban"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"strconv"
)

func main()  {
	ctx := context.Background()
	isDebug,_ := strconv.ParseBool(os.Getenv("IS_DEBUG"))
	e := echo.New()
	err := kanban.InitKanbanClient(ctx,isDebug)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/kanban/set",controller.SetKanban)
	e.POST("/kanban/write",controller.WriteKanban)

	err = e.Start(":1323")
	if err != nil {
		panic(err)
	}

	quitCh := make(chan os.Signal,1)
	go func() {
		for {
			select {
			case <-quitCh:
				err := e.Close()
				if err != nil {
					panic(err)
				}
			}
		}

	}()
}
