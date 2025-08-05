run-all: stop-all
	$(MAKE) run-infra
	sleep 10
	$(MAKE) db-migrateup-all
	$(MAKE) srv-build-run-all

stop-all: srv-stop-all stop-infra
	
run-infra:
	docker compose -f deployments/infrastructure/docker-compose.yml up -d

stop-infra:
	docker compose -f deployments/infrastructure/docker-compose.yml down

srv-build-all:
	docker compose -f deployments/docker-compose.yml build

srv-run-all:
	docker compose -f deployments/docker-compose.yml up -d

srv-build-run-all: srv-build-all srv-run-all

srv-stop-all:
	docker compose -f deployments/docker-compose.yml down

db-migrateup-all:
	deployments/create_db_schema.sh && \
	goose -dir ./receivingsrv/deployments/migrations -env ./receivingsrv/deployments/.env up && \
	goose -dir ./inventorysrv/deployments/migrations -env ./inventorysrv/deployments/.env up && \
	goose -dir ./fulfillmentsrv/deployments/migrations -env ./fulfillmentsrv/deployments/.env up

# Receiving srv
build-receiving-srv:
	docker compose -f deployments/docker-compose.yml build receiving-srv

run-receiving-srv:
	docker compose -f deployments/docker-compose.yml up receiving-srv

build-run-receiving-srv: build-receiving-srv run-receiving-srv

# Inventory srv
build-inventory-srv:
	docker compose -f deployments/docker-compose.yml build inventory-srv

run-inventory-srv:
	docker compose -f deployments/docker-compose.yml up inventory-srv

build-run-inventory-srv: build-inventory-srv run-inventory-srv

# Inventory Integrator
build-inventory-intg:
	docker compose -f deployments/docker-compose.yml build inventory-integrator

run-inventory-intg:
	docker compose -f deployments/docker-compose.yml up inventory-integrator

build-run-inventory-intg: build-inventory-intg run-inventory-intg

# Fulfillment srv
build-fulfillment-srv:
	docker compose -f deployments/docker-compose.yml build fulfillment-srv

run-fulfillment-srv:
	docker compose -f deployments/docker-compose.yml up fulfillment-srv

build-run-fulfillment-srv: build-fulfillment-srv run-fulfillment-srv

# Fulfillment Integrator
build-fulfillment-intg:
	docker compose -f deployments/docker-compose.yml build fulfillment-integrator

run-fulfillment-intg:
	docker compose -f deployments/docker-compose.yml up fulfillment-integrator

build-run-fulfillment-intg: build-fulfillment-intg run-fulfillment-intg
