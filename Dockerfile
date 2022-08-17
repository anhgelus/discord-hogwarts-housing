FROM golang:1.19-alpine

WORKDIR /app

ENV TOKEN = ""

COPY . .

RUN go build .

CMD ./discord-hogwarts-housing $TOKEN
