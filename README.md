## how to run the project.
### 1. Getting the my sql image
``` bash docker pull mysql:latest
```
### 2. run the docker image
``` bash docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest
```
### 3. run the main.go
``` bash go run main.go
```
