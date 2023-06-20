package server

import (
	"github.com/TrHung-297/chat-v2/biz/core"
	"github.com/TrHung-297/chat-v2/module/account"
	"github.com/TrHung-297/fountain/baselib/kafka_client"
	"github.com/TrHung-297/fountain/baselib/redis_client"
	"github.com/gofiber/fiber/v2"
)

type OpenIDServer struct {
	AccountModule *account.AccountModule
	controllers []core.CoreController
}

func InitializeApp(app *fiber.App) error {
	redis_client.InstallRedisClientManager()
	//sql_client.InstallSQLClientManager()
	//elastic_client.InstallElasticClientManager()
	//
	//mqtt_client.InstanceMQTTClientManager()
	kafka_client.InstallKafkaClient()
	//amqp_client.InstanceAMQPClientManager()
	//
	//g_etcd.InstanceEtcdManger("Server.Discovery")
	//
	//f_dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())

	core.InstallCoreControllers()

	account.Initialize(app)

	return nil
}
