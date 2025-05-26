# 📚 Table of Contents

## Loops in Go

- [Loop (While)](#loop-while)
- [Loop (Int Range)](#loop-int-range)
- [Loop (List or Array)](#loop-list-or-array)
- [Loop (Dictionary or Map)](#loop-dictionary-or-map)

## If Statements in Go

- [If](#if)
- [If and (&&)](#if-and)
- [If or (or)](#if-or)
- [If and/or (Nested Logic)](#if-andor-nested-logic)
- [If else if](#if-else-if)
- [If else](#if-else)

## Switch Statements
- [Basic Switch](#basic-switch)
- [Switch With Multiple Values Per Case](#switch-with-multiple-values-per-case)
- [Switch Without Expression](#switch-without-expression)
- [Switch With Fallthrough](#switch-with-fallthrough)
- [Switch With Type Assertion (Type Switch)](#switch-with-type-assertion-type-switch)

## Dictionaries in Go (Maps)

- [Create Empty Dict](#create-empty-dict)
- [Create Populated Dict](#create-populated-dict)
- [Reference Key Value in Dict](#reference-key-value-in-dict)
- [Remove Key Value Pair](#remove-key-value-pair)

## Arrays in Go (Slices)

- [Create Empty Array](#create-empty-array)
- [Create Populated Array](#create-populated-array)
- [Reference First Value in Array](#reference-first-value-in-array)
- [Reference Last Value in Array](#reference-last-value-in-array)
- [Reference Second Last Value in Array](#reference-second-last-value-in-array)
- [Reference Nested Array Item](#reference-nested-array-item)
- [Remove Referenced Value](#remove-referenced-value)

## File Handling in Go

- [Get File Contents](#get-file-contents)
- [Loop Through Lines](#loop-through-lines)
- [Convert JSON File to Dict](#convert-json-file-to-dict)

## Base64 Handling in Go

- [Decode Base64 to String](#decode-base64-to-string)
- [Encode String to Base64](#encode-string-to-base64)
- [Decode Base64 JSON to Dict](#decode-base64-json-to-dict)
- [Encode Dict to JSON then Base64](#encode-dict-to-json-then-base64)

## HTTP Requests in Go

- [Basic HTTP Client Setup](#basic-http-client-setup)
- [Headers](#headers)
- [Auth (Basic Authentication)](#auth-basic-authentication)
- [GET](#get)
- [POST](#post)
- [PUT](#put)
- [PATCH](#patch)
- [DELETE](#delete)

## Routes in Go
- [Setup](#setup)
- [Add Auth Requirements per Route](#add-auth-requirements-per-route)
- [GET](#get-1)
- [POST](#post-1)
- [PATCH](#patch-1)
- [PUT](#put-1)
- [DELETE](#delete-1)
- [URL Parameters and Parse to Dict](#url-parameters-and-parse-to-dict)
- [500, 400, 401... Exceptions](#500-400-401-exceptions)

## Enumerators
- [Define Basic Enumerator Using iota](#define-basic-enumerator-using-iota)
- [Enumerator With Explicit Values](#enumerator-with-explicit-values)
- [Enumerator With Skipped Values](#enumerator-with-skipped-values)
- [Typed Enum (Using Custom Type)](#typed-enum-using-custom-type)
- [Enum to String Conversion (Stringer Pattern)](#enum-to-string-conversion-stringer-pattern)
- [Enum From String (Reverse Lookup)](#enum-from-string-reverse-lookup)
- [Use Enum in Switch Statement](#use-enum-in-switch-statement)
- [Use Enum in Struct Fields](#use-enum-in-struct-fields)
- [Enum Validation and Zero Value Handling](#enum-validation-and-zero-value-handling)
- [Grouped Constant Blocks (Multiple Enums)](#grouped-constant-blocks-multiple-enums)
- [Enum With Flags (Bitmasking via iota << n)](#enum-with-flags-bitmasking-via-iota--n)

## Bash Commands and Run Bash Scripts in Go
- [Run a Simple Bash Command](#run-a-simple-bash-command)
- [Run a Bash Script File](#run-a-bash-script-file)
- [Pass Arguments to a Bash Script](#pass-arguments-to-a-bash-script)
- [Capture Standard Output and Error Separately](#capture-standard-output-and-error-separately)
- [Run Bash Commands with Environment Variables](#run-bash-commands-with-environment-variables)

## Mutex (Queuing) in Go
- [Basic Mutex Lock/Unlock](#basic-mutex-lockunlock)
- [Mutex With Goroutines (Safe Queuing)](#mutex-with-goroutines-safe-queuing)
- [Read and Modify Shared Data Safely](#read-and-modify-shared-data-safely)

## MQTT (Mosquitto) in Go
- [Setup](#setup-1)
- [Add Message](#add-message)
- [Listen for Topic](#listen-for-topic)
- [Get Pending List](#get-pending-list)
- [Remove Item from List](#remove-item-from-pending)
- [Usage Example](#usage-example)

## Loops in Go

Go uses a single looping construct: `for`. It can behave like a `while`, traditional `for`, or range-based loop. There is no `while` or `foreach` keyword in Go—`for` handles all cases.

### Loop (While)

This mimics the `while` loop from other languages using `for` with a condition.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println("i is", i)
        i++
    }
}

```

### Loop (Int Range)

Use a basic `for` loop with initialisation, condition, and post statement to iterate over a numeric range.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Number:", i)
    }
}

```

### Loop (List / Array)

Use `range` to iterate over slices or arrays. You receive both the index and the value.

```go
package main

import "fmt"

func main() {
    items := []string{"apple", "banana", "cherry"}

    for idx, val := range items {
        fmt.Printf("Index %d: %s\n", idx, val)
    }
    
    //You can ignore the index with _ if only the value is needed:
    for _, val := range items {
    	fmt.Println(val)
    }

}

```

### Loop (Dictionary / Map)

Use `range` to iterate over maps. Each iteration provides a key-value pair.

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Australia": "Canberra",
        "Japan":     "Tokyo",
        "France":    "Paris",
    }

    for country, capital := range capitals {
        fmt.Printf("The capital of %s is %s\n", country, capital)
    }
    // You can ignore the value if only keys are needed:
    for country := range capitals {
        fmt.Println("Country:", country)
    }
}

```

## **🔍 If Statements in Go**

Go supports standard `if`, `if-else`, and `if-else if` conditional logic using concise syntax. Conditions must evaluate to a `bool` and do not require parentheses. Curly braces `{}` are mandatory.

### If

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}

```

### If and (`&&`)

```go
package main

import "fmt"

func main() {
    a := 7
    if a > 5 && a < 10 {
        fmt.Println("a is between 6 and 9")
    }
}

```

### If or (`||`)

```go
package main

import "fmt"

func main() {
    b := 3
    if b < 5 || b > 10 {
        fmt.Println("b is either less than 5 or greater than 10")
    }
}

```

### If and/or (nested logic)

```go
package main

import "fmt"

func main() {
    x, y := 8, 3
    if (x > 5 && y < 5) || x == 10 {
        fmt.Println("Condition met")
    }
}

```

### If else if

```go
package main

import "fmt"

func main() {
    score := 85

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    }
}

```

### If else

```go
package main

import "fmt"

func main() {
    isActive := true

    if isActive {
        fmt.Println("The system is active")
    } else {
        fmt.Println("The system is inactive")
    }
}

```

## Switch Statements
Go's switch statement provides a concise way to express conditional logic. Unlike some other languages, Go’s switch cases do not require break—each case exits automatically unless you use fallthrough.
### Basic Switch
```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("Almost weekend")
    default:
        fmt.Println("Midweek day")
    }
}

```
### Switch With Multiple Values Per Case
```go
package main

import "fmt"

func main() {
    char := 'a'

    switch char {
    case 'a', 'e', 'i', 'o', 'u':
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }
}

```
### Switch Without Expression (acts like if-else if chain)
```go
package main

import "fmt"

func main() {
    num := 42

    switch {
    case num < 0:
        fmt.Println("Negative number")
    case num == 0:
        fmt.Println("Zero")
    case num > 0:
        fmt.Println("Positive number")
    }
}

```
### Switch With Fallthrough
```go
package main

import "fmt"

func main() {
    value := 1

    switch value {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two (also printed because of fallthrough)")
    default:
        fmt.Println("Default case")
    }
}

```
### Switch With Type Assertion (Type Switch)
```go
package main

import "fmt"

func identify(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    identify(10)
    identify("hello")
    identify(true)
}

```


## Dictionaries in Go (Maps)

Go provides a built-in `map` type to represent key-value pairs. Keys must be comparable types (e.g. `string`, `int`, etc.), and values can be of any type.

### Create Empty Dict

```go
package main

import "fmt"

func main() {
    dict := make(map[string]string)
    fmt.Println("Empty dictionary:", dict)
    // Alternatively, declare with var and initialise later:
    var dict map[string]string
    dict = make(map[string]string)
}

```

### Create Populated Dict

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Australia": "Canberra",
        "Japan":     "Tokyo",
        "France":    "Paris",
    }

    fmt.Println("Capitals:", capitals)
}

```

### Reference Key Value in Dict

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Australia": "Canberra",
        "Japan":     "Tokyo",
    }

    capital := capitals["Japan"]
    fmt.Println("Capital of Japan is", capital)

    // Optional existence check
    if val, ok := capitals["France"]; ok {
        fmt.Println("Capital of France is", val)
    } else {
        fmt.Println("France not found")
    }
}

```

### Remove Key Value Pair

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Australia": "Canberra",
        "Japan":     "Tokyo",
        "France":    "Paris",
    }

    delete(capitals, "France")
    fmt.Println("After deletion:", capitals)
}

```

## Arrays in Go (Slices)

### Create Empty Array (strings, int, floats, nested arrays/lists)

```go
package main

import "fmt"

func main() {
    // Empty slices (preferred over arrays in Go)
    var strSlice []string
    var intSlice []int
    var floatSlice []float64

    // Nested slices
    var nestedIntSlice [][]int
    var nestedStrSlice [][]string

    fmt.Println("Empty string slice:", strSlice)
    fmt.Println("Empty int slice:", intSlice)
    fmt.Println("Empty float slice:", floatSlice)
    fmt.Println("Nested int slice:", nestedIntSlice)
    fmt.Println("Nested string slice:", nestedStrSlice)
}

```

### Create Populated Array

```go
package main

import "fmt"

func main() {
    names := []string{"Alice", "Bob", "Charlie"}
    numbers := []int{10, 20, 30, 40}
    matrix := [][]int{{1, 2}, {3, 4}}

    fmt.Println("Names:", names)
    fmt.Println("Numbers:", numbers)
    fmt.Println("Matrix:", matrix)
}

```

### Reference First Value in Array

```go
package main

import "fmt"

func main() {
    names := []string{"Alice", "Bob", "Charlie"}
    fmt.Println("First name:", names[0])
}

```

### Reference Last Value in Array

```go
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40}
    fmt.Println("Last number:", numbers[len(numbers)-1])
}

```

### Reference Second Last Value in Array

```go
package main

import "fmt"

func main() {
    items := []string{"red", "green", "blue", "yellow"}
    fmt.Println("Second last item:", items[len(items)-2])
}

```

### Reference Nested Array Item in Array

```go
package main

import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    fmt.Println("Item at [1][2]:", matrix[1][2]) // Outputs 6
}

```

### Remove Referenced Value

Go doesn't have built-in deletion from slices; use slicing and `append()` idiomatically.

```go
package main

import "fmt"

func main() {
    fruits := []string{"apple", "banana", "cherry", "date"}

    // Remove "banana" (index 1)
    indexToRemove := 1
    fruits = append(fruits[:indexToRemove], fruits[indexToRemove+1:]...)

    fmt.Println("After removal:", fruits)
}

```

## File Handling in Go

File reading in Go is handled through the `os`, `bufio`, and `io/ioutil` or `os.ReadFile` packages, while JSON decoding is handled via the `encoding/json` package.

### Get File Contents

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("File contents:\n", string(data))
}

```

### Loop Through Lines

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println("Line:", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Scanner error:", err)
    }
}

```

### Convert JSON File to Dict (Map)

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("data.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(data, &result); err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return
    }

    fmt.Println("JSON as dictionary:", result)
}

```

## Base64 Handling in Go

The `encoding/base64` and `encoding/json` packages are used to encode and decode data between base64 and Go types.

### Decode Base64 to String

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    encoded := "SGVsbG8gd29ybGQh" // "Hello world!"
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Decode error:", err)
        return
    }

    fmt.Println("Decoded string:", string(decodedBytes))
}

```

### Encode String to Base64

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    raw := "Hello world!"
    encoded := base64.StdEncoding.EncodeToString([]byte(raw))

    fmt.Println("Base64 encoded string:", encoded)
}

```

### Decode Base64 String (JSON) to Dict

```go
package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
)

func main() {
    encoded := "eyJmb28iOiAiYmFyIiwgIm51bSI6IDQyfQ==" // {"foo": "bar", "num": 42}
    
    decodedJSON, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Base64 decode error:", err)
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(decodedJSON, &result); err != nil {
        fmt.Println("JSON unmarshal error:", err)
        return
    }

    fmt.Println("Decoded JSON to dict:", result)
}

```

### Encode Dict to JSON, then Base64

```go
package main

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
)

func main() {
    data := map[string]interface{}{
        "foo": "bar",
        "num": 42,
    }

    jsonBytes, err := json.Marshal(data)
    if err != nil {
        fmt.Println("JSON marshal error:", err)
        return
    }

    encoded := base64.StdEncoding.EncodeToString(jsonBytes)
    fmt.Println("Base64-encoded JSON string:", encoded)
}

```

## HTTP Requests (Go)

Go provides robust HTTP client and server capabilities in the `net/http` package. Below are examples for common request types and features.

### Basic HTTP Client Setup

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://example.com")
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response:", string(body))
}

```

### Headers

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://example.com", nil)
    if err != nil {
        fmt.Println("Request creation error:", err)
        return
    }

    req.Header.Set("User-Agent", "GoClient/1.0")
    req.Header.Set("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    fmt.Println("Response with headers:", string(body))
}

```

### Auth (Basic Authentication)

```go
package main

import (
    "encoding/base64"
    "fmt"
    "io"
    "net/http"
)

func main() {
    req, err := http.NewRequest("GET", "https://example.com", nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    username := "user"
    password := "pass"
    auth := username + ":" + password
    encoded := base64.StdEncoding.EncodeToString([]byte(auth))
    req.Header.Set("Authorization", "Basic "+encoded)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("Authenticated response:", string(body))
}

```

### GET

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        fmt.Println("GET request error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("GET response:", string(body))
}

```

### POST

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "net/http"
)

func main() {
    jsonData := []byte(`{"name": "Alice", "age": 30}`)
    resp, err := http.Post("https://api.example.com/create", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("POST request error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("POST response:", string(body))
}

```

### PUT

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "net/http"
)

func main() {
    client := &http.Client{}
    data := []byte(`{"id": 1, "name": "Updated Name"}`)

    req, err := http.NewRequest("PUT", "https://api.example.com/resource/1", bytes.NewBuffer(data))
    if err != nil {
        fmt.Println("Error creating PUT request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("PUT request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("PUT response:", string(body))
}

```

### PATCH

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "net/http"
)

func main() {
    client := &http.Client{}
    data := []byte(`{"name": "Patched Name"}`)

    req, err := http.NewRequest("PATCH", "https://api.example.com/resource/1", bytes.NewBuffer(data))
    if err != nil {
        fmt.Println("Error creating PATCH request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("PATCH request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("PATCH response:", string(body))
}

```

### DELETE

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", "https://api.example.com/resource/1", nil)
    if err != nil {
        fmt.Println("Error creating DELETE request:", err)
        return
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("DELETE request failed:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("DELETE response:", string(body))
}

```

## Routes in Go

For production-grade routing and middleware, it is common to use third-party routers like `gorilla/mux`, `chi`, or `gin`. However, these examples will use only the Go standard library (`net/http`) to remain tool-agnostic.

### Setup

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

    http.ListenAndServe(":8080", nil)
}

```

### Add Auth Requirements per Route

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/secure", requireAuth(secureHandler))
    http.ListenAndServe(":8080", nil)
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") != "Bearer mysecrettoken" {
            http.Error(w, "Unauthorised", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Authenticated access granted")
}

```

### GET

```go
func handleGet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w, "GET request handled")
}

```

### POST

```go
func handlePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w, "POST request handled")
}

```

### PATCH

```go
func handlePatch(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w, "PATCH request handled")
}

```

### PUT

```go
func handlePut(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w, "PUT request handled")
}

```

### DELETE

```go
func handleDelete(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w, "DELETE request handled")
}

```

### URL Paramaters and parse to Dict

```go
func handleQueryParams(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    params := map[string]string{
        "id":   query.Get("id"),
        "name": query.Get("name"),
    }

    fmt.Fprintf(w, "Parsed query params: %+v\n", params)
}

```

### 500, 400, 401, ... exceptions

```go
func errorHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Query().Get("status") {
    case "400":
        http.Error(w, "Bad Request", http.StatusBadRequest)
    case "401":
        http.Error(w, "Unauthorised", http.StatusUnauthorized)
    case "403":
        http.Error(w, "Forbidden", http.StatusForbidden)
    case "404":
        http.Error(w, "Not Found", http.StatusNotFound)
    case "500":
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    default:
        http.Error(w, "Unknown error code", http.StatusTeapot) // 418 for humour
    }
}

```

## Enumerators

### Define Basic Enumerator Using iota
```go
package main

import "fmt"

const (
    Red = iota
    Green
    Blue
)

func main() {
    fmt.Println("Red:", Red)
    fmt.Println("Green:", Green)
    fmt.Println("Blue:", Blue)
}

```

### Enumerator With Explicit Values
```go
package main

import "fmt"

const (
    Sunday = 0
    Monday = 1
    Tuesday = 2
)

func main() {
    fmt.Println("Monday:", Monday)
}

```
### Enumerator With Skipped Values
```go
package main

import "fmt"

const (
    A = iota
    _ // skip value 1
    C
)

func main() {
    fmt.Println("A:", A)
    fmt.Println("C:", C)
}

```
### Typed Enum (Using Custom Type)
```go
package main

import "fmt"

type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

func main() {
    var s Status = Approved
    fmt.Println("Status:", s)
}

```
### Enum to String Conversion (Stringer Pattern)
```go
package main

import "fmt"

type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

func (s Status) String() string {
    switch s {
    case Pending:
        return "Pending"
    case Approved:
        return "Approved"
    case Rejected:
        return "Rejected"
    default:
        return "Unknown"
    }
}

func main() {
    fmt.Println("String:", Approved.String())
}

```
### Enum From String (Reverse Lookup)
```go
package main

import (
    "errors"
    "fmt"
)

type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

func ParseStatus(s string) (Status, error) {
    switch s {
    case "Pending":
        return Pending, nil
    case "Approved":
        return Approved, nil
    case "Rejected":
        return Rejected, nil
    default:
        return 0, errors.New("invalid status")
    }
}

func main() {
    s, err := ParseStatus("Approved")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Parsed status:", s)
    }
}

```
### Use Enum in Switch Statement
```go
package main

import "fmt"

type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

func describe(s Status) {
    switch s {
    case Pending:
        fmt.Println("Still processing")
    case Approved:
        fmt.Println("Approved!")
    case Rejected:
        fmt.Println("Rejected.")
    default:
        fmt.Println("Unknown status")
    }
}

func main() {
    describe(Pending)
}

```
### Use Enum in Struct Fields
```go
package main

import "fmt"

type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

type Application struct {
    ID     int
    Name   string
    Status Status
}

func main() {
    app := Application{ID: 1, Name: "Loan", Status: Approved}
    fmt.Printf("Application: %+v\n", app)
}

```
### Enum Validation and Zero Value Handling
```go
package main

import "fmt"

type Status int

const (
    Unknown Status = iota
    Pending
    Approved
    Rejected
)

func IsValidStatus(s Status) bool {
    return s >= Pending && s <= Rejected
}

func main() {
    var s Status
    fmt.Println("Initial zero value is valid?", IsValidStatus(s))
}

```
### Grouped Constant Blocks (Multiple Enums)
```go
package main

import "fmt"

type Priority int
type State int

const (
    Low Priority = iota
    Medium
    High
)

const (
    Open State = iota
    InProgress
    Closed
)

func main() {
    fmt.Println("Priority High:", High)
    fmt.Println("State Closed:", Closed)
}

```
### Enum With Flags (Bitmasking via iota << n)
```go
package main

import "fmt"

type Permission uint

const (
    Read Permission = 1 << iota
    Write
    Execute
)

func main() {
    var p Permission = Read | Execute
    fmt.Println("Permissions:", p)

    if p&Read != 0 {
        fmt.Println("Has read permission")
    }
    if p&Write != 0 {
        fmt.Println("Has write permission")
    }
    if p&Execute != 0 {
        fmt.Println("Has execute permission")
    }
}

```

## Bash Commands and Run Bash Scripts in Go
The os/exec package allows Go programs to run external commands, including bash scripts and inline shell commands.

### Run a Simple Bash Command
```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("bash", "-c", "echo Hello from Bash")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Command failed:", err)
        return
    }
    fmt.Println("Output:", string(output))
}

```
### Run a Bash Script File
```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("bash", "/path/to/script.sh")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Script execution failed:", err)
        return
    }
    fmt.Println("Script output:", string(output))
}

```
### Pass Arguments to a Bash Script
```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("bash", "/path/to/script.sh", "arg1", "arg2")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Script execution failed:", err)
        return
    }
    fmt.Println("Output with args:", string(output))
}

```
### Capture Standard Output and Error Separately
```go
package main

import (
    "bytes"
    "fmt"
    "os/exec"
)

func main() {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", "ls -la /nonexistent")
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()
    fmt.Println("STDOUT:", stdout.String())
    fmt.Println("STDERR:", stderr.String())

    if err != nil {
        fmt.Println("Error:", err)
    }
}

```
### Run Bash Commands with Environment Variables
```go
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("bash", "-c", "echo $MY_VAR")
    cmd.Env = append(cmd.Env, "MY_VAR=HelloFromGo")

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Command failed:", err)
        return
    }
    fmt.Println("Output with env var:", string(output))
}

```

## Mutex (Queuing) in Go
The sync.Mutex type provides mutual exclusion, allowing only one goroutine to access a critical section at a time. It is commonly used to protect shared resources like queues, maps, and counters.

> **Avoiding Deadlocks**
> Always ensure:
> - Lock() and Unlock() are matched in all code paths.
> - Avoid calling external or blocking functions while holding a lock.
> - Use defer mu.Unlock() immediately after locking to ensure unlock even on panic or return.

### Basic Mutex Lock/Unlock
```go 
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    count := 0

    mu.Lock()
    count++
    mu.Unlock()

    fmt.Println("Final count:", count)
}

```
### Mutex With Goroutines (Safe Queuing)
```go 
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    queue := []string{}
    wg := sync.WaitGroup{}

    enqueue := func(item string) {
        mu.Lock()
        queue = append(queue, item)
        mu.Unlock()
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            enqueue(fmt.Sprintf("task-%d", i))
        }(i)
    }

    wg.Wait()

    mu.Lock()
    fmt.Println("Queue contents:", queue)
    mu.Unlock()
}

```
### Read and Modify Shared Data Safely
```go 
package main

import (
    "fmt"
    "sync"
)

type SafeQueue struct {
    mu    sync.Mutex
    items []string
}

func (q *SafeQueue) Enqueue(item string) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
}

func (q *SafeQueue) Dequeue() (string, bool) {
    q.mu.Lock()
    defer q.mu.Unlock()
    if len(q.items) == 0 {
        return "", false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func main() {
    queue := &SafeQueue{}

    queue.Enqueue("alpha")
    queue.Enqueue("beta")

    if item, ok := queue.Dequeue(); ok {
        fmt.Println("Dequeued:", item)
    }
}

```

## MQTT (Mosquitto) in Go
Install the required package using:
```bash
go get github.com/eclipse/paho.mqtt.golang
```
Import it in your code:
```go
import mqtt "github.com/eclipse/paho.mqtt.golang"
```
### Setup
```go
package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "time"
)

func connectBroker() mqtt.Client {
    opts := mqtt.NewClientOptions().
        AddBroker("tcp://localhost:1883").
        SetClientID("go-client").
        SetCleanSession(true)

    client := mqtt.NewClient(opts)
    token := client.Connect()
    if token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    fmt.Println("Connected to MQTT broker")
    return client
}

```
### Add Message
```go
func publishMessage(client mqtt.Client, topic, payload string) {
    token := client.Publish(topic, 0, false, payload)
    token.Wait()
    fmt.Println("Published:", payload)
}

```
### Listen for Topic
```go
func subscribeToTopic(client mqtt.Client, topic string) {
    handler := func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), string(msg.Payload()))
    }

    token := client.Subscribe(topic, 0, handler)
    token.Wait()
    fmt.Println("Subscribed to:", topic)
}

```
### Get Pending List
You need to manage your own pending list (e.g. a slice or queue) because MQTT brokers do not expose queue contents directly. Here’s an example using a local slice:
```go
var pending []string
var mu sync.Mutex

func onMessage(client mqtt.Client, msg mqtt.Message) {
    mu.Lock()
    pending = append(pending, string(msg.Payload()))
    mu.Unlock()
}

```

```go
func getPendingList() []string {
    mu.Lock()
    defer mu.Unlock()
    return append([]string(nil), pending...) // return a copy
}

```
### Remove Item from List
```go
func removeItemFromPending(value string) {
    mu.Lock()
    defer mu.Unlock()
    for i, v := range pending {
        if v == value {
            pending = append(pending[:i], pending[i+1:]...)
            break
        }
    }
}

```
### Usage Example
```go
func main() {
    client := connectBroker()

    // Subscribe and listen
    subscribeToTopic(client, "readyfleet/updates")

    // Publish a message
    publishMessage(client, "readyfleet/updates", `{"id": 123, "task": "load"}`)

    // Keep program running to receive messages
    select {}
}

```
