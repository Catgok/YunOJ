FROM ubuntu:22.04

WORKDIR /app
COPY ./gateway ./
RUN chmod +x ./gateway && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates

CMD ["sh", "-c", "/app/gateway -f /app/etc/gateway.yaml"]