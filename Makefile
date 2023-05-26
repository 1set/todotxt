SHELL=bash

test_pull:
	@docker pull alpine:latest
	@docker pull golang:1.18-alpine
	@docker pull golang:1.19-alpine
	@docker pull golang:alpine
	@docker pull golangci/golangci-lint:latest
	docker compose --file ./.github/docker-compose.yml pull

test_build: test_pull
	docker compose --file ./.github/docker-compose.yml build --no-cache

test: test_118 test_119 test_latest lint
test_118:
	@echo -n "* Running unit test on go 1.18 ... "
	@docker compose --file ./.github/docker-compose.yml run --rm v1_18
test_119:
	@echo -n "* Running unit test on go 1.19 ... "
	@docker compose --file ./.github/docker-compose.yml run --rm v1_19
test_latest:
	@echo -n "* Running unit test on latest go ... "
	@docker compose --file ./.github/docker-compose.yml run --rm latest

lint:
	@echo -n "* Running golangci-lint ... "
	@docker compose --file ./.github/docker-compose.yml run --rm lint && \
	echo "ok"

clean:
	@echo '! Cleaning up all docker images, containers, and networks (runs `docker system prune -a`)'
	@read -p "Are you sure?(press y or n) " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		echo ' ... Cleaning up ...'; \
		docker system prune -af && \
		echo 'Done! Run `make test_pull` to pull all images again'; \
	fi; \
	echo;
