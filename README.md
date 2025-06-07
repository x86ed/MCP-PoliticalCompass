# MCP-PoliticalCompass

A Model Context Protocol (MCP) server that provides an interactive Political Compass quiz for baselining the political beliefs of AI agents and users. This server implements the complete Political Compass questionnaire with 62 questions across economic and social dimensions.

## Overview

The Political Compass is a multi-axis political model that maps political ideologies across two dimensions:

- **Economic Axis**: Left (collectivist) to Right (individualist)
- **Social Axis**: Libertarian (freedom-focused) to Authoritarian (order-focused)

This MCP server allows AI agents to take the complete Political Compass quiz and receive a detailed analysis of their political positioning across four quadrants:

- **Libertarian Left** (Left-wing, freedom-oriented)
- **Libertarian Right** (Right-wing, freedom-oriented)
- **Authoritarian Left** (Left-wing, order-oriented)
- **Authoritarian Right** (Right-wing, order-oriented)

## Features

### Tools Available

- **`political_compass`**: Interactive quiz tool that presents randomized political questions
- **`reset_quiz`**: Resets quiz progress to start fresh
- **`quiz_status`**: Shows current quiz progress and statistics

### Quiz Capabilities

- **62 authentic questions** from the Political Compass dataset
- **Randomized question order** for each quiz session
- **Real-time progress tracking** with current scores and status display
- **Response distribution analytics** showing breakdown by response type
- **Authentic scoring algorithm** that matches the original Political Compass methodology
- **Detailed final analysis** with quadrant placement and scores
- **Interactive SVG visualization** of results on political compass grid

### Visualization Features

- **Color-coded quadrants**: Each political quadrant has a distinct background color
- **Position marker**: Red dot shows exact location on the political compass
- **Grid lines and labels**: Clear axis markings and quadrant labels
- **Coordinate display**: Shows precise numerical position (Economic, Social)
- **Professional styling**: Clean, readable design suitable for analysis

### Example Interactions

**Starting the quiz:**

```yaml
Tool: political_compass
Response: ""
```

**Answering questions:**

```yaml
Tool: political_compass
Response: "Strongly Agree"
```

**Checking quiz progress:**

```yaml
Tool: quiz_status
```

**Quiz completion provides:**

- Final Economic Score (e.g., -1.23)
- Final Social Score (e.g., 2.45)
- Political Quadrant placement
- **Interactive SVG visualization** showing position on political compass grid
- Total questions answered

## Installation & Setup

### Option 1: Download Pre-built Binaries (Recommended)

The easiest way to get started is to download a pre-built binary for your platform:

**Latest Release**: [Download from GitHub Releases](https://github.com/x86ed/MCP-PoliticalCompass/releases/latest)

**Supported Platforms**:

- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)  
- **Windows**: amd64
- **FreeBSD**: amd64, arm64

**Installation Steps**:

1. Download the appropriate archive for your platform
2. Extract the binary:

   ```bash
   # For Linux/macOS
   tar -xzf mcp-political-compass-*-your-platform.tar.gz
   
   # For Windows
   unzip mcp-political-compass-*-windows-amd64.zip
   ```

3. Make executable (Linux/macOS only):

   ```bash
   chmod +x mcp-political-compass-*
   ```

4. Run the server:

   ```bash
   ./mcp-political-compass-*
   ```

**Verifying Downloads**:

Each release includes SHA256 checksums. To verify your download:

```bash
# Download checksums.txt from the release
curl -L -o checksums.txt https://github.com/x86ed/MCP-PoliticalCompass/releases/latest/download/checksums.txt

# Verify your downloaded file (example for Linux amd64)
sha256sum -c checksums.txt --ignore-missing
```

The server will start and listen for MCP protocol connections via standard input/output.

### Option 2: Build from Source

#### Prerequisites

- **Go 1.23.4+** (or compatible version)
- **Git** for cloning the repository

#### Quick Build

```bash
git clone https://github.com/x86ed/MCP-PoliticalCompass.git
cd MCP-PoliticalCompass
go mod download
go build -o mcp-political-compass .
./mcp-political-compass
```

#### Cross-Platform Build

To build for multiple platforms using the included build script:

```bash
# Build for all supported platforms
./build.sh

# Build with specific version
./build.sh v2.1.0

# Artifacts will be available in the dist/ directory
```

**Supported build targets**:

- linux/amd64, linux/arm64
- darwin/amd64, darwin/arm64 (macOS)
- windows/amd64
- freebsd/amd64, freebsd/arm64

