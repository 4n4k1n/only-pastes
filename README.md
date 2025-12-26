# Only Pastes

A lightweight, self-hosted pastebin service built with Go and PostgreSQL. Share code snippets with automatic expiration, syntax highlighting, and a clean, minimal interface.

## Features

- **Simple Paste Creation**: Create and share code snippets with a clean web interface
- **Syntax Highlighting**: Support for 18+ programming languages using highlight.js
- **Automatic Expiration**: Set pastes to expire after 1 hour, 1 day, 1 week, or never
- **Unique URLs**: Generate short, unique slugs for easy sharing
- **View Counter**: Track how many times a paste has been viewed
- **REST API**: Full API with Swagger documentation
- **Docker Ready**: Complete containerized deployment with Docker Compose
- **Automatic HTTPS**: Caddy reverse proxy with automatic SSL certificates

## Tech Stack

**Backend:**
- Go 1.24.4
- Gin web framework
- PostgreSQL database
- Swagger/OpenAPI documentation

**Frontend:**
- HTML5, CSS3, Vanilla JavaScript
- highlight.js for syntax highlighting

**Deployment:**
- Docker & Docker Compose
- Caddy v2 reverse proxy
- Multi-stage Docker builds

## Quick Start

### Prerequisites

- Docker and Docker Compose
- (Optional) Make for build automation

### Installation

1. Clone the repository:
```bash
git clone https://github.com/4n4k1n/only-pastes.git
cd pastes
```

2. Create environment configuration:
```bash
cp .env.example .env
```

3. Edit `.env` with your configuration:
```env
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=pastes
POSTGRES_HOST=db
POSTGRES_PORT=5432
PORT=8080
BASE_URL=http://localhost:8080
```

4. Start the services:
```bash
make start
# or
docker-compose up -d
```

5. Access the application:
- Web Interface: `http://localhost:8080`
- API Documentation: `http://localhost:8080/swagger/index.html`

## Usage

### Creating a Paste

1. Navigate to the homepage
2. Enter your code snippet
3. Select the programming language
4. Choose expiration time
5. Click "Create Paste"
6. Share the generated URL

### API Endpoints

#### Create Paste
```bash
POST /api/paste
Content-Type: application/json

{
  "content": "console.log('Hello World');",
  "language": "javascript",
  "expiresIn": "1d"
}
```

#### Retrieve Paste
```bash
GET /api/paste/:slug
```

**Expiration Options:**
- `1h` - 1 hour
- `1d` - 1 day
- `1w` - 1 week
- `never` - No expiration

## Development

### Build Commands

```bash
# Start all services
make start

# Build Docker images
make build

# Stop services
make down

# Restart with rebuild
make restart

# Clean build artifacts
make clean

# Install Go dependencies
make install

# Start database only
make db-start
```

### Local Development (Without Docker)

1. Start PostgreSQL:
```bash
make db-start
```

2. Install dependencies:
```bash
make install
```

3. Set up `.env` file with local configuration

4. Run the application:
```bash
go run main.go
```

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `POSTGRES_USER` | PostgreSQL username | - |
| `POSTGRES_PASSWORD` | PostgreSQL password | - |
| `POSTGRES_DB` | Database name | `pastes` |
| `POSTGRES_HOST` | Database host | `db` |
| `POSTGRES_PORT` | Database port | `5432` |
| `PORT` | Application port | `8080` |
| `BASE_URL` | Base URL for paste links | - |

### Supported Languages

JavaScript, Python, Go, Java, C, C++, C#, Ruby, PHP, Swift, Kotlin, Rust, TypeScript, HTML, CSS, SQL, Bash, and Plain Text.

## Project Structure

```
.
├── main.go                 # Application entry point
├── handlers/
│   └── paste.go           # HTTP request handlers
├── models/
│   └── paste.go           # Data models
├── database/
│   └── db.go              # Database connection and migrations
├── static/
│   ├── index.html         # Create paste interface
│   ├── view.html          # View paste interface
│   ├── app.js             # Frontend logic for creating
│   ├── view.js            # Frontend logic for viewing
│   └── style.css          # Styling
├── docs/                  # Swagger documentation
├── docker-compose.yml     # Docker orchestration
├── Dockerfile             # Multi-stage build
├── Caddyfile             # Reverse proxy configuration
├── Makefile              # Build automation
└── .env.example          # Configuration template
```

## Deployment

### Production Deployment with Caddy

The project includes Caddy configuration for automatic HTTPS:

1. Update `Caddyfile` with your domain
2. Configure DNS to point to your server
3. Update `BASE_URL` in `.env`
4. Run `docker-compose up -d`

Caddy will automatically obtain and renew SSL certificates from Let's Encrypt.

## License

[Add your license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions, please use the GitHub issue tracker.