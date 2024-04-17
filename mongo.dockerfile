FROM mongo:7.0-rc 
ADD scripts/db-create-mongo.sh /docker-entrypoint-initdb.d/
