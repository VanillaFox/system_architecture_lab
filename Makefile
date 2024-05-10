swag-user:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./users
	mv ./docs/docs.go ./docs/users
	mv ./docs/swagger.* ./docs/users

swag-conferences:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./conferences 
	mv ./docs/docs.go ./docs/conferences
	mv ./docs/swagger.* ./docs/conferences

run:
	docker-compose up --build --no-deps --force-recreate -d

stop: 
	docker-compose down

db-fill:
	sh scripts/db-fill.sh

mongo-db-fill:
	sh scripts/db-fill-mongo.sh

build-conferences:
	go build -o /dev/null ./cmd/conferences/

build-user:
	go build -o /dev/null ./cmd/users/