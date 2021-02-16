run-go:
	go run cmd/server/main.go

up:
	sudo docker-compose up -d

run-docker:
	sudo docker-compose up

run-build:
	sudo docker-compose up --build

down:
	sudo docker-compose down