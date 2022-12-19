package main

import (
	"fmt"
	"os"
	"time"

	mediaDevices "github.com/davidventura/go-media-devices-state"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	broker := "iot.labs"
	port := 1883
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(host)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mediaDevices.InitDevices()
	lastCameraValue := false
	lastMicValue := false
	micTopic := fmt.Sprintf("/microphone/%s", host)
	camTopic := fmt.Sprintf("/camera/%s", host)
	valueMap := map[bool]string{
		true:  "1",
		false: "0",
	}
	for {
		isCameraOn, err := mediaDevices.IsCameraOn()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			//fmt.Println("Is camera on:", isCameraOn)
		}

		isMicrophoneOn, err := mediaDevices.IsMicrophoneOn()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			//fmt.Println("Is microphone on:", isMicrophoneOn)
		}

		if isCameraOn != lastCameraValue {
			lastCameraValue = isCameraOn
			token := client.Publish(camTopic, 0, false, valueMap[isCameraOn])
			token.Wait()
		}

		if isMicrophoneOn != lastMicValue {
			lastMicValue = isMicrophoneOn
			token := client.Publish(micTopic, 0, false, valueMap[isMicrophoneOn])
			token.Wait()
		}

		time.Sleep(20 * time.Millisecond)
	}
}
