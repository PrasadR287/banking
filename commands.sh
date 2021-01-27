curl http://localhost:8000/greet

curl --header "Content-Type: application/json" http://localhost:8000/customers
curl --header "Content-Type: application/xml" http://localhost:8000/customers
# Not working: curl --header "Content-Type: text/csv" http://localhost:8000/customers 

# Kill process
lsof -i :8000 <port number>
kill -9 23020 <pid>