# Generating the Proto File
protoc -IdoingSSH/proto --go_out=doingSSH/ --go_opt=module=github.com/chaijmin/gRPC/doingSSH --go-grpc_out=doingSSH/ --go-grpc_opt=module=github.com/chaijmin/gRPC/doingSSH doingSSH/proto/doingSSH.proto

# For Testing 
go run doingSSH/server/*.go
go run doingSSH/client/*.go
