package main

import (
	"fmt"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/kafka_client"
	"github.com/spf13/viper"
	"math/rand"
	"testing"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	env.SetupConfigEnv()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`Debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	} else {
		fmt.Println("Service RUN on PRODUCTION mode")
	}
}

type DataTmp struct {
	Data int64 `json:"data,omitempty"`
}

type ConsumerInstace struct{}

// ErrorCallback func;
func (c *ConsumerInstace) ErrorCallback(err error) {
	fmt.Printf("errCallback::Error - %+v\n", err)
}

// MessageCallback func;
func (c *ConsumerInstace) MessageCallback(msgObj kafka_client.MessageKafka) {
	fmt.Printf("procCallback::msgObj - %s\n", base.JSONDebugDataString(msgObj))
}

// Test that TestProducerPushMessage works
func TestProducerPushMessage(t *testing.T) {
	instance := kafka_client.InstallKafkaClient()
	if instance == nil {
		t.Errorf("TestKafkaConsumer - Error can not create kafka instance")
	}

	randInt := rand.Int63()
	tmp := DataTmp{
		Data: randInt,
	}

	dataJSON := base.JSONDebugDataString(tmp)
	msg := kafka_client.MessageKafka{
		Event:      "Test",
		ObjectJSON: dataJSON,
	}

	kafClient := kafka_client.GetKafkaClientInstance()

	consumerCallback := &ConsumerInstace{}
	go kafClient.InstallConsumerGroup(consumerCallback, "gtv_test", "Test")

	par, off, err := instance.ProducerPushMessage("gtv_test", msg)
	if err != nil {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage Error %+v while result expect nil", err)
	}

	if off == 0 {
		t.Errorf("TestProducerPushMessage - ProducerPushMessage offset is 0 while result expect greater 0")
	}

	_ = par
}