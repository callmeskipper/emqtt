package main

import (
	"os"
	"time"
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}


func main(){
	opts := mqtt.NewClientOptions().AddBroker("118.25.65.20:1883")
	opts.SetClientID("clientA")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	time.Sleep(3 * time.Minute)
	c.Unsubscribe("go-mqtt/sample")
	c.Disconnect(250)
}