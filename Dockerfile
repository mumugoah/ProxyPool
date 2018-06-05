FROM golang

COPY . /go/src/github.com/mumugoah/ProxyPool
WORKDIR /go/src/github.com/mumugoah/ProxyPool

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go get ./ \
    && go build

EXPOSE 8080

CMD ["./ProxyPool"]