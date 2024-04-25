FROM ubuntu:22.04

WORKDIR /app
COPY ./problem ./
RUN chmod +x ./problem && \
    apt-get -qq update && \
    apt-get -qq install -y --no-install-recommends ca-certificates

CMD ["sh", "-c", "/app/problem -f /app/etc/problem.yaml"]