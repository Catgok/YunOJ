FROM ubuntu:22.04

WORKDIR /app
COPY ./contest ./
RUN chmod +x ./contest && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates

CMD ["sh", "-c", "/app/contest -f /app/etc/contest.yaml"]