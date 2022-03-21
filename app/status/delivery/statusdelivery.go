package delivery

import (
	"net/http"

	"Assignment3/app/status"

	"github.com/gin-gonic/gin"
)

type HttpDelivery struct {
	status.StatusUsecase
}

func NewHttpDelivery(r *gin.Engine, statusUC status.StatusUsecase) {
	handler := HttpDelivery{
		StatusUsecase: statusUC,
	}

	r.GET("/status", handler.updateStatus)
}

func (h *HttpDelivery) updateStatus(c *gin.Context) {
	data, err := h.UpdateStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	datas := StatusBahaya{
		Water:       data.Water,
		Wind:        data.Wind,
		StatusAir:   data.StatusAir,
		StatusAngin: data.StatusAngin,
	}

	c.HTML(
		http.StatusOK,
		"header.html",
		gin.H{
			"title": "Status Page",
			"datas": datas,
		},
	)

}
