package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	//opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts := MQTT.NewClientOptions().AddBroker("tcp://39.98.156.243:1883")
	opts.SetCleanSession(true)
	opts.SetClientID("vm02")
	opts.SetUsername("vm02")
	opts.SetPassword("123") // 需要mqtt服务器支持

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer c.Disconnect(250)

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for {
		var input string
		fmt.Scanln(&input)
		token := c.Publish("go-mqtt/sample", 0, false, input)
		token.Wait()
	}

	// time.Sleep(3 * time.Second)

}
