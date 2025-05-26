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












