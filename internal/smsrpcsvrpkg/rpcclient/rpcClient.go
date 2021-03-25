package rpcclient

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"fmt"

	pb "github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/proto_datasvr"
)

func RpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8781", grpc.WithInsecure())
	if err != nil {
		// log.Logger.Error("did not connect: %v", err)
		fmt.Println("aa bb cc")
	}
	defer conn.Close()

	c := pb.NewEmailServiceClient(conn)

	subject := "MySubject"
	body := "This is body."
	address := []string{"2483777000@qq.com"}
	fmt.Println("0 " + body)
	r, err := c.SendEmail(context.Background(), &pb.SendEmailRequest{Address: address, Subject: subject, Body: body})
	if err != nil {
		// log.Logger.Error("could not send sms: %v", err)
		fmt.Printf("4 could not send sms: %v \n\n", err)
	}
	fmt.Println("r.code:", r.Code)
}
