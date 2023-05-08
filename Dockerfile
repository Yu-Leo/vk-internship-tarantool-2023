FROM golang:1.19

EXPOSE 8080

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/ -v ./...
RUN chmod a+x /usr/local/bin/app

CMD [ "/usr/local/bin/app" ]
