package initialize

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/global"
	"thingue-launcher/server/web/handler"
)

func initMqttServer() {
	handler.MqttHandler.SetConfig("ws", provider.AppConfig.LocalServer.BindAddr, Server.listen)
	opts := new(mqtt.Options)
	opts.InlineClient = true
	mqttServer := mqtt.New(opts)
	_ = mqttServer.AddHook(new(auth.AllowHook), nil)
	_ = mqttServer.AddListener(handler.MqttHandler)
	_ = mqttServer.Serve()
	global.MQTT_SERVER = mqttServer
}
