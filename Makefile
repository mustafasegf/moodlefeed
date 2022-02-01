run:
	go run main.go

watch:
	air -c watcher.conf

build:
	go build -o ./bin/main main.go

up:
	docker-compose up -d
	docker-compose logs -f

down:
	docker-compose down

updb:
	docker-compose up -d db
	docker-compose logs -f

