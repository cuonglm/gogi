FROM alpine:latest
MAINTAINER "Cuong Manh Le <cuong.manhle.vn@gmail.com>"

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*

COPY gogi /bin/gogi

ENTRYPOINT ["/bin/gogi"]
