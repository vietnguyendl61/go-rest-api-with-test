test: docker_build_test
	docker-compose up -d
	docker-compose exec -T user-service go test ./...
	docker-compose down

docker_build_test:
	docker build . -t service_test --target=test