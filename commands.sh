curl http://localhost:8000/greet

curl --header "Content-Type: application/json" http://localhost:8000/customers
curl --header "Content-Type: application/xml" http://localhost:8000/customers
# Not working: curl --header "Content-Type: text/csv" http://localhost:8000/customers 
curl --header "Content-Type: application/xml" http://localhost:8000/customers/123

# POSTing with curl’s -d option will include a default header that looks like: Content-Type: application/x-www-form-urlencoded
curl -X POST http://localhost:8000/customers
curl -d "user=user1&pass=abcd" -X POST http://localhost:8000/customers 


# Kill process
lsof -i :8000 <port number>
kill -9 23020 <pid>

# Hexagonal architecture
# curl --header "Content-Type: application/json" http://localhost:8000/customers

curl http://localhost:8000/customers
curl http://localhost:8000/customers/2000
curl http://localhost:8000/customers/20
curl http://localhost:8000/customers/rrr


# Assignment 2
curl http://localhost:8000/customers
curl "http://localhost:8000/customers?status=active"
curl "http://localhost:8000/customers?status=inactive"


# sqlx
curl http://localhost:8000/customers
curl http://localhost:8000/customers/2000
curl http://localhost:8000/customers/20

# DTO
curl http://localhost:8000/customers/2000


# Application configuration
SERVER_ADDRESS=localhost SERVER_PORT=8020 go run main.go
curl http://localhost:8020/customers
curl http://Donocalhost:8000/customers
curl http://localhost:8000/customers
SERVER_PORT=8020 go run main.go 
# save env variable
export SERVER_ADDRESS=localhost
SERVER_PORT=8020 go run main.go 
DB_USER=root DB_PASSWD=shrish287 DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking SERVER_ADDRESS=localhost SERVER_PORT=8020 go run main.go
curl http://localhost:8020/customers


# New Bank Account Part 2
DB_USER=root DB_PASSWD=shrish287 DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking SERVER_ADDRESS=localhost SERVER_PORT=8020 go run main.go
curl --header "Content-Type: application/json" -X POST -d '{"account_type":"saving","amount":5000.23}' http://localhost:8020/customers/2000/account
curl --header "Content-Type: application/json" -X POST -d '{"account_type":"saving","amount":1000.23}' http://localhost:8020/customers/2000/account


# Making a transaction
DB_USER=root DB_PASSWD=shrish287 DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking SERVER_ADDRESS=localhost SERVER_PORT=8020 go run main.go
curl --header "Content-Type: application/json" -X POST -d '{"transaction_type":"deposit","amount":5000.23}' http://localhost:8020/customers/2000/account/95474
curl --header "Content-Type: application/json" -X POST -d '{"transaction_type":"depositaa","amount":5000.23}' http://localhost:8020/customers/2000/account/95474
curl --header "Content-Type: application/json" -X POST -d '{"transaction_type":"deposit","amount":-5000.23}' http://localhost:8020/customers/2000/account/95474
curl --header "Content-Type: application/json" -X POST -d '{"transaction_type":"withdrawal","amount":5000.23}' http://localhost:8020/customers/2000/account/95474

