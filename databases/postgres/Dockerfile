FROM postgres:latest

ENV POSTGRES_PASSWORD=mysecretpasswordvalue
COPY ecuador.sql /docker-entrypoint-initdb.d/1.sql
COPY procedures.sql /docker-entrypoint-initdb.d/2.sql
