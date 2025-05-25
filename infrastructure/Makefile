.PHONY: up down migrate-up migrate-down new-migration psql redis-cli

up:
	docker compose up -d postgres redis

down:
	docker compose down

migrate-up:
	docker compose run --rm dbmate up

migrate-down:
	docker compose run --rm dbmate down

new-migration:
	mkdir -p migrations
	docker compose run --rm dbmate new $(name)

psql:
	docker exec -it chat_postgres psql -U postgres -d chat

redis-cli:
	docker exec -it chat_redis redis-cli

status:
	docker compose ps

logs:
	docker compose logs -f