FROM golang:latest

RUN apt-get update && apt-get install -y mariadb-client

WORKDIR /app

COPY . .

ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_USERNAME=root
ENV DB_PASSWORD=supersecretpassword
ENV DB_NAME=kampus

EXPOSE 8080

CMD ["go", "run", "main.go"]
