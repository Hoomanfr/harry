PHONY: migrateup migratedown

migrateup:
	goose -dir ./deployments/migrations -env ./deployments/.env up

migratedown:
	goose -dir ./deployments/migrations -env ./deployments/.env down