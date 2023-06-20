package chat

import (
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/kafka_client"
)

func PublishToTopic(topic string, message string) error {
	client := kafka_client.GetKafkaClientInstance()

	kafkaMsg := kafka_client.MessageKafka{
		Event:      topic,
		ObjectJSON: message,
	}

	g_log.V(3).Infof("[kafka-v2] - Publish message, Topic:%s, Message:%s", topic, message)
	_, _, err := client.ProducerPushMessage(topic, kafkaMsg)
	if err != nil {
		g_log.V(3).Errorf("[kafka-v2] - Publish message to topic exception: %s, Topic: %s, Msg: %s", err, topic, message)
	}
	return err
}
