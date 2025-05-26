# Table of Contents

- [Files](#files)
- [Base64](#base64)
- [MQTT](#mqtt)
- [HTTP](#http)
- [JSON](#json)
- [Bash](#bash)
- [Mutex](#mutex)
- [Logging](#logging)
  - [Tedge](#tedge)

---

## Files

| Function | Description |
|----------|-------------|
| `ReadFile` | Reads the entire content of a file and returns it as a string. |
| `ReadLines` | Reads a file line-by-line and returns the contents as a slice of strings. |
| `WriteFile` | Writes the specified string content to a file, overwriting it if it exists. |
| `AppendToFile` | Appends string content to a file, or creates the file if it doesn't exist. |
| `FileExists` | Returns `true` if a file or directory exists at the given path, otherwise `false`. |
| `ListFiles` | Lists all file names (excluding directories) in a specified directory. |
| `FindFilesByExtension` | Recursively searches a directory for files matching a given file extension. |

---

## Base64

| Function | Description |
|----------|-------------|
| `EncodeStringToBase64` | Encodes a plain string to a Base64 string. |
| `DecodeBase64ToString` | Decodes a Base64 string to a plain string. |
| `EncodeJSONToBase64` | Marshals a Go value to JSON and encodes it as Base64. |
| `DecodeBase64JSONToMap` | Decodes a Base64-encoded JSON string into a `map[string]interface{}`. |
| `DecodeBase64JSONToStruct` | Decodes a Base64 JSON string into a struct pointer. |

---

## MQTT

| Function | Description |
|----------|-------------|
| `ConnectMQTT` | Connects to an MQTT broker and returns a client wrapper. |
| `Publish` | Publishes a message to a given topic with QoS and retain options. |
| `Subscribe` | Subscribes to a topic with a callback for handling incoming messages. |
| `Disconnect` | Cleanly disconnects from the MQTT broker. |

---

## HTTP

| Function | Description |
|----------|-------------|
| `HTTPRequest` | Makes an HTTP request with custom method, headers, and body. |
| `HTTPGet` | Sends an HTTP GET request. |
| `HTTPPost` | Sends an HTTP POST request with JSON body. |
| `HTTPPut` | Sends an HTTP PUT request with JSON body. |
| `HTTPPatch` | Sends an HTTP PATCH request with JSON body. |
| `HTTPDelete` | Sends an HTTP DELETE request. |

---

## JSON

| Function | Description |
|----------|-------------|
| `ToJSON` | Marshals a Go value into JSON bytes. |
| `ToJSONString` | Marshals a Go value into a pretty-printed JSON string. |
| `FromJSON` | Unmarshals JSON bytes into a Go value (pointer required). |
| `FromJSONString` | Unmarshals a JSON string into a Go value (pointer required). |
| `FromJSONToMap` | Converts JSON bytes into a map. |
| `SafeGetString` | Safely retrieves a string from a map. |
| `SafeGetNumber` | Safely retrieves a float64 from a map. |
| `SafeGetBool` | Safely retrieves a bool from a map. |
| `FlattenJSON` | Flattens nested JSON maps into key-value string pairs. |

---

## Bash

| Function | Description |
|----------|-------------|
| `RunCommand` | Executes a single Bash command string. |
| `RunScript` | Executes a Bash script with optional arguments. |
| `RunCommandWithEnv` | Executes a Bash command with additional environment variables. |
| `RunAndCapture` | Executes a command and returns separated stdout and stderr. |

---

## Mutex

| Function | Description |
|----------|-------------|
| `NewSafeQueue` | Creates a new thread-safe queue. |
| `Enqueue` | Adds an item to the queue. |
| `Dequeue` | Removes and returns the first item in the queue. |
| `Peek` | Returns the first item without removing it. |
| `Length` | Returns the number of items in the queue. |
| `Clear` | Empties the queue. |
| `Values` | Returns a copy of the current queue contents. |

---

## Logging

| Function | Description |
|----------|-------------|
| `NewLogger` | Creates a new namespaced logger. |
| `Infof` | Logs an info-level message. |
| `Debugf` | Logs a debug-level message. |
| `Warnf` | Logs a warning message. |
| `Errorf` | Logs an error message. |
| `Criticalf` | Logs a critical error message. |

---

### Tedge

| Function | Description |
|----------|-------------|
| `NewTedgePublisher` | Connects to the local tedge MQTT broker. |
| `PublishAlarm` | Publishes an alarm message to the tedge broker. |
| `PublishEvent` | Publishes an event message to the tedge broker. |
| `PublishMeasurement` | Publishes measurement data as a nested structure. |
| `publishJSON` | Internal method to publish JSON-formatted data to MQTT. |
