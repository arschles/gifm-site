.PHONY: start-db
start-db:
	docker-compose -p gifmsite up -d dev-env

.PHONY: stop-db
stop-db:
	docker-compose -p gifmsite down
