build:
	docker build -t call-billing:1.0.0 --progress=plain .
deploy:
	docker compose down
	docker compose up -d
remove:
	docker compose down