# Load environment variables from .env file
include .env

# Export all variables from .env file
export $(shell sed 's/=.*//' .env)

createmigrate:
	# migrate create -ext sql -dir internal/db/migration -seq recipe
	migrate create -ext sql -dir internal/db/migration -seq user

sqlc:
	sqlc generate

create:
	docker exec -it recipedb createdb --username=root --owner=root recipe

drop:
	docker exec -it telegramdb dropdb --username=root recipe

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:kaak@localhost:5432/recipe?sslmode=disable" -verbose up

run:
	go run cmd/main.go