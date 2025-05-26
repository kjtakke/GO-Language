package utils

import (
    "fmt"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTClient wraps an active MQTT client instance.
type MQTTClient struct {
    client mqtt.Client
}

// ConnectMQTT establishes a connection to the MQTT broker and returns an MQTTClient.
func ConnectMQTT(brokerURL string, clientID string) (*MQTTClient, error) {
    opts := mqtt.NewClientOptions().
        AddBroker(brokerURL).
        SetClientID(clientID).
        SetCleanSession(true).
        SetConnectTimeout(5 * time.Second)

    client := mqtt.NewClient(opts)
    token := client.Connect()
    if token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }

    return &MQTTClient{client: client}, nil
}

// Publish sends a message to the given topic with optional QoS and retain flag.
func (m *MQTTClient) Publish(topic string, payload string, qos byte, retain bool) error {
    token := m.client.Publish(topic, qos, retain, payload)
    token.Wait()
    return token.Error()
}

// Subscribe registers a handler for messages received on a topic.
func (m *MQTTClient) Subscribe(topic string, qos byte, callback func(topic string, payload string)) error {
    handler := func(client mqtt.Client, msg mqtt.Message) {
        callback(msg.Topic(), string(msg.Payload()))
    }

    token := m.client.Subscribe(topic, qos, handler)
    token.Wait()
    return token.Error()
}

// Disconnect cleanly disconnects from the broker.
func (m *MQTTClient) Disconnect() {
    m.client.Disconnect(250)
    fmt.Println("Disconnected from MQTT broker")
}
