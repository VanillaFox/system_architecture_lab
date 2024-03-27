swag-user:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./users

run:
	docker-compose up --build --no-deps --force-recreate -d

stop: 
	docker-compose down

db-fill:
	sh scripts/db-fill.sh