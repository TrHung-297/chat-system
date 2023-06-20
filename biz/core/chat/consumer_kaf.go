package chat

import (
	"fmt"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/kafka_client"
)

func (ctrl *ChatController) InstallKafkaConsumer() {
	client := kafka_client.GetKafkaClientInstance()
	go client.InstallConsumerGroup(ctrl, "Test-kafka", "chat-message")
}

func (ctrl *ChatController) MessageCallback(messageObj kafka_client.MessageKafka) {
	switch messageObj.Topic {
	case "chat-topic":
		fmt.Println("aaa")
	}
}

func (ctrl *ChatController) ErrorCallback(err error) {
	g_log.V(3).Infof("ChatController::ErrorCallback - Kafka client error: %+v", err)
}