api/
├── cmd/
│   ├── api/
│   │   └── main.go           # API server entry point
│   └── cli/
│       └── main.go           # CLI tools entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── user.go       # User domain entity
│   │   │   ├── product.go    # Product domain entity
│   │   │   └── subscription.go # Subscription domain entity
│   │   └── repositories/
│   │       ├── user_repository.go         # User repository interface
│   │       ├── product_repository.go      # Product repository interface
│   │       └── subscription_repository.go # Subscription repository interface
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── connection.go    # Database connection
│   │   │   ├── migrations.go    # Database initialization
│   │   │   └── repositories/
│   │   │       ├── mongo_user_repository.go         # MongoDB user implementation
│   │   │       ├── mongo_product_repository.go      # MongoDB product implementation
│   │   │       └── mongo_subscription_repository.go # MongoDB subscription implementation
│   │   └── web/
│   │       ├── router.go        # Route definitions
│   │       ├── middleware/
│   │       │   └── cors.go      # CORS middleware
│   │       └── handlers/
│   │           ├── user_handler.go         # User HTTP handlers
│   │           ├── product_handler.go      # Product HTTP handlers
│   │           └── subscription_handler.go # Subscription HTTP handlers
│   └── usecases/
│       ├── user_usecase.go         # User business logic
│       ├── product_usecase.go      # Product business logic
│       └── subscription_usecase.go # Subscription business logic
├── pkg/
│   ├── errors/
│   │   └── errors.go           # Custom error types
│   └── utils/
│       └── response.go         # HTTP response helpers
├── go.mod
├── go.sum
└── README.md