FROM ubuntu:22.04

WORKDIR /app
COPY ./user ./

RUN chmod +x ./user && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates

CMD ["sh", "-c", "/app/user -f /app/etc/user.yaml"]