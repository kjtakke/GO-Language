# Build Script

This Bash script is designed to build, detect, and run the GO application across multiple platforms and architectures. It supports a wide range of operating systems and CPU architectures, making it highly versatile for cross-platform development and deployment.

## Table of Contents

- [Usage](#usage)
- [Supported Targets](#supported-targets)
- [Build Process](#build-process)
- [Target Detection](#target-detection)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Usage

The script accepts the following command-line arguments:

- `--build`: Build all target binaries to the `./bin/` directory.
- `--which`: Display the appropriate binary for the current host.
- `--run`: Detect and execute the matching binary.
- `--help` or `-h`: Show the help message.

You can use these options in any combination. If no arguments are provided, the script will show the help message by default.

## Supported Targets

The script supports the following target combinations:

- linux/amd64
- linux/386
- linux/arm64
- linux/arm
- linux/ppc64le
- linux/s390x
- linux/mipsle
- linux/mips
- linux/mips64le
- linux/mips64
- linux/riscv64
- linux/loong64
- windows/amd64
- windows/386
- windows/arm64

## Build Process

When the `--build` flag is used, the script will:

1. Create a `bin` directory if it doesn't exist.
2. Iterate through the supported targets.
3. For each target, it sets the appropriate `GOOS` and `GOARCH` environment variables.
4. It then uses `go build` to compile the application for the specified target, saving the output to the `bin` directory.

The naming convention for the output files is `readyfleet-tac-{os}-{arch}`, with `.exe` appended for Windows targets.

## Target Detection

The script includes logic to detect the current system's operating system and architecture. It uses this information to:

- Determine the appropriate target when using the `--which` or `--run` flags.
- Check if a matching binary exists in the `bin` directory.

The detection process uses `uname -s` for the OS and `uname -m` for the architecture, mapping these to the corresponding Go environment variables.

## Examples

### Building for all targets

```bash
./build.sh --build
```

This will compile the application for all supported targets and save the binaries in the `bin` directory.

### Finding the appropriate binary for the current system

```bash
./build.sh --which
```

This will output the name of the binary that matches the current system's OS and architecture.

### Building and running the application

```bash
./build.sh --build --run
```

This will first build all targets, then detect and run the appropriate binary for the current system.

### Running without building

```bash
./build.sh --run
```

This will attempt to detect and run the appropriate binary, assuming it has already been built.

## Contributing

Contributions to this script are welcome. If you'd like to add support for new targets or improve existing functionality, please:

1. Fork the repository.
2. Make your changes.
3. Submit a pull request with a detailed description of your changes.

## License

This script is released under the MIT License. See the LICENSE file for more details.

