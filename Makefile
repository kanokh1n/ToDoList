include .env
export

.DEFAULT_GOAL := help

help:
	@echo "Команды:"
	@echo "  make build         - Собрать Docker образы"
	@echo "  make up            - Запустить сервисы"
	@echo "  make down          - Остановить сервисы"
	@echo "  make logs          - Логи всех сервисов"
	@echo "  make restart       - Перезапустить"
	@echo "  make clean         - Удалить всё"
	@echo "  make test          - Тест API"
	@echo "  make migrate-up    - Накатить миграции"
	@echo "  make migrate-down  - Откатить миграции"
	@echo "  make redis-cli     - Подключиться к Redis CLI"
	@echo ""

build:
	sudo docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down -v

logs:
	docker-compose logs -f

logs-api:
	docker-compose logs -f api-service

logs-db:
	docker-compose logs -f db-service

logs-postgres:
	docker-compose logs -f postgres

logs-redis:
	docker-compose logs -f redis

restart:
	docker-compose restart

restart-api:
	docker-compose restart api-service

restart-db:
	docker-compose restart db-service

rebuild:
	docker-compose down -v
	sudo docker-compose build --no-cache
	docker-compose up -d

rebuild-db:
	docker-compose build db-service
	docker-compose up -d db-service
clean:
	docker-compose down -v

psql:
	docker exec -it todolist-postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

redis-cli:
	docker exec -it todolist-redis redis-cli -a $(REDIS_PASSWORD)

shell-api:
	docker exec -it checklist-api-service sh

shell-db:
	docker exec -it checklist-db-service sh

migrate-up:
	docker exec -it todolist-db-service migrate -path=/app/migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DB)?sslmode=disable" up

migrate-down:
	docker exec -it todolist-db-service migrate -path=/app/migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DB)?sslmode=disable" down
