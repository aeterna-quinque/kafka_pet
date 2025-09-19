COMPOSE_FILE := compose.yaml

.PHONY: build
build:
	docker-compose -f $(COMPOSE_FILE) build

.PHONY: up
up:
	docker-compose -f $(COMPOSE_FILE) up

.PHONY: start
start:
	docker-compose -f $(COMPOSE_FILE) start

.PHONY: stop
stop:
	docker-compose -f $(COMPOSE_FILE) stop

.PHONY: down
down:
	docker-compose -f $(COMPOSE_FILE) down

.PHONY: clear
clear:
	docker-compose -f $(COMPOSE_FILE) down -v --rmi all

.PHONY: restart
restart:
	docker-compose -f $(COMPOSE_FILE) restart

.PHONY: ps
ps:
	docker-compose -f $(COMPOSE_FILE) ps
