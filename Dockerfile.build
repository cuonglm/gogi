FROM golang:alpine
MAINTAINER "Cuong Manh Le <cuong.manhle.vn@gmail.com>"

RUN apk update && \
    apk add git build-base && \
    rm -rf /var/cache/apk/* && \
    go get -v -d github.com/cuonglm/gogi && \
    cd "$GOPATH/src/github.com/cuonglm/gogi/gogi" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o /gogi

COPY Dockerfile.run /
CMD tar -cf - -C / Dockerfile.run gogi
