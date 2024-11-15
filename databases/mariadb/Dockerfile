FROM mariadb:latest

ENV MARIADB_ROOT_PASSWORD=123
COPY ecuador.sql /docker-entrypoint-initdb.d/1.sql
COPY procedures.sql /docker-entrypoint-initdb.d/2.sql
