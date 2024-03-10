__DIR__ := $(realpath ./)

hello:
	@echo "init project cmd -> make startup"
	@echo "start docker cmd -> make dockerup" 
	@echo "stop docker cmd  -> make dockerdown"
	
dockerup:
	docker compose -f ./dev.docker-compose.yml up --build -d

dockerdown:
	docker compose -f ./dev.docker-compose.yml down -v

startup: 
	go run main.go
