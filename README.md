# MCP-PoliticalCompass

A Model Context Protocol (MCP) server that provides interactive political quiz tools for baselining the political beliefs of AI agents and users. This server implements three comprehensive political questionnaires: the classic Political Compass (62 questions), the detailed 8values quiz (70 questions), and the multilingual Politiscales quiz (117 questions).

## Version 3.0+ Migration

**Version 3.0.0** represents a major migration from the legacy `metoro-io/mcp-golang` library to the modern `mark3labs/mcp-go` library. This migration includes:

- **New MCP Library**: Migrated to `mark3labs/mcp-go` for better performance and maintainability
- **Explicit Schemas**: Removed reflection-based schemas in favor of explicit tool definitions
- **Enhanced Error Handling**: Improved error reporting and validation
- **Full Test Coverage**: All 150+ tests updated and passing
- **Breaking Changes**: Module path updated to `/v3` following Go module versioning conventions

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

The Politiscales quiz offers the most comprehensive political analysis with 117 questions across multiple political dimensions, faithfully implementing the original PolitiScales methodology:

#### Core Features

- **117 authentic questions** from the original PolitiScales questionnaire
- **17 political axes** with sophisticated scoring algorithms
- **Complete multilingual support**: English, French, Spanish, Italian, Arabic, Russian, and Chinese
- **Authentic scoring implementation** matching the original TypeScript/Vue.js implementation
- **Professional SVG visualizations** with accurate colors, positioning, and labels

#### Political Dimensions

**Paired Axes (with opposing values):**

- **Identity**: Constructivism ↔ Essentialism
- **Justice**: Rehabilitative ↔ Punitive  
- **Culture**: Progressive ↔ Conservative
- **Globalism**: Internationalism ↔ Nationalism
- **Economy**: Communism ↔ Capitalism
- **Markets**: Regulation ↔ Laissez-faire
- **Environment**: Ecology ↔ Production
- **Change**: Revolution ↔ Reform

**Unpaired Axes (specialized indicators with thresholds):**

- **Anarchism** (threshold: 0.9) - Anti-state tendencies
- **Feminism** (threshold: 0.9) - Gender equality advocacy
- **Conspiracism** (threshold: 0.9) - Conspiracy theory acceptance
- **Pragmatism** (threshold: 0.5) - Practical over ideological approaches
- **Veganism** (threshold: 0.5) - Animal rights and dietary ethics
- **Monarchism** (threshold: 0.5) - Support for monarchical systems
- **Religion** (threshold: 0.5) - Religious influence in politics

#### Advanced Scoring Features

- **Paired axis normalization** ensuring balanced representation
- **Neutral value calculation** for incomplete ideological positions
- **Threshold-based badge system** highlighting specific political tendencies
- **Unified badge and slogan logic** sourced directly from module data
- **Dynamic language switching** with complete UI translations

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

- **`politiscales`**: Interactive politiscales quiz tool with 117 questions across 17 political axes
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

- **117 comprehensive questions** covering 17 distinct political axes
- **Multilingual support** with 7 languages: English, French, Spanish, Italian, Arabic, Russian, Chinese
- **Randomized question order** for unbiased assessment
- **Five response options**: Strongly Disagree, Disagree, Neutral, Agree, Strongly Agree
- **Sophisticated scoring algorithm** with paired axis normalization matching the original TypeScript implementation
- **17 political axes** including both paired and standalone dimensions:
  - **Paired axes**: Identity (Constructivism vs Essentialism), Justice (Rehabilitative vs Punitive), Culture (Progressive vs Conservative), Globalism (Internationalism vs Nationalism), Economy (Communism vs Capitalism), Markets (Regulation vs Laissez-faire), Environment (Ecology vs Production), Change (Revolution vs Reform)
  - **Unpaired axes (badges)**: Anarchism (0.9), Pragmatism (0.5), Feminism (0.9), Conspiracism (0.9), Veganism (0.5), Monarchism (0.5), Religion (0.5)
- **Enhanced results display** showing both paired values and neutral states for comprehensive analysis
- **Threshold-based badge system** that highlights specific political tendencies with proper colors and labels sourced from module data
- **Consistent SVG and text results** with unified badge logic and neutral state calculations
- **Real-time progress tracking** with language and completion status
- **Response distribution analytics** with detailed breakdown by response type
- **Dynamic language switching** (only available before starting the quiz)
- **Authentic implementation** faithfully reproducing the original PolitiScales methodology and user experience

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

#### Political Compass Usage Example

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

#### 8values Usage Example

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

#### Politiscales Usage Example

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

#### Political Compass Completion Results

- Final Economic Score (e.g., -1.23)
- Final Social Score (e.g., 2.45)
- Political Quadrant placement
- **Interactive SVG compass visualization** showing position on political grid
- Total questions answered

#### 8values Completion Results

- Economic Axis percentage and classification (Socialist/Capitalist)
- Diplomatic Axis percentage and classification (Internationalist/Nationalist)
- Government Axis percentage and classification (Libertarian/Authoritarian)
- Society Axis percentage and classification (Progressive/Traditional)
- **Interactive SVG bar chart visualization** showing all four axis scores
- Total questions answered and response distribution

