swag-user:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./users
	mv ./docs/docs.go ./docs/users
	mv ./docs/swagger.* ./docs/users

swag-conferences:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./conferences 
	mv ./docs/docs.go ./docs/conferences
	mv ./docs/swagger.* ./docs/conferences

swag-gateway:
	swag i -g swagger_info.go --parseInternal --parseDependency --dir ./api-gateway 
	mv ./docs/docs.go ./docs/apigateway
	mv ./docs/swagger.* ./docs/apigateway

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

build-gateway:
	go build -o /dev/null ./cmd/api-gateway/

prune:
	docker system prune -f --all