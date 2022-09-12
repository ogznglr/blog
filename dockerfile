FROM golang:alpine3.16
WORKDIR /go/src/goweb

COPY . .
RUN chmod 777 /go/src/goweb/blog
CMD ["/go/src/goweb/blog"]