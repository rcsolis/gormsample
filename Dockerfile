FROM golang:1.20-buster as build

WORKDIR /app

COPY . ./

RUN make build

FROM debian:buster-slim as release
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN groupadd -r golang && useradd -r -g golang golang

COPY --from=build --chown=golang:golang /app/bin/gormsample .

CMD ["/gormsample"]