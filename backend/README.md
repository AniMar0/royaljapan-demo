# Royal Japan Demo - Backend API

A lightweight, high-performance REST API backend built with Go and Gin framework.

## Tech Stack

- **Language**: Go 1.23+
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) v1.11.0
- **HTTP Server**: Go's built-in `net/http`
- **Port**: 8080

## Why Gin Framework?

Gin was selected as the web framework for this project for several key reasons:

### 1. **High Performance**
- One of the fastest Go web frameworks available
- Uses `httprouter` under the hood for ultra-fast routing
- Minimal overhead and excellent throughput

### 2. **Lightweight & Fast**
- Small memory footprint
- Fast compilation and startup times
- Efficient request handling

### 3. **Developer-Friendly**
- Clean, intuitive API design
- Easy-to-use middleware system
- Clear error messages and debugging support

### 4. **Built-in JSON Support**
- Native JSON validation and rendering
- `gin.H` shorthand for JSON responses
- Automatic content negotiation

### 5. **Production-Ready**
- Includes panic recovery middleware
- Built-in logging and monitoring
- Battle-tested in production environments

### 6. **Active Community**
- Well-maintained with regular updates
- Extensive documentation and examples
- Large ecosystem of middleware and plugins

## API Endpoints

### Health Check
```
GET /healthz
```

Returns the service status and application version.

**Response** (200 OK):
```json
{
  "status": "ok",
  "version": "0.1.0"
}
```

## Version Management

The application supports multiple version sources with the following priority:

1. **VERSION file** - Plain text file containing version string
2. **pom.xml** - Maven configuration file (reads `<version>` tag)
3. **package.json** - Node.js package file (reads `"version"` field)
4. **APP_VERSION** - Environment variable
5. **Default** - Fallback to `"1.0.0"`

### Setting Version

Choose one of these methods:

**Option 1: VERSION file**
```bash
echo "1.2.3" > VERSION
```

**Option 2: package.json**
```json
{
  "version": "1.2.3"
}
```

**Option 3: Environment variable**
```bash
export APP_VERSION="1.2.3"
go run main.go
```

**Option 4: pom.xml** (for Maven projects)
```xml
<project>
  <version>1.2.3</version>
</project>
```

## Getting Started

### Prerequisites
- Go 1.23 or higher
- Git (optional, for cloning)

### Installation

1. **Clone the repository** (if not already done):
```bash
cd backend
```

2. **Install dependencies**:
```bash
go mod download
```

3. **Run the server**:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Testing the API

Test the health endpoint:
```bash
curl http://localhost:8080/healthz
```

Expected response:
```json
{
  "status": "ok",
  "version": "0.1.0"
}
```

## Project Structure

```
backend/
├── main.go           # Application entry point and API routes
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
├── package.json      # Version configuration (Node.js style)
└── README.md         # This file
```

## Development

### Running in Development Mode
```bash
go run main.go
```

### Building for Production
```bash
go build -o server main.go
./server
```

### Running with Custom Version
```bash
APP_VERSION="2.0.0" go run main.go
```

## Code Documentation

The codebase follows Go documentation standards:

- **Package-level documentation**: Describes the overall application purpose and architecture
- **Function documentation**: Each exported function includes detailed comments explaining parameters, return values, and behavior
- **Inline comments**: Complex logic is explained with inline comments

View documentation:
```bash
go doc -all
```

## Dependencies

Main dependencies listed in `go.mod`:

- `github.com/gin-gonic/gin` v1.11.0 - Web framework
- `github.com/gin-contrib/sse` - Server-Sent Events support
- `github.com/go-playground/validator/v10` - Struct validation
- Standard library packages for JSON, XML, HTTP, etc.

## Contact

[Email](mailto:mr.zakariakahlaoui@gmail.com)
