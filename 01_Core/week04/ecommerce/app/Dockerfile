FROM golang:latest

RUN mkdir app

WORKDIR /app

RUN \ 
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gorilla/mux && \
    go get github.com/joho/godotenv

ADD ./ ./

CMD ["go", "run", "main.go"]

