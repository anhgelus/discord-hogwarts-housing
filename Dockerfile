FROM golang:1.19-alpine

WORKDIR /app

ENV TOKEN = ""
ENV PREFIX = "!"

COPY . .

RUN go build .

CMD ./discord-hogwarts-housing $TOKEN $PREFIX
