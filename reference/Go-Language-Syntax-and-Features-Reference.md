
## Table of Contents

| Feature Category             | Description                                                   |
|------------------------------|---------------------------------------------------------------|
| [Variables & Constants](#variables--constants) | Declaring, initialising, and using variables and constants. |
| [Data Types](#data-types)                   | Built-in types: int, float, string, bool, etc.             |
| [Type Conversion](#type-conversion)         | Converting between compatible types.                      |
| [Operators](#operators)                     | Arithmetic, comparison, logical, and bitwise operators.   |
| [Control Flow](#control-flow)               | `if`, `else`, `switch`, `select`, `goto`.                 |
| [Loops](#loops)                             | `for`, range loops, and infinite loops.                   |
| [Functions](#functions)                     | Function declaration, parameters, return values.          |
| [Multiple Return Values](#multiple-return-values) | Returning and handling multiple values.              |
| [Named Return Values](#named-return-values) | Using named return parameters in functions.               |
| [Defer, Panic, and Recover](#defer-panic-and-recover) | Error handling and cleanup logic.                |
| [Pointers](#pointers)                       | Using pointers, dereferencing, and pointer arithmetic.    |
| [Arrays](#arrays)                           | Fixed-length sequences of values.                         |
| [Slices](#slices)                           | Dynamic views over arrays.                                |
| [Maps](#maps)                               | Key-value stores.                                         |
| [Structs](#structs)                         | Custom composite types.                                   |
| [Methods](#methods)                         | Attaching methods to types.                               |
| [Interfaces](#interfaces)                   | Polymorphism and abstraction.                             |
| [Anonymous Functions & Closures](#anonymous-functions--closures) | Functions without names and closures.     |
| [Variadic Functions](#variadic-functions)   | Functions accepting a variable number of arguments.       |
| [Packages & Imports](#packages--imports)    | Using and organising code across multiple files.          |
| [Modules](#modules)                         | Working with `go.mod` and modules.                        |
| [Initialization](#initialisation)           | `init` functions and variable initialisation order.       |
| [Concurrency](#concurrency)                 | Goroutines and channels.                                  |
| [Select Statement](#select-statement)       | Multiplexing over multiple channel operations.            |
| [Type Assertions](#type-assertions)         | Extracting concrete values from interfaces.               |
| [Type Switch](#type-switch)                 | Switch specifically for interface type checks.            |
| [Tags and Reflection](#tags-and-reflection) | Struct tags and use of the `reflect` package.             |

# Variables & Constants

The `var` and `const` keywords in Go are used to declare variables and constants respectively. Variables may be explicitly typed or type-inferred. Constants are immutable and must be assigned a compile-time constant expression.

## Variable Declaration

| Syntax                                 | Description                                          |
|----------------------------------------|------------------------------------------------------|
| `var x int`                            | Declares `x` with type `int` and default value `0`.  |
| `var y = 10`                           | Declares `y` with inferred type `int`.               |
| `var z, w = 1, "hello"`                | Multiple variable declarations with inference.       |
| `var a, b int = 1, 2`                  | Multiple variable declarations with explicit type.   |
| `x := 5`                               | Short-hand declaration inside a function (inferred). |

> 🔹 Short-hand `:=` is only valid inside function bodies.

## Constant Declaration

| Syntax                                 | Description                                          |
|----------------------------------------|------------------------------------------------------|
| `const Pi = 3.14`                      | Declares a constant with inferred type.              |
| `const Greeting string = "hello"`      | Declares a typed constant.                           |
| `const (X = 1; Y = 2)`                 | Declares multiple constants in a block.              |
| `const (A = iota; B; C)`               | Uses `iota` to create enumerated constants.          |

## Notes

- Variables declared at package level are initialised in declaration order.
- Constants must be assigned values known at compile time.
- `iota` is reset to `0` with each `const` block and increments automatically.

# Data Types

Go is a statically typed language with built-in support for a wide variety of basic and composite types. You can also define your own types using `type`.

## Built-in Basic Types

| Category         | Types                                                  | Description                                           |
|------------------|--------------------------------------------------------|-------------------------------------------------------|
| Boolean          | `bool`                                                 | Represents true or false.                            |
| Integer (signed) | `int`, `int8`, `int16`, `int32`, `int64`               | Signed integers of varying sizes.                    |
| Integer (unsigned) | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` | Unsigned integers and pointer-sized uint.       |
| Floating Point   | `float32`, `float64`                                   | IEEE-754 floating point types.                       |
| Complex Numbers  | `complex64`, `complex128`                              | Complex numbers with float32 or float64 components.  |
| Strings          | `string`                                               | Immutable sequence of bytes (UTF-8 encoded).         |
| Runes            | `rune` (alias for `int32`)                             | Represents a Unicode code point.                     |
| Bytes            | `byte` (alias for `uint8`)                             | Used to represent raw binary data.                   |

## Composite Types

| Type       | Description                                                             |
|------------|-------------------------------------------------------------------------|
| `array`    | Fixed-size list of elements of the same type.                          |
| `slice`    | Dynamically-sized view into an array.                                   |
| `map`      | Unordered collection of key-value pairs.                                |
| `struct`   | Collection of named fields (like records or objects).                   |
| `pointer`  | Holds the memory address of another variable.                           |
| `function` | First-class type that can be assigned and passed.                       |
| `interface`| Defines method sets for polymorphism.                                   |
| `channel`  | Used for communication between goroutines.                              |

## Type Aliases

| Alias    | Underlying Type | Use Case                                      |
|----------|------------------|-----------------------------------------------|
| `byte`   | `uint8`          | Often used for raw binary data.               |
| `rune`   | `int32`          | Used for Unicode code points.                 |

## Type Inference

| Example            | Description                             |
|--------------------|-----------------------------------------|
| `x := 42`          | Compiler infers `int`.                  |
| `y := "hello"`     | Compiler infers `string`.               |
| `z := 3.14`        | Compiler infers `float64`.              |


# Type Conversion

Go does not allow implicit type conversions between different types. You must use explicit conversion syntax even for compatible types.

## Syntax

Where `T` is the target type and `expression` is the value to convert.

## Common Type Conversions

| From        | To           | Syntax Example             | Notes                                            |
|-------------|--------------|----------------------------|--------------------------------------------------|
| `int`       | `float64`    | `float64(i)`               | Converts integer to float.                       |
| `float64`   | `int`        | `int(f)`                   | Truncates decimal portion.                      |
| `int`       | `string`     | `string(i)`                | Converts to Unicode character (e.g., 'A').      |
| `string`    | `[]byte`     | `[]byte(s)`                | Converts to byte slice.                         |
| `[]byte`    | `string`     | `string(b)`                | Converts byte slice to string.                  |
| `rune`      | `string`     | `string(r)`                | Converts single rune to string.                 |
| `CustomType1` | `CustomType2` | `CustomType2(v)`         | If both have the same underlying type.          |

## Invalid Conversions (Compile-Time Errors)

| Example             | Reason                                           |
|----------------------|--------------------------------------------------|
| `int("123")`         | Cannot convert string to int directly.          |
| `[]int("123")`       | Cannot convert string to slice of integers.     |
| `map[int]string` to `map[string]string` | Incompatible key types.         |

## Notes

- Conversion differs from type assertion (used with interfaces).
- Data loss may occur when converting types (e.g., `float64` to `int`, `int64` to `int8`).
- Use the `strconv` package for converting between strings and numeric types.

# Operators

Go provides a rich set of operators for arithmetic, comparison, logical operations, and more. Operators work on basic types and follow strict type compatibility.

## Arithmetic Operators

| Operator | Description          | Example     |
|----------|----------------------|-------------|
| `+`      | Addition              | `x + y`     |
| `-`      | Subtraction           | `x - y`     |
| `*`      | Multiplication        | `x * y`     |
| `/`      | Division              | `x / y`     |
| `%`      | Modulus (remainder)   | `x % y`     |

## Comparison Operators

| Operator | Description           | Example     |
|----------|-----------------------|-------------|
| `==`     | Equal                 | `x == y`    |
| `!=`     | Not equal             | `x != y`    |
| `<`      | Less than             | `x < y`     |
| `<=`     | Less than or equal    | `x <= y`    |
| `>`      | Greater than          | `x > y`     |
| `>=`     | Greater than or equal | `x >= y`    |

## Logical Operators

| Operator | Description         | Example         |
|----------|---------------------|------------------|
| `&&`     | Logical AND         | `x > 0 && y > 0` |
| `||`     | Logical OR          | `x > 0 || y > 0` |
| `!`      | Logical NOT         | `!flag`          |

## Bitwise Operators (Integers Only)

| Operator | Description             | Example     |
|----------|-------------------------|-------------|
| `&`      | AND                     | `x & y`     |
| `|`      | OR                      | `x | y`     |
| `^`      | XOR                     | `x ^ y`     |
| `&^`     | AND NOT (bit clear)     | `x &^ y`    |
| `<<`     | Left shift              | `x << 2`    |
| `>>`     | Right shift             | `x >> 2`    |

## Assignment Operators

| Operator | Description                   | Equivalent To     |
|----------|-------------------------------|-------------------|
| `=`      | Assign                        | `x = y`           |
| `+=`     | Add and assign                | `x = x + y`       |
| `-=`     | Subtract and assign           | `x = x - y`       |
| `*=`     | Multiply and assign           | `x = x * y`       |
| `/=`     | Divide and assign             | `x = x / y`       |
| `%=`     | Modulus and assign            | `x = x % y`       |
| `&=`     | Bitwise AND and assign        | `x = x & y`       |
| `|=`     | Bitwise OR and assign         | `x = x | y`       |
| `^=`     | Bitwise XOR and assign        | `x = x ^ y`       |
| `<<=`    | Left shift and assign         | `x = x << 2`      |
| `>>=`    | Right shift and assign        | `x = x >> 2`      |
| `&^=`    | AND NOT and assign            | `x = x &^ y`      |

## Other Operators

| Operator | Description                       | Example       |
|----------|-----------------------------------|---------------|
| `&`      | Address of variable               | `&x`          |
| `*`      | Dereference pointer               | `*ptr`        |
| `<-`     | Channel send/receive              | `ch <- v`, `<-ch` |

# Control Flow

Go provides a range of control flow structures to conditionally execute code blocks, loop over data, or perform branching logic.

## if / else Statements

| Syntax                          | Description                                    |
|----------------------------------|------------------------------------------------|
| `if condition { ... }`           | Executes block if condition is true.           |
| `if condition { ... } else { ... }` | Adds an alternative path.                  |
| `if init; condition { ... }`     | Allows a short statement before the condition.|

## switch Statement

| Syntax                          | Description                                           |
|----------------------------------|-------------------------------------------------------|
| `switch x { case a: ... }`       | Compares `x` against `case` values.                  |
| `switch { case cond: ... }`      | Condition-based switch (like chained `if`).          |
| `switch x := expr; x { ... }`    | Includes a short statement before the switch.        |
| `case a, b:`                     | Multiple values in a single case.                    |
| `default:`                       | Executed when no other case matches.                 |
| Fallthrough                     | Proceeds to next case unconditionally (optional).    |

## select Statement (Concurrency)

| Syntax         | Description                                               |
|----------------|-----------------------------------------------------------|
| `select { ... }` | Waits on multiple channel operations. Executes one that proceeds. |

## goto Statement

| Syntax           | Description                                     |
|------------------|-------------------------------------------------|
| `goto label`     | Jumps to the labelled statement.                |
| `label:`         | Defines a label within the function.            |

> ⚠️ Use `goto` sparingly; mostly used for early exit in deeply nested logic.

## Notes

- `if`, `switch`, and `select` do **not** require parentheses around conditions.
- Braces `{}` are mandatory.
- `select` is only valid when working with channels and goroutines.

# Loops

Go has only one looping construct: the `for` loop. It is flexible enough to serve as a `while`, `do-while`, or traditional `for` loop.

## Loop Forms

| Form                          | Description                                     |
|-------------------------------|-------------------------------------------------|
| `for init; condition; post {}`| Traditional `for` loop.                        |
| `for condition {}`            | Acts like a `while` loop.                      |
| `for {}`                      | Infinite loop (manually break when needed).    |
| `for index, value := range collection {}` | Iterates over arrays, slices, strings, maps, or channels. |

## Examples

| Syntax                              | Description                                |
|-------------------------------------|--------------------------------------------|
| `for i := 0; i < 10; i++ {}`        | Loops 10 times.                            |
| `for x < 5 {}`                      | Loops while condition is true.             |
| `for { break }`                     | Infinite loop with manual exit.            |
| `for i, v := range slice {}`        | Iterates over slice elements.              |
| `for k, v := range mapVar {}`       | Iterates over map keys and values.         |
| `for i := range arr {}`             | Iterates over indices only.                |

## Loop Control

| Keyword     | Description                                   |
|-------------|-----------------------------------------------|
| `break`     | Exits the nearest enclosing `for`, `switch`, or `select`. |
| `continue`  | Skips the current iteration and proceeds to the next.    |
| `goto label`| Jumps to a labelled statement (use cautiously).         |

## Notes

- `range` returns one or two values: index/key and element.
- `continue` inside `range` affects the current iteration.
- Always use braces `{}` even for single-line loop bodies.

# Functions

Functions in Go are reusable blocks of code that encapsulate logic and can be executed with a set of input parameters. They support multiple return values, named return values, variadic arguments, and closures.

## Function Declarations

| Element                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Function Name            | Identifier used to invoke the function.                          |
| Parameters               | Comma-separated list of typed input variables.                   |
| Return Type              | Optional single or multiple return types.                        |
| Function Body            | Block of statements executed when the function is called.        |

## Multiple Return Values

| Concept                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Return Tuple             | Functions may return more than one value.                        |
| Use in Assignment        | Caller can unpack multiple return values.                        |
| Common Use Case          | Often used to return a result and an error.                      |

## Named Return Values

| Concept                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Named Parameters         | Return variables can be named in the function signature.         |
| Implicit Return          | A bare `return` uses the named variables.                        |
| Improves Readability     | Useful for self-documenting function signatures.                 |

## Parameter Passing

| Concept                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Pass by Value            | All arguments are passed by value (copied into function).        |
| Pass by Reference        | Use pointers to allow function to modify caller’s value.         |

## Variadic Functions

| Feature                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Variadic Parameter       | Allows a function to accept zero or more arguments of a type.    |
| Use Case                 | Useful for operations on a variable number of inputs.            |
| Internally Treated As    | Treated as a slice inside the function.                          |

## Anonymous Functions

| Feature                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| No Name                  | Functions without a declared identifier.                         |
| Assigned to Variable     | Can be stored and invoked via a variable.                        |
| Used Inline              | Commonly used for short-lived behaviour, callbacks, or goroutines.|

## Closures

| Concept                  | Description                                                      |
|--------------------------|------------------------------------------------------------------|
| Captures Scope           | Anonymous functions that access variables from outer scope.      |
| Persistent State         | Useful for retaining state between function calls.               |

## Notes

- Function declarations are not nested in Go.
- Functions are first-class citizens and can be assigned to variables, passed as arguments, and returned from other functions.
- Go encourages small, pure functions with single responsibility.

# Defer, Panic, and Recover

Go provides built-in mechanisms for managing control flow in exceptional and finalisation scenarios: `defer`, `panic`, and `recover`.

## Defer

| Concept                  | Description                                                             |
|--------------------------|-------------------------------------------------------------------------|
| Postponed Execution      | Defers the execution of a function until the surrounding function returns. |
| LIFO Order               | Deferred calls are executed in last-in, first-out (stack) order.        |
| Use Cases                | Commonly used for resource cleanup, file closing, unlocking mutexes, etc. |
| Evaluated Immediately    | Arguments to deferred functions are evaluated when the defer is declared. |

## Panic

| Concept                  | Description                                                             |
|--------------------------|-------------------------------------------------------------------------|
| Unexpected Failure       | Aborts the normal control flow of the program.                          |
| Call Stack Unwinding     | Deferred functions are still executed during a panic.                   |
| Fatal Conditions         | Typically used when encountering unrecoverable errors.                 |
| Avoid Overuse            | Should be used sparingly; not for ordinary error handling.             |

## Recover

| Concept                  | Description                                                             |
|--------------------------|-------------------------------------------------------------------------|
| Trap a Panic             | Regains control of a panicking goroutine.                               |
| Only in Deferred Call    | Must be called within a deferred function to stop the panic.            |
| Return to Normal Flow    | If recover is successful, execution continues normally after the deferred block. |

## Common Patterns

| Pattern                  | Description                                                             |
|--------------------------|-------------------------------------------------------------------------|
| Defer for Cleanup        | Always run specific logic regardless of how a function exits.           |
| Panic and Recover        | Used together to safely manage unexpected runtime errors.               |
| Graceful Degradation     | Prevents application crashes by catching panics and returning safe values. |

## Notes

- A panic stops the ordinary execution of a function.
- Recover only works when directly called within a deferred function.
- Use `defer` carefully in performance-critical code, as it adds overhead.

# Pointers

Pointers in Go store the memory address of a value. They allow functions and methods to modify the original data without copying it.

## Core Concepts

| Concept                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Address-of Operator      | The `&` operator is used to get the memory address of a variable.         |
| Dereferencing            | The `*` operator accesses the value at the address a pointer refers to.   |
| Zero Value               | The zero value of a pointer is `nil`.                                     |
| Typed Pointers           | A pointer has a specific type and must match the type it points to.       |
| Pointer Comparison       | Pointers can be compared to other pointers or to `nil`.                   |

## Use Cases

| Scenario                  | Description                                                             |
|---------------------------|-------------------------------------------------------------------------|
| Function Arguments        | Pass a pointer to allow the function to modify the original value.      |
| Struct Modification       | Use pointers to structs to avoid copying large values.                  |
| Efficient Memory Use      | Prevent large copies of data structures by passing references.          |
| Linked Data Structures    | Essential for implementing linked lists, trees, and graphs.             |

## Rules and Behaviour

| Rule                        | Description                                                            |
|-----------------------------|------------------------------------------------------------------------|
| No Pointer Arithmetic       | Go does not support arithmetic on pointers (unlike C/C++).             |
| No Implicit Dereference     | Must use `*` explicitly to read or write to pointer values.            |
| Automatic Dereferencing     | Methods with value receivers can be called on pointer receivers and vice versa in many cases. |

## Common Pitfalls

| Issue                       | Description                                                            |
|-----------------------------|------------------------------------------------------------------------|
| Nil Pointer Dereference     | Accessing a `nil` pointer causes a runtime panic.                      |
| Escaping to Heap            | Variables may be allocated on the heap if captured by pointers.        |
| Uninitialised Pointers      | Must assign a valid address before use.                                |

## Notes

- Pointers are a fundamental concept in Go, used extensively in function parameters and data structures.
- While Go does not have pointer arithmetic, it retains the power of safe referencing and mutation.
- The language's design encourages safe and idiomatic use of pointers without exposing low-level memory operations.


# Arrays

Arrays in Go are fixed-size sequences of elements of the same type. They are value types, meaning assigning one array to another copies all elements.

## Core Characteristics

| Feature                | Description                                                              |
|------------------------|--------------------------------------------------------------------------|
| Fixed Size             | Array size is part of its type and cannot change after declaration.      |
| Homogeneous Elements   | All elements in an array must be of the same type.                       |
| Zero Values            | Elements are automatically initialised to the zero value of the type.    |
| Value Semantics        | Arrays are copied on assignment or function passing.                     |

## Declaration

| Pattern                        | Description                                       |
|--------------------------------|---------------------------------------------------|
| With Length                   | Declares an array with a specified length.        |
| With Initial Values           | Declares and initialises using literal syntax.    |
| Compiler Inferred Length      | Uses `...` to let the compiler determine the size.|

## Indexing & Access

| Concept                       | Description                                       |
|-------------------------------|---------------------------------------------------|
| Index-based Access            | Elements are accessed using zero-based indices.  |
| Out-of-bounds Access          | Accessing an invalid index causes a runtime panic.|

## Iteration

| Method                        | Description                                       |
|-------------------------------|---------------------------------------------------|
| Traditional For Loop          | Iterates using explicit indices.                 |
| `range` Loop                  | Iterates over index and value pairs.             |

## Comparison and Equality

| Rule                          | Description                                       |
|-------------------------------|---------------------------------------------------|
| Comparable Elements           | Arrays can be compared using `==` if their elements are comparable. |
| Full Match Required           | Two arrays are equal only if all corresponding elements are equal. |

## Use Cases

| Scenario                      | Description                                       |
|-------------------------------|---------------------------------------------------|
| Fixed-size Data               | Use when the size is known and fixed at compile-time. |
| Performance-sensitive Code    | Avoid dynamic resizing overhead of slices.       |

## Notes

- Arrays are rarely used directly in Go; slices are generally preferred due to their flexibility.
- Copying arrays is an intentional operation, not a side effect.
- Understanding arrays is fundamental to understanding slices, which are built on top of them.

# Slices

Slices are flexible, dynamically-sized views into arrays. They are one of the most commonly used data structures in Go and provide powerful tools for working with sequential collections.

## Core Characteristics

| Feature                | Description                                                              |
|------------------------|--------------------------------------------------------------------------|
| Backed by Arrays       | A slice references a segment of an underlying array.                     |
| Variable Length        | The length of a slice can grow or shrink dynamically.                    |
| Value Semantics        | Slices are reference types; copying a slice creates a new slice header pointing to the same underlying array. |
| Zero Value             | The zero value of a slice is `nil`, with length and capacity of 0.       |

## Declaration

| Method                         | Description                                                      |
|--------------------------------|------------------------------------------------------------------|
| Slice Literal                  | Declares and initialises a slice in one step.                    |
| From Array                    | Uses slicing syntax to create a slice from an array.             |
| `make([]T, len)`              | Allocates a slice of a specific length (zeroed elements).        |
| `make([]T, len, cap)`         | Allocates a slice with specific length and capacity.             |

## Properties

| Property             | Description                                           |
|----------------------|-------------------------------------------------------|
| Length (`len`)       | Number of elements in the slice.                     |
| Capacity (`cap`)     | Maximum number of elements before reallocation.      |

## Operations

| Operation              | Description                                           |
|------------------------|-------------------------------------------------------|
| Indexing               | Access or modify elements using `[i]`.               |
| Slicing                | Extracts sub-slices using `[low:high]` syntax.       |
| Appending              | Adds elements using the built-in `append` function.  |
| Copying                | Copies contents using the built-in `copy` function.  |

## Behaviour

| Concept                | Description                                           |
|------------------------|-------------------------------------------------------|
| Shared Backing Array   | Multiple slices may share the same underlying array. |
| Append Reallocation    | Appending beyond capacity allocates a new array.     |
| Nil Slice              | Represents zero-length, zero-capacity slice.         |

## Use Cases

| Scenario                | Description                                           |
|------------------------|-------------------------------------------------------|
| Dynamic Collections     | When number of items is not known at compile-time.   |
| Sub-list Processing     | When working on segments of data.                    |

## Notes

- Always consider whether appending may overwrite shared data.
- Slices provide the simplicity of arrays with added flexibility.
- Use slices instead of arrays in most practical applications.

# Maps

Maps in Go are unordered collections of key-value pairs. They provide fast lookups, inserts, and deletions, and are implemented as hash tables.

## Core Characteristics

| Feature                | Description                                                              |
|------------------------|--------------------------------------------------------------------------|
| Dynamic Size           | Maps grow automatically as new keys are added.                          |
| Unordered              | Iteration order is not guaranteed and may vary between executions.       |
| Reference Type         | Maps are reference types; assigning a map to a new variable copies the reference, not the data. |
| Key Uniqueness         | Each key must be unique; assigning to an existing key overwrites it.     |

## Declaration

| Method                        | Description                                       |
|-------------------------------|---------------------------------------------------|
| Map Literal                  | Creates a map and initialises it with values.     |
| `make(map[KeyType]ValueType)`| Allocates and returns an empty map.              |
| `var m map[KeyType]ValueType`| Declares a nil map (must be initialised before use). |

## Operations

| Operation           | Description                                            |
|---------------------|--------------------------------------------------------|
| Insertion           | Assigning a value to a key adds or updates an entry.   |
| Retrieval           | Accessing a key returns the value and a boolean for presence. |
| Deletion            | The `delete` built-in removes a key-value pair.        |
| Presence Check      | Second return value from retrieval indicates existence.|

## Key Requirements

| Requirement         | Description                                            |
|---------------------|--------------------------------------------------------|
| Comparable          | Keys must be of a type that supports equality (`==`).  |
| Examples            | Valid: strings, integers, booleans, pointers. Invalid: slices, maps, functions. |

## Iteration

| Concept             | Description                                            |
|---------------------|--------------------------------------------------------|
| `range` Loop        | Iterates over key-value pairs in unspecified order.    |
| Key/Value Access    | Loop can access keys only, or both keys and values.    |

## Behaviour

| Concept             | Description                                            |
|---------------------|--------------------------------------------------------|
| Accessing Missing Key | Returns zero value of the value type.              |
| Deleting Missing Key | Safe and has no effect.                             |
| Modifying During Iteration | Allowed, but discouraged and unpredictable.   |

## Use Cases

| Scenario            | Description                                            |
|---------------------|--------------------------------------------------------|
| Fast Lookup         | Ideal for indexing and deduplication tasks.            |
| Configuration       | Suitable for key-based settings or metadata.           |

## Notes

- Always check the second return value when key presence is uncertain.
- A `nil` map behaves like an empty map for reads, but panics on write.
- Maps do not guarantee stability in key order; avoid relying on it.

# Structs

Structs in Go are composite types that group together zero or more fields with varying types under a single name. They are the foundation for creating custom data types.

## Core Characteristics

| Feature                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Named Fields             | Each field has a name and a type.                                         |
| Value Type               | Structs are value types and are copied when assigned or passed.           |
| Zero Values              | Fields are initialised to their zero values when a struct is created.     |
| Field Access             | Fields are accessed using dot notation.                                   |

## Declaration

| Pattern                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Named Struct Type        | Struct declared with a type name and used to create instances.            |
| Anonymous Struct         | Declared without a named type, often for temporary use.                   |
| Field Tags               | Optional metadata for fields, typically used for encoding and reflection. |

## Initialisation

| Method                   | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Named Fields             | Fields can be assigned by name in any order.                              |
| Positional Initialisation| Fields must be provided in declared order, without names.                 |
| Zero Initialisation      | Declaring without initial values sets all fields to zero values.          |

## Embedded Fields (Struct Composition)

| Concept                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Promoted Fields          | An embedded struct’s fields are accessible via the outer struct.          |
| Composition              | Allows struct reuse without inheritance.                                  |

## Methods with Struct Receivers

| Receiver Types           | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Value Receiver           | Method receives a copy of the struct.                                     |
| Pointer Receiver         | Method can modify the original struct; avoids copying large structs.      |
| Method Set               | Determines which methods are accessible on values and pointers.           |

## Comparison

| Rule                     | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Comparable Fields        | Structs can be compared using `==` if all fields are comparable.          |
| Deep Comparison          | Use `reflect.DeepEqual` for nested structs with non-comparable fields.    |

## Use Cases

| Scenario                 | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Data Modelling           | Structs represent entities, configurations, requests, etc.                |
| Method Grouping          | Associate behaviour with data via methods.                               |

## Notes

- Struct field names must be exported (start with an uppercase letter) to be accessible from other packages.
- Use composition and interfaces together to model complex systems idiomatically in Go.
- Structs provide a lightweight and efficient alternative to classes in object-oriented languages.

# Methods

Methods in Go are functions with a special receiver argument. They are used to associate behaviour with user-defined types such as structs.

## Core Characteristics

| Feature                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Receiver Argument        | Defined between the `func` keyword and method name.                       |
| Type Association         | Methods are associated with named types (usually structs).                |
| Multiple Methods         | A type can have multiple methods.                                         |
| No Classes               | Go does not have classes; methods provide similar encapsulation.          |

## Receiver Types

| Type                     | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Value Receiver           | The method operates on a copy of the receiver.                           |
| Pointer Receiver         | The method operates on the original value and can modify it.             |

## Method Sets

| Rule                     | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Value Type               | Has methods with value receivers.                                         |
| Pointer Type             | Has methods with both value and pointer receivers.                        |
| Interface Satisfaction   | A type implements an interface if it defines all required methods.        |

## Declaration Rules

| Rule                     | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Same Package             | Methods must be defined in the same package as the receiver type.         |
| Named Type Only          | Receiver must be a named type (not a pointer, slice, or map directly).    |

## Use Cases

| Scenario                 | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Behaviour Definition     | Encapsulate functionality related to a type.                              |
| API Design               | Create fluent and structured APIs.                                        |
| Interface Implementation | Satisfy interfaces with methods to enable polymorphism.                   |

## Notes

- Use pointer receivers when the method needs to modify the receiver or avoid copying large values.
- Both pointer and value receivers are valid; consistency across methods on the same type is encouraged.
- Methods can be attached to any user-defined type, not just structs.

# Interfaces

Interfaces in Go define a set of method signatures. A type implements an interface implicitly by defining all required methods. This enables polymorphism without inheritance.

## Core Characteristics

| Feature                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Method Set Definition    | An interface specifies behaviour through method signatures.               |
| Implicit Implementation  | A type satisfies an interface by implementing its methods.                |
| Polymorphism             | Enables writing functions and structs that work with any conforming type. |
| Zero Value               | The zero value of an interface is `nil`.                                  |

## Declaration

| Pattern                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Interface Type           | Defined using the `interface` keyword with method signatures.             |
| Empty Interface          | `interface{}` matches any value (like `any`).                             |
| Embedding                | Interfaces can embed other interfaces to compose new ones.                |

## Method Sets and Satisfaction

| Rule                     | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Exact Match              | The concrete type must define all methods in the interface.               |
| Pointer vs Value Receiver | Method set depends on the receiver; pointer types satisfy more interfaces.|
| Structural Typing        | Satisfaction is based on method signatures, not declared relationships.   |

## Type Assertion

| Feature                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Single Return            | Extracts the underlying value; panics if the type does not match.         |
| Two-value Return         | Returns the value and a boolean indicating success.                       |

## Type Switch

| Feature                  | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Multi-type Handling      | Switches on the dynamic type held by an interface.                        |
| Safe                     | Allows different code paths based on concrete types.                      |

## Use Cases

| Scenario                 | Description                                                               |
|--------------------------|---------------------------------------------------------------------------|
| Abstraction              | Write generic logic against behaviour, not concrete types.                |
| Dependency Injection     | Swap implementations easily in testing and production.                    |
| Plugin Architectures     | Dynamically load and use external behaviours.                             |

## Notes

- Interfaces promote loose coupling and code reusability.
- The `io` and `fmt` packages define widely used interfaces like `Reader`, `Writer`, and `Stringer`.
- Interfaces cannot contain fields—only method signatures.

# Anonymous Functions & Closures

Go supports anonymous functions—functions defined without a name. When these functions capture and reference variables from their surrounding scope, they become closures.

## Anonymous Functions

| Feature                    | Description                                                              |
|----------------------------|--------------------------------------------------------------------------|
| No Name                    | Defined without an identifier.                                           |
| Assigned to Variables      | Can be stored in variables for reuse.                                   |
| Passed as Arguments        | Commonly used for callbacks and functional parameters.                   |
| Declared Inline            | Often defined within expressions or function calls.                      |

## Closures

| Feature                    | Description                                                              |
|----------------------------|--------------------------------------------------------------------------|
| Captures Environment       | Retains access to variables declared in outer scopes.                   |
| Persistent State           | Maintains variable state across multiple calls.                          |
| Scope Sharing              | Multiple closures can share and mutate the same captured variables.      |

## Behaviour and Lifetime

| Rule                       | Description                                                              |
|----------------------------|--------------------------------------------------------------------------|
| Variables by Reference     | Captured variables are shared by reference, not copied.                 |
| Execution Context          | Closure retains the context in which it was declared.                   |
| Use in Goroutines          | Care must be taken with closures that run asynchronously.               |

## Use Cases

| Scenario                   | Description                                                              |
|----------------------------|--------------------------------------------------------------------------|
| Functional Programming     | Create reusable, concise behaviours.                                     |
| Callback Functions         | Used in event-driven or concurrent code.                                 |
| Encapsulation              | Hide helper logic inside another function.                               |
| Delayed Execution          | Used with `defer` or returned as part of a function.                     |

## Notes

- Closures can simplify code by eliminating the need to define named helper functions.
- Be cautious when closures capture loop variables, especially in goroutines.
- Proper use of closures promotes cleaner and more maintainable code.

# Variadic Functions

Variadic functions in Go accept a variable number of arguments of the same type. They are useful when the exact number of inputs is unknown or flexible.

## Core Characteristics

| Feature                      | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Variable Number of Arguments | Accepts zero or more values of a specified type.                          |
| Single Variadic Parameter    | Only the final parameter in a function can be variadic.                   |
| Internally a Slice           | The variadic argument is represented as a slice within the function.      |
| Flexible Calls               | Can be called with individual values or a slice (using the `...` operator).|

## Declaration Rules

| Rule                         | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Syntax                       | Uses ellipsis `...` before the type (e.g., `...int`).                     |
| Must Be Last Parameter       | No parameters may follow a variadic one.                                  |
| Optional Arguments           | A variadic function can be called with zero values.                       |

## Behaviour

| Behaviour                    | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Slice Conversion             | A variadic argument behaves like a slice inside the function.             |
| Argument Expansion           | A slice can be expanded into variadic parameters using `...`.             |
| Argument Length              | The number of variadic values can be checked using `len()`.               |

## Use Cases

| Scenario                     | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Logging                      | Accept arbitrary message parts or fields.                                 |
| Formatting Functions         | Build dynamic output with custom logic.                                   |
| Mathematical Aggregation     | Accept any number of numeric values to compute totals or averages.        |
| API Wrappers                 | Simplify function calls with flexible inputs.                             |

## Notes

- Avoid using multiple variadic functions together in a single call chain to reduce confusion.
- Though convenient, variadic functions should not be overused in performance-critical code paths.
- Always document the intended use and behaviour of variadic parameters for clarity.

# Packages & Imports

Packages are the basic unit of code organisation and reuse in Go. Every Go source file belongs to a package, and packages can be imported to use code defined elsewhere.

## Core Characteristics

| Feature                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Organisational Unit        | Groups related Go files under a common namespace.                        |
| One Package per Directory  | Each directory defines exactly one package.                              |
| Entry Point                | The `main` package defines the executable entry point with `func main()`.|
| Visibility                 | Identifiers starting with an uppercase letter are exported (public).     |

## Importing Packages

| Pattern                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Standard Import            | Uses the `import` keyword to include external or standard packages.       |
| Multiple Imports           | Can be grouped in a parenthesised list.                                   |
| Aliased Import             | Allows renaming of the imported package identifier.                       |
| Blank Identifier (`_`)     | Imports a package solely for its side-effects (initialisation only).      |
| Dot Import (`.`)           | Imports all exported identifiers into the local namespace (not recommended).|

## Import Paths

| Type                       | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Standard Library           | Uses the built-in package name (e.g., `fmt`, `time`).                     |
| Module-based Import        | Refers to packages in your module using relative paths or module paths.   |
| Third-party Packages       | Fetched via their repository paths (e.g., `github.com/user/repo`).        |

## Package Initialisation

| Concept                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| `init()` Function          | Automatically invoked before `main()` or test execution.                 |
| One per File               | Each Go file may include one `init()` function.                          |
| Multiple Init Functions    | All `init()` functions run in file and import dependency order.           |

## Use Cases

| Scenario                   | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Code Reuse                 | Share common functionality across multiple projects.                      |
| Dependency Management      | Organise logic and abstract third-party integrations.                     |
| Encapsulation              | Hide internal implementation details behind exported functions/types.     |

## Notes

- Avoid circular imports; they are not permitted in Go.
- Organise packages by responsibility or functional domain.
- Prefer small, well-named packages with a clear purpose and minimal dependencies.

# Modules

Modules are Go's unit of dependency management. They define the root of a collection of related Go packages and enable versioning, reproducible builds, and dependency resolution.

## Core Characteristics

| Feature                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Root-Level Unit            | A module is the top-level project or library, often a VCS repository.     |
| Declared in `go.mod`       | Modules are configured with a `go.mod` file in the root directory.        |
| Contains Multiple Packages | A single module can contain many packages organised in subdirectories.    |

## `go.mod` File

| Element                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| `module` Directive         | Declares the module's path or import identifier.                          |
| `go` Directive             | Specifies the Go language version used.                                   |
| `require`                  | Lists dependencies and their versions.                                    |
| `replace`                  | Replaces one module path with another, e.g., local override.              |
| `exclude`                  | Prevents a specific module version from being used.                       |

## Initialisation & Commands

| Command                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| `go mod init <module>`     | Creates a new `go.mod` file.                                              |
| `go mod tidy`              | Adds missing and removes unused modules.                                  |
| `go get <package>`         | Adds a new dependency to `go.mod`.                                       |
| `go mod vendor`            | Copies dependencies to a local `vendor` folder.                           |
| `go list -m all`           | Lists all modules in the build list.                                      |

## Versioning

| Concept                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Semantic Versioning        | Modules use `vMAJOR.MINOR.PATCH` format (e.g., `v1.2.3`).                 |
| Compatibility Rules        | Major version upgrades (e.g., v1 → v2) require path changes.              |
| Proxy Support              | Go modules use a proxy by default to cache and serve modules.             |

## Dependency Management

| Concept                    | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Transitive Imports         | Automatically includes required sub-dependencies.                        |
| Minimal Version Selection  | Chooses the lowest version that satisfies all requirements.               |
| Replace and Override       | Redirects dependencies during development or testing.                     |

## Use Cases

| Scenario                   | Description                                                               |
|----------------------------|---------------------------------------------------------------------------|
| Project Isolation          | Define clear boundaries and dependencies per project.                     |
| Reproducible Builds        | Ensure consistent dependency resolution across environments.              |
| Open Source Distribution   | Publish reusable libraries with version control.                          |

## Notes

- Use `go.work` files to manage multiple modules during multi-repo or mono-repo development.
- `go.sum` records cryptographic checksums to verify dependency integrity.
- Avoid committing unnecessary dependencies; use `go mod tidy` to clean up.

# Initialisation

Initialisation in Go refers to the order and behaviour of how variables, constants, packages, and `init()` functions are prepared before program execution begins.

## Core Concepts

| Concept                     | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| Zero Value Initialisation   | All variables and struct fields are automatically initialised to their zero value. |
| `init()` Function           | A special function that runs before `main()` or tests.                    |
| Package-Level Initialisers  | Variables and constants can be assigned computed values at the package level. |

## `init()` Function

| Feature                     | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| Automatic Execution         | Runs before `main()` and before any other function in the same package.   |
| No Arguments or Returns     | Cannot accept parameters or return values.                                |
| One Per File Allowed        | Each Go source file may contain one or more `init()` functions.           |
| Multiple Initialisers       | Run in the order they appear within a file, across all files in dependency order. |

## Initialisation Order

| Stage                       | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| Dependency Order            | Packages are initialised in the order of their import dependencies.       |
| File Order                  | Files within a package are initialised in lexical filename order.         |
| Variable Order              | Variables are initialised top to bottom within each file.                 |
| `init()` Execution          | Runs after all package-level variables are initialised.                   |

## Variable Initialisation

| Concept                     | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| Constants                   | Compile-time values assigned using expressions.                           |
| Variables                   | Can be initialised using values or function calls.                        |
| Dependency Handling         | Go automatically detects and initialises dependent variables first.       |

## Use Cases

| Scenario                    | Description                                                               |
|-----------------------------|---------------------------------------------------------------------------|
| One-Time Setup              | Prepare shared state, register types, or validate configuration.           |
| Resource Allocation         | Open connections, create caches, or seed random generators.               |
| Side Effects                | Initialise frameworks or external dependencies before usage.              |

## Notes

- Avoid heavy logic in `init()`; keep it focused and predictable.
- `init()` is commonly used in conjunction with package registration patterns.
- Use care when initialising global state to avoid race conditions or tight coupling.

# Concurrency

Go has built-in support for concurrency, allowing multiple functions or processes to run independently using goroutines and channels. Concurrency in Go is based on the CSP (Communicating Sequential Processes) model.

## Core Concepts

| Concept                    | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Goroutines                 | Lightweight threads managed by the Go runtime, created with the `go` keyword. |
| Channels                   | Typed conduits for communication between goroutines.                        |
| Blocking Behaviour         | Channel operations block until both sender and receiver are ready.         |
| Synchronisation            | Can coordinate execution and state sharing without explicit locks.         |

## Goroutines

| Feature                    | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Lightweight                | Thousands of goroutines can be created with minimal overhead.               |
| Concurrent Execution       | Runs functions asynchronously.                                             |
| Independent Stack          | Each has a dynamic stack managed by the runtime.                           |
| Run with `go` Keyword      | Prefixing a function call with `go` runs it in a new goroutine.            |

## Channels

| Feature                    | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Bidirectional              | By default, channels can send and receive values.                          |
| Directional Channels       | Functions can declare send-only or receive-only channels.                 |
| Buffered Channels          | Can hold a fixed number of values before blocking.                        |
| Unbuffered Channels        | Block until both sender and receiver are ready.                           |
| Closing Channels           | Use `close()` to indicate no more values will be sent.                    |

## Channel Operations

| Operation                  | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Send                       | Sends a value into a channel.                                               |
| Receive                    | Retrieves a value from a channel.                                           |
| Close                      | Terminates the sending side of the channel.                                |
| Range                      | Iterates over received values until the channel is closed.                 |

## Synchronisation Tools (via `sync` Package)

| Tool                       | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Mutex                      | Ensures mutual exclusion for critical sections.                            |
| WaitGroup                  | Waits for a collection of goroutines to finish.                            |
| Once                       | Guarantees a function is only executed once.                               |
| Cond                       | Coordinates goroutines using condition variables.                          |

## Use Cases

| Scenario                   | Description                                                                 |
|----------------------------|-----------------------------------------------------------------------------|
| Parallel Processing        | Perform tasks simultaneously to reduce runtime.                            |
| Worker Pools               | Manage concurrent jobs with fixed worker counts.                           |
| Pipelines                  | Process data through a series of concurrent stages.                        |
| Asynchronous I/O           | Handle tasks like HTTP requests or file operations concurrently.           |

## Notes

- Avoid sharing memory by communicating; instead, communicate by sharing channels.
- Goroutines that are not synchronised or terminated properly may lead to leaks.
- Use `select` to wait on multiple channel operations concurrently.

# Select Statement

The `select` statement in Go is used to wait on multiple channel operations. It allows a goroutine to block until one of several possible communication operations is ready.

## Core Characteristics

| Feature                      | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Multiple Cases               | Each `case` specifies a send or receive operation on a channel.           |
| First Ready Runs             | Only one case is executed—whichever is ready first.                       |
| Random Selection             | If multiple cases are ready, one is selected at random.                   |
| Blocking Behaviour           | If no case is ready, `select` blocks until one becomes available.         |
| Non-blocking Option          | A `default` case allows non-blocking behaviour if no channels are ready. |

## Syntax Elements

| Element                      | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| `case <-ch:`                 | Receives from channel `ch`.                                               |
| `case ch <- val:`            | Sends `val` into channel `ch`.                                            |
| `default:`                   | Executes if no other case is ready.                                       |

## Behaviour Rules

| Rule                         | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| One Case Executed            | Only a single matching case is run per `select`.                          |
| No Ready Case (No Default)   | If no channels are ready and no `default` exists, it blocks.              |
| Channel Closed               | A receive from a closed channel returns immediately with zero value.      |
| Send on Closed Channel       | Causes a runtime panic.                                                   |

## Use Cases

| Scenario                     | Description                                                               |
|------------------------------|---------------------------------------------------------------------------|
| Timeout Handling             | Use with `time.After` to add timeout logic to channel operations.         |
| Fan-in / Fan-out Patterns    | Combine multiple input channels or distribute work across outputs.        |
| Graceful Shutdown            | Listen for control signals to terminate goroutines cleanly.               |
| Multiplexing Streams         | Process input from multiple sources simultaneously.                       |

## Notes

- `select` is only applicable to channel operations.
- It is the cornerstone of Go's CSP-based concurrency model.
- Avoid busy waiting with `default` unless intentional for performance tuning.

# Type Assertions

Type assertions in Go provide access to the concrete value stored in an interface. They allow safe or unchecked extraction of the underlying type from an interface value.

## Core Characteristics

| Feature                        | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| Extract Concrete Type          | Retrieves the original type stored in an interface.                     |
| Must Match Runtime Type        | The asserted type must match the dynamic type held in the interface.    |
| Panics on Failure              | An invalid assertion without safety checking causes a runtime panic.    |
| Two-Value Form                 | Provides a boolean to safely test whether the assertion succeeded.       |

## Syntax

| Form                            | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| `value := iface.(T)`             | Asserts `iface` holds type `T`. Panics if incorrect.                   |
| `value, ok := iface.(T)`         | Returns `value` and `ok == true` if successful; `ok == false` otherwise.|

## Behaviour Rules

| Rule                             | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| Static Check Only for Interface | Type assertions can only be used on interfaces.                         |
| Cannot Assert from Non-Interface| Using a type assertion on a non-interface causes a compile error.       |
| Interface to Interface          | Permitted if the asserted interface is a subset of the original.        |

## Use Cases

| Scenario                         | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| Type Inspection                  | Determine the specific type stored in a generic interface.              |
| API Wrapping                     | Cast back to concrete types after abstraction.                          |
| Conditional Logic                | Vary behaviour based on the runtime type of interface values.           |

## Notes

- Type assertions are not the same as type conversions.
- Avoid excessive use of assertions; they may indicate missing abstractions or improper interface design.
- Prefer the two-value form for robust, panic-safe code.

# Type Switch

A type switch in Go is a construct that allows execution of different code paths based on the dynamic type of an interface value. It is an extension of the type assertion mechanism.

## Core Characteristics

| Feature                        | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| Switch on Type                | Evaluates the dynamic type held by an interface.                        |
| Compact Syntax                | Provides concise branching based on multiple possible types.            |
| No Panic Risk                 | Safe by design; no runtime panic for unmatched types.                   |
| Fallthrough Not Allowed       | Each case is exclusive; Go does not allow fallthrough in type switches. |

## Syntax Elements

| Element                         | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| `switch x := v.(type)`          | Begins a type switch; `x` holds the concrete value.                    |
| `case T1:`                      | Matches if `v` holds type `T1`.                                        |
| `case T2, T3:`                  | Matches multiple types in a single case.                               |
| `default:`                      | Runs if no other case matches.                                         |

## Behaviour Rules

| Rule                             | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| Evaluates Once                   | Evaluates the type of the interface value a single time.                |
| First Match Wins                 | Executes only the first matching case.                                  |
| Concrete Value Access            | The variable declared in the switch gets the asserted value.            |
| Required Interface               | Only values of interface type can be used in a type switch.            |

## Use Cases

| Scenario                         | Description                                                             |
|----------------------------------|-------------------------------------------------------------------------|
| Runtime Type Handling            | Handle multiple concrete types from a single interface.                 |
| Custom Marshalling               | Choose logic based on data type during serialisation or deserialisation.|
| Interface Implementations        | Implement behaviours for varying implementations of an interface.       |

## Notes

- Type switches improve readability and reduce repetitive type assertion logic.
- The declared variable (`x := v.(type)`) is available in all matching case blocks.
- Avoid abusing type switches when polymorphism via interface methods is sufficient.

# Tags and Reflection

Tags and reflection in Go allow runtime inspection and manipulation of types, values, and metadata. Struct tags provide a lightweight way to attach metadata to fields, commonly used by packages such as `encoding/json`.

## Struct Tags

| Feature                        | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| Inline Metadata                | Tags are string literals associated with struct fields.                 |
| Convention-Based Parsing       | Parsed and used by libraries (e.g. `json`, `xml`, `validate`).          |
| Single Source of Truth         | Enables declarative configuration within the struct definition.         |

## Syntax Rules

| Rule                           | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| String Literal Format          | Enclosed in backticks using key:"value" syntax.                         |
| Multiple Tags per Field        | Space-separated within the same backtick string.                        |
| Not Interpreted by Compiler    | Tags are metadata only—ignored unless accessed via reflection.          |

## Use Cases

| Scenario                       | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| Data Serialisation             | Tags guide how fields are encoded and decoded.                          |
| Validation                     | Specify constraints for form or API validation.                         |
| ORM/DB Mapping                 | Define table and column metadata for persistence layers.                |
| Documentation and Introspection| Used by code generation tools and UI builders.                         |

## Reflection (via `reflect` Package)

| Feature                        | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| Runtime Type Inspection        | Inspect values, types, and struct field tags at runtime.                |
| Kind-Based Introspection       | Identify general type category (e.g. `Slice`, `Struct`, `Ptr`).         |
| Value Modification             | Use reflection to read or modify values (must be settable).             |
| Type Metadata                  | Retrieve field names, types, and tags from struct definitions.          |

## Key Functions

| Function / Method              | Description                                                             |
|--------------------------------|-------------------------------------------------------------------------|
| `reflect.TypeOf()`            | Returns the `reflect.Type` of a value.                                  |
| `reflect.ValueOf()`           | Returns the `reflect.Value` of a value.                                 |
| `Field.Tag.Get("key")`        | Accesses a specific tag value by key.                                   |
| `CanSet`, `Set`               | Used to determine and assign values to fields via reflection.           |

## Notes

- Struct tags must be kept in sync with consuming logic; incorrect tags can cause silent failures.
- Reflection enables powerful dynamic behaviour but comes with performance overhead.
- Avoid overusing reflection in performance-critical code paths; prefer static design where possible.
