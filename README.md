# txListener-go

### Setup

1- Rename `conf.json.copy` to `conf.json`

2- update configuration variables in conf.json file

3- vault setup : https://www.vaultproject.io/ 

4- run "export VAULT_ADDR='http://127.0.0.1:8200'"

5- run "export VAULT_TOKEN='your token here' "

6- store mysql db password and admin pk in vault 

7- use `ddl.sql` file to create database tables. 

8- run code using `go run main.go`

9- Run Rest server using go run rest/server.go


