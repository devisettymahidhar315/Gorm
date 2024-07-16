## CRUD operations in GORM(using Mysql) implementing using hexagonal architecture.

## how to run the project.
### 1. Get the my sql image
``` bash
  docker pull mysql:latest
```
### 2. Run the docker image
``` bash
  docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest
```
### 3. run the main.go
``` bash
  go run main.go
```
## End Points for this Project
## CREATE
### if you are using the postman
### For Create, select the post method.
#### For Create a data ```http:localhost:8080/id/value```

## DELETE
### For delete, select the delete method
#### For delete a key ```http://localhost:8080/:id ```
#### For delete entire database ```http://localhost:8080/all```

## PRINTING
### For printing , select the get method
#### For Printing a key ```http:localhost:8080/:id```
#### For Printing all ```http:localhost:8080/print```
