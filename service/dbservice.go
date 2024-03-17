package service

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	model "warehouse/model"
)

func DatebaseInit() *sql.DB {
	fmt.Println("SET UP")
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env " + err.Error())
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database: " + err.Error())

		return nil
	}

	return db
}

func GetAllProduct() []model.Product {
	db := DatebaseInit()
	products, err := db.Query("SELECT * FROM warehouse.product")
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return nil
	}
	defer products.Close()

	productList := []model.Product{}

	for products.Next() {
		var product model.Product
		err := products.Scan(&product.ID, &product.Title, &product.Category, &product.Price, &product.Amount)
		if err != nil {
			fmt.Println("Error scanning product:", err.Error())
			return nil
		}
		productList = append(productList, product)
	}

	return productList

}

func GetAllOrder() []model.Order {
	db := DatebaseInit()
	results, err := db.Query("SELECT warehouse.transaction.idtransaction, warehouse.product.product_name, warehouse.customer.name, warehouse.transaction.amount FROM warehouse.transaction INNER JOIN customer ON transaction.customerid = customer.idcustomer INNER JOIN product ON transaction.productid = product.idproduct")
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return nil
	}

	defer results.Close()

	orderList := []model.Order{}

	for results.Next() {
		var order model.Order
		err := results.Scan(&order.ID, &order.ProductName, &order.CustomerName, &order.Amount)
		if err != nil {
			fmt.Println("Error scanning order ", err.Error())
			return nil
		}
		orderList = append(orderList, order)
	}
	return orderList
}

func GetOrder(id string) *model.Order {
	db := DatebaseInit()
	results, err := db.Query("SELECT warehouse.transaction.idtransaction, warehouse.product.product_name, warehouse.customer.name, warehouse.transaction.amount FROM warehouse.transaction INNER JOIN customer ON transaction.customerid = customer.idcustomer INNER JOIN product ON transaction.productid = product.idproduct where idtransaction=?", id)
	if err != nil {
		fmt.Println("Error querying database order: ", err.Error())
		return nil
	}
	defer results.Close()

	order := &model.Order{}

	if results.Next() {
		err := results.Scan(&order.ID, &order.ProductName, &order.CustomerName, &order.Amount)
		if err != nil {
			fmt.Println("Error scanning order: ", err.Error())
			return nil
		}
	}
	return order

}

func GetProduct(id string) *model.Product {
	db := DatebaseInit()
	product, err := db.Query("SELECT * FROM warehouse.product where idproduct=?", id)
	if err != nil {
		fmt.Println("Error querying database:", err.Error())
		return nil
	}
	defer product.Close()

	productTarget := &model.Product{}

	if product.Next() {
		err := product.Scan(&productTarget.ID, &productTarget.Title, &productTarget.Category, &productTarget.Price, &productTarget.Amount)
		if err != nil {
			fmt.Println("Error scanning product:", err.Error())
			return nil
		}
	}

	return productTarget

}

func GetAllCustomer() []model.Customer {
	db := DatebaseInit()
	customers, err := db.Query("SELECT * FROM warehouse.customer")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil
	}
	defer customers.Close()

	customerList := []model.Customer{}

	for customers.Next() {
		var customer model.Customer
		err := customers.Scan(&customer.ID, &customer.CustomerName)
		if err != nil {
			fmt.Println("Error scanning product:", err)
			return nil
		}
		customerList = append(customerList, customer)
	}

	return customerList

}
