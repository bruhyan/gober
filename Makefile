APP_NAME=gober
DOCKER_IMAGE=${APP_NAME}:latest

docker-build:
	@echo "Building Docker image..."
	docker build --tag $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run --publish 8080:8080 gober