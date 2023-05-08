FROM golang:1.19

EXPOSE 8080

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY config /app/config
COPY internal /app/internal
COPY main.go /app

RUN go build -o /go-server

CMD [ "/go-server" ]
