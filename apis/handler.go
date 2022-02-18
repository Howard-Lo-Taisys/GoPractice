package apis

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	. "crawler2/models"
)

func GetDataApi(c *gin.Context) {
	var d Attributes
	datas, err := d.GetData()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, &datas)
}

func PostDataApi(c *gin.Context) {
	ranks := c.Request.FormValue("numb")
	addr := c.Request.FormValue("address")
	amount := c.Request.FormValue("amount")
	percentages := c.Request.FormValue("percentage")

	rank, _ := strconv.Atoi(ranks)
	percentage, _ := strconv.ParseFloat(percentages, 64)

	d := Attributes{Rank: rank, Address: addr, Amount: amount, Percentage: percentage}
   
	ra, err := d.AddData()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeleteDataApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	d := Attributes{Id: id}
	ra, err := d.DeleteData()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete id %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}