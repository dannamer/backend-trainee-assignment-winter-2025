FROM golang:1.22

WORKDIR ${GOPATH}/avito-shop/
COPY . ${GOPATH}/avito-shop/

RUN go build -o /build ./cmd/service \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]