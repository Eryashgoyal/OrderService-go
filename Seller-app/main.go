package main

import (
  
	"orderService/controllers"
	"orderService/initializers"

  "github.com/gin-gonic/gin"
  
)
func init(){

	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/Add-orders", controllers.AddOrder)
	r.POST("/Update-orders/:id", controllers.UpdateOrder)
	r.GET("/List-orders", controllers.ListOrders)
	r.GET("/Show-orders/:id", controllers.ShowOrders)

	r.run(":3000")
}


