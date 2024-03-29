.PHONY: start-db
start-db:
	docker-compose -p gifmsite up -d dev-env
	sleep 5
	buffalo db migrate up

.PHONY: stop-db
stop-db:
	docker-compose -p gifmsite down

DOCKER_NAME := docker.io/arschles/gifm-site

.PHONY: docker-name
docker-name:
	@echo ${DOCKER_NAME}

.PHONY: docker
docker: docker-build docker-push

.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_NAME} .

docker-run:
	docker run -p 3000:3000 -e SESSION_SECRET=localsecret ${DOCKER_NAME}

.PHONY: docker-push
docker-push:
	docker push ${DOCKER_NAME}
