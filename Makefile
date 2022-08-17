restart:
	docker compose down
	docker compose build
	docker compose up -d
stop:
	docker compose down
start:
	docker compose build
	docker compose up -d