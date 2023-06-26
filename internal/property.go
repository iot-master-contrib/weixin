package internal

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
	"time"
	"weixin/influx"
)

func SubscribeProperty(c mqtt.Client) {
	//报警消息
	c.Subscribe("alarm/+/+", 0, func(client mqtt.Client, message mqtt.Message) {
		topics := strings.Split(message.Topic(), "/")
		pid := topics[2]
		id := topics[3]
		var properties map[string]interface{}
		err := json.Unmarshal(message.Payload(), &properties)
		if err != nil {
			log.Panicln(err)
			return
		}
		tm := time.Now()
		influx.Insert(pid, id, properties, tm)
	})
}
