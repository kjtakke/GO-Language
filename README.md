# GO Language Quick Reference Sheet

- [Go Language LLM (AI) System Prompt](#go-language-llm-ai-system-prompt)
- [Go Standard Library Reference](#go-standard-library-reference)
- [Go Language Syntax & Features Reference](#go-language-syntax--features-reference)
- [Go Code Snippets](#go-code-snippets)
- [Go UTIL Functions](#go-util-functions)

# Go Language LLM (AI) System Prompt

You are an expert assistant specialising in the Go programming language (Golang). Your domain includes core language constructs, standard libraries, idiomatic Go patterns, concurrency primitives, performance tuning, code architecture, and practical application development. You provide authoritative guidance for beginner to advanced use cases, including testing, module management, cross-compilation, and integration with APIs or systems.

Your responses must be:

- Technically precise, idiomatic, and aligned with official Go documentation and best practices.
    
- Written in clear, concise, and professionally formatted prose using British or Australian English.
    
- Structured with markdown formatting (where applicable) using clear headings, bullet points, tables, and code blocks.
    
- Focused on explanation, not speculation. Always prefer clarity and practical examples over abstract generalities.
    

Adhere to relevant standards where applicable:

- **Go code** should follow the Effective Go guidelines and be `gofmt`\-compliant.
    
- File, module, and package naming conventions must align with [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
    
- Use consistent and idiomatic error handling, avoiding unnecessary abstractions.
    

**Excluded Topics and Boundaries:**

- Do not include commentary on application-level security or encryption unless explicitly asked.
    
- Avoid comparisons with other languages unless specifically requested.
    
- Refrain from generating speculative architectural advice without context.
    

Remain tool-agnostic (e.g. refrain from assuming use of VS Code, JetBrains GoLand, etc.) unless explicitly specified.

Your goal is to equip developers with reliable, production-grade Go knowledge suitable for real-world development and collaboration.

# Go Standard Library Reference

## Standard Library Index

| Package | Description |
|--------|-------------|
| [fmt](#fmt-package) | Implements formatted I/O (e.g., printing to stdout, formatting strings). |
| [os](#os-package) | Provides platform-independent interface to operating system functionality. |
| [io](#io-package) | Provides basic interfaces to I/O primitives (e.g., Reader, Writer). |
| [bufio](#bufio-package) | Buffered I/O for reading and writing with efficiency. |
| [strings](#strings-package) | Implements simple functions to manipulate UTF-8 encoded strings. |
| [strconv](#strconv-package) | Provides conversions to and from string representations of basic types. |
| [math](#math-package) | Provides basic constants and mathematical functions. |
| [math/rand](#mathrand-package) | Implements pseudo-random number generators. |
| [time](#time-package) | Provides functionality for measuring and displaying time. |
| [net](#net-package) | Provides a portable interface for network I/O, including TCP/IP. |
| [net/http](#nethttp-package) | Provides HTTP client and server implementations. |
| [encoding/json](#encodingjson-package) | Implements JSON encoding and decoding. |
| [encoding/xml](#encodingxml-package) | Implements XML parsing and encoding. |
| [path](#path-package) | Implements utility routines for manipulating slash-separated paths. |
| [path/filepath](#pathfilepath-package) | Implements utility routines for manipulating filename paths. |
| [regexp](#regexp-package) | Implements regular expressions using re2 syntax. |
| [sync](#sync-package) | Provides basic synchronisation primitives (e.g., Mutex, WaitGroup). |
| [sync/atomic](#syncatomic-package) | Provides low-level atomic memory primitives. |
| [context](#context-package) | Defines context for carrying deadlines, cancel signals, and other values. |
| [log](#log-package) | Provides simple logging package with configurable output. |
| [flag](#flag-package) | Implements command-line flag parsing. |
| [errors](#errors-package) | Defines functions for creating and inspecting errors. |
| [reflect](#reflect-package) | Implements run-time reflection, allowing inspection of types and values. |
| [testing](#testing-package) | Supports automated testing of Go packages. |
| [sort](#sort-package) | Provides primitives for sorting slices and user-defined collections. |
| [bytes](#bytes-package) | Implements functions for manipulating byte slices. |
| [unicode](#unicode-package) | Provides data and functions for working with Unicode characters. |
| [unicode/utf8](#unicodeutf8-package) | Provides functions for UTF-8 encoding and decoding. |
| [runtime](#runtime-package) | Contains operations that interact with Go runtime system (e.g., GC info). |
| [runtime/debug](#runtimedebug-package) | Provides facilities for debugging Go programs. |
| [encoding/base64](#encodingbase64-package) | Implements base64 encoding and decoding. |

# fmt Package

The `fmt` package implements formatted I/O with functions analogous to C's `printf` and `scanf`.

## Functions and Descriptions

| Function               | Description                                                                                   | Parameters                                                                                         |
|------------------------|-----------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------|
| `Print`                | Prints to standard output. Returns the number of bytes written and any write error.           | `a ...interface{}`                                                                                 |
| `Printf`               | Formats according to a format specifier and writes to standard output.                        | `format string, a ...interface{}`                                                                  |
| `Println`              | Prints to standard output with spaces and a newline.                                          | `a ...interface{}`                                                                                 |
| `Sprint`               | Returns a formatted string without printing it.                                               | `a ...interface{}`                                                                                 |
| `Sprintf`              | Returns a formatted string based on the format specifier.                                     | `format string, a ...interface{}`                                                                  |
| `Sprintln`             | Returns a formatted string with spaces and a newline.                                         | `a ...interface{}`                                                                                 |
| `Fprint`               | Writes to the given writer (e.g., file, buffer). Returns bytes written and any error.         | `w io.Writer, a ...interface{}`                                                                    |
| `Fprintf`              | Like `Printf`, but writes to a given writer.                                                  | `w io.Writer, format string, a ...interface{}`                                                     |
| `Fprintln`             | Like `Println`, but writes to a given writer.                                                 | `w io.Writer, a ...interface{}`                                                                    |
| `Errorf`               | Formats and returns an error instead of writing to output.                                    | `format string, a ...interface{}`                                                                  |
| `Scan`                 | Scans input from standard input, storing values pointed to by arguments.                      | `a ...interface{}`                                                                                 |
| `Scanf`                | Scans input from standard input based on format string.                                       | `format string, a ...interface{}`                                                                  |
| `Scanln`               | Scans a line from standard input, stopping at newline.                                        | `a ...interface{}`                                                                                 |
| `Sscan`                | Scans from a string.                                                                          | `str string, a ...interface{}`                                                                     |
| `Sscanf`               | Scans formatted text from a string.                                                           | `str string, format string, a ...interface{}`                                                      |
| `Sscanln`              | Scans a line from a string.                                                                   | `str string, a ...interface{}`                                                                     |
| `Fscan`                | Scans from a reader.                                                                          | `r io.Reader, a ...interface{}`                                                                    |
| `Fscanf`               | Scans formatted input from a reader.                                                          | `r io.Reader, format string, a ...interface{}`                                                     |
| `Fscanln`              | Scans a line from a reader.                                                                   | `r io.Reader, a ...interface{}`                                                                    |

## Common Format Specifiers

| Specifier | Meaning                                 |
|-----------|------------------------------------------|
| `%v`      | Default format                          |
| `%+v`     | Adds field names to structs             |
| `%#v`     | Go syntax representation                |
| `%T`      | Type of the value                       |
| `%t`      | Boolean                                 |
| `%d`      | Decimal integer                         |
| `%b`      | Binary                                  |
| `%x`, `%X`| Hexadecimal                             |
| `%f`      | Decimal floating point                  |
| `%s`      | String                                  |
| `%q`      | Quoted string                           |
| `%p`      | Pointer                                 |

# os Package

The `os` package provides a platform-independent interface to operating system functionality such as file and process management.

## Functions and Descriptions

| Function / Type             | Description                                                                                 | Parameters / Fields                                                                                   |
|-----------------------------|---------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| `Open`                      | Opens a file for reading. Returns a `*File` and error.                                      | `name string`                                                                                          |
| `Create`                    | Creates a file, truncating if it already exists. Returns a `*File` and error.              | `name string`                                                                                          |
| `Remove`                    | Removes (deletes) the named file or directory.                                              | `name string`                                                                                          |
| `Rename`                    | Renames (moves) a file or directory.                                                       | `oldpath string, newpath string`                                                                       |
| `Mkdir`                     | Creates a new directory with specified permissions.                                         | `name string, perm FileMode`                                                                           |
| `MkdirAll`                  | Creates a directory and all necessary parents.                                              | `path string, perm FileMode`                                                                           |
| `Stat`                      | Returns a `FileInfo` describing the named file.                                             | `name string`                                                                                          |
| `Lstat`                     | Like `Stat` but does not follow symbolic links.                                             | `name string`                                                                                          |
| `Chmod`                     | Changes the permissions of a named file.                                                    | `name string, mode FileMode`                                                                           |
| `Getwd`                     | Returns a rooted path name corresponding to the current directory.                         | –                                                                                                      |
| `Chdir`                     | Changes the current working directory to the named directory.                              | `dir string`                                                                                           |
| `Executable`               | Returns the path name for the executable that started the current process.                 | –                                                                                                      |
| `Exit`                      | Exits the program with the given status code.                                               | `code int`                                                                                             |
| `Getenv`                    | Gets the value of an environment variable.                                                  | `key string`                                                                                           |
| `Setenv`                    | Sets the value of an environment variable.                                                  | `key string, value string`                                                                             |
| `Unsetenv`                  | Deletes an environment variable.                                                            | `key string`                                                                                           |
| `Environ`                   | Returns a copy of strings representing the environment, in "key=value" form.                | –                                                                                                      |
| `TempDir`                   | Returns the default directory to use for temporary files.                                   | –                                                                                                      |
| `ReadDir`                   | Reads the directory named by dirname and returns a list of directory entries.              | `name string`                                                                                          |
| `Getpid`                    | Returns the process ID of the caller.                                                       | –                                                                                                      |
| `Getppid`                   | Returns the parent process ID of the caller.                                                | –                                                                                                      |
| `StartProcess`              | Starts a new process with the given arguments.                                              | `name string, argv []string, attr *ProcAttr`                                                           |
| `FindProcess`               | Finds a process by its PID.                                                                | `pid int`                                                                                              |

## Common Types

| Type        | Description                                                                                              |
|-------------|----------------------------------------------------------------------------------------------------------|
| `File`      | Represents an open file descriptor.                                                                      |
| `FileInfo`  | Interface describing a file object (used with `Stat`, etc.).                                             |
| `Process`   | Represents a running process.                                                                            |
| `ProcAttr`  | Specifies attributes for a new process.                                                                  |
| `FileMode`  | Bitmask describing a file’s mode and permission bits.                                                    |

# io Package

The `io` package provides basic interfaces to I/O primitives, as well as some fundamental I/O utility functions.

## Functions and Descriptions

| Function / Type         | Description                                                                 | Parameters / Fields                                                                 |
|-------------------------|-----------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| `Copy`                  | Copies from `src` to `dst` until `EOF` or error. Returns bytes copied and error. | `dst Writer, src Reader`                                                             |
| `CopyN`                 | Copies `n` bytes from `src` to `dst`.                                        | `dst Writer, src Reader, n int64`                                                   |
| `CopyBuffer`            | Like `Copy`, but uses a provided buffer.                                     | `dst Writer, src Reader, buf []byte`                                                |
| `ReadAll`               | Reads from `r` until `EOF` and returns the data.                             | `r Reader`                                                                          |
| `ReadFull`              | Reads exactly `len(buf)` bytes into `buf`.                                   | `r Reader, buf []byte`                                                              |
| `WriteString`           | Writes the contents of string `s` to `w`.                                    | `w Writer, s string`                                                                |
| `Pipe`                  | Creates a synchronous in-memory pipe. Returns connected `PipeReader` and `PipeWriter`. | –                                                                            |
| `TeeReader`             | Returns a `Reader` that writes to `w` what it reads from `r`.                | `r Reader, w Writer`                                                                |
| `SectionReader`         | Creates a `Reader` that reads from `r` starting at `off` and stopping at `off+n`. | `r ReaderAt, off int64, n int64`                                             |
| `LimitReader`           | Returns a `Reader` that reads from `r` but stops after `n` bytes.            | `r Reader, n int64`                                                                 |
| `MultiReader`           | Returns a `Reader` that's the logical concatenation of `r1`, `r2`, etc.      | `readers ...Reader`                                                                 |
| `MultiWriter`           | Returns a `Writer` that duplicates writes to all provided writers.           | `writers ...Writer`                                                                 |
| `NopCloser`             | Wraps a `Reader` with a `ReadCloser` with a no-op `Close` method.            | `r Reader`                                                                          |

## Interfaces

| Interface      | Description                                                                 |
|----------------|-----------------------------------------------------------------------------|
| `Reader`       | Basic interface for reading: `Read(p []byte) (n int, err error)`            |
| `Writer`       | Basic interface for writing: `Write(p []byte) (n int, err error)`           |
| `Closer`       | Interface with a `Close()` method                                           |
| `ReadCloser`   | Combines `Reader` and `Closer`                                              |
| `WriteCloser`  | Combines `Writer` and `Closer`                                              |
| `Seeker`       | Interface for seeking in a stream: `Seek(offset int64, whence int)`         |
| `ReaderAt`     | Reads at specific offsets: `ReadAt(p []byte, off int64)`                    |
| `WriterAt`     | Writes at specific offsets: `WriteAt(p []byte, off int64)`                  |
| `ByteReader`   | Interface for reading a single byte: `ReadByte() (byte, error)`             |
| `ByteWriter`   | Interface for writing a single byte: `WriteByte(byte) error`                |
| `RuneReader`   | Reads a single UTF-8 encoded Unicode character                             |
| `StringWriter` | Interface for writing a string: `WriteString(s string) (n int, err error)`  |

# bufio Package

The `bufio` package implements buffered I/O. It wraps `io.Reader` and `io.Writer` objects with buffering to reduce the number of system calls.

## Functions and Descriptions

| Function / Type          | Description                                                                 | Parameters / Fields                                                                 |
|--------------------------|-----------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| `NewReader`              | Returns a new buffered reader that reads from `r`.                          | `r io.Reader`                                                                       |
| `NewReaderSize`          | Like `NewReader`, but allows specifying the buffer size.                    | `r io.Reader, size int`                                                             |
| `NewWriter`              | Returns a new buffered writer that writes to `w`.                           | `w io.Writer`                                                                       |
| `NewWriterSize`          | Like `NewWriter`, but with specified buffer size.                           | `w io.Writer, size int`                                                             |
| `NewScanner`             | Returns a new `Scanner` to read from `r` token by token.                    | `r io.Reader`                                                                       |

## Types and Methods

### Reader

| Method                | Description                                                                 | Parameters / Returns                                                                 |
|-----------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `Read`                | Reads data into `p`.                                                        | `p []byte` → `(n int, err error)`                                                   |
| `ReadByte`            | Reads and returns a single byte.                                            | → `(byte, error)`                                                                   |
| `ReadRune`            | Reads and returns a single UTF-8 encoded rune.                              | → `(rune, size int, err error)`                                                     |
| `Peek`                | Returns the next `n` bytes without advancing the reader.                    | `n int` → `([]byte, error)`                                                         |
| `ReadLine`            | Reads a single line (deprecated – use Scanner).                             | → `(line []byte, isPrefix bool, err error)`                                         |
| `ReadString`          | Reads until the first occurrence of `delim` byte and returns a string.      | `delim byte` → `(string, error)`                                                    |
| `ReadBytes`           | Reads until the first occurrence of `delim` and returns a byte slice.       | `delim byte` → `([]byte, error)`                                                    |
| `Discard`             | Discards the next `n` bytes.                                                | `n int` → `(discarded int, err error)`                                              |
| `Reset`               | Resets the reader to read from a new source.                                | `r io.Reader`                                                                       |

### Writer

| Method                | Description                                                                 | Parameters / Returns                                                                 |
|-----------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `Write`               | Writes data to the buffer.                                                  | `p []byte` → `(n int, err error)`                                                   |
| `WriteByte`           | Writes a single byte to the buffer.                                         | `c byte` → `error`                                                                   |
| `WriteRune`           | Writes a single UTF-8 encoded rune.                                         | `r rune` → `(size int, err error)`                                                  |
| `WriteString`         | Writes a string to the buffer.                                              | `s string` → `(n int, err error)`                                                   |
| `Flush`               | Writes any buffered data to the underlying writer.                          | → `error`                                                                            |
| `Available`           | Returns how many bytes can be written without allocation.                   | → `int`                                                                              |
| `Buffered`            | Returns the number of bytes that have been written into the buffer.         | → `int`                                                                              |
| `Reset`               | Resets the writer to write to a new destination.                            | `w io.Writer`                                                                        |

### Scanner

| Method                | Description                                                                 | Parameters / Returns                                                                 |
|-----------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `Scan`                | Advances the Scanner to the next token.                                     | → `bool`                                                                             |
| `Text`                | Returns the most recent token as a string.                                  | → `string`                                                                           |
| `Bytes`               | Returns the most recent token as a byte slice.                              | → `[]byte`                                                                           |
| `Err`                 | Returns the first non-EOF error encountered.                                | → `error`                                                                            |
| `Split`               | Sets the split function for the scanner.                                    | `split SplitFunc`                                                                    |
| `Buffer`              | Sets an initial buffer and maximum size for the scanner.                    | `buf []byte, max int`                                                                |

# strings Package

The `strings` package provides functions to manipulate UTF-8 encoded strings.

## Functions and Descriptions

| Function                  | Description                                                                 | Parameters / Returns                                                                 |
|---------------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `Compare`                 | Lexicographically compares two strings.                                     | `a, b string` → `int` (−1 if a < b, 0 if a == b, +1 if a > b)                        |
| `Contains`               | Checks if a substring is within a string.                                   | `s, substr string` → `bool`                                                          |
| `ContainsAny`           | Checks if any Unicode code points in `chars` are within `s`.                 | `s, chars string` → `bool`                                                           |
| `ContainsRune`          | Checks if a specific rune exists in a string.                                | `s string, r rune` → `bool`                                                          |
| `Count`                  | Counts non-overlapping instances of `substr` in `s`.                         | `s, substr string` → `int`                                                           |
| `EqualFold`              | Reports whether two strings are equal under Unicode case folding.            | `s, t string` → `bool`                                                               |
| `Fields`                 | Splits a string around white space.                                          | `s string` → `[]string`                                                              |
| `FieldsFunc`             | Splits a string using a custom function.                                     | `s string, f func(rune) bool` → `[]string`                                           |
| `HasPrefix`              | Checks whether the string begins with prefix.                                | `s, prefix string` → `bool`                                                          |
| `HasSuffix`              | Checks whether the string ends with suffix.                                  | `s, suffix string` → `bool`                                                          |
| `Index`                  | Returns the index of the first occurrence of `substr` in `s`.                | `s, substr string` → `int`                                                           |
| `IndexAny`               | Returns the index of the first Unicode code point in `chars` found in `s`.   | `s, chars string` → `int`                                                            |
| `IndexByte`              | Returns the index of the first instance of byte `c` in `s`.                  | `s string, c byte` → `int`                                                           |
| `IndexRune`              | Returns the index of the first instance of rune `r` in `s`.                  | `s string, r rune` → `int`                                                           |
| `Join`                   | Concatenates elements of `a` with `sep` in between.                          | `a []string, sep string` → `string`                                                  |
| `LastIndex`              | Returns the index of the last occurrence of `substr` in `s`.                 | `s, substr string` → `int`                                                           |
| `Map`                    | Applies a mapping function to each rune in the string.                       | `mapping func(rune) rune, s string` → `string`                                       |
| `Repeat`                 | Repeats the string `count` times.                                            | `s string, count int` → `string`                                                     |
| `Replace`                | Replaces `n` occurrences of `old` with `new` in `s`.                         | `s, old, new string, n int` → `string`                                               |
| `ReplaceAll`             | Replaces all occurrences of `old` with `new`.                                | `s, old, new string` → `string`                                                      |
| `Split`                  | Splits a string around each instance of `sep`.                               | `s, sep string` → `[]string`                                                         |
| `SplitAfter`             | Splits after each instance of `sep`.                                         | `s, sep string` → `[]string`                                                         |
| `SplitN`                 | Splits `s` into at most `n` substrings.                                      | `s, sep string, n int` → `[]string`                                                  |
| `SplitAfterN`            | Splits after each instance of `sep`, limited to `n` substrings.              | `s, sep string, n int` → `[]string`                                                  |
| `Title`                  | Returns a copy of `s` with all words title-cased. (Deprecated in Go 1.18)    | `s string` → `string`                                                                |
| `ToLower`                | Returns a lowercase version of the string.                                   | `s string` → `string`                                                                |
| `ToTitle`                | Returns an upper-case version of the string using Unicode rules.             | `s string` → `string`                                                                |
| `ToUpper`                | Returns an upper-case version of the string.                                 | `s string` → `string`                                                                |
| `Trim`                   | Trims all leading and trailing Unicode code points in `cutset`.              | `s, cutset string` → `string`                                                        |
| `TrimLeft`               | Trims all leading Unicode code points in `cutset`.                           | `s, cutset string` → `string`                                                        |
| `TrimRight`              | Trims all trailing Unicode code points in `cutset`.                          | `s, cutset string` → `string`                                                        |
| `TrimPrefix`             | Removes `prefix` from the start of `s`, if present.                          | `s, prefix string` → `string`                                                        |
| `TrimSuffix`             | Removes `suffix` from the end of `s`, if present.                            | `s, suffix string` → `string`                                                        |
| `TrimSpace`              | Trims leading and trailing white space.                                      | `s string` → `string`                                                                |


# strconv Package

The `strconv` package implements conversions to and from string representations of basic data types such as integers, floats, booleans, and more.

## Functions and Descriptions

| Function                  | Description                                                                 | Parameters / Returns                                                                 |
|---------------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `Atoi`                    | Converts a string to an integer.                                            | `s string` → `(int, error)`                                                          |
| `Itoa`                    | Converts an integer to a string.                                            | `i int` → `string`                                                                   |
| `ParseBool`               | Converts a string to a boolean. Accepts 1, t, T, TRUE, true, True etc.      | `s string` → `(bool, error)`                                                         |
| `FormatBool`              | Converts a boolean to a string ("true" or "false").                         | `b bool` → `string`                                                                  |
| `ParseInt`                | Parses a signed integer from a string in a given base (2 to 36).            | `s string, base int, bitSize int` → `(int64, error)`                                 |
| `ParseUint`               | Parses an unsigned integer from a string.                                   | `s string, base int, bitSize int` → `(uint64, error)`                                |
| `FormatInt`               | Converts an integer to a string in the given base.                          | `i int64, base int` → `string`                                                       |
| `FormatUint`              | Converts an unsigned integer to a string in the given base.                 | `i uint64, base int` → `string`                                                      |
| `AppendInt`               | Appends the string form of an integer to a byte slice.                      | `dst []byte, i int64, base int` → `[]byte`                                           |
| `AppendUint`              | Appends the string form of an unsigned integer to a byte slice.             | `dst []byte, i uint64, base int` → `[]byte`                                          |
| `ParseFloat`              | Converts a string to a floating-point number.                               | `s string, bitSize int` → `(float64, error)`                                         |
| `FormatFloat`             | Converts a floating-point number to a string in specified format and precision. | `f float64, fmt byte, prec, bitSize int` → `string`                              |
| `AppendFloat`             | Appends the string form of a float to a byte slice.                         | `dst []byte, f float64, fmt byte, prec, bitSize int` → `[]byte`                      |
| `Quote`                   | Returns a double-quoted Go string literal safely escaped.                   | `s string` → `string`                                                                |
| `QuoteToASCII`            | Quotes and escapes non-ASCII characters.                                    | `s string` → `string`                                                                |
| `QuoteRune`               | Returns a single-quoted Go character literal safely escaped.                | `r rune` → `string`                                                                  |
| `QuoteRuneToASCII`        | Like `QuoteRune`, but escapes non-ASCII characters.                         | `r rune` → `string`                                                                  |
| `CanBackquote`            | Reports whether a string can be represented unchanged as a raw string literal. | `s string` → `bool`                                                               |
| `Unquote`                 | Interprets a quoted Go string literal and returns the unquoted value.       | `s string` → `(string, error)`                                                       |
| `UnquoteChar`             | Unquotes a single character (used when parsing quoted literals).            | `s string, quote byte` → `(value rune, multibyte bool, tail string, err error)`      |

# math Package

The `math` package provides basic constants and mathematical functions for floating-point and integer math.

## Constants

| Constant       | Description                                     |
|----------------|-------------------------------------------------|
| `Pi`           | The ratio of a circle’s circumference to diameter (π ≈ 3.14159) |
| `E`            | Euler’s number (e ≈ 2.71828)                     |
| `Sqrt2`        | Square root of 2                                |
| `SqrtE`        | Square root of e                                |
| `SqrtPi`       | Square root of π                                |
| `Ln2`          | Natural logarithm of 2                          |
| `Ln10`         | Natural logarithm of 10                         |
| `Log2E`        | Base-2 logarithm of e                          |
| `Log10E`       | Base-10 logarithm of e                         |
| `MaxFloat32`   | Largest float32 value                           |
| `MaxFloat64`   | Largest float64 value                           |
| `SmallestNonzeroFloat32` | Smallest non-zero float32 value           |
| `SmallestNonzeroFloat64` | Smallest non-zero float64 value           |

## Functions and Descriptions

| Function             | Description                                                           | Parameters / Returns                                       |
|----------------------|-----------------------------------------------------------------------|------------------------------------------------------------|
| `Abs`                | Absolute value of a float                                             | `x float64` → `float64`                                    |
| `Acos`               | Inverse cosine                                                        | `x float64` → `float64`                                    |
| `Asin`               | Inverse sine                                                          | `x float64` → `float64`                                    |
| `Atan`               | Inverse tangent                                                       | `x float64` → `float64`                                    |
| `Atan2`              | Arc tangent of y/x, result in correct quadrant                        | `y, x float64` → `float64`                                 |
| `Ceil`               | Smallest integer ≥ x                                                  | `x float64` → `float64`                                    |
| `Cos`                | Cosine of x (x in radians)                                            | `x float64` → `float64`                                    |
| `Cosh`               | Hyperbolic cosine                                                     | `x float64` → `float64`                                    |
| `Exp`                | e raised to the power of x                                            | `x float64` → `float64`                                    |
| `Exp2`               | 2 raised to the power of x                                            | `x float64` → `float64`                                    |
| `Expm1`              | e^x − 1, more accurate for small x                                    | `x float64` → `float64`                                    |
| `Floor`              | Largest integer ≤ x                                                   | `x float64` → `float64`                                    |
| `Hypot`              | sqrt(p² + q²), the length of the hypotenuse                          | `p, q float64` → `float64`                                 |
| `Log`                | Natural logarithm (base e)                                            | `x float64` → `float64`                                    |
| `Log10`              | Base-10 logarithm                                                     | `x float64` → `float64`                                    |
| `Log1p`              | Log(1 + x), more accurate for small x                                 | `x float64` → `float64`                                    |
| `Log2`               | Base-2 logarithm                                                      | `x float64` → `float64`                                    |
| `Max`                | Larger of x or y                                                      | `x, y float64` → `float64`                                 |
| `Min`                | Smaller of x or y                                                     | `x, y float64` → `float64`                                 |
| `Mod`                | Modulo operation (remainder of division)                              | `x, y float64` → `float64`                                 |
| `Pow`                | x raised to the power y                                               | `x, y float64` → `float64`                                 |
| `Pow10`              | 10^n                                                                  | `n int` → `float64`                                        |
| `Remainder`          | IEEE 754 remainder of x / y                                           | `x, y float64` → `float64`                                 |
| `Round`              | Rounds to nearest integer, halves away from zero                      | `x float64` → `float64`                                    |
| `RoundToEven`        | Rounds to nearest even integer on ties                                | `x float64` → `float64`                                    |
| `Signbit`            | Returns true if x is negative or negative zero                       | `x float64` → `bool`                                       |
| `Sin`                | Sine of x (x in radians)                                              | `x float64` → `float64`                                    |
| `Sinh`               | Hyperbolic sine                                                       | `x float64` → `float64`                                    |
| `Sqrt`               | Square root                                                           | `x float64` → `float64`                                    |
| `Tan`                | Tangent of x (x in radians)                                           | `x float64` → `float64`                                    |
| `Tanh`               | Hyperbolic tangent                                                    | `x float64` → `float64`                                    |
| `Trunc`              | Integer value of x toward zero                                        | `x float64` → `float64`                                    |
| `Modf`               | Splits number into integer and fractional parts                       | `x float64` → `(int float64, frac float64)`                |
| `IsNaN`              | Checks if a number is NaN                                             | `x float64` → `bool`                                       |
| `IsInf`              | Checks if a number is infinite (±∞)                                   | `f float64, sign int` → `bool`                             |
| `Frexp`              | Breaks float into mantissa and exponent                               | `f float64` → `(frac float64, exp int)`                    |
| `Ldexp`              | Inverse of Frexp: frac × 2^exp                                        | `frac float64, exp int` → `float64`                        |

# math/rand Package

The `math/rand` package provides pseudo-random number generators suitable for simulations and non-cryptographic use.

## Functions and Descriptions

| Function                | Description                                                                 | Parameters / Returns                                            |
|-------------------------|-----------------------------------------------------------------------------|-----------------------------------------------------------------|
| `Int`                   | Returns a non-negative pseudo-random int.                                  | → `int`                                                         |
| `Intn`                  | Returns, as an int, a non-negative pseudo-random number less than `n`.     | `n int` → `int`                                                 |
| `Int31`                 | Returns a non-negative 31-bit pseudo-random int.                           | → `int32`                                                       |
| `Int31n`                | Returns a non-negative pseudo-random int32 less than `n`.                  | `n int32` → `int32`                                             |
| `Int63`                 | Returns a non-negative 63-bit pseudo-random int.                           | → `int64`                                                       |
| `Int63n`                | Returns a non-negative pseudo-random int64 less than `n`.                  | `n int64` → `int64`                                             |
| `Float32`              | Returns a pseudo-random float32 in [0.0, 1.0).                              | → `float32`                                                     |
| `Float64`              | Returns a pseudo-random float64 in [0.0, 1.0).                              | → `float64`                                                     |
| `ExpFloat64`           | Returns an exponentially distributed float64 with mean 1.0.                 | → `float64`                                                     |
| `NormFloat64`          | Returns a normally distributed float64 (standard normal distribution).      | → `float64`                                                     |
| `Perm`                  | Returns a pseudo-random permutation of integers [0, n).                    | `n int` → `[]int`                                               |
| `Shuffle`               | Randomises the order of elements in a slice.                               | `n int, swap func(i, j int)`                                   |
| `Seed`                  | Sets the seed for the default Source.                                      | `seed int64`                                                    |
| `New`                   | Returns a new Rand using the provided Source.                              | `src Source` → `*Rand`                                          |
| `NewSource`             | Returns a new Source seeded with the provided value.                       | `seed int64` → `Source`                                        |

## Types and Methods

### Type: `Rand`

A source of pseudo-random numbers.

| Method                  | Description                                                                 | Parameters / Returns                                            |
|-------------------------|-----------------------------------------------------------------------------|-----------------------------------------------------------------|
| `Int`                   | Returns a non-negative int                                                  | → `int`                                                         |
| `Intn`                  | Returns int in [0, n)                                                       | `n int` → `int`                                                 |
| `Float32`              | Returns float32 in [0.0, 1.0)                                               | → `float32`                                                     |
| `Float64`              | Returns float64 in [0.0, 1.0)                                               | → `float64`                                                     |
| `ExpFloat64`           | Returns exponentially distributed float64                                   | → `float64`                                                     |
| `NormFloat64`          | Returns normally distributed float64                                        | → `float64`                                                     |
| `Perm`                  | Returns permutation [0, n)                                                  | `n int` → `[]int`                                               |
| `Shuffle`               | Shuffles a slice                                                            | `n int, swap func(i, j int)`                                   |
| `Seed`                  | Sets the seed                                                               | `seed int64`                                                    |

## Interfaces

| Interface | Description                                                  |
|-----------|--------------------------------------------------------------|
| `Source`  | Represents a source of uniformly-distributed pseudo-random numbers |
| `Source64`| Like `Source`, but has `Uint64()` method as well             |

> ⚠️ Note: Not safe for cryptographic use. Use `crypto/rand` for security-sensitive randomness.

# time Package

The `time` package provides functionality for measuring and displaying time. It includes types for durations, instants, and formatting.

## Constants

| Constant     | Description                             |
|--------------|-----------------------------------------|
| `Nanosecond` | 1 nanosecond                            |
| `Microsecond`| 1000 nanoseconds                        |
| `Millisecond`| 1000 microseconds                       |
| `Second`     | 1 second                                |
| `Minute`     | 60 seconds                              |
| `Hour`       | 60 minutes                              |

## Functions and Descriptions

| Function             | Description                                                                 | Parameters / Returns                                      |
|----------------------|-----------------------------------------------------------------------------|-----------------------------------------------------------|
| `Now`                | Returns the current local time.                                              | → `Time`                                                  |
| `Sleep`              | Pauses the current goroutine for the specified duration.                    | `d Duration`                                              |
| `Since`              | Returns time elapsed since `t`.                                             | `t Time` → `Duration`                                     |
| `Until`              | Returns duration until `t`.                                                 | `t Time` → `Duration`                                     |
| `Parse`              | Parses a formatted string and returns the time value.                       | `layout, value string` → `(Time, error)`                  |
| `ParseInLocation`    | Like `Parse`, but uses the provided Location.                               | `layout, value string, loc *Location` → `(Time, error)`   |
| `Unix`               | Returns local time corresponding to Unix time (sec, nsec).                  | `sec int64, nsec int64` → `Time`                          |
| `Date`               | Returns a Time from year, month, day, etc.                                  | `year int, month Month, day, hour, min, sec, nsec int, loc *Location` → `Time` |
| `FixedZone`          | Returns a Location that uses a fixed offset.                                | `name string, offset int` → `*Location`                   |
| `LoadLocation`       | Loads a Location by name (e.g., "Europe/London").                           | `name string` → `(*Location, error)`                      |
| `Tick`               | Returns a channel that delivers ticks every `Duration`.                     | `d Duration` → `<-chan Time`                              |
| `After`              | Returns a channel that receives the current time after a duration.          | `d Duration` → `<-chan Time`                              |
| `NewTicker`          | Returns a new Ticker that ticks periodically.                               | `d Duration` → `*Ticker`                                  |
| `NewTimer`           | Returns a Timer that fires once after duration.                             | `d Duration` → `*Timer`                                   |
| `AfterFunc`          | Runs a function once after duration.                                        | `d Duration, f func()` → `*Timer`                         |

## Types and Key Methods

### Type: `Time`

Represents an instant in time with nanosecond precision.

| Method                 | Description                                              | Returns                            |
|------------------------|----------------------------------------------------------|------------------------------------|
| `Add`                  | Adds a duration                                          | `Time`                             |
| `Sub`                  | Returns the duration between two times                   | `Duration`                         |
| `Before`, `After`      | Compares time values                                     | `bool`                             |
| `Equal`                | Reports whether two times are equal                      | `bool`                             |
| `IsZero`               | Reports whether the time is the zero time                | `bool`                             |
| `Unix`, `UnixNano`     | Returns Unix timestamp in seconds or nanoseconds         | `int64`                            |
| `Format`               | Formats time using a layout string                       | `string`                           |
| `In`                   | Returns time in a specified location                     | `Time`                             |
| `Location`             | Returns the Location associated with the time            | `*Location`                        |
| `Round`, `Truncate`    | Rounds/truncates time to nearest multiple of duration    | `Time`                             |

### Type: `Duration`

Represents elapsed time between two instants.

| Method          | Description                                          | Returns           |
|------------------|------------------------------------------------------|-------------------|
| `Hours`          | Duration in hours                                   | `float64`         |
| `Minutes`        | Duration in minutes                                 | `float64`         |
| `Seconds`        | Duration in seconds                                 | `float64`         |
| `Milliseconds`   | Duration in milliseconds                            | `float64`         |
| `Microseconds`   | Duration in microseconds                            | `float64`         |
| `Nanoseconds`    | Duration in nanoseconds (integer)                   | `int64`           |
| `String`         | Returns the string representation (e.g., "2h45m")   | `string`          |

# net Package

The `net` package provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix sockets.

## Functions and Descriptions

| Function                  | Description                                                                  | Parameters / Returns                                      |
|---------------------------|------------------------------------------------------------------------------|-----------------------------------------------------------|
| `Dial`                    | Connects to a network address (TCP, UDP, etc.).                              | `network, address string` → `(Conn, error)`               |
| `DialTimeout`             | Like `Dial` but with a timeout.                                              | `network, address string, timeout time.Duration` → `(Conn, error)` |
| `DialTCP`                 | Connects to a TCP address.                                                   | `network string, laddr, raddr *TCPAddr` → `(TCPConn, error)` |
| `Listen`                  | Creates a generic listener for a network.                                    | `network, address string` → `(Listener, error)`           |
| `ListenTCP`               | Listens for incoming TCP connections.                                        | `network string, laddr *TCPAddr` → `(TCPListener, error)` |
| `ListenUDP`               | Listens for incoming UDP datagrams.                                          | `network string, laddr *UDPAddr` → `(UDPConn, error)`     |
| `ListenPacket`            | Creates a connectionless packet listener.                                    | `network, address string` → `(PacketConn, error)`         |
| `ResolveIPAddr`           | Resolves a host/IP to an IPAddr struct.                                      | `network, address string` → `(IPAddr, error)`             |
| `ResolveTCPAddr`          | Resolves a TCP address.                                                      | `network, address string` → `(TCPAddr, error)`            |
| `ResolveUDPAddr`          | Resolves a UDP address.                                                      | `network, address string` → `(UDPAddr, error)`            |
| `LookupAddr`              | Performs a reverse lookup for an IP address.                                | `addr string` → `([]string, error)`                       |
| `LookupCNAME`             | Returns the canonical DNS name for a host.                                  | `host string` → `(string, error)`                         |
| `LookupHost`              | Looks up hostnames and returns IP addresses.                                | `host string` → `([]string, error)`                       |
| `LookupIP`                | Looks up IP addresses for a host.                                           | `host string` → `([]IP, error)`                           |
| `LookupPort`              | Looks up the port number for a network and service.                         | `network, service string` → `(int, error)`                |
| `Interfaces`              | Returns a list of network interfaces.                                       | → `([]Interface, error)`                                  |

## Interfaces

| Interface    | Description                                                                              |
|--------------|------------------------------------------------------------------------------------------|
| `Conn`       | Represents a generic stream-oriented network connection (e.g., TCP).                    |
| `Listener`   | Represents a generic network listener for stream-oriented protocols.                    |
| `PacketConn` | Represents a connectionless network protocol (e.g., UDP).                               |
| `Addr`       | Represents a network end point address (implemented by `TCPAddr`, `UDPAddr`, etc.).     |

## Types and Key Methods

### Type: `TCPConn`, `UDPConn`

Represents a TCP or UDP connection.

| Method          | Description                                      | Returns            |
|------------------|--------------------------------------------------|--------------------|
| `Read`, `Write`  | I/O operations for sending/receiving data       | standard `io.Reader` / `io.Writer` methods |
| `Close`          | Closes the connection                            | `error`            |
| `LocalAddr`      | Returns the local address                        | `Addr`             |
| `RemoteAddr`     | Returns the remote address                       | `Addr`             |
| `SetDeadline`    | Sets read/write timeout                          | `error`            |

### Type: `TCPAddr`, `UDPAddr`, `IPAddr`

Represents structured addresses used by TCP, UDP, and IP.

| Field      | Description                      |
|------------|----------------------------------|
| `IP`       | IP address (type `net.IP`)       |
| `Port`     | Port number                      |
| `Zone`     | Scope zone for IPv6 addresses    |

# net/http Package

The `net/http` package provides HTTP client and server implementations. It supports building web servers and sending HTTP requests.

## Server-Side Functions and Types

| Function / Type          | Description                                                                | Parameters / Returns                                        |
|--------------------------|----------------------------------------------------------------------------|-------------------------------------------------------------|
| `ListenAndServe`         | Starts an HTTP server on an address with a handler.                        | `addr string, handler Handler` → `error`                    |
| `ListenAndServeTLS`      | Like `ListenAndServe` but with TLS support.                                | `addr, certFile, keyFile string, handler Handler` → `error` |
| `Handle`                 | Registers a handler for a given pattern.                                   | `pattern string, handler Handler`                           |
| `HandleFunc`             | Like `Handle` but takes a function instead of a Handler.                   | `pattern string, handler func(ResponseWriter, *Request)`    |
| `Serve`                  | Serves HTTP requests on a given listener.                                  | `l net.Listener, handler Handler` → `error`                 |
| `ServeTLS`               | Like `Serve` but supports TLS.                                             | `l net.Listener, certFile, keyFile string, handler Handler` |

### Interface: `Handler`

| Method         | Description                                      |
|----------------|--------------------------------------------------|
| `ServeHTTP`    | Handles HTTP requests                            |
| `w ResponseWriter, r *Request` → `void`                           |

### Type: `Request`

Represents an HTTP request.

| Field             | Description                                      |
|-------------------|--------------------------------------------------|
| `Method`          | Request method (GET, POST, etc.)                 |
| `URL`             | Parsed URL (`*url.URL`)                          |
| `Header`          | Map of header keys to values                     |
| `Body`            | Request body (`io.ReadCloser`)                   |
| `Form`            | Parsed form values                               |
| `PostForm`        | Parsed POST or PUT form values                   |
| `RemoteAddr`      | Client address                                   |
| `ContentLength`   | Length of body in bytes                          |

### Type: `ResponseWriter`

Interface used by HTTP handlers to construct responses.

| Method               | Description                                      |
|----------------------|--------------------------------------------------|
| `Header()`           | Returns response headers                         |
| `Write([]byte)`      | Writes response body                             |
| `WriteHeader(int)`   | Sets status code (e.g., 200, 404)                |

## Client-Side Functions and Types

| Function             | Description                                                              | Parameters / Returns                                      |
|----------------------|--------------------------------------------------------------------------|-----------------------------------------------------------|
| `Get`                | Sends a GET request.                                                     | `url string` → `(resp *Response, err error)`              |
| `Post`               | Sends a POST request.                                                    | `url, contentType string, body io.Reader` → `(resp *Response, err error)` |
| `PostForm`           | Sends a POST request with form data.                                     | `url string, data url.Values` → `(resp *Response, err error)`             |
| `Head`               | Sends a HEAD request.                                                    | `url string` → `(resp *Response, err error)`              |
| `NewRequest`         | Creates a new HTTP request.                                              | `method, url string, body io.Reader` → `(*Request, error)` |
| `DefaultClient`      | The default `*Client` used for requests                                  | –                                                         |

### Type: `Client`

Represents an HTTP client.

| Method             | Description                                                              | Returns                        |
|--------------------|--------------------------------------------------------------------------|--------------------------------|
| `Do`               | Sends an HTTP request and returns a response.                           | `req *Request` → `(*Response, error)` |

### Type: `Response`

Represents the HTTP response.

| Field             | Description                                     |
|-------------------|-------------------------------------------------|
| `Status`          | HTTP status string (e.g., "200 OK")            |
| `StatusCode`      | HTTP status code                               |
| `Header`          | Response headers                               |
| `Body`            | Response body (`io.ReadCloser`)                |
| `ContentLength`   | Length of the response body                    |

# encoding/json Package

The `encoding/json` package implements encoding and decoding of JSON as defined in RFC 7159. It supports converting between Go values and JSON representations.

## Core Functions

| Function                | Description                                                                | Parameters / Returns                                       |
|-------------------------|----------------------------------------------------------------------------|------------------------------------------------------------|
| `Marshal`               | Converts a Go value to a JSON-encoded byte slice.                         | `v interface{}` → `([]byte, error)`                        |
| `MarshalIndent`         | Like `Marshal`, but formats the output with indentation.                 | `v interface{}, prefix, indent string` → `([]byte, error)` |
| `Unmarshal`             | Parses JSON-encoded data into a Go value.                                | `data []byte, v interface{}` → `error`                     |
| `Valid`                 | Reports whether the input is a valid JSON encoding.                      | `data []byte` → `bool`                                     |
| `Compact`               | Compacts JSON-encoded data by removing whitespace.                       | `dst *bytes.Buffer, src []byte` → `error`                  |
| `Indent`                | Indents JSON for readability.                                            | `dst *bytes.Buffer, src []byte, prefix, indent string` → `error` |
| `HTMLEscape`            | Escapes JSON for safe use in HTML contexts.                              | `dst *bytes.Buffer, src []byte`                            |

## Types and Methods

### Type: `Decoder`

Used for streaming decoding from an `io.Reader`.

| Method              | Description                                                                | Returns                             |
|---------------------|----------------------------------------------------------------------------|-------------------------------------|
| `Decode`            | Reads the next JSON-encoded value from input.                             | `v interface{}` → `error`           |
| `Buffered`          | Returns buffered data not yet consumed.                                   | `io.Reader`                         |
| `More`              | Reports whether there is another element in the current array or object. | `bool`                              |
| `Token`             | Returns the next JSON token in the input stream.                          | `(Token, error)`                    |
| `UseNumber`         | Causes the Decoder to unmarshal numbers into `json.Number`.              | –                                   |
| `DisallowUnknownFields` | Causes Decoder to throw an error if unknown fields are present.        | –                                   |

### Type: `Encoder`

Used for streaming encoding to an `io.Writer`.

| Method              | Description                                                  | Returns      |
|---------------------|--------------------------------------------------------------|--------------|
| `Encode`            | Writes the JSON encoding of `v` to the stream.               | `v interface{}` → `error` |
| `SetIndent`         | Sets indentation for encoded output.                         | `prefix, indent string`  |
| `SetEscapeHTML`     | Sets whether HTML characters should be escaped.              | `on bool`                |

### Type: `RawMessage`

| Description                                     |
|-------------------------------------------------|
| A raw encoded JSON value used for delayed decoding. Implements `json.Marshaler` and `json.Unmarshaler`. |

### Type: `Number`

| Method         | Description                             | Returns        |
|----------------|-----------------------------------------|----------------|
| `String`       | Returns the string form of the number. | `string`       |
| `Int64`        | Converts to `int64`.                    | `(int64, error)` |
| `Float64`      | Converts to `float64`.                  | `(float64, error)` |

## Struct Tags

| Tag Syntax                         | Effect                                                                |
|------------------------------------|-----------------------------------------------------------------------|
| `json:"fieldname"`                 | Maps the Go struct field to a JSON key.                              |
| `json:"fieldname,omitempty"`      | Omits the field if it has a zero value.                              |
| `json:"-"`                         | Field is never encoded/decoded.                                      |

# encoding/xml Package

The `encoding/xml` package implements a simple XML 1.0 parser and encoder. It supports marshaling and unmarshaling of Go structs to and from XML.

## Core Functions

| Function                | Description                                                                 | Parameters / Returns                                         |
|-------------------------|-----------------------------------------------------------------------------|--------------------------------------------------------------|
| `Marshal`               | Encodes a Go value to XML.                                                  | `v interface{}` → `([]byte, error)`                          |
| `MarshalIndent`         | Like `Marshal`, but with custom indentation and prefix.                    | `v interface{}, prefix, indent string` → `([]byte, error)`   |
| `Unmarshal`             | Parses XML-encoded data into the provided Go value.                        | `data []byte, v interface{}` → `error`                       |
| `EscapeText`            | Escapes text for safe inclusion in XML.                                    | `w io.Writer, s []byte` → `error`                            |
| `Escape`                | Escapes both text and attribute values.                                    | `w io.Writer, s []byte` → `error`                            |
| `CopyToken`             | Copies a token to another writer.                                          | `dst TokenWriter, t Token` → `error`                         |
| `NewDecoder`            | Creates a new XML decoder from an `io.Reader`.                             | `r io.Reader` → `*Decoder`                                   |
| `NewEncoder`            | Creates a new XML encoder for an `io.Writer`.                              | `w io.Writer` → `*Encoder`                                   |

## Types and Methods

### Type: `Decoder`

| Method                | Description                                                        | Returns                        |
|------------------------|--------------------------------------------------------------------|--------------------------------|
| `Decode`              | Reads the next XML value into `v`.                                 | `v interface{}` → `error`      |
| `Token`               | Returns the next XML token in the input.                           | `(Token, error)`               |
| `InputOffset`         | Returns the input stream offset of the current token.              | `int64`                        |
| `Skip`                | Skips over the next XML element.                                   | `error`                        |
| `RawToken`            | Like `Token` but includes comments and directives.                 | `(Token, error)`               |

### Type: `Encoder`

| Method                | Description                                                        | Returns                        |
|------------------------|--------------------------------------------------------------------|--------------------------------|
| `Encode`              | Writes XML-encoded value of `v` to the stream.                     | `v interface{}` → `error`      |
| `EncodeElement`       | Writes a single XML element with custom start tag.                 | `v interface{}, start StartElement` → `error` |
| `Flush`               | Flushes any buffered data.                                         | `error`                        |

### Type: `Token` Interface and Implementations

| Token Type        | Description                                |
|-------------------|--------------------------------------------|
| `StartElement`    | XML start tag with attributes              |
| `EndElement`      | XML end tag                                |
| `CharData`        | Text data between tags                     |
| `Comment`         | XML comments                               |
| `ProcInst`        | Processing instruction                     |
| `Directive`       | XML directives (e.g., `<!DOCTYPE>`)        |

## Struct Tags

| Tag Syntax                        | Description                                                            |
|-----------------------------------|------------------------------------------------------------------------|
| `xml:"name"`                      | Sets the element or attribute name                                     |
| `xml:"name,attr"`                | Sets the field as an attribute                                         |
| `xml:",chardata"`                | Sets the field as character data inside an element                     |
| `xml:",innerxml"`               | Stores raw XML inside the field                                       |
| `xml:",comment"`                | Interprets the field as an XML comment                                |
| `xml:"-"`                         | Field is ignored during encoding/decoding                             |

# path Package

The `path` package provides utility functions for manipulating slash-separated paths, typically used for URLs. It is safe for purely textual path handling and does **not** interact with the file system.

> 📝 For OS-specific file paths, use `path/filepath`.

## Functions and Descriptions

| Function             | Description                                                                 | Parameters / Returns                             |
|----------------------|-----------------------------------------------------------------------------|--------------------------------------------------|
| `Base`               | Returns the last element of the path.                                       | `path string` → `string`                         |
| `Dir`                | Returns all but the last element of the path.                               | `path string` → `string`                         |
| `Ext`                | Returns the file extension, including the dot.                              | `path string` → `string`                         |
| `IsAbs`              | Reports whether the path is absolute (i.e., starts with a slash).           | `path string` → `bool`                           |
| `Join`               | Joins any number of path elements into a single, clean path.                | `elem ...string` → `string`                      |
| `Clean`              | Cleans a path by simplifying redundant separators and dot elements.         | `path string` → `string`                         |
| `Split`              | Splits path into directory and file name components.                        | `path string` → `(dir, file string)`             |

# path/filepath Package

The `path/filepath` package provides functions to manipulate filename paths in a way that is compatible with the target operating system (Windows, Unix, etc.).

## Functions and Descriptions

| Function               | Description                                                                 | Parameters / Returns                                 |
|------------------------|-----------------------------------------------------------------------------|------------------------------------------------------|
| `Abs`                  | Returns an absolute representation of the path.                            | `path string` → `(string, error)`                    |
| `Base`                 | Returns the last element of the path.                                      | `path string` → `string`                             |
| `Clean`                | Simplifies a path by removing redundant separators and elements.           | `path string` → `string`                             |
| `Dir`                  | Returns all but the last element of the path.                              | `path string` → `string`                             |
| `EvalSymlinks`         | Returns the path after evaluating symbolic links.                          | `path string` → `(string, error)`                    |
| `Ext`                  | Returns the file extension, including the dot.                             | `path string` → `string`                             |
| `FromSlash`            | Converts slashes to OS-specific separators.                                | `path string` → `string`                             |
| `Glob`                 | Returns the names of files matching a pattern.                             | `pattern string` → `([]string, error)`               |
| `HasPrefix`            | Reports whether `path` starts with `prefix`, respecting path boundaries.   | `path, prefix string` → `bool`                       |
| `IsAbs`                | Reports whether the path is absolute.                                      | `path string` → `bool`                              |
| `Join`                 | Joins any number of path elements into a single path.                      | `elem ...string` → `string`                          |
| `Match`                | Reports whether the name matches the shell pattern.                        | `pattern, name string` → `(bool, error)`             |
| `Rel`                  | Returns a relative path from basepath to targpath.                         | `basepath, targpath string` → `(string, error)`      |
| `Split`                | Splits path into directory and file name.                                  | `path string` → `(dir, file string)`                 |
| `SplitList`            | Splits a list of paths joined by the OS-specific ListSeparator.            | `path string` → `[]string`                           |
| `ToSlash`              | Converts OS-specific path separators to slashes.                           | `path string` → `string`                             |
| `VolumeName`           | Returns the leading volume name.                                           | `path string` → `string`                             |
| `Walk`                 | Traverses the file tree rooted at root, calling walkFn for each file/dir.  | `root string, walkFn WalkFunc` → `error`             |
| `WalkDir`              | Like `Walk`, but uses `fs.DirEntry` instead of `os.FileInfo`.              | `root string, fn fs.WalkDirFunc` → `error`           |

## Interfaces and Types

| Type / Interface       | Description                                                                 |
|------------------------|-----------------------------------------------------------------------------|
| `WalkFunc`             | Function called by `Walk` for each visited file or directory.               |
| `WalkDirFunc`          | Like `WalkFunc`, but receives `fs.DirEntry`.                                |

# regexp Package

The `regexp` package implements regular expressions using RE2 syntax. It provides functions and types for compiling and matching regular expressions efficiently.

## Core Functions

| Function             | Description                                                                 | Parameters / Returns                                 |
|----------------------|-----------------------------------------------------------------------------|------------------------------------------------------|
| `Compile`            | Compiles a regular expression string into a `*Regexp`.                      | `expr string` → `(*Regexp, error)`                   |
| `CompilePOSIX`       | Like `Compile`, but uses POSIX leftmost-longest matching.                  | `expr string` → `(*Regexp, error)`                   |
| `MustCompile`        | Like `Compile`, but panics on error.                                        | `expr string` → `*Regexp`                            |
| `MustCompilePOSIX`   | Like `CompilePOSIX`, but panics on error.                                   | `expr string` → `*Regexp`                            |
| `QuoteMeta`          | Escapes all regular expression metacharacters in a string.                  | `s string` → `string`                                |

## Type: `Regexp`

A `Regexp` represents a compiled regular expression.

| Method                   | Description                                                           | Returns / Parameters                                 |
|--------------------------|-----------------------------------------------------------------------|------------------------------------------------------|
| `Match`                  | Tests whether the Regexp matches the byte slice.                     | `b []byte` → `bool`                                  |
| `MatchString`            | Tests whether the Regexp matches the string.                         | `s string` → `bool`                                  |
| `Find`                   | Finds the leftmost match in the input.                               | `b []byte` → `[]byte`                                |
| `FindString`             | Same as `Find`, but returns a string.                                | `s string` → `string`                                |
| `FindIndex`              | Returns the start and end positions of the match.                    | `b []byte` → `[]int`                                 |
| `FindStringIndex`        | Same as `FindIndex`, but input is a string.                          | `s string` → `[]int`                                 |
| `FindAll`                | Finds all matches in the input.                                      | `b []byte, n int` → `[][]byte`                       |
| `FindAllString`          | Finds all matching substrings.                                       | `s string, n int` → `[]string`                       |
| `FindAllIndex`           | Finds positions of all matches.                                      | `b []byte, n int` → `[][]int`                        |
| `FindAllStringIndex`     | Same as `FindAllIndex`, but with string input.                       | `s string, n int` → `[][]int`                        |
| `SubexpNames`            | Returns the names of subexpressions in the Regexp.                   | → `[]string`                                         |
| `NumSubexp`              | Returns the number of subexpressions.                                | → `int`                                              |
| `ReplaceAll`             | Replaces all matches with a byte slice.                              | `src, repl []byte` → `[]byte`                        |
| `ReplaceAllString`       | Replaces all matches with a string.                                  | `src, repl string` → `string`                        |
| `ReplaceAllFunc`         | Applies a function to all matches for replacement.                   | `src []byte, repl func([]byte) []byte` → `[]byte`    |
| `ReplaceAllStringFunc`   | Applies a string function to all matches.                            | `src string, repl func(string) string` → `string`    |
| `Split`                  | Splits input around all matches.                                     | `s string, n int` → `[]string`                       |
| `Longest`                | Sets the Regexp to prefer leftmost-longest matches (like POSIX).     | → `void`                                             |

## Syntax Highlights (RE2)

| Syntax         | Description                              |
|----------------|------------------------------------------|
| `.`            | Any character except newline             |
| `^`, `$`       | Start and end of a line                  |
| `*`, `+`, `?`  | Quantifiers: zero/more, one/more, zero/one |
| `{n}`, `{n,}`  | Exact and at least n repetitions         |
| `[...]`        | Character class                          |
| `( )`          | Grouping                                 |
| `|`            | Alternation (OR)                         |
| `\`            | Escape character                         |
| `(?P<name>...)`| Named capture group                      |

# sync Package

The `sync` package provides basic synchronisation primitives such as mutual exclusion locks. It is essential for managing shared resources across goroutines safely.

## Types and Methods

### Type: `Mutex`

A mutual exclusion lock.

| Method     | Description                                  | Returns     |
|------------|----------------------------------------------|-------------|
| `Lock`     | Locks the mutex. Blocks if already locked.   | –           |
| `Unlock`   | Unlocks the mutex.                           | –           |

### Type: `RWMutex`

A reader/writer mutual exclusion lock. Multiple readers can hold the lock, but only one writer at a time.

| Method        | Description                                         | Returns     |
|---------------|-----------------------------------------------------|-------------|
| `Lock`        | Locks the mutex for writing.                        | –           |
| `Unlock`      | Unlocks the writer lock.                            | –           |
| `RLock`       | Locks the mutex for reading.                        | –           |
| `RUnlock`     | Unlocks a reader lock.                              | –           |

### Type: `WaitGroup`

Waits for a collection of goroutines to finish.

| Method     | Description                                                  | Returns     |
|------------|--------------------------------------------------------------|-------------|
| `Add`      | Adds delta to the WaitGroup counter.                         | `delta int` |
| `Done`     | Decrements the counter by 1 (usually deferred in goroutines).| –           |
| `Wait`     | Blocks until the counter is zero.                            | –           |

### Type: `Once`

Ensures a function is only executed once.

| Method     | Description                                                  | Returns     |
|------------|--------------------------------------------------------------|-------------|
| `Do`       | Executes the function only once, even if called multiple times. | `f func()`  |

### Type: `Cond`

A condition variable, used with a `Locker` (e.g. `*Mutex`) to wait for or signal events.

| Method        | Description                                                  | Returns     |
|---------------|--------------------------------------------------------------|-------------|
| `Wait`        | Waits for a signal while releasing the lock.                | –           |
| `Signal`      | Wakes one goroutine waiting on the condition.              | –           |
| `Broadcast`   | Wakes all goroutines waiting on the condition.             | –           |
| `NewCond`     | Creates a new `*Cond` using a `Locker`.                    | `l Locker` → `*Cond` |

## Interfaces

| Interface | Description                           |
|-----------|---------------------------------------|
| `Locker`  | Any type that has `Lock` and `Unlock` methods (e.g. `*Mutex`, `*RWMutex`) |

# sync/atomic Package

The `sync/atomic` package provides low-level atomic memory primitives for synchronisation. These operations are useful for implementing lock-free data structures and thread-safe counters.

## Core Functions

### Integer Operations

| Function              | Description                                                       | Parameters / Returns                                |
|-----------------------|-------------------------------------------------------------------|-----------------------------------------------------|
| `AddInt32`            | Atomically adds delta to *int32 and returns the new value.        | `addr *int32, delta int32` → `int32`               |
| `AddInt64`            | Atomically adds delta to *int64.                                  | `addr *int64, delta int64` → `int64`               |
| `LoadInt32`           | Atomically loads *int32.                                          | `addr *int32` → `int32`                            |
| `LoadInt64`           | Atomically loads *int64.                                          | `addr *int64` → `int64`                            |
| `StoreInt32`          | Atomically stores val into *int32.                                | `addr *int32, val int32` → `void`                  |
| `StoreInt64`          | Atomically stores val into *int64.                                | `addr *int64, val int64` → `void`                  |
| `SwapInt32`           | Atomically swaps *int32 with new value.                           | `addr *int32, new int32` → `int32`                 |
| `SwapInt64`           | Atomically swaps *int64 with new value.                           | `addr *int64, new int64` → `int64`                 |
| `CompareAndSwapInt32` | Atomically compares and swaps if current value equals old.        | `addr *int32, old, new int32` → `bool`             |
| `CompareAndSwapInt64` | Same as above for `int64`.                                        | `addr *int64, old, new int64` → `bool`             |

### Unsigned and Pointer Operations

| Function                  | Description                                                         | Parameters / Returns                                |
|---------------------------|---------------------------------------------------------------------|-----------------------------------------------------|
| `AddUint32`, `AddUint64`  | Atomically adds delta to *uint32/uint64.                           | `addr *uintX, delta uintX` → `uintX`                |
| `LoadUint32`, `LoadUint64`| Loads the value atomically.                                        | `addr *uintX` → `uintX`                             |
| `StoreUint32`, `StoreUint64`| Stores the value atomically.                                    | `addr *uintX, val uintX`                            |
| `SwapUint32`, `SwapUint64`| Atomically swaps *uintX.                                           | `addr *uintX, new uintX` → `uintX`                  |
| `CompareAndSwapUint32`    | CAS operation for *uint32.                                         | `addr *uint32, old, new uint32` → `bool`            |
| `CompareAndSwapUint64`    | CAS operation for *uint64.                                         | `addr *uint64, old, new uint64` → `bool`            |
| `LoadPointer`             | Atomically loads *unsafe.Pointer.                                  | `addr *unsafe.Pointer` → `unsafe.Pointer`           |
| `StorePointer`            | Atomically stores value into *unsafe.Pointer.                      | `addr *unsafe.Pointer, val unsafe.Pointer`          |
| `SwapPointer`             | Atomically swaps the value.                                        | `addr *unsafe.Pointer, new unsafe.Pointer` → `unsafe.Pointer` |
| `CompareAndSwapPointer`   | CAS operation for pointers.                                        | `addr *unsafe.Pointer, old, new unsafe.Pointer` → `bool`       |

## Notes

- These operations are performed without locks and are safe for concurrent use.
- They ensure proper memory ordering on all supported Go architectures.
- For 64-bit atomics on 32-bit systems, variables must be 64-bit aligned.

# context Package

The `context` package defines the `Context` type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines.

## Core Functions

| Function             | Description                                                                 | Parameters / Returns                                                     |
|----------------------|-----------------------------------------------------------------------------|--------------------------------------------------------------------------|
| `Background`         | Returns an empty root context (never cancelled, no values, no deadline).   | → `Context`                                                              |
| `TODO`               | Returns a context intended for future use or placeholders.                 | → `Context`                                                              |
| `WithCancel`         | Returns a copy of the parent context and a `CancelFunc` to cancel it.      | `parent Context` → `(Context, CancelFunc)`                              |
| `WithDeadline`       | Returns a context with a deadline. Cancels automatically when the deadline is reached. | `parent Context, d time.Time` → `(Context, CancelFunc)`         |
| `WithTimeout`        | Returns a context with a timeout (shorthand for `WithDeadline`).           | `parent Context, timeout time.Duration` → `(Context, CancelFunc)`       |
| `WithValue`          | Returns a copy of parent context with a new key-value pair.                | `parent Context, key, val any` → `Context`                              |

## Type: `Context`

An immutable object for controlling the lifecycle of a request.

| Method          | Description                                                    | Returns                     |
|------------------|----------------------------------------------------------------|-----------------------------|
| `Deadline`       | Returns the time when the context will be cancelled (if set). | `(time.Time, bool)`         |
| `Done`           | Returns a channel that is closed when the context is cancelled. | `<-chan struct{}`          |
| `Err`            | Returns a non-nil error if context is cancelled or deadline exceeded. | `error`               |
| `Value`          | Retrieves a value from the context by key.                    | `interface{}`               |

## Type: `CancelFunc`

| Description                                  |
|----------------------------------------------|
| A function that cancels a context. Always call it to release resources when no longer needed. |

## Use Cases

- HTTP request cancellation.
- Deadline/timeouts for slow operations.
- Contextual data propagation (e.g. request IDs).
- Managing parent/child relationships in goroutines.

# log Package

The `log` package provides simple logging functions for writing to standard error or any custom `io.Writer`. It supports formatted output, prefixes, and flags for time and source file info.

## Core Functions

| Function         | Description                                                               | Parameters / Returns                                 |
|------------------|---------------------------------------------------------------------------|------------------------------------------------------|
| `Print`          | Logs a message using `fmt.Sprint`.                                        | `v ...interface{}` → `void`                          |
| `Printf`         | Logs a formatted message using `fmt.Sprintf`.                             | `format string, v ...interface{}` → `void`           |
| `Println`        | Logs a message with a newline using `fmt.Sprintln`.                       | `v ...interface{}` → `void`                          |
| `Fatal`          | Logs a message and calls `os.Exit(1)`.                                    | `v ...interface{}` → `void`                          |
| `Fatalf`         | Logs a formatted message and exits.                                       | `format string, v ...interface{}` → `void`           |
| `Fatalln`        | Logs a message with newline and exits.                                    | `v ...interface{}` → `void`                          |
| `Panic`          | Logs a message and panics.                                                | `v ...interface{}` → `void`                          |
| `Panicf`         | Logs a formatted message and panics.                                      | `format string, v ...interface{}` → `void`           |
| `Panicln`        | Logs a message with newline and panics.                                   | `v ...interface{}` → `void`                          |
| `SetOutput`      | Sets the destination for log output (default is `os.Stderr`).             | `w io.Writer`                                        |
| `SetFlags`       | Sets output control flags (e.g. timestamp, file info).                    | `flag int`                                           |
| `SetPrefix`      | Sets the output prefix for all log entries.                               | `prefix string`                                      |
| `Flags`          | Returns the current output flags.                                         | → `int`                                              |
| `Prefix`         | Returns the current output prefix.                                        | → `string`                                           |

## Constants (Flags)

| Constant           | Description                                            |
|--------------------|--------------------------------------------------------|
| `Ldate`            | Date in local time (2006/01/02)                        |
| `Ltime`            | Time in local time (15:04:05)                          |
| `Lmicroseconds`    | Microsecond resolution: 15:04:05.000000                |
| `Llongfile`        | Full file name and line number (e.g. /a/b/c/d.go:23)  |
| `Lshortfile`       | Final file name element and line number (e.g. d.go:23)|
| `LUTC`             | Use UTC rather than local time                        |
| `Lmsgprefix`       | Move the prefix before the log entry instead of after |
| `LstdFlags`        | Equivalent to `Ldate | Ltime`                          |

## Notes

- By default, logs are written to `os.Stderr`.
- `Fatal` and `Panic` variants exit or panic respectively after logging.
- Use `log.New()` to create custom loggers with their own output and formatting.

# flag Package

The `flag` package provides a command-line flag parsing utility. It supports defining and reading flags of various types from `os.Args`.

## Core Functions

| Function         | Description                                                                  | Parameters / Returns                                              |
|------------------|------------------------------------------------------------------------------|-------------------------------------------------------------------|
| `Bool`           | Defines a `bool` flag with specified name, default value, and usage.         | `name string, default bool, usage string` → `*bool`               |
| `Int`            | Defines an `int` flag.                                                       | `name string, default int, usage string` → `*int`                 |
| `Int64`          | Defines an `int64` flag.                                                     | `name string, default int64, usage string` → `*int64`             |
| `Uint`           | Defines a `uint` flag.                                                       | `name string, default uint, usage string` → `*uint`               |
| `Uint64`         | Defines a `uint64` flag.                                                     | `name string, default uint64, usage string` → `*uint64`           |
| `String`         | Defines a `string` flag.                                                     | `name string, default string, usage string` → `*string`           |
| `Float64`        | Defines a `float64` flag.                                                    | `name string, default float64, usage string` → `*float64`         |
| `Duration`       | Defines a `time.Duration` flag (e.g., "2h45m").                              | `name string, default time.Duration, usage string` → `*Duration`  |
| `Var`            | Defines a flag with a custom `Value` interface implementation.               | `value Value, name, usage string`                                 |
| `Parse`          | Parses the command-line flags from `os.Args`.                               | – → `void`                                                        |
| `Args`           | Returns non-flag command-line arguments.                                    | → `[]string`                                                      |
| `Arg`            | Returns the ith non-flag argument.                                          | `i int` → `string`                                                |
| `NArg`           | Returns the number of non-flag arguments.                                   | → `int`                                                           |
| `NFlag`          | Returns the number of flags that have been set.                             | → `int`                                                           |
| `Visit`          | Visits all flags that have been set.                                        | `fn func(*Flag)`                                                  |
| `VisitAll`       | Visits all defined flags, set or not.                                       | `fn func(*Flag)`                                                  |
| `Set`            | Sets the value of the named flag.                                           | `name string, value string` → `error`                             |
| `Lookup`         | Looks up a flag by name.                                                    | `name string` → `*Flag`                                           |
| `PrintDefaults`  | Prints default values of all flags.                                         | – → `void`                                                        |
| `Usage`          | A variable containing the function called when flag parsing fails.          | `var Usage func()` (can be overridden)                            |

## Types and Interfaces

### Type: `Flag`

| Field      | Description                           |
|------------|---------------------------------------|
| `Name`     | Name as it appears on command line    |
| `Usage`    | Help text                             |
| `Value`    | Value of the flag                     |
| `DefValue` | Default value (as string)             |

### Interface: `Value`

| Method           | Description                        |
|------------------|------------------------------------|
| `String()`       | Returns the flag’s value as string |
| `Set(string)`    | Sets the flag’s value              |

## Notes

- Flags must be defined before `flag.Parse()` is called.
- Custom flag types can implement the `flag.Value` interface.

# errors Package

The `errors` package provides simple error handling primitives. It includes functions for creating and inspecting errors, and supports error wrapping and unwrapping.

## Core Functions

| Function         | Description                                                             | Parameters / Returns                             |
|------------------|-------------------------------------------------------------------------|--------------------------------------------------|
| `New`            | Creates a new error with the specified message.                         | `text string` → `error`                          |
| `Unwrap`         | Returns the result of calling `Unwrap()` on the error, if supported.    | `err error` → `error`                            |
| `Is`             | Reports whether any error in the chain matches the target error.        | `err, target error` → `bool`                     |
| `As`             | Finds the first error in the chain that matches a specific type.        | `err error, target interface{}` → `bool`         |
| `Join`           | Combines multiple errors into a single `error`.                         | `errs ...error` → `error`                        |

## Usage and Behaviour

| Concept           | Description                                                                 |
|-------------------|-----------------------------------------------------------------------------|
| Error Wrapping    | Go 1.13+ allows errors to wrap others using `fmt.Errorf("%w", err)`         |
| `Unwrap` Support  | Custom error types can implement `Unwrap() error` for use with `errors.Is` and `errors.As` |
| Chain Traversal   | `errors.Is` and `errors.As` check recursively through wrapped errors        |

## Notes

- `errors.New` is for creating simple, static error messages.
- To include contextual information and preserve original errors, use wrapping via `fmt.Errorf`.
- `errors.Join` (Go 1.20+) allows you to aggregate multiple errors into one, which is useful for multi-part operations or validations.

# reflect Package

The `reflect` package implements run-time reflection, allowing a program to inspect and manipulate objects with arbitrary types.

## Core Functions

| Function            | Description                                                                 | Parameters / Returns                         |
|---------------------|-----------------------------------------------------------------------------|----------------------------------------------|
| `TypeOf`            | Returns the reflection `Type` of a value.                                   | `i interface{}` → `reflect.Type`             |
| `ValueOf`           | Returns the reflection `Value` of a value.                                  | `i interface{}` → `reflect.Value`            |
| `Zero`              | Returns a zero `Value` for the given type.                                  | `t reflect.Type` → `reflect.Value`           |
| `Indirect`          | Returns the value pointed to by a pointer.                                  | `v reflect.Value` → `reflect.Value`          |
| `DeepEqual`         | Reports whether two values are deeply equal.                                | `a, b interface{}` → `bool`                  |
| `MakeSlice`         | Creates a new slice of a specified type, length, and capacity.              | `t reflect.Type, len, cap int` → `reflect.Value` |
| `MakeMap`           | Creates a new map of the specified type.                                    | `t reflect.Type` → `reflect.Value`           |
| `MakeChan`          | Creates a new channel of the specified type and buffer size.                | `t reflect.Type, buffer int` → `reflect.Value` |
| `Append`            | Appends elements to a slice.                                                | `s reflect.Value, x ...reflect.Value` → `reflect.Value` |
| `New`               | Allocates a new value of the specified type.                                | `t reflect.Type` → `reflect.Value`           |

## Type: `reflect.Type`

Describes the type of a value.

| Method            | Description                                          | Returns            |
|-------------------|------------------------------------------------------|--------------------|
| `Kind`            | Returns the specific kind (e.g. reflect.Int, Slice) | `reflect.Kind`     |
| `Name`            | Returns the type's name                             | `string`           |
| `PkgPath`         | Returns the import path of the package              | `string`           |
| `Size`            | Returns the number of bytes occupied                | `uintptr`          |
| `NumField`        | For structs, returns number of fields               | `int`              |
| `Field`           | Returns a `StructField` by index                    | `i int` → `StructField` |
| `FieldByName`     | Returns a `StructField` by name                     | `name string` → `(StructField, bool)` |
| `NumMethod`       | Returns number of methods                           | `int`              |
| `Method`          | Returns method by index                             | `i int` → `Method` |
| `Implements`      | Checks if type implements interface type            | `u Type` → `bool`  |

## Type: `reflect.Value`

Represents the value of a Go object.

| Method              | Description                                     | Returns             |
|---------------------|-------------------------------------------------|---------------------|
| `Kind`              | Returns the kind of the value                   | `reflect.Kind`      |
| `Type`              | Returns the Type of the value                   | `reflect.Type`      |
| `IsValid`           | Checks if value is valid                        | `bool`              |
| `IsNil`             | Checks if value is nil (if applicable)          | `bool`              |
| `Elem`              | Gets the value pointed to                       | `reflect.Value`     |
| `CanSet`            | Returns true if the value can be changed        | `bool`              |
| `Set`               | Sets a new value                                | `v reflect.Value`   |
| `Interface`         | Returns the value as an interface{}             | `interface{}`       |
| `Convert`           | Converts to another type                        | `t reflect.Type` → `reflect.Value` |
| `Call`              | Invokes the function                            | `in []reflect.Value` → `[]reflect.Value` |
| `Len`               | Returns length (for arrays, slices, maps, strings) | `int`              |
| `Index`             | Gets element at index                           | `i int` → `reflect.Value` |

## Enums

### Kind (from `reflect.Kind`)

| Kind            | Description          |
|------------------|----------------------|
| `Invalid`        | No value             |
| `Bool`           | Boolean              |
| `Int`, `Int8`... | Integer types        |
| `Uint`, `Uint8`...| Unsigned integers   |
| `Float32`, `Float64` | Floating point    |
| `Complex64`, `Complex128` | Complex numbers |
| `Array`, `Chan`, `Func`, `Interface`, `Map`, `Ptr`, `Slice`, `Struct`, `String`, `UnsafePointer` | Composite types |

# testing Package

The `testing` package provides support for automated testing of Go packages. It defines the `*testing.T` type and related functionality used in test files.

## Core Concepts

- Test files must end with `_test.go`.
- Each test function must begin with `Test`, take a `*testing.T` parameter, and be exported.
- Benchmarks start with `Benchmark`, and examples with `Example`.

## Functions and Types

| Function / Type       | Description                                                                | Parameters / Returns                      |
|------------------------|----------------------------------------------------------------------------|-------------------------------------------|
| `T`                   | Type used to manage test state and report failures.                        | –                                         |
| `B`                   | Type used for writing benchmarks.                                          | –                                         |
| `M`                   | Used in `TestMain(m *testing.M)` to control test execution.                | –                                         |

### Methods of `*testing.T`

| Method               | Description                                                          | Returns     |
|----------------------|----------------------------------------------------------------------|-------------|
| `Error`              | Logs error message, marks test as failed, but continues.             | `v ...interface{}` |
| `Errorf`             | Like `Error`, but formatted.                                         | `format string, v ...interface{}` |
| `Fail`               | Marks the test as failed.                                            | –           |
| `FailNow`            | Marks the test as failed and stops execution immediately.            | –           |
| `Fatal`              | Logs message and calls `FailNow`.                                    | `v ...interface{}` |
| `Fatalf`             | Formatted fatal log.                                                 | `format string, v ...interface{}` |
| `Log`                | Logs output without failing.                                         | `v ...interface{}` |
| `Logf`               | Formatted log output.                                                | `format string, v ...interface{}` |
| `Name`               | Returns the name of the running test.                                | `string`    |
| `Skip`               | Skips the test.                                                      | `v ...interface{}` |
| `Skipf`              | Formatted skip log.                                                  | `format string, v ...interface{}` |
| `SkipNow`            | Marks test as skipped immediately.                                   | –           |
| `Skipped`            | Reports if the test was skipped.                                     | `bool`      |
| `Helper`             | Marks calling function as helper (adjusts line number in logs).       | –           |
| `Run`                | Runs a subtest.                                                      | `name string, f func(t *T)` → `bool` |

### Methods of `*testing.B`

| Method               | Description                                                          | Returns     |
|----------------------|----------------------------------------------------------------------|-------------|
| `StartTimer`         | Starts timing (if stopped).                                          | –           |
| `StopTimer`          | Stops timing.                                                        | –           |
| `ResetTimer`         | Clears elapsed time.                                                 | –           |
| `ReportAllocs`       | Enables allocation tracking.                                         | –           |
| `SetBytes`           | Reports processed bytes per operation.                              | `n int64`   |
| `Run`                | Runs a sub-benchmark.                                                | `name string, f func(b *B)` → `bool` |

## Environment

- `go test` compiles and runs tests.
- `TestMain(m *testing.M)` can be used to run setup/teardown logic before/after tests.
- Flags like `-v`, `-run`, and `-bench` control test execution.

# sort Package

The `sort` package provides primitives for sorting slices and user-defined collections.

## Core Functions

| Function              | Description                                                                   | Parameters / Returns                                  |
|-----------------------|-------------------------------------------------------------------------------|-------------------------------------------------------|
| `Ints`                | Sorts a slice of `int` in increasing order.                                  | `a []int` → `void`                                    |
| `Float64s`            | Sorts a slice of `float64` in increasing order (not safe for NaNs).          | `a []float64` → `void`                                |
| `Strings`             | Sorts a slice of `string` in increasing lexicographical order.               | `a []string` → `void`                                 |
| `IntsAreSorted`       | Reports whether a slice of `int` is sorted.                                  | `a []int` → `bool`                                    |
| `Float64sAreSorted`   | Reports whether a slice of `float64` is sorted.                              | `a []float64` → `bool`                                |
| `StringsAreSorted`    | Reports whether a slice of `string` is sorted.                               | `a []string` → `bool`                                 |
| `Search`              | Binary search for `x` in sorted slice `f(i)`.                                | `n int, f func(int) bool` → `int`                     |
| `SearchInts`          | Binary search for `x` in sorted `[]int`.                                     | `a []int, x int` → `int`                              |
| `SearchFloat64s`      | Binary search for `x` in sorted `[]float64`.                                 | `a []float64, x float64` → `int`                      |
| `SearchStrings`       | Binary search for `x` in sorted `[]string`.                                  | `a []string, x string` → `int`                        |

## Interface: `sort.Interface`

Types that implement this interface can be sorted using `sort.Sort`.

| Method        | Description                                                   |
|---------------|---------------------------------------------------------------|
| `Len()`       | Returns the number of elements                                |
| `Less(i, j)`  | Reports whether element i should sort before element j        |
| `Swap(i, j)`  | Swaps the elements at indexes i and j                         |

## Other Functions for Custom Sorting

| Function       | Description                                                                  | Parameters / Returns                           |
|----------------|------------------------------------------------------------------------------|------------------------------------------------|
| `Sort`         | Sorts the data using provided sort.Interface implementation.                | `data Interface` → `void`                      |
| `Stable`       | Like `Sort`, but keeps equal elements in their original order.              | `data Interface` → `void`                      |
| `Reverse`      | Returns a reverse order wrapper for `Interface`.                            | `data Interface` → `Interface`                 |
| `Slice`        | Sorts a slice using a comparison function. (Go 1.8+)                        | `x any, less func(i, j int) bool`              |
| `SliceStable`  | Like `Slice`, but keeps equal elements in original order.                   | `x any, less func(i, j int) bool`              |
| `SliceIsSorted`| Reports whether a slice is sorted using the provided less function.         | `x any, less func(i, j int) bool` → `bool`     |

# bytes Package

The `bytes` package implements functions for the manipulation of byte slices (`[]byte`). It mirrors many functions found in the `strings` package.

## Core Functions

| Function               | Description                                                               | Parameters / Returns                               |
|------------------------|---------------------------------------------------------------------------|----------------------------------------------------|
| `Compare`              | Lexicographically compares two byte slices.                              | `a, b []byte` → `int`                              |
| `Equal`                | Reports whether two slices are equal.                                    | `a, b []byte` → `bool`                             |
| `EqualFold`            | Reports whether two slices are equal under Unicode case-folding.         | `a, b []byte` → `bool`                             |
| `Contains`             | Checks whether `sub` is within `s`.                                      | `s, sub []byte` → `bool`                           |
| `ContainsAny`          | Checks whether any of the characters in `chars` are in `s`.              | `s []byte, chars string` → `bool`                  |
| `Count`                | Counts occurrences of `sep` in `s`.                                      | `s, sep []byte` → `int`                            |
| `HasPrefix`            | Checks whether `s` begins with `prefix`.                                 | `s, prefix []byte` → `bool`                        |
| `HasSuffix`            | Checks whether `s` ends with `suffix`.                                   | `s, suffix []byte` → `bool`                        |
| `Index`                | Returns the index of the first instance of `sep` in `s`.                 | `s, sep []byte` → `int`                            |
| `IndexAny`             | Returns index of any Unicode code point from `chars` in `s`.             | `s []byte, chars string` → `int`                   |
| `LastIndex`            | Returns the index of the last instance of `sep` in `s`.                  | `s, sep []byte` → `int`                            |
| `Repeat`               | Returns a new byte slice consisting of `count` copies of `b`.            | `b []byte, count int` → `[]byte`                   |
| `Replace`              | Replaces occurrences of `old` with `new` in `s`.                         | `s, old, new []byte, n int` → `[]byte`             |
| `Split`                | Splits `s` around each instance of `sep`.                                | `s, sep []byte` → `[][]byte`                       |
| `SplitAfter`           | Splits `s` after each instance of `sep`.                                 | `s, sep []byte` → `[][]byte`                       |
| `SplitN`               | Splits `s` at most `n` times around `sep`.                               | `s, sep []byte, n int` → `[][]byte`                |
| `SplitAfterN`          | Like `SplitN`, but after each instance of `sep`.                         | `s, sep []byte, n int` → `[][]byte`                |
| `ToLower`              | Returns a copy of `s` with all Unicode letters mapped to lower case.     | `s []byte` → `[]byte`                              |
| `ToUpper`              | Returns a copy of `s` with all Unicode letters mapped to upper case.     | `s []byte` → `[]byte`                              |
| `Trim`                 | Removes all leading and trailing Unicode code points in `cutset`.        | `s []byte, cutset string` → `[]byte`               |
| `TrimSpace`            | Removes leading and trailing white space.                                | `s []byte` → `[]byte`                              |
| `Fields`               | Splits `s` around each instance of white space.                          | `s []byte` → `[][]byte`                            |
| `Join`                 | Joins slices using `sep`.                                                 | `s [][]byte, sep []byte` → `[]byte`                |
| `Runes`                | Converts byte slice to slice of runes.                                   | `s []byte` → `[]rune`                              |

## Buffer Type

### Type: `Buffer`

A variable-sized buffer of bytes with `io.Reader`/`io.Writer` interface methods.

| Method         | Description                                             | Returns                         |
|----------------|---------------------------------------------------------|----------------------------------|
| `Write`        | Writes data to the buffer.                             | `p []byte` → `(n int, err error)`|
| `WriteString`  | Writes a string to the buffer.                         | `s string` → `(n int, err error)`|
| `Read`         | Reads data into `p`.                                   | `p []byte` → `(n int, err error)`|
| `Bytes`        | Returns the contents as a byte slice.                  | → `[]byte`                        |
| `String`       | Returns the contents as a string.                      | → `string`                        |
| `Reset`        | Clears the buffer.                                     | –                                |
| `Len`          | Returns the number of bytes in the buffer.             | → `int`                          |
| `Cap`          | Returns the capacity of the buffer.                    | → `int`                          |
| `Grow`         | Grows the buffer's capacity.                           | `n int` → `void`                 |
| `Truncate`     | Truncates the buffer to length `n`.                    | `n int` → `void`                 |
| `Next`         | Reads and advances the buffer by `n` bytes.            | `n int` → `[]byte`               |

# unicode Package

The `unicode` package provides data and functions to test and transform Unicode characters (runes). It supports categories, cases, and character properties.

## Core Functions

| Function         | Description                                                                 | Parameters / Returns                      |
|------------------|-----------------------------------------------------------------------------|-------------------------------------------|
| `Is`             | Reports whether the rune is in a given table (range table).                 | `rangeTab *RangeTable, r rune` → `bool`   |
| `IsControl`      | Reports whether the rune is a control character.                            | `r rune` → `bool`                         |
| `IsDigit`        | Reports whether the rune is a decimal digit.                                | `r rune` → `bool`                         |
| `IsGraphic`      | Reports whether the rune is defined as a graphic.                           | `r rune` → `bool`                         |
| `IsLetter`       | Reports whether the rune is a letter.                                       | `r rune` → `bool`                         |
| `IsLower`        | Reports whether the rune is a lowercase letter.                             | `r rune` → `bool`                         |
| `IsMark`         | Reports whether the rune is a mark character.                               | `r rune` → `bool`                         |
| `IsNumber`       | Reports whether the rune is a number.                                       | `r rune` → `bool`                         |
| `IsPrint`        | Reports whether the rune is printable.                                      | `r rune` → `bool`                         |
| `IsPunct`        | Reports whether the rune is a punctuation mark.                             | `r rune` → `bool`                         |
| `IsSpace`        | Reports whether the rune is a white space character.                        | `r rune` → `bool`                         |
| `IsSymbol`       | Reports whether the rune is a symbol.                                       | `r rune` → `bool`                         |
| `IsTitle`        | Reports whether the rune is a title-case letter.                            | `r rune` → `bool`                         |
| `IsUpper`        | Reports whether the rune is an uppercase letter.                            | `r rune` → `bool`                         |
| `To`             | Converts a rune to a different case using a case constant.                  | `case int, r rune` → `rune`               |
| `ToLower`        | Converts a rune to lower case.                                              | `r rune` → `rune`                         |
| `ToTitle`        | Converts a rune to title case.                                              | `r rune` → `rune`                         |
| `ToUpper`        | Converts a rune to upper case.                                              | `r rune` → `rune`                         |
| `SimpleFold`     | Returns the next equivalent rune in Unicode simple case folding.            | `r rune` → `rune`                         |

## Constants

| Constant         | Description                           |
|------------------|---------------------------------------|
| `UpperCase`      | Case conversion constant for `To`     |
| `LowerCase`      | Case conversion constant for `To`     |
| `TitleCase`      | Case conversion constant for `To`     |

## Types

### Type: `RangeTable`

| Description                          |
|--------------------------------------|
| Represents a set of Unicode code points (used with `Is`) |

## Predefined Tables

| Table Name       | Category Description             |
|------------------|----------------------------------|
| `unicode.Letter` | All letter characters            |
| `unicode.Digit`  | All decimal digit characters     |
| `unicode.Space`  | All space characters             |
| `unicode.Punct`  | All punctuation characters       |
| `unicode.Symbol` | All symbol characters            |

# unicode/utf8 Package

The `unicode/utf8` package provides functions for encoding and decoding UTF-8 byte sequences. It supports processing of Unicode code points in UTF-8 encoded data.

## Constants

| Constant          | Description                                              |
|-------------------|----------------------------------------------------------|
| `RuneError`       | The Unicode replacement character (U+FFFD) used for errors. |
| `RuneSelf`        | Runes below this can be represented as a single byte.    |
| `MaxRune`         | Maximum valid Unicode code point.                        |
| `UTFMax`          | Maximum number of bytes a rune may occupy in UTF-8.      |

## Core Functions

| Function              | Description                                                                  | Parameters / Returns                                 |
|-----------------------|------------------------------------------------------------------------------|------------------------------------------------------|
| `DecodeRune`          | Decodes the first rune in a byte slice.                                     | `p []byte` → `(r rune, size int)`                   |
| `DecodeLastRune`      | Decodes the last rune in a byte slice.                                      | `p []byte` → `(r rune, size int)`                   |
| `EncodeRune`          | Encodes a rune into UTF-8.                                                   | `p []byte, r rune` → `int`                          |
| `FullRune`            | Reports whether the byte slice contains a full UTF-8 encoding.              | `p []byte` → `bool`                                 |
| `FullRuneInString`    | Like `FullRune`, but for strings.                                            | `s string` → `bool`                                 |
| `RuneCount`           | Returns the number of runes in the byte slice.                              | `p []byte` → `int`                                  |
| `RuneCountInString`   | Returns the number of runes in the string.                                  | `s string` → `int`                                  |
| `RuneLen`             | Returns the number of bytes required to encode the rune.                    | `r rune` → `int`                                    |
| `Valid`               | Reports whether the byte slice is valid UTF-8.                              | `p []byte` → `bool`                                 |
| `ValidString`         | Reports whether the string is valid UTF-8.                                  | `s string` → `bool`                                 |

## Notes

- UTF-8 encoding is variable-width (1–4 bytes per rune).
- Most ASCII characters are 1 byte.
- Use `RuneError` to replace invalid encodings.

# runtime Package

The `runtime` package contains operations that interact with Go’s runtime system, such as managing goroutines, memory statistics, and system-level behaviour.

## Core Functions

| Function               | Description                                                                    | Parameters / Returns                                 |
|------------------------|--------------------------------------------------------------------------------|------------------------------------------------------|
| `Goexit`               | Terminates the current goroutine without affecting others.                    | – → `void`                                           |
| `GOMAXPROCS`           | Sets or returns the maximum number of CPUs used simultaneously.               | `n int` → `int`                                      |
| `NumCPU`               | Returns the number of logical CPUs usable.                                     | – → `int`                                            |
| `NumGoroutine`         | Returns the number of currently active goroutines.                             | – → `int`                                            |
| `GoroutineProfile`     | Writes a snapshot of the stack traces for all goroutines.                      | `[]StackRecord` → `int`                              |
| `Stack`                | Copies the current goroutine’s stack trace into `buf`.                         | `buf []byte, all bool` → `int`                       |
| `Caller`               | Returns file and line number information for a function invocation.            | `skip int` → `(pc uintptr, file string, line int, ok bool)` |
| `Callers`              | Fills a slice with return PCs of function invocations on the calling goroutine’s stack. | `skip int, pc []uintptr` → `int`         |
| `ReadMemStats`         | Populates a `MemStats` struct with memory allocator statistics.                | `m *MemStats` → `void`                               |
| `MemProfileRate`       | Memory profiling sampling rate (set manually).                                 | `int`                                                |
| `GC`                   | Forces a garbage collection.                                                   | – → `void`                                           |
| `SetFinalizer`         | Sets a finaliser function to run when an object becomes unreachable.           | `obj interface{}, finalizer interface{}`             |
| `KeepAlive`            | Ensures that an object is not garbage-collected before the point of call.     | `interface{}` → `void`                               |
| `Breakpoint`           | Invokes a breakpoint trap for debuggers.                                       | – → `void`                                           |

## Types

### Type: `MemStats`

Contains memory allocation statistics collected by the runtime.

| Field             | Description                                  |
|-------------------|----------------------------------------------|
| `Alloc`           | Bytes of allocated heap objects              |
| `TotalAlloc`      | Total bytes allocated, even if freed         |
| `Sys`             | Total bytes obtained from the system         |
| `Mallocs`         | Number of heap objects allocated             |
| `Frees`           | Number of heap objects freed                 |
| ...               | Many other low-level runtime memory stats    |

## Notes

- Use `runtime` for debugging, profiling, and performance tuning.
- Not typically needed in application-level code.

# runtime/debug Package

The `runtime/debug` package provides functions for debugging Go programs. It includes facilities for printing stack traces, controlling garbage collection, and reading build information.

## Core Functions

| Function             | Description                                                                 | Parameters / Returns                               |
|----------------------|-----------------------------------------------------------------------------|----------------------------------------------------|
| `PrintStack`         | Prints the stack trace of the current goroutine to standard error.          | – → `void`                                         |
| `FreeOSMemory`       | Forces the garbage collector to return memory to the OS immediately.        | – → `void`                                         |
| `SetGCPercent`       | Sets the garbage collection target percentage and returns the previous value. | `percent int` → `int`                             |
| `SetMaxStack`        | Sets the maximum stack size for new goroutines (not recommended).           | `bytes int` → `void`                               |
| `SetMaxThreads`      | Sets the maximum number of OS threads the Go runtime can use.               | `threads int` → `int`                              |
| `ReadBuildInfo`      | Reads build information embedded in the binary.                            | – → `*BuildInfo, bool`                             |

## Type: `BuildInfo`

Holds build-related metadata for the binary, if available.

| Field        | Description                         |
|--------------|-------------------------------------|
| `Path`       | Module path of the main module      |
| `Main`       | Module info for the main module     |
| `Deps`       | Direct module dependencies          |
| `Settings`   | Key-value build settings (e.g., `-ldflags`) |

## Notes

- `PrintStack` is helpful for logging during panic or debugging deadlocks.
- `ReadBuildInfo` is available for Go 1.12+ binaries built with module support.
- These tools are useful during development and performance troubleshooting but should be avoided in production paths unless necessary.

Let me know when you're ready to continue with `encoding/base64`, or if you'd like to cover other packages from your initial list like `context`, `io/ioutil`, etc.

# encoding/base64 Package

The `encoding/base64` package provides functions for encoding and decoding data using the standard and URL-compatible Base64 encodings, as defined in RFC 4648.

## Core Types and Functions

### Type: `Encoding`

Represents a Base64 encoding scheme (standard or URL-compatible).

| Method                | Description                                                        | Parameters / Returns                                 |
|------------------------|--------------------------------------------------------------------|------------------------------------------------------|
| `Encode`              | Encodes `src` into Base64 and writes to `dst`.                     | `dst, src []byte` → `void`                          |
| `EncodeToString`      | Encodes `src` and returns the Base64 result as a string.           | `src []byte` → `string`                             |
| `EncodedLen`          | Returns the number of bytes needed to encode `n` bytes of data.    | `n int` → `int`                                     |
| `Decode`              | Decodes Base64 `src` into `dst`. Returns number of bytes written.  | `dst, src []byte` → `(int, error)`                 |
| `DecodeString`        | Decodes a Base64-encoded string.                                   | `s string` → `[]byte, error`                        |
| `DecodedLen`          | Returns the max length of the output when decoding `n` bytes.      | `n int` → `int`                                     |
| `Strict`              | Returns an Encoding that requires padded input.                    | – → `*Encoding`                                     |
| `WithPadding`         | Returns a new Encoding with the specified padding character.       | `pad rune` → `*Encoding`                            |
| `NewEncoder`          | Returns a new Base64 stream encoder that writes to `io.Writer`.    | `enc *Encoding, w io.Writer` → `io.WriteCloser`     |
| `NewDecoder`          | Returns a new Base64 stream decoder that reads from `io.Reader`.   | `enc *Encoding, r io.Reader` → `io.Reader`          |

## Predefined Encodings

| Variable             | Description                                                  |
|----------------------|--------------------------------------------------------------|
| `StdEncoding`        | Standard Base64 encoding with `+` and `/`, padded with `=`   |
| `URLEncoding`        | URL-safe Base64 with `-` and `_`, padded with `=`            |
| `RawStdEncoding`     | Standard Base64 without padding                              |
| `RawURLEncoding`     | URL-safe Base64 without padding                              |

## Notes

- Base64 is commonly used for encoding binary data as ASCII, particularly in web and email contexts.
- `NewEncoder` and `NewDecoder` are useful for streaming large data.
- Use `Raw*` variants when padding is not required (e.g. JWT headers, URL tokens).














# Go Language Syntax & Features Reference


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










# Go Code Snippets

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

# Go UTIL Functions

- [Files](#files)
- [Base64](#base64)
- [MQTT](#mqtt)
- [HTTP](#http)
- [JSON](#json)
- [Bash](#bash)
- [Mutex](#mutex)
- [Logging](#logging)
  
## Files

`utils/file.go`

## Base64

`utils/base64.go`

```go
package utils

import (
    "encoding/base64"
    "encoding/json"
    "errors"
)

// EncodeStringToBase64 returns the Base64-encoded representation of a string.
func EncodeStringToBase64(input string) string {
    return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeBase64ToString decodes a Base64-encoded string into a plain string.
func DecodeBase64ToString(encoded string) (string, error) {
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return "", err
    }
    return string(decodedBytes), nil
}

// EncodeJSONToBase64 takes a Go map or struct, marshals it to JSON, and encodes it in Base64.
func EncodeJSONToBase64(data interface{}) (string, error) {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

// DecodeBase64JSONToMap decodes a Base64 string assumed to contain JSON and parses it into a map.
func DecodeBase64JSONToMap(encoded string) (map[string]interface{}, error) {
    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return nil, err
    }

    var result map[string]interface{}
    if err := json.Unmarshal(decodedBytes, &result); err != nil {
        return nil, err
    }

    return result, nil
}

// DecodeBase64JSONToStruct decodes a Base64 JSON string into a given struct pointer.
func DecodeBase64JSONToStruct(encoded string, target interface{}) error {
    if target == nil {
        return errors.New("target must be a non-nil pointer to a struct")
    }

    decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        return err
    }

    return json.Unmarshal(decodedBytes, target)
}

```

## MQTT

`utils/mqtt.go`

```go
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

```

## HTTP

`utils/http.go`

```go
package utils

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
)

// HTTPClient provides a shared client with timeout settings.
var HTTPClient = &http.Client{
    Timeout: 10 * time.Second,
}

// HTTPRequest performs a generic HTTP request with method, headers, and optional body.
func HTTPRequest(method, url string, headers map[string]string, body []byte) ([]byte, int, error) {
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, 0, err
    }

    for key, val := range headers {
        req.Header.Set(key, val)
    }

    resp, err := HTTPClient.Do(req)
    if err != nil {
        return nil, 0, err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, resp.StatusCode, err
    }

    if resp.StatusCode >= 400 {
        return responseBody, resp.StatusCode, errors.New(fmt.Sprintf("HTTP error %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode)))
    }

    return responseBody, resp.StatusCode, nil
}

// HTTPGet performs an HTTP GET request.
func HTTPGet(url string, headers map[string]string) ([]byte, int, error) {
    return HTTPRequest(http.MethodGet, url, headers, nil)
}

// HTTPPost performs an HTTP POST request with a JSON body.
func HTTPPost(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPost, url, headers, body)
}

// HTTPPut performs an HTTP PUT request with a JSON body.
func HTTPPut(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPut, url, headers, body)
}

// HTTPPatch performs an HTTP PATCH request with a JSON body.
func HTTPPatch(url string, headers map[string]string, body []byte) ([]byte, int, error) {
    return HTTPRequest(http.MethodPatch, url, headers, body)
}

// HTTPDelete performs an HTTP DELETE request.
func HTTPDelete(url string, headers map[string]string) ([]byte, int, error) {
    return HTTPRequest(http.MethodDelete, url, headers, nil)
}

```

## JSON

`utils/json.go`

```go
package utils

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ToJSON converts any Go value (struct, map, slice) into a JSON-formatted byte slice.
func ToJSON(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

// ToJSONString converts any Go value into a formatted JSON string.
func ToJSONString(data interface{}) (string, error) {
    bytes, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// FromJSON parses a JSON byte slice into the provided destination (must be a pointer).
func FromJSON(jsonBytes []byte, target interface{}) error {
    if target == nil {
        return errors.New("target must be a non-nil pointer")
    }
    return json.Unmarshal(jsonBytes, target)
}

// FromJSONString parses a JSON string into the provided destination (must be a pointer).
func FromJSONString(jsonStr string, target interface{}) error {
    return FromJSON([]byte(jsonStr), target)
}

// FromJSONToMap decodes JSON bytes into a generic map[string]interface{}.
func FromJSONToMap(jsonBytes []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := json.Unmarshal(jsonBytes, &result)
    return result, err
}

// SafeGetString retrieves a string from a decoded JSON map.
func SafeGetString(data map[string]interface{}, key string) (string, bool) {
    val, ok := data[key]
    str, valid := val.(string)
    return str, ok && valid
}

// SafeGetNumber retrieves a float64 from a decoded JSON map.
func SafeGetNumber(data map[string]interface{}, key string) (float64, bool) {
    val, ok := data[key]
    num, valid := val.(float64)
    return num, ok && valid
}

// SafeGetBool retrieves a boolean from a decoded JSON map.
func SafeGetBool(data map[string]interface{}, key string) (bool, bool) {
    val, ok := data[key]
    b, valid := val.(bool)
    return b, ok && valid
}

// FlattenJSON extracts all keys and their values recursively from a nested map.
// Returns a slice of key-value pairs as [][2]string.
func FlattenJSON(data map[string]interface{}) [][2]string {
    var result [][2]string
    walkJSON(data, "", &result)
    return result
}

// walkJSON is a recursive helper for FlattenJSON.
func walkJSON(data map[string]interface{}, prefix string, acc *[][2]string) {
    for k, v := range data {
        fullKey := k
        if prefix != "" {
            fullKey = prefix + "." + k
        }

        switch val := v.(type) {
        case map[string]interface{}:
            walkJSON(val, fullKey, acc)
        default:
            *acc = append(*acc, [2]string{fullKey, fmt.Sprintf("%v", val)})
        }
    }
}

```

## Bash

`utils/bash.go`

```go
package utils

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// RunCommand executes an inline Bash command string and returns combined stdout/stderr output.
func RunCommand(command string) (string, error) {
    cmd := exec.Command("bash", "-c", command)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunScript executes a Bash script file located at scriptPath with optional args.
// Returns combined stdout/stderr output.
func RunScript(scriptPath string, args ...string) (string, error) {
    cmd := exec.Command("bash", append([]string{scriptPath}, args...)...)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunCommandWithEnv runs a Bash command with additional environment variables.
func RunCommandWithEnv(command string, envVars map[string]string) (string, error) {
    cmd := exec.Command("bash", "-c", command)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    env := os.Environ()
    for k, v := range envVars {
        env = append(env, fmt.Sprintf("%s=%s", k, v))
    }
    cmd.Env = env

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunAndCapture splits output and error separately.
func RunAndCapture(command string) (stdout string, stderr string, err error) {
    cmd := exec.Command("bash", "-c", command)
    var outBuf, errBuf bytes.Buffer
    cmd.Stdout = &outBuf
    cmd.Stderr = &errBuf

    err = cmd.Run()
    return strings.TrimSpace(outBuf.String()), strings.TrimSpace(errBuf.String()), err
}

```

## Mutex

`utils/mutex.go`

```go
package utils

import (
    "errors"
    "sync"
)

// SafeQueue is a thread-safe FIFO queue backed by a slice and protected by a mutex.
type SafeQueue struct {
    mu    sync.Mutex
    items []interface{}
}

// NewSafeQueue creates and returns a new SafeQueue.
func NewSafeQueue() *SafeQueue {
    return &SafeQueue{
        items: make([]interface{}, 0),
    }
}

// Enqueue adds an item to the end of the queue.
func (q *SafeQueue) Enqueue(item interface{}) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
// Returns an error if the queue is empty.
func (q *SafeQueue) Dequeue() (interface{}, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.items) == 0 {
        return nil, errors.New("queue is empty")
    }

    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

// Peek returns the first item in the queue without removing it.
func (q *SafeQueue) Peek() (interface{}, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.items) == 0 {
        return nil, errors.New("queue is empty")
    }

    return q.items[0], nil
}

// Length returns the number of items in the queue.
func (q *SafeQueue) Length() int {
    q.mu.Lock()
    defer q.mu.Unlock()
    return len(q.items)
}

// Clear removes all items from the queue.
func (q *SafeQueue) Clear() {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = []interface{}{}
}

// Values returns a copy of the queue's items.
func (q *SafeQueue) Values() []interface{} {
    q.mu.Lock()
    defer q.mu.Unlock()
    return append([]interface{}(nil), q.items...)
}

```

## Logging

`utils/logging.go`

```go
package utils

import (
    "fmt"
    "log"
    "os"
    "time"
)

// LogLevel defines supported log levels.
type LogLevel string

const (
    LogInfo    LogLevel = "INFO"
    LogWarn    LogLevel = "WARN"
    LogError   LogLevel = "ERROR"
    LogDebug   LogLevel = "DEBUG"
    LogCrit    LogLevel = "CRITICAL"
)

// Logger provides namespaced logging for systemd-compatible environments.
type Logger struct {
    namespace string
    subject   string
    logger    *log.Logger
}

// NewLogger returns a new logger for a given namespace and subject.
// These appear as LOG_SUBJECT and LOG_NAMESPACE in journald.
func NewLogger(namespace, subject string) *Logger {
    prefix := fmt.Sprintf("LOG_NAMESPACE=%s LOG_SUBJECT=%s", namespace, subject)
    return &Logger{
        namespace: namespace,
        subject:   subject,
        logger:    log.New(os.Stdout, prefix+" ", 0),
    }
}

// logf outputs a log message with timestamp and level, systemd-readable.
func (l *Logger) logf(level LogLevel, format string, args ...interface{}) {
    timestamp := time.Now().Format(time.RFC3339)
    msg := fmt.Sprintf(format, args...)
    logLine := fmt.Sprintf("LEVEL=%s TIMESTAMP=%s MESSAGE=%q", level, timestamp, msg)
    l.logger.Println(logLine)
}

func (l *Logger) Infof(format string, args ...interface{}) {
    l.logf(LogInfo, format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
    l.logf(LogDebug, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
    l.logf(LogWarn, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
    l.logf(LogError, format, args...)
}

func (l *Logger) Criticalf(format string, args ...interface{}) {
    l.logf(LogCrit, format, args...)
}

```

```go
package utils

import (
    "bufio"
    "errors"
    "io/fs"
    "os"
    "path/filepath"
)

// ReadFile reads the contents of a file and returns it as a string.
func ReadFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// ReadLines reads a file line-by-line and returns a slice of strings.
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil
}

// WriteFile writes the given content to a file, overwriting if it exists.
func WriteFile(path string, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

// AppendToFile appends content to the end of a file or creates it if it doesn't exist.
func AppendToFile(path string, content string) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(content)
    return err
}

// FileExists checks if a file or directory exists at the given path.
func FileExists(path string) bool {
    _, err := os.Stat(path)
    return !errors.Is(err, os.ErrNotExist)
}

// ListFiles returns a list of file names (not directories) in the given folder.
func ListFiles(dir string) ([]string, error) {
    var files []string

    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    for _, entry := range entries {
        if !entry.IsDir() {
            files = append(files, entry.Name())
        }
    }

    return files, nil
}

// FindFilesByExtension recursively finds files with a given extension.
func FindFilesByExtension(root, ext string) ([]string, error) {
    var matched []string

    err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && filepath.Ext(path) == ext {
            matched = append(matched, path)
        }
        return nil
    })

    return matched, err
}
```
