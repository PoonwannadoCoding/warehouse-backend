package service

import (
	"warehouse/model"

	"github.com/gin-gonic/gin"
)

func Init() {

	app := gin.Default()

	app.GET("/hello/:id", respondID)
	app.GET("/all/product", getAllProduct)
	app.GET("/product/:id", getProduct)
	app.GET("/all/customer", getAllCustomer)
	app.GET("/all/order", getAllOrder)
	app.GET("/order/:id", getOrder)
	app.POST("/add/product", postProduct)
	app.POST("/add/customer", postCustomer)

	app.Run()

}

func respondID(c *gin.Context) {

	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "hello " + id,
	})
}

func postProduct(c *gin.Context) {
	var product model.Product

	if err := c.BindJSON(&product); err != nil {
		return
	}
	InsertProduct(product)
	newProduct := GetProduct(product.ID)
	c.JSON(200, gin.H{
		"message": newProduct,
	})
}

func postCustomer(c *gin.Context) {
	var customer model.Customer

	if err := c.BindJSON(&customer); err != nil {
		return
	}
	InsertCustomer(customer)

}

func getAllProduct(c *gin.Context) {
	productList := GetAllProduct()

	c.JSON(200, gin.H{
		"message": productList,
	})
}

func getAllOrder(c *gin.Context) {
	orderList := GetAllOrder()
	c.JSON(200, gin.H{
		"message": orderList,
	})
}

func getOrder(c *gin.Context) {
	id := c.Param("id")
	order := GetOrder(id)

	c.JSON(200, gin.H{
		"message": order,
	})
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	product := GetProduct(id)

	c.JSON(200, gin.H{
		"message": product,
	})
}

func getAllCustomer(c *gin.Context) {
	customerList := GetAllCustomer()

	c.JSON(200, gin.H{
		"message": customerList,
	})
}
