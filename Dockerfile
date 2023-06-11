FROM docker.mirror.hashicorp.services/golang:alpine
LABEL maintainer="Nitro Agility S.r.l. Team <opensource@nitroagility.com>"

RUN apk add --no-cache git bash openssh

COPY . /app
WORKDIR /app

RUN /bin/bash ./scripts/build.sh

ENTRYPOINT ["./opsinsights-exporter"]