FROM postgres:16.2
ADD scripts/db-create.sql /docker-entrypoint-initdb.d/
