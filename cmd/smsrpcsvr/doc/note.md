

`cd proto/`  



`protoc --go_out=plugins=grpc:../proto_datasvr *.proto`   



`protoc -I=/home/chan/Desktop/chanprogo/somemodule --go_out=/home/chan/Desktop/chanprogo/somemodule internal/smsrpcsvrpkg/proto/*.proto`



`protoc -I=/home/chan/Desktop/chanprogo/somemodule --go_out=plugins=grpc:/home/chan/Desktop/chanprogo/somemodule internal/smsrpcsvrpkg/proto/*.proto`


protoc   

--go_out=plugins=grpc:../proto_datasvr   

*.proto  