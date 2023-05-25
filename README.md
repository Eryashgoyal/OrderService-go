# OrderService-go

This service provide the functionality for Adding,Updating,Find the orders

In order to run the services you must have some Service Dependency

Mysql database 
go-lang
Install Go package dependencies by running the below commands from order-service folder

go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm

go get -u gorm.io/driver/postgres

go get github.com/thedevsaddam/govalidator

Setting up DB Migrations for Order service
go run migrations/migrate.go

Starting the Service
Start Go order service locally. Run the following command from the directory in which main.go is located in order-service folder. This command will run the executable to start the service.

go run main.go

Play with order endpoints to add/update/get order.

The endpoints can be invoked using the popular curl command or any REST client such as POSTMAN, etc.

Create a new order

HTTP METHOD : POST

URL : http://localhost:5050/Add-orders

Content-Type : application/json; charset=utf-8

REQUEST PAYLOAD :

{ "status": "PENDING_INVOICE", "items" : [{ "id": "1234", "description": "some description", "price": 10, "quantity": 1 }, { "id": "1234", "description": "any description", "price": 29, "quantity": 1 }], "total": 46, "currency_unit": "USD" }

Fetch all existing orders

HTTP METHOD : GET

URL : http://localhost:5050/List-orders

Filter orders

HTTP METHOD : GET

URL : http://localhost:5050/orders?status=PENDING_INVOICE&order_by=id&sort_order=desc

or

URL : http://localhost:5050/show-orders?id=2

Update an existing order

HTTP METHOD : PUT

URL : http://localhost:5050/update-orders/{order_id}

SAMPLE URL : http://localhost:5050/update-orders/1

Content-Type : application/json; charset=utf-8

REQUEST PAYLOAD :

{ "status": "SUCCESS"}
