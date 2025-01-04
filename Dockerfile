FROM golang:1.23-alpine3.21
LABEL authors="osami"

WORKDIR /app

RUN go mod download
RUN go build -o floof cmd/floof-db/main.go

ENTRYPOINT ["top", "-b"]

CMD ./floof