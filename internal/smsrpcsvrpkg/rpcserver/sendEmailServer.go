package rpcserver

import (
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/proto_datasvr"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/service/mailService"

	"golang.org/x/net/context"
)

type SendEmailServer struct {
}

// 发送邮件通知
func (t *SendEmailServer) SendEmail(ctx context.Context, req *proto_datasvr.SendEmailRequest) (*proto_datasvr.SendEmailResponse, error) {
	sendEmail := &mailService.SendEmail{}
	code, err := sendEmail.SendEmail(req.Address, req.Subject, req.Body)

	rsp := &proto_datasvr.SendEmailResponse{}

	if code == 0 {
		rsp.Code = 0
		rsp.Msg = "success!"
		return rsp, nil
	}

	rsp.Code = 1
	rsp.Msg = "fail!"
	return rsp, err
}
