FROM golang:1.19-alpine

WORKDIR /app

ENV TOKEN = ""
ENV PREFIX = "!"
ENV REDIS_HOST = "redis"

COPY . .

RUN go build .

CMD ./discord-hogwarts-housing $TOKEN $PREFIX $REDIS_HOST