The build script creates compressed archives (.tar.gz for Unix-like systems, .zip for Windows) and generates SHA256 checksums for verification.

## Development

### Project Structure

```markdown
MCP-PoliticalCompass/
├── main.go                # Server setup and configuration
├── tool.go                # Political compass quiz logic
├── *_test.go              # Comprehensive test suite
├── political-compass/     # Political compass data and interfaces
│   ├── interface.go       # Question and response definitions
│   ├── questions.go       # Complete dataset of 62 questions
│   └── interface_test.go  # Data integrity tests
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── VERSION                # Current version tracking
└── README.md              # This file
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out .
go tool cover -html=coverage.out -o coverage.html
```

### Test Coverage

The project maintains **85.7% code coverage** with comprehensive testing:

- **100% coverage** on all core business logic (`tool.go`)
- **100% coverage** on data structures (`politicalcompass` package)
- **50+ individual test cases** covering edge cases, boundary conditions, and error handling
- **Integration tests** for complete quiz workflows
- **SVG generation tests** for visualization features
- **Quiz status functionality tests** for progress tracking

### Key Dependencies

- **[mcp-golang](https://github.com/metoro-io/mcp-golang)**: Official Go MCP SDK
- **Standard Go libraries**: No external runtime dependencies

### Release Process

This project uses automated GitHub Actions to build and release binaries:

1. **Automated Builds**: Every push triggers cross-platform builds
2. **Tagged Releases**: Creating a git tag (e.g., `v2.1.0`) triggers a GitHub release
3. **Manual Releases**: Use `gh release create` or the GitHub web interface

**Creating a Release**:

```bash
# Tag and push a new version
git tag v2.1.0
git push origin v2.1.0

# Or create manually with GitHub CLI
gh release create v2.1.0 dist/* --title "Release v2.1.0" --notes "Release notes here"
```

**Build Pipeline**:

- Runs tests across all packages
- Cross-compiles for all supported platforms  
- Creates compressed archives with embedded version info
- Generates SHA256 checksums
- Uploads artifacts to GitHub Releases

## Algorithm Details

### Scoring Methodology

The server implements the authentic Political Compass scoring algorithm:

1. **Response Mapping**:
   - Strongly Disagree: Index 0
   - Disagree: Index 1
   - Agree: Index 2
   - Strongly Agree: Index 3

2. **Score Calculation**:
   - Each question has Economic and Social score arrays `[4]float64`
   - User response index determines which score to apply
   - Scores accumulate across all 62 questions

3. **Final Position Calculation**:

   ```go
   Economic Position = (Total Economic Score / 8.0) + 0.38
   Social Position = (Total Social Score / 19.5) + 2.41
   ```

4. **Quadrant Determination**:
   - Economic > 0, Social > 0: **Libertarian Left**
   - Economic > 0, Social ≤ 0: **Authoritarian Left**
   - Economic ≤ 0, Social > 0: **Libertarian Right**
   - Economic ≤ 0, Social ≤ 0: **Authoritarian Right**

## API Reference

### political_compass Tool

**Purpose**: Present political compass questions and process responses

**Arguments**:

- `response` (string, required): One of:
  - `"Strongly Disagree"`
  - `"Disagree"`
  - `"Agree"`
  - `"Strongly Agree"`
  - `""` (empty string to start quiz)

**Returns**: Tool response with question text, progress, and current scores

### reset_quiz Tool

**Purpose**: Reset quiz progress to start fresh

**Arguments**: None

**Returns**: Confirmation message that quiz has been reset

### quiz_status Tool

**Purpose**: Display current quiz progress and statistics

**Arguments**: None

**Returns**: Tool response with:
- Current progress (questions answered / total questions)
- Progress percentage
- Current economic and social scores (if in progress)
- Response distribution statistics
- Overall quiz state information

## Version History

- **v2.2.0**: Added quiz status tool with progress tracking and thread safety improvements
- **v1.0.5**: Added interactive SVG visualization for quiz results
- **v1.0.4**: Fixed README.md markdown formatting issues
- **v1.0.3**: Significantly improved test coverage to 88.3%
- **v1.0.2**: Fixed scoring algorithm to match original Political Compass
- **v1.0.1**: Initial bug fixes and improvements
- **v1.0.0**: Initial release

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes with tests
4. Ensure tests pass (`go test ./...`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Political Compass questions and methodology based on the original [Political Compass](https://www.politicalcompass.org/) project
- Built using the [Model Context Protocol](https://modelcontextprotocol.io/) specification
- Powered by [mcp-golang](https://github.com/metoro-io/mcp-golang) SDK
