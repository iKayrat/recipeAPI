# # Load environment variables from .env file
# include .env

# # Export all variables from .env file
# export $(shell sed 's/=.*//' .env)

psql:
	docker exec -it db psql -U root recipe

createmigrate:
	# migrate create -ext sql -dir internal/db/migration -seq recipe
	# migrate create -ext sql -dir internal/db/migration -seq user

sqlc:
	sqlc generate

create:
	docker exec -it db createdb --username=root --owner=root recipe

drop:
	docker exec -it db dropdb --username=root recipe

migrateup:
	migrate -path internal/db/migration -database $(DB_URL) -verbose up

run:
	docker-compose up

stop:
	docker-compose down