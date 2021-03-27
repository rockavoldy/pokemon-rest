FROM golang:1.15.7-alpine

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux go build -v -o restapi .

ARG DB_HOST
ENV DB_HOST=$DB_HOST
ENV GIN_MODE release
ENV HOST 0.0.0.0
ENV PORT 8080

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.0

RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

CMD dockerize -wait tcp://$DB_HOST:3306 -timeout 60m ./restapi