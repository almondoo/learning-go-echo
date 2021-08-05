#-----------------------------------------------------------
# 初期セットアップ
#-----------------------------------------------------------
start-up:
	@make build
	@make up

#-----------------------------------------------------------
# 個々のコマンド
#-----------------------------------------------------------
build:
	docker-compose build

build-no-cache:
	docker-compose build --no-cache

up:
	docker-compose up -d redis mysql phpmyadmin
	@make up-w

up-w:
	docker-compose up workspace

up-d:
	docker-compose up -d mysql phpmyadmin redis workspace

down:
	docker-compose down

restart:
	@make down
	@make up

bash:
	docker-compose exec workspace bash

redis-bash:
	docker-compose exec redis ash

ps:
	docker-compose ps

logs-c:
	docker-compose logs

logs-f:
	docker-compose logs -f
