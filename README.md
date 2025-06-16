# MCP-PoliticalCompass

A Model Context Protocol (MCP) server that provides interactive political quiz tools for baselining the political beliefs of AI agents and users. This server implements three comprehensive political questionnaires: the classic Political Compass (62 questions), the detailed 8values quiz (70 questions), and the multilingual Politiscales quiz (104 questions).

## Overview

This MCP server provides three distinct political assessment tools:

### Political Compass Quiz

The Political Compass is a multi-axis political model that maps political ideologies across two dimensions:

- **Economic Axis**: Left (collectivist) to Right (individualist)
- **Social Axis**: Libertarian (freedom-focused) to Authoritarian (order-focused)

Results place users in one of four quadrants:

- **Libertarian Left** (Left-wing, freedom-oriented)
- **Libertarian Right** (Right-wing, freedom-oriented)
- **Authoritarian Left** (Left-wing, order-oriented)
- **Authoritarian Right** (Right-wing, order-oriented)

### 8values Quiz

The 8values quiz provides a more detailed analysis across four political axes with 70 questions:

- **Economic Axis**: Socialist (equality-focused) to Capitalist (market-focused)
- **Diplomatic Axis**: Internationalist (globalist) to Nationalist (isolationist)
- **Government Axis**: Libertarian (freedom-focused) to Authoritarian (order-focused)
- **Society Axis**: Progressive (change-oriented) to Traditional (stability-focused)

Each axis provides percentage scores and ideological classifications based on your responses.

### Politiscales Quiz

The Politiscales quiz offers the most comprehensive political analysis with 104 questions across multiple political dimensions:

- **17 political axes** including paired axes like Economy (Communism vs Capitalism), Culture (Progressive vs Conservative), and Identity (Constructivism vs Essentialism)
- **Multiple language support**: English, French, Spanish, Italian, Arabic, Russian, and Chinese
- **Specialized indicators** for specific political tendencies like Anarchism, Feminism, and Pragmatism
- **Sophisticated scoring algorithm** with paired axis normalization and threshold-based indicators

## Features

### Tools Available

#### Political Compass Tools

- **`political_compass`**: Interactive quiz tool that presents randomized political questions
- **`reset_quiz`**: Resets Political Compass quiz progress to start fresh
- **`quiz_status`**: Shows current Political Compass quiz progress and statistics

#### 8values Tools

- **`eight_values`**: Interactive 8values quiz tool with 70 questions across four political axes
- **`reset_eight_values`**: Resets 8values quiz progress to start fresh
- **`eight_values_status`**: Shows current 8values quiz progress and statistics

#### Politiscales Tools

- **`politiscales`**: Interactive politiscales quiz tool with 104 questions across 17 political axes
- **`reset_politiscales`**: Resets politiscales quiz progress to start fresh
- **`politiscales_status`**: Shows current politiscales quiz progress and statistics
- **`set_politiscales_language`**: Sets the language for the politiscales quiz (supports: en, fr, es, it, ar, ru, zh)

### Quiz Capabilities

#### Political Compass Features

- **62 authentic questions** from the Political Compass dataset
- **Randomized question order** for each quiz session
- **Real-time progress tracking** with current scores and status display
- **Response distribution analytics** showing breakdown by response type
- **Authentic scoring algorithm** that matches the original Political Compass methodology
- **Detailed final analysis** with quadrant placement and scores
- **Interactive SVG compass visualization** showing position on the classic political grid

#### 8values Features

- **70 comprehensive questions** covering four distinct political axes
- **Randomized question order** for unbiased assessment
- **Five response options**: Strongly Disagree, Disagree, Neutral, Agree, Strongly Agree
- **Real-time progress tracking** with percentage completion
- **Authentic 8values scoring algorithm** matching the original implementation
- **Detailed axis analysis** with percentage scores and ideological classifications
- **Interactive SVG bar chart visualization** showing position on all four axes
- **Response distribution analytics** with comprehensive breakdown

