# FROM mcr.microsoft.com/mssql/server:2017-latest
#
# ENV ACCEPT_EULA=Y
# ENV MSSQL_SA_PASSWORD=yourStrong()Password
#
# # Create a config directory
# RUN mkdir -p /usr/config
# WORKDIR /usr/config
#
# # Bundle config source
# COPY . /usr/config
#
# # Grant permissions for to our scripts to be executable
# RUN chmod +x /usr/config/entrypoint.sh
# RUN chmod +x /usr/config/configure-db.sh
#
# ENTRYPOINT ["./entrypoint.sh"]
FROM mcr.microsoft.com/mssql/server:2022-latest
ENV ACCEPT_EULA=Y
ENV MSSQL_SA_PASSWORD=yourStrong()Password

COPY ecuador.sql / 
COPY procedures.sql / 

# RUN /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P 'yourStrong()Password' -No -Q ':r /ecuador.sql'
# RUN /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P 'yourStrong()Password' -No -Q ':r /procedures.sql'
