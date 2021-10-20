run:
	go run main.go

watch:
	air -c watcher.conf

build:
	go build -o ./bin/main main.go

up:
	docker-compose -f docker-compose-prod.yml up

upb:
	docker-compose -f docker-compose-prod.yml up --build

down:
	docker-compose -f docker-compose-prod.yml down