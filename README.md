## Simple gRPC in Go Language
---
### Steps to create gRPC API
- create main function
- create a domain user along with its proto file
- message name in proto cannot be the same as struct in todo
- run the genproto.sh file by running command
    ```bash
    ./genproto.sh . todo/todo.proto
    ```
- create the GetTodo function by adjusting the one in the todo_grpc file
- run the server
    ```bash
    go run cmd/main.go
    ```

### How to call gRPC API
- open Postman
- click file -> new
- choose gRPC
- fill in the URL with your gRPC server address
- in the select method section, select import file. then select the proto file in your project (in this case, select todo.proto)
- click next, and import (if you don't want to rename API's title)
- select the method you want to call
- invoke

![postman-grpc](https://user-images.githubusercontent.com/68763385/246731724-f6c60bf3-a200-403c-9de6-c74fb7264ee1.png)