#### Politiscales Features

- **104 comprehensive questions** covering 17 distinct political axes
- **Multilingual support** with 7 languages: English, French, Spanish, Italian, Arabic, Russian, Chinese
- **Randomized question order** for unbiased assessment
- **Five response options**: Strongly Disagree, Disagree, Neutral, Agree, Strongly Agree
- **Sophisticated scoring algorithm** with paired axis normalization matching the original TypeScript implementation
- **17 political axes** including both paired and standalone dimensions:
  - **Paired axes**: Identity (Constructivism vs Essentialism), Justice (Rehabilitative vs Punitive), Culture (Progressive vs Conservative), Globalism (Internationalism vs Nationalism), Economy (Communism vs Capitalism), Markets (Regulation vs Laissez-faire), Environment (Ecology vs Production), Radicalism (Revolution vs Reform), Perspective (Materialism vs Idealism), Development (Sustainability vs Growth)
  - **Standalone indicators**: Anarchism, Pragmatism, Feminism, Complotism, Veganism, Monarchism, Religion
- **Threshold-based special indicators** that highlight specific political tendencies
- **Real-time progress tracking** with language and completion status
- **Response distribution analytics** with detailed breakdown by response type
- **Dynamic language switching** (only available before starting the quiz)

### Visualization Features

#### Political Compass Visualization

- **Color-coded quadrants**: Each political quadrant has a distinct background color
- **Position marker**: Red dot shows exact location on the political compass
- **Grid lines and labels**: Clear axis markings and quadrant labels
- **Coordinate display**: Shows precise numerical position (Economic, Social)
- **Professional styling**: Clean, readable design suitable for analysis

#### 8values Visualization

- **Bar chart format**: Horizontal bars showing percentage scores for each axis
- **Color-coded axes**: Each political axis has a distinct color scheme
- **Percentage labels**: Clear numerical values for each axis score
- **Ideological classifications**: Text labels showing political tendencies
- **Professional styling**: Clean, modern design optimized for readability

### Example Interactions

#### Political Compass Quiz

**Starting the Political Compass quiz:**

```yaml
Tool: political_compass
Response: ""
```

**Answering Political Compass questions:**

```yaml
Tool: political_compass
Response: "strongly_agree"
```

**Checking Political Compass progress:**

```yaml
Tool: quiz_status
```

#### 8values Quiz

**Starting the 8values quiz:**

```yaml
Tool: eight_values
Response: ""
```

**Answering 8values questions:**

```yaml
Tool: eight_values
Response: "agree"
```

**Checking 8values progress:**

```yaml
Tool: eight_values_status
```

#### Politiscales Quiz

**Setting language for Politiscales quiz:**

```yaml
Tool: set_politiscales_language
Language: "fr"
```

**Starting the Politiscales quiz:**

```yaml
Tool: politiscales
Response: ""
```

**Answering Politiscales questions:**

```yaml
Tool: politiscales
Response: "strongly_agree"
```

**Checking Politiscales progress:**

```yaml
Tool: politiscales_status
```

### Quiz Results

#### Political Compass completion provides:

- Final Economic Score (e.g., -1.23)
- Final Social Score (e.g., 2.45)
- Political Quadrant placement
- **Interactive SVG compass visualization** showing position on political grid
- Total questions answered

#### 8values completion provides:

- Economic Axis percentage and classification (Socialist/Capitalist)
- Diplomatic Axis percentage and classification (Internationalist/Nationalist)
- Government Axis percentage and classification (Libertarian/Authoritarian)
- Society Axis percentage and classification (Progressive/Traditional)
- **Interactive SVG bar chart visualization** showing all four axis scores
- Total questions answered and response distribution

#### Politiscales completion provides:

