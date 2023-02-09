```
> protoc -I="."  --go_out=. --go_opt=module="github.com/defeng-hub/mcube-demo" --go-grpc_out=. --go-grpc_opt=module="github.com/defeng-hub/mcube-demo" apps/book/pb/*.proto

> protoc-go-inject-tag -input=apps/*/*.pb.go

> mcube generate enum -p -m apps/*/*.pb.go
```