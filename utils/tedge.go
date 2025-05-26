package utils

import (
    "encoding/json"
    "fmt"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TedgePublisher wraps an MQTT client for tedge publishing.
type TedgePublisher struct {
    Client   mqtt.Client
    DeviceID string
}

// NewTedgePublisher connects to the local tedge MQTT broker.
func NewTedgePublisher(deviceID string) (*TedgePublisher, error) {
    opts := mqtt.NewClientOptions().
        AddBroker("tcp://localhost:1883").
        SetClientID("tedge-util-" + deviceID).
        SetCleanSession(true).
        SetConnectTimeout(5 * time.Second)

    client := mqtt.NewClient(opts)
    token := client.Connect()
    if token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }

    return &TedgePublisher{
        Client:   client,
        DeviceID: deviceID,
    }, nil
}

// PublishAlarm publishes an alarm (severity: "critical", "major", "minor", "warning").
func (t *TedgePublisher) PublishAlarm(alarmType, severity, text string) error {
    payload := map[string]interface{}{
        "text":     text,
        "severity": severity,
        "time":     time.Now().Format(time.RFC3339),
    }
    topic := fmt.Sprintf("te/device/%s/alarm/%s", t.DeviceID, alarmType)
    return t.publishJSON(topic, payload)
}

// PublishEvent publishes an event.
func (t *TedgePublisher) PublishEvent(eventType, text string) error {
    payload := map[string]interface{}{
        "text": text,
        "time": time.Now().Format(time.RFC3339),
    }
    topic := fmt.Sprintf("te/device/%s/event/%s", t.DeviceID, eventType)
    return t.publishJSON(topic, payload)
}

// PublishMeasurement publishes a measurement group (e.g. "temperature") and value.
func (t *TedgePublisher) PublishMeasurement(group string, measurements map[string]float64) error {
    payload := map[string]interface{}{
        "time": time.Now().Format(time.RFC3339),
        group: map[string]map[string]float64{},
    }

    nested := map[string]float64{}
    for k, v := range measurements {
        nested[k] = v
    }

    payload[group] = map[string]interface{}{}
    for k, v := range measurements {
        payload[group].(map[string]interface{})[k] = map[string]float64{"value": v}
    }

    topic := fmt.Sprintf("te/device/%s/measurements", t.DeviceID)
    return t.publishJSON(topic, payload)
}

// publishJSON publishes the given payload to the topic as JSON.
func (t *TedgePublisher) publishJSON(topic string, payload interface{}) error {
    bytes, err := json.Marshal(payload)
    if err != nil {
        return err
    }
    token := t.Client.Publish(topic, 0, false, bytes)
    token.Wait()
    return token.Error()
}
