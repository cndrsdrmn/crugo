# Crugo

A simple and clean **RESTful CRUD API** built with **Go**. Crugo demonstrates how to structure a Go project with full CRUD operations using common frameworks and best practices.

## Features

- RESTful API with CRUD operations
- Clean project structure following Go best practices
- Use of popular libraries:
  - [Gin](https://github.com/gin-gonic/gin) for HTTP routing
  - [GORM](https://gorm.io/) for ORM (with SQLite)
  - [Viper](https://github.com/spf13/viper) for environment variable management
- Docker support for easy deployment
- Basic request validation and error handling

## Getting Started

### Prerequisites

- Go 1.20+
- Docker (optional for containerization)
- SQLite (used as the default database)

### Installation

```bash
git clone https://github.com/cndrsdrmn/crugo.git
cd crugo
go mod tidy
```

## Tech Stack

- **Language:** Go
- **Web Framework:** Gin
- **Database:** SQLite (via GORM)
- **Configuration:** Viper
- **Deployment:** Docker

## Contributing

1. Fork the repo
2. Create a new branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## Acknowledgments

- Thanks to the creators of Gin, GORM, Viper, and the Go community for their awesome tools and resources.

## Contact

Made with ❤️ by [Candra Sudirman](https://github.com/cndrsdrmn). Feel free to reach out for questions or collaboration!