#### Politiscales Completion Results

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

- **[mcp-go](https://github.com/mark3labs/mcp-go)**: Modern Go MCP SDK (v3.0+ migration)
- **Standard Go libraries**: No external runtime dependencies

### Version 3.0+ Installation

For the latest version (v3.0+), ensure you're using the correct module path:

```go
import "github.com/x86ed/MCP-PoliticalCompass/v3"
```

### Migration from v2.x

If upgrading from v2.x, note the breaking changes in v3.0+:

- Module path changed to `/v3`
- Migrated to `mark3labs/mcp-go` library
- Updated tool schemas and handler signatures
- See the [v3.0.0 release notes](https://github.com/x86ed/MCP-PoliticalCompass/releases/tag/v3.0.0) for full details

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

### v3.2.1 (2025-06-19)

- **Test Coverage Improvement**: Significantly improved test coverage from 85.6% to 90.5%
- **Enhanced EightValues Testing**: Increased `handleEightValues` coverage from 55.7% to 98.4%
- **Better Quality Assurance**: Added comprehensive tests for quiz completion scenarios and edge cases
- **Improved Reliability**: Enhanced testing of score calculations, label generation, and progress tracking

### v3.2.0 (2025-06-19)

- **Enhanced Visualization**: Added SVG chart rendering to Politiscales quiz completion
- **Consistent User Experience**: All three quizzes (Political Compass, 8Values, Politiscales) now include visual charts
- **Improved Instructions**: Added clear SVG rendering instructions for better user experience
- **Test Coverage**: Updated tests to verify SVG inclusion in all quiz completions

### v3.1.0 (2025-06-19)

- **Documentation Update**: Comprehensive README update with v3.0+ migration information
- **Test Suite**: Re-enabled all previously disabled test files
- **Library Documentation**: Updated all references from `metoro-io/mcp-golang` to `mark3labs/mcp-go`
- **Installation Guide**: Added v3.0+ installation instructions and migration notes
- **Quality**: All 150+ tests passing with full coverage

### v3.0.0 (2025-06-19) - Major Migration

- **🚨 BREAKING CHANGE**: Migrated from `metoro-io/mcp-golang` to `mark3labs/mcp-go`
- **Module Path**: Updated to `/v3` following Go module versioning conventions
- **Explicit Schemas**: Removed reflection-based schemas for explicit tool definitions
- **Handler Signatures**: Updated all tool handlers to new API patterns
- **Error Handling**: Enhanced error reporting and validation
- **Performance**: Improved performance with modern MCP library
- **Testing**: All test files updated and verified working
- **Architecture**: Complete rewrite of server initialization and tool registration

### v2.x Series Highlights

- **v2.8.9**: Restored original badge thresholds for unpaired axes in PolitiScales (anarchism, feminism, conspiracism: 0.9; pragmatism, veganism, monarchism, religion: 0.5)
- **v2.8.8**: Enhanced PolitiScales implementation with refactored SVG and text results, unified badge/slogan logic sourced from module data, and cleaned up codebase
- **v2.8.7**: Major PolitiScales overhaul - complete scoring algorithm refactor to match original TypeScript implementation, improved neutral value calculations, enhanced SVG visualization with proper layout and colors
- **v2.8.6**: Fixed PolitiScales scoring algorithm with proper paired axis normalization and neutral state handling matching the original Vue.js implementation
- **v2.8.5**: Added comprehensive PolitiScales testing suite with 100+ test scenarios, improved badge threshold accuracy, and enhanced multilingual support validation
- **v2.8.4**: Enhanced politiscales functionality with improved neutral state handling, unified badge system sourced from module data, and comprehensive paired axis display in both SVG and text results
- **v2.8.3**: Complete internationalization of politiscales with full UI and question translations for Arabic, Spanish, French, Italian, Russian, and Chinese
- **v2.3.0**: Added complete 8values quiz implementation with 70 questions, improved test coverage to 90.9%, and comprehensive SVG visualizations
- **v2.2.0**: Added quiz status tool with progress tracking and thread safety improvements

### v1.x Series

- **v1.0.5**: Added interactive SVG visualization for quiz results
- **v1.0.4**: Fixed README.md markdown formatting issues
- **v1.0.3**: Significantly improved test coverage to 88.3%
- **v1.0.2**: Fixed scoring algorithm to match original Political Compass
- **v1.0.1**: Initial bug fixes and improvements
- **v1.0.0**: Initial release

### Migration Notes

**Upgrading to v3.0+:**

- Update import paths to include `/v3`: `github.com/x86ed/MCP-PoliticalCompass/v3`
- No functional changes to quiz behavior or API
- Enhanced performance and reliability with new MCP library
- All existing tool names and parameters remain the same

**From v2.x to v3.0+:**

- Module path change required due to major version bump
- Internal architecture updated but external API unchanged
- Recommended for all new installations

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
- Politiscales questions and methodology based on the original [Politiscales](https://www.politiscales.net/) project
- Built using the [Model Context Protocol](https://modelcontextprotocol.io/) specification
- Powered by [mcp-go](https://github.com/mark3labs/mcp-go) SDK (v3.0+)
