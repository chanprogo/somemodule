
```
cd proto/
mkdir ../proto_datasvr
protoc --go_out=plugins=grpc:../proto_datasvr *.proto
```