FROM golang:1.13.1-alpine3.10

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    git

COPY . /go/src/github.com/ygt1qa/todo_backend

RUN go get github.com/go-sql-driver/mysql
RUN go get -u -t github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/volatiletech/null

WORKDIR /go/src/github.com/ygt1qa/todo_backend/

RUN go build -o /go/bin/todo_backend ./cmd/main.go

CMD ["todo_backend"]