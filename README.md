Please make sure that you have Postgres installed in your system.

Clone the repo
Go to repository root, 
Run the program: 
$ go run main.go

Please note that a table named `employee` will get created in local Postgres setup (localhost:5432).
And use some tool like pgAdmin to insert 1 or more records to the table.

Try the request:
GET http://localhost:3002/sk-em/employee
Sample response:
[{"id":1,"name":"Arun","age":25,"address":"Palakkad, Kerala","designation":"Software Developer","joining_date":"2021-10-01T00:00:00Z"}]
