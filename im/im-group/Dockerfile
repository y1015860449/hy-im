FROM alpine:3.11

WORKDIR /data
COPY im-group ./app
COPY config*.yaml ./

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD ["/data/app"]