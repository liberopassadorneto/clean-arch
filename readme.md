# Challenge: Clean Architecture

## How to run the project

### Prerequisites
* In the root directory, run `docker-compose up -d`
* Run docker `exec -it {container_id} bash` to access the mysql container
* Run `mysql -u root -p` and enter the password (root)
* Select `orders` database by running `use orders` 
* Run the following SQL command to create the orders table
```sql
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))
```
## Run the project
* Go to `cmd/ordersystem` directory 
* Run `go run main.go wire_gen.go`

### Via REST API
#### Create a new order via `api.http` file
* Open `api.http` file
* Click on `POST .../order` and then click on `Send Request`
* Alternate on `id`, `price` and `tax` fields to create a new order

#### Get all orders via `api.http` file
* Open `api.http` file
* Click on `GET .../order` and then click on `Send Request`

### Via GRPC
* On the terminal run `evans -r repl` to access the GRPC client
#### Create a new order
* Run `call CreateOrder` to create a new order
* Enter the `id`, `price` and `tax` fields to create a new order
#### List all orders
* Run `call ListOrders` to get all orders

### Via GraphQL
* Open `http://localhost:8080/` on the browser
#### Create a new order
* Run the following mutation to create a new order
```graphql
mutation createOrder {
    createOrder(input: {id: "abc123", Price: 10.0, Tax: 2.0}) {
        id
        Price
        Tax
        FinalPrice
    }
}
```

#### List all orders
* Run the following query to get all orders
```graphql
query listOrders {
    listOrders {
        id
        Price
        FinalPrice
    }
}
```