- **Paired axis results** with dominant tendency for each political dimension
- **Special indicator scores** for specific political tendencies (Anarchism, Feminism, etc.)
- **Multilingual results** displayed in the selected language
- **Sophisticated scoring** using paired axis normalization and threshold analysis
- **Comprehensive political profile** across 17 different political dimensions
- Total questions answered and response distribution by language

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
├── tool.go                # Political quiz logic (both compass and 8values)
├── *_test.go              # Comprehensive test suite
├── 8values.js             # Reference implementation for 8values scoring
├── political-compass/     # Political compass data and interfaces
│   ├── interface.go       # Question and response definitions
│   ├── questions.go       # Complete dataset of 62 questions
│   └── interface_test.go  # Data integrity tests
├── eightvalues/           # 8values quiz data and interfaces
│   ├── eightvalues.go     # Question definitions and constants
│   └── questions.go       # Complete dataset of 70 questions
├── politiscales/          # PolitiScales framework (future implementation)
│   └── politiscales.go    # Basic structure definitions
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

The project maintains **90.9% code coverage** with comprehensive testing:

- **100% coverage** on core business logic functions
- **100% coverage** on data structures (`politicalcompass` and `eightvalues` packages)
- **80+ individual test cases** covering edge cases, boundary conditions, and error handling
- **Integration tests** for complete quiz workflows (both Political Compass and 8values)
- **SVG generation tests** for both compass and bar chart visualizations
- **Quiz status functionality tests** for progress tracking on both quiz types
- **Edge case testing** for boundary conditions and error scenarios
- **Complete quiz completion tests** ensuring full workflow integrity

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

### eight_values Tool

**Purpose**: Present 8values questions and process responses across four political axes

**Arguments**:

- `response` (string, required): One of:
  - `"strongly_disagree"`
  - `"disagree"`
  - `"neutral"`
  - `"agree"`
  - `"strongly_agree"`
  - `""` (empty string to start quiz)

**Returns**: Tool response with question text, progress, and current scores

### reset_eight_values Tool

**Purpose**: Reset 8values quiz progress to start fresh

**Arguments**: None

**Returns**: Confirmation message that 8values quiz has been reset

### eight_values_status Tool

**Purpose**: Display current 8values quiz progress and statistics

**Arguments**: None

**Returns**: Tool response with:

- Current progress (questions answered / total questions)
- Progress percentage
- Current axis scores and classifications (if complete)
- Response distribution statistics
- Overall quiz state information

## Algorithm Details

### Political Compass Scoring Methodology

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

### 8values Scoring Methodology

The server implements the authentic 8values scoring algorithm:

1. **Response Mapping**:
   - Strongly Disagree: -1.0
   - Disagree: -0.5
   - Neutral: 0.0
   - Agree: 0.5
   - Strongly Agree: 1.0

2. **Score Calculation**:
   - Each question has effect values for four axes: Economic, Diplomatic, Government, Society
   - User response multiplier is applied to each effect value
   - Scores accumulate across all 70 questions

3. **Final Percentage Calculation**:

   ```go
   Axis Percentage = (100 * (Max + Score) / (2 * Max))
   ```

   Where `Max` is the maximum possible absolute score for that axis.

4. **Classification Determination**:
   - Economic > 50%: **Socialist**, ≤ 50%: **Capitalist**
   - Diplomatic > 50%: **Internationalist**, ≤ 50%: **Nationalist**
   - Government > 50%: **Libertarian**, ≤ 50%: **Authoritarian**
   - Society > 50%: **Progressive**, ≤ 50%: **Traditional**

## Version History

- **v2.3.0**: Added complete 8values quiz implementation with 70 questions, improved test coverage to 90.9%, and comprehensive SVG visualizations
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
- 8values questions and methodology based on the original [8values](https://8values.github.io/) project
- Built using the [Model Context Protocol](https://modelcontextprotocol.io/) specification
- Powered by [mcp-golang](https://github.com/metoro-io/mcp-golang) SDK
