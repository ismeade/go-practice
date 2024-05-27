package main

import (
	"fmt"
	"os"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var c MQTT.OnConnectHandler = func(client MQTT.Client) {
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := client.Subscribe("go-mqtt/#", 2, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("subscribe: go-mqtt/#")
}

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	//opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts := MQTT.NewClientOptions().AddBroker("tcp://39.98.156.243:1883")
	opts.SetCleanSession(true) // false 离线时的消息也能接受
	opts.SetClientID("vm01")
	opts.SetUsername("vm01")
	opts.SetPassword("123") // 需要mqtt服务器支持
	opts.SetDefaultPublishHandler(f)
	opts.SetOnConnectHandler(c)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer c.Disconnect(250)

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	//if token := c.Subscribe("go-mqtt/#", 2, nil); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}

	var sw sync.WaitGroup
	sw.Add(1)
	sw.Wait()

}
