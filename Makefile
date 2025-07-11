.PHONY: docker-cleanup docker-cleanup-all db-up db-down db-restart db-logs db-ps db-clean db-reset db-build db-shell api-run api-deps api-init-db api-clean-db api-setup

# Docker general cleanup
docker-cleanup:
	docker system prune -f

docker-cleanup-all:
	docker system prune -a -f
	docker volume prune -f

# Database management
db-up:
	docker-compose up -d

db-down:
	docker-compose down

db-restart:
	docker-compose restart

db-logs:
	docker-compose logs -f

db-ps:
	docker-compose ps

# Clean database (stops containers and removes volumes)
db-clean:
	docker-compose down -v
	docker-compose rm -f

# Complete reset (stops, removes containers, volumes, and images)
db-reset:
	docker-compose down -v --rmi all
	docker-compose rm -f

# Build and start fresh
db-build:
	docker-compose up --build -d

# Execute commands in database container
db-shell:
	docker-compose exec mongo mongosh --username root --password password --authenticationDatabase admin

# API management
api-run:
	cd api && go run cmd/api/main.go

api-deps:
	cd api && go mod tidy

api-init-db:
	cd api && go run cmd/cli/main.go init-db

api-clean-db:
	cd api && go run cmd/cli/main.go clean-db

api-setup:
	@echo "🔧 Setting up development environment..."
	$(MAKE) db-up
	@echo "⏳ Waiting for MongoDB to be ready..."
	sleep 5
	$(MAKE) api-init-db