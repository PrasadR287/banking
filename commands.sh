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
