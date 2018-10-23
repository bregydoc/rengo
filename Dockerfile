FROM golang:1.11

WORKDIR /go/src/rengo
COPY . .
RUN go get -v 
RUN go build
CMD ["rengo"]
