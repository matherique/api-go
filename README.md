# REST API in Go

## Folder Struct

### `cmd/api`
This folder has all the necessary code to create the RESTFull API. Its
receive the request, parse it, pass to app function to handle it and return
the error or success in json

### `internal/app`
All business logic, receive the necessary data and return the response
or error 

### `internal/storage` 
Contains all the function that handle the communication with the
database using, all the DQL's, DML's, etc

### `internal/models` 
Data models used in the `internal/app` and `internal/storage`


### `pkg/database` 
contains the database connection and helpers 

### `pkg/logger` 
Create a new logger with a prefix and the application standard form
