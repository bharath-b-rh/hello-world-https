FROM quay.io/centos/centos:stream

WORKDIR /app

COPY bin/hello-world-https /app/

CMD ["/app/hello-world-https"]

EXPOSE 8082 8083
