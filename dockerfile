FROM golang:alpine3.16
WORKDIR /go/src/goweb
COPY . .
CMD ["/go/src/goweb/blog","+rwx"]