up:
	docker-compose up -d
down:
	docker-compose down
start-client:
	go run ./client/...
start-worker:
	go run ./worker/...
start-monitoring:
	go run ./monitoring/...
redis-cli:
	docker-compose exec redis redis-cli
