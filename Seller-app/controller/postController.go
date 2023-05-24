package controller

import (

	"github.com/gin-gonic/gin"
	"encoding/json"
	"orderService/initializers"
	"orderService/models"
	"strconv"
)
func PostCreate(c *gin*Context){
	c.JSON(200,gin.H{

		// var body struct{
		// 	id string
		// 	status string
		// 	items []string
		// 	total int
		// 	currencyUnit string
		// }
		// c.Bind(%body)
		// order :=models.Post{id: body.id, status: body.status,items: body.items,total: body.total,currencyUnit: body.currencyUnit}

		orderRequest := models.OrderRequest{}
		ctx.Bind(&orderRequest)
		itemsString, err := json.Marshal(orderRequest.Items)
		order := models.Order{
			Status:       orderReq.Status,
			Items:        string(itemsString),
			Total:        total,
			CurrencyUnit: orderReq.CurrencyUnit,
		}

		result :=initializers.DB.Create(&order)
		if result.Error != nil{

			c.Status(400)
			return
		}

		c.JSON(200,gin.H{
			"Order " : order,
		})

	})
}

func ListOrder(c *gin*Context){
    var orders []models.Order
	initializers.DB.Find(&orders)

	
	c.JSON(200,gin.H{
		"list of orders " : orders,
	})

}


func ShowOrders(c *gin*Context){

id :=c.Param("id")
var order models.Order
initializers.DB.First(&order,id)

c.JSON(200,gin.H{

	"order" : order
})

}



func updateOrder(c *gin.Context){

	id :=c.Param("id")
	var body struct{
		status string
	}
	c.Bind(&body)

	var order models.Order	
	initializers.DB.First(&order,id)
	initializers.DB.Model(&order).Update(models.Order{ satus: body.status})
	c.JSON(200,gin.H{

		"order" : order
	})
	
}
