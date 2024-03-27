#!/bin/sh

first_names=("Ivan" "Sergey" "Anatholy" "Vladimir" "Maxim")
last_names=("Ivanov" "Sergeev" "Anatholyev" "Vladimirov" "Maximov")

for i in {1..10}
do
	rand_name=$(( RANDOM % 4 ))
	rand_last_name=$(( RANDOM % 4 ))
	curl -X 'POST' \
	'http://localhost:8080/api/v1/users' \
	-H 'accept: application/json' \
	-H 'Content-Type: application/json' \
	-d "{ \"full_name\": \"${first_names[rand_name]} ${last_names[rand_last_name]}\", \"password\": \"$(openssl rand -hex 5)\", \"username\": \"$(openssl rand -hex 5)\" }"
done
