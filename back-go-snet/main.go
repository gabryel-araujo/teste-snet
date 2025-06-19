package main

import (
	"net/http"

	"github.com/gabryel-araujo/back-go-snet/db"
	"github.com/gabryel-araujo/back-go-snet/handler"
	"github.com/gabryel-araujo/back-go-snet/repository"
	"github.com/gabryel-araujo/back-go-snet/service"
	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()

	d := db.Connect()
	defer d.Close()

	estabRepo := &repository.EstabRepository{DB: d}
	estabService := &service.EstabService{Repo: estabRepo}
	estabHandler := &handler.EstabHandler{Service: estabService}

	storeRepo := &repository.StoreRepository{DB: d}
	storeService := &service.StoreService{Repo: storeRepo}
	storeHandler := &handler.StoreHandler{Service: storeService}

	e.GET("/",func(c echo.Context)error{
		return c.String(http.StatusOK,"Api rodando")
	})

	e.POST("/estabelecimentos",estabHandler.CreateEstab)
	e.GET("/estabelecimentos",estabHandler.GetAllEstabs)
	e.GET("/estabelecimentos/:id",estabHandler.GetOneEstab)
	e.PUT("/estabelecimentos/:id",estabHandler.UpdateEstab)
	e.DELETE("/estabelecimentos/:id",estabHandler.DeleteEstab)

	e.POST("/lojas",storeHandler.CreateStore)
	e.GET("/lojas",storeHandler.GetAllStores)
	e.GET("/lojas/:id",storeHandler.GetOneStore)
	e.PUT("/lojas/:id",storeHandler.UpdateStore)
	e.DELETE("/lojas/:id",storeHandler.DeleteStore)

	e.Logger.Fatal(e.Start(":1323"))
}