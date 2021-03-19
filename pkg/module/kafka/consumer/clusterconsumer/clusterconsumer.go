package clusterconsumer

import (
	"fmt"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
)

var ordRespConsumer *cluster.Consumer

func KafkaInit(addrs []string, groupID string, topicId string, superReboot bool) error {

	OrdRespTopic := fmt.Sprintf("order_resp_%s", topicId)

	var err error

	config := cluster.NewConfig()

	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	if superReboot {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	ordRespConsumer, err = cluster.NewConsumer(addrs, groupID, []string{OrdRespTopic}, config)
	if err != nil {
		glog.Errorln(err)
		return err
	}
	go chkConsumer(ordRespConsumer)
	return nil
}

func chkConsumer(c *cluster.Consumer) {
	errors := c.Errors()
	noti := c.Notifications()
	for {
		select {
		case err := <-errors:
			if err != nil {
				glog.Error("consumer error", err)
			}
		case <-noti:
		}
	}
}

func Close() {
	ordRespConsumer.Close()
}

func GetOrdRespChan() <-chan *sarama.ConsumerMessage {
	return ordRespConsumer.Messages()
}
