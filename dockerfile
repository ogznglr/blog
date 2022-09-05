FROM golang:alpine3.16
WORKDIR /go/src/goweb
RUN chmod 777 /go/src/goweb
COPY . .
CMD ["/go/src/goweb/blog"]