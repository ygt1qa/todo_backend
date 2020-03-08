FROM golang:1.13.1-alpine3.10

RUN apk update && \
    apk upgrade && \
    apk add --update --no-cache ca-certificates git \
    bash \
    git

COPY . /go/src/github.com/ygt1qa/todo_backend

WORKDIR /go/src/github.com/ygt1qa/todo_backend/

RUN go mod download

RUN go build -o /go/bin/todo_backend ./cmd/main.go

CMD ["todo_backend"]