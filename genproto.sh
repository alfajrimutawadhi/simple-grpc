# ref : https://protobuf.dev/reference/go/go-generated/
mkdir -p $1
protoc --go_out=$1 --go_opt=paths=source_relative \
    --go-grpc_out=$1 --go-grpc_opt=paths=source_relative \
    $2

# param 1 = path of project
# param 2 = proto file

# for this example, run below code to compile todo:
# ./genproto.sh . todo/todo.proto