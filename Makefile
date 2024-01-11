DC := docker-compose -f ./docker-compose.yml

all:
	@mkdir -p /home/postgres-data
	@$(DC) up

down:
	@$(DC) down

re: clean all

clean:
	@$(DC) down
	@rm -rf /home/postgres-data
	docker system prune -a -f

db:
	@mkdir -p /home/postgres-data
	@$(DC) up -d --build postgres

app:
	@mkdir -p /home/postgres-data
	@$(DC) up -d --build app

.PHONY: all down re clean db app