# 🎯 Subscription Manager API

A clean, scalable subscription management API built with Go, following Clean Architecture principles.

## 🏗️ Architecture

This project follows Clean Architecture and SOLID principles to ensure maintainability and testability.

### 📁 Project Structure

```
api/
├── cmd/
│   ├── api/
│   │   └── main.go           # 🚀 API server entry point
│   └── cli/
│       └── main.go           # 🛠️ CLI tools entry point
├── internal/
│   ├── config/
│   │   └── config.go         # ⚙️ Configuration management
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── user.go       # 👤 User domain entity
│   │   │   ├── product.go    # 📦 Product domain entity
│   │   │   └── subscription.go # 💳 Subscription domain entity
│   │   └── repositories/
│   │       ├── user_repository.go         # 👤 User repository interface
│   │       ├── product_repository.go      # 📦 Product repository interface
│   │       └── subscription_repository.go # 💳 Subscription repository interface
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── connection.go    # 🔗 Database connection
│   │   │   ├── migrations.go    # 📋 Database initialization
│   │   │   └── repositories/
│   │   │       ├── mongo_user_repository.go         # 🍃 MongoDB user implementation
│   │   │       ├── mongo_product_repository.go      # 🍃 MongoDB product implementation
│   │   │       └── mongo_subscription_repository.go # 🍃 MongoDB subscription implementation
│   │   └── web/
│   │       ├── router.go        # 🌐 Route definitions
│   │       ├── middleware/
│   │       │   └── cors.go      # 🔒 CORS middleware
│   │       └── handlers/
│   │           ├── user_handler.go         # 👤 User HTTP handlers
│   │           ├── product_handler.go      # 📦 Product HTTP handlers
│   │           └── subscription_handler.go # 💳 Subscription HTTP handlers
│   └── usecases/
│       ├── user_usecase.go         # 👤 User business logic
│       ├── product_usecase.go      # 📦 Product business logic
│       └── subscription_usecase.go # 💳 Subscription business logic
├── pkg/
│   ├── errors/
│   │   └── errors.go           # ❌ Custom error types
│   └── utils/
│       └── response.go         # 📤 HTTP response helpers
├── go.mod
├── go.sum
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Make

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/subs-manager.git
   cd subs-manager
   ```

2. **Start the environment**
   ```bash
   make api-setup
   ```

3. **Run the API**
   ```bash
   make api-run
   ```

4. **Test the API**
   ```bash
   curl http://localhost:8080/ping
   ```

## 🛠️ Development Commands

### Database Management

| Command | Description |
|---------|-------------|
| `make db-up` | Start MongoDB container |
| `make db-down` | Stop MongoDB container |
| `make db-clean` | Remove containers and volumes |
| `make db-logs` | View database logs |
| `make db-shell` | Access MongoDB shell |

### API Management

| Command | Description |
|---------|-------------|
| `make api-run` | Start the API server |
| `make api-deps` | Install dependencies |
| `make api-init-db` | Initialize database with sample data |
| `make api-clean-db` | Reset database to default state |
| `make api-setup` | Complete setup (DB + initialization) |

### Docker Management

| Command | Description |
|---------|-------------|
| `make docker-cleanup` | Clean unused Docker resources |
| `make docker-cleanup-all` | Clean all Docker resources |

## 🔌 API Endpoints

### Health Check
- `GET /ping` - Health check endpoint

### Subscriptions
- `GET /api/v1/subscriptions` - Get all subscriptions
- `GET /api/v1/subscriptions/:id` - Get subscription by ID
- `POST /api/v1/subscriptions` - Create new subscription
- `PUT /api/v1/subscriptions/:id` - Update subscription
- `DELETE /api/v1/subscriptions/:id` - Delete subscription

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user

### Products
- `GET /api/v1/products` - Get all products
- `GET /api/v1/products/:id` - Get product by ID
- `POST /api/v1/products` - Create new product

## 📚 Data Models

### User
```json
{
  "id": "ObjectId",
  "username": "string",
  "email": "string",
  "status": "active|inactive",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Product
```json
{
  "id": "ObjectId",
  "name": "string",
  "description": "string",
  "price": "number",
  "billing_type": "monthly|yearly",
  "category": "string",
  "status": "active|inactive",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Subscription
```json
{
  "id": "ObjectId",
  "user_id": "ObjectId",
  "product_id": "ObjectId",
  "status": "active|cancelled|expired",
  "start_date": "timestamp",
  "end_date": "timestamp",
  "next_billing": "timestamp",
  "price_at_start": "number",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

## 🧪 Testing

Test the API using the provided HTTP file:

```bash
# View available requests
cat example.http

# Or use curl
curl -X GET http://localhost:8080/api/v1/subscriptions
```

## 🏛️ Architecture Principles

### Clean Architecture Layers

1. **Entities** (`internal/domain/entities/`)
   - Core business objects
   - Independent of external concerns

2. **Use Cases** (`internal/usecases/`)
   - Business logic and rules
   - Orchestrates data flow

3. **Interface Adapters** (`internal/infrastructure/`)
   - Database implementations
   - Web handlers and routing

4. **Frameworks & Drivers** (`cmd/`, `pkg/`)
   - Entry points and utilities
   - External libraries integration

### SOLID Principles

- **Single Responsibility**: Each layer has one reason to change
- **Open/Closed**: Open for extension, closed for modification
- **Liskov Substitution**: Interfaces can be substituted
- **Interface Segregation**: Small, focused interfaces
- **Dependency Inversion**: Depend on abstractions, not concretions

## 🔧 Configuration

The API uses environment variables for configuration:

```bash
# Database
MONGO_URI=mongodb://root:password@localhost:27017
MONGO_DB_NAME=subs-db

# Server
PORT=8080
GIN_MODE=release
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

---

**Built with ❤️ using Go, Gin, and MongoDB**