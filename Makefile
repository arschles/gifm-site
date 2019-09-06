.PHONY: start-db
start-db:
	docker-compose -p gifmsite up -d dev-env

.PHONY: stop-db
stop-db:
	docker-compose -p gifmsite down

.PHONY: docker
docker: docker-build docker-push

.PHONY: docker-build
docker-build:
	docker build -t arschles/gifm-site .

docker-run:
	docker run -p 3000:3000 -e SESSION_SECRET=localsecret arschles/gifm-site

.PHONY: docker-push
docker-push:
	docker push arschles/gifm-site

