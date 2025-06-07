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

### Quiz Capabilities

- **62 authentic questions** from the Political Compass dataset
- **Randomized question order** for each quiz session
- **Real-time progress tracking** with current scores
- **Authentic scoring algorithm** that matches the original Political Compass methodology
- **Detailed final analysis** with quadrant placement and scores

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

**Quiz completion provides:**

- Final Economic Score (e.g., -1.23)
- Final Social Score (e.g., 2.45)
- Political Quadrant placement
- Total questions answered

## Installation & Setup

### Prerequisites

- **Go 1.23.4+** (or compatible version)
- **Git** for cloning the repository

### 1. Clone the Repository

```bash
git clone https://github.com/x86ed/MCP-PoliticalCompass.git
cd MCP-PoliticalCompass
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build the Server

```bash
go build -o political-compass-server ./cmd
```

### 4. Run the Server

```bash
./political-compass-server
```

The server will start and listen for MCP protocol connections via standard input/output.

## Development

### Project Structure

```markdown
MCP-PoliticalCompass/
├── cmd/                    # Main application
│   ├── main.go            # Server setup and configuration
│   ├── tool.go            # Political compass quiz logic
│   ├── prompt.go          # Additional prompt handlers
│   └── *_test.go          # Comprehensive test suite
├── polcomp/               # Political compass data and interfaces
│   ├── interface.go       # Question and response definitions
│   ├── questions.go       # Complete dataset of 62 questions
│   └── interface_test.go  # Data integrity tests
├── go.mod                 # Go module definition
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
go test -coverprofile=coverage.out ./cmd
go tool cover -html=coverage.out -o coverage.html
```

### Test Coverage

The project maintains **88.3% code coverage** with comprehensive testing:

- **100% coverage** on all core business logic (`tool.go`)
- **100% coverage** on prompt handlers (`prompt.go`)
- **45+ individual test cases** covering edge cases, boundary conditions, and error handling
- **Integration tests** for complete quiz workflows

### Key Dependencies

- **[mcp-golang](https://github.com/metoro-io/mcp-golang)**: Official Go MCP SDK
- **Standard Go libraries**: No external runtime dependencies

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

## Version History

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
