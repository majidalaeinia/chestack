package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

// The route is something like: endpoint?filter={"q":"مجید علائی نیا"}

type QueryParam struct {
	Filter string `form:"filter"`
}

type FilterQueryValue struct {
	Q string `json:"q"`
}

func Q(c *gin.Context) (fqv FilterQueryValue) {
	var qp QueryParam
	err := c.BindQuery(&qp)
	if err != nil {
		log.Println(err)
	}
	qj := qp.Filter
	qb := []byte(qj)
	err = json.Unmarshal(qb, &fqv)
	if err != nil {
		log.Println(err)
	}
	return
}
