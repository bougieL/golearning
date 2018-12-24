FROM mysql:5.5

ENV MYSQL_ROOT_PASSWORD=root DB_MYSQL_HOST=host DB_MYSQL_PORT=port DB_MYSQL_USERNAME=username DB_MYSQL_PASSWORD=password DB_MYSQL_NAME=database
RUN mysqldump -h ${DB_MYSQL_HOST} -u ${} -p ${DB_MYSQL_NAME} > /${DB_MYSQL_NAME}.sql
# VOLUME [ "/data" ]
RUN mysql -uroot -proot golearning < ../docs/sql.sql