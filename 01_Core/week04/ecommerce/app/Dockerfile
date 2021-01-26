FROM golang:latest

RUN mkdir app

WORKDIR /app

RUN \ 
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gorilla/mux

ADD ./ ./

CMD ["go", "run", "main.go"]

