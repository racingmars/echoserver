FROM golang:latest

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . .
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD ["go-wrapper", "run"]
