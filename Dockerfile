FROM golang:alpine3.15

WORKDIR /app

COPY . .

CMD sleep 1d