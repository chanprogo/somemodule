package rpcclient

import (
	"testing"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/proto_datasvr"
	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/rpcserver"
	"github.com/chanprogo/somemodule/pkg/conf/bconf"
	"github.com/chanprogo/somemodule/pkg/log"
)

func init() {
	bconf.NewConfiger("../../../cmd/smsrpcsvr/app.conf")
	log.NewLogger(bconf.Configer.ApiConf.LogPath, bconf.Configer.ApiConf.LogLevel)
	app.NewRPCSvr(bconf.Configer.ApiConf.RpcPort)
	proto_datasvr.RegisterEmailServiceServer(app.RPCSvr, &rpcserver.SendEmailServer{})
}

func TestRpcClient(t *testing.T) {
	RpcClient()
}
