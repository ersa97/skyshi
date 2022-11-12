FROM mariadb:10
COPY ./migrations/*.up.sql /docker-entrypoint-initdb.d/
CMD ["mysqld"]