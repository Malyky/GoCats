FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build cats.go

CMD ["/myapp/cats"]
EXPOSE 8080
