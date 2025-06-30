#!/usr/bin/env bash
set -e

# -----------------------------------------------
# Target matrix (with inline comments for clarity)
# -----------------------------------------------
targets=(
  "linux/amd64"         # Intel/AMD 64-bit
  "linux/386"           # Legacy 32-bit x86
  "linux/arm64"         # ARMv8+, used in SBCs & servers (Pi 4/5, Jetson, AWS Graviton)
  "linux/arm"           # ARMv6/v7 (Pi Zero, Pi 3, older embedded systems)
  "linux/ppc64le"       # IBM POWER8/9/10 (LE) - used in enterprise servers
  "linux/s390x"         # IBM z Systems (mainframe)
  "linux/mipsle"        # Little-endian MIPS 32-bit (used in routers, OpenWRT)
  "linux/mips"          # Big-endian MIPS 32-bit
  "linux/mips64le"      # Little-endian MIPS 64-bit
  "linux/mips64"        # Big-endian MIPS 64-bit
  "linux/riscv64"       # RISC-V 64-bit boards (HiFive, VisionFive, etc.)
  "linux/loong64"       # LoongArch64 (mainland China use)
  "windows/amd64"       # Standard 64-bit Windows
  "windows/386"         # 32-bit Windows (legacy)
  "windows/arm64"       # Windows ARM64 (Surface Pro X, some Dev Boards)
)

# -----------------------------------------------
# Argument flags
# -----------------------------------------------
DO_BUILD=false
DO_RUN=false
DO_WHICH=false

# Show help if no arguments were passed
if [[ $# -eq 0 ]]; then
  set -- --help
fi

# Parse arguments in any order
for arg in "$@"; do
  case "$arg" in
    --build) DO_BUILD=true ;;
    --which) DO_WHICH=true ;;
    --run) DO_RUN=true ;;
    --help|-h|"")
      echo "Usage: ./build.sh [--build] [--which] [--run] [--help]"
      echo
      echo "  --build     Build all target binaries to ./bin/"
      echo "  --which     Display the appropriate binary for this host"
      echo "  --run       Detect and execute the matching binary"
      echo "  --help, -h  Show this help message"
      exit 0
      ;;
    *)
      echo "❌ Unknown option: $arg"
      exit 1
      ;;
  esac
done

# -----------------------------------------------
# Build logic
# -----------------------------------------------
if $DO_BUILD; then
  mkdir -p bin
  for target in "${targets[@]}"; do
    os="${target%/*}"
    arch="${target#*/}"
    output="bin/readyfleet-tac-${os}-${arch}"
    [[ "$os" == "windows" ]] && output="${output}.exe"
    echo "🔧 Building: $target → $output"
    GOOS=$os GOARCH=$arch go build -o "$output" .
  done
fi

# -----------------------------------------------
# Target detection logic
# -----------------------------------------------
detect_target() {
  case "$(uname -s)" in
    Linux)   GOOS="linux" ;;
    Darwin)  GOOS="darwin" ;;
    FreeBSD) GOOS="freebsd" ;;
    OpenBSD) GOOS="openbsd" ;;
    NetBSD)  GOOS="netbsd" ;;
    MINGW*|MSYS*|CYGWIN*|Windows_NT) GOOS="windows" ;;
    *) echo "❌ Unsupported OS: $(uname -s)"; exit 1 ;;
  esac

  case "$(uname -m)" in
    x86_64)      GOARCH="amd64" ;;
    i386|i686)   GOARCH="386" ;;
    armv7l)      GOARCH="arm" ;;
    aarch64)     GOARCH="arm64" ;;
    ppc64le)     GOARCH="ppc64le" ;;
    s390x)       GOARCH="s390x" ;;
    mips)        GOARCH="mips" ;;
    mipsel)      GOARCH="mipsle" ;;
    mips64)      GOARCH="mips64" ;;
    mips64el)    GOARCH="mips64le" ;;
    riscv64)     GOARCH="riscv64" ;;
    loongarch64) GOARCH="loong64" ;;
    *) echo "❌ Unsupported architecture: $(uname -m)"; exit 1 ;;
  esac

  TARGET="${GOOS}/${GOARCH}"
}

# Scan ./bin for built targets
get_built_targets() {
  BUILT_TARGETS=()
  while IFS= read -r file; do
    name=$(basename "$file")
    suffix="${name#readyfleet-tac-}"
    suffix="${suffix%.exe}" # strip .exe
    BUILT_TARGETS+=("${suffix//-//}") # dash-to-slash for GOOS/GOARCH
  done < <(find bin -type f -name 'readyfleet-tac-*')
}

# ----------------------------------------
# BUILD
# ----------------------------------------
if $DO_BUILD; then
  echo "🔧 Building binaries..."
  mkdir -p bin
  for target in \
    linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x \
    linux/mips linux/mipsle linux/mips64 linux/mips64le linux/riscv64 linux/loong64 \
    windows/amd64 windows/386 windows/arm64; do
    os="${target%/*}"
    arch="${target#*/}"
    outfile="bin/readyfleet-tac-${os}-${arch}"
    [[ "$os" == "windows" ]] && outfile="${outfile}.exe"
    echo "→ $target → $outfile"
    GOOS=$os GOARCH=$arch go build -o "$outfile" .
  done
fi

# ----------------------------------------
# WHICH / RUN
# ----------------------------------------
if $DO_WHICH || $DO_RUN; then
  detect_target
  get_built_targets

  if printf '%s\n' "${BUILT_TARGETS[@]}" | grep -qx "$TARGET"; then
    BIN="bin/readyfleet-tac-${GOOS}-${GOARCH}"
    [[ "$GOOS" == "windows" ]] && BIN="${BIN}.exe"

    echo "✅ Detected: $TARGET"
    echo "➡️  Binary: $BIN"
    if $DO_RUN; then
      echo "🚀 Running..."
      exec "$BIN"
    fi
  else
    echo "❌ No built binary found for: $TARGET"
    echo "   (Check that you have built it via --build)"
    exit 1
  fi
fi
