FROM alpine:3.12

LABEL org.opencontainers.image.source="https://github.com/g1ibby/eth-address-generator"

WORKDIR /app

COPY . .

ENTRYPOINT [ "bin/eth-address-generator" ]
