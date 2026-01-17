include .env
export

service-run:
	@go run main.go
migrate-up:
	migrate -path migrations -database ${DATABASE_URL} up
migrate-down:
	migrate -path migrations -database ${DATABASE_URL} down

service-deploy:
	docker compose up todolist
service-undeploy:
	docker compose down todolist
service-build:
	docker build -t todolist .