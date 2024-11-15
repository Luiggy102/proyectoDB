FROM mcr.microsoft.com/mssql/server:2022-latest

ENV ACCEPT_EULA=Y
ENV MSSQL_SA_PASSWORD=123
COPY ecuador.sql /docker-entrypoint-initdb.d/1.sql
COPY procedures.sql /docker-entrypoint-initdb.d/2.sql
