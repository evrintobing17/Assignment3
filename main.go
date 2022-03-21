package main

import (
	"fmt"
	"log"
	"net/http"

	orderdelivery "Assignment3/app/status/delivery"
	orderUC "Assignment3/app/status/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	OrderUc := orderUC.NewStatusUsecase()
	orderdelivery.NewHttpDelivery(r, OrderUc)

	r.LoadHTMLFiles("html/header.html")

	fmt.Println("Listening to port 8081")

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
