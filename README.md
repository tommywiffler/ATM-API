# ATM-API

## Run Program

Necessary Tools Installed
- Git
- Golang

To run this program, clone this repository to your local machine, navigate to this repository in your working directory, and run following command in your terminal:
```
go run main.go
```

## Test

To test the endpoints, use an API Client such as Postman.

The URI definitions and methods are listed in the /handlers/routes.go file.

Each endpoint in your API Client should begin with:
```
http://127.0.0.1:8080/
```
Append the relevant URI definition for the endpoint you wish to access.

2 Example Users with Accounts have been made for testing. The sample data is hardcoded into the database.go file.

For User Login, the request body should read:
```
{
    "id": "1234"
}
```
or 
```
{
    "id": "4141"
}
```

To view the user's bank balance, place the user's 6 digit account number at the end of the Bank Balance endpoint.


For the Bank Transaction related endpoints, the request body should read:
```
{
    "number": "<AccNumber returned from user login>",
    "balance": <amount to deposit or withdraw>
}
```
