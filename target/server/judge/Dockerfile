FROM ubuntu:22.04

WORKDIR /app
COPY ./judge ./
RUN chmod +x ./judge && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates

CMD ["sh", "-c", "/app/judge -f /app/etc/judge.yaml"]