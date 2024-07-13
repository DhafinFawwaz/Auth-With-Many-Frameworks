FROM golang:latest

RUN apt-get update && \
    apt-get install -y mariadb-client mariadb-server && \
    service mariadb start && \
    mysql -e "CREATE DATABASE IF NOT EXISTS kampus;"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_USERNAME=user
ENV DB_PASSWORD=supersecretpassword
ENV DB_NAME=kampus
ENV MYSQL_ROOT_PASSWORD=supersecretpassword
ENV MYSQL_DATABASE=kampus
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=supersecretpassword

EXPOSE 8080

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
