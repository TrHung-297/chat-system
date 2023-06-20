package chat

type ChatController struct {

}

func (ctrl *ChatController) AfterInstalledDone() {
	ctrl.InstallKafkaConsumer()
}

func (ctrl *ChatController) RegisterCallback(anotherController interface{}) {
}

