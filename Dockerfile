
FROM golang:1.12-alpine3.9 AS minimum-base
LABEL MAINTAINER 'amaya <mail@sapphire.in.net>'

RUN set -eux && \
    apk add --no-cache \
        curl bash make git
WORKDIR /go/src/amaya382/go-api-server-template
EXPOSE 80

# ---------- #

FROM minimum-base AS debug-base
LABEL MAINTAINER 'amaya <mail@sapphire.in.net>'

RUN set -eux && \
    go get github.com/derekparker/delve/cmd/dlv
ENV CGO_ENABLED=0

# ---------- #

FROM minimum-base AS release-base
LABEL MAINTAINER 'amaya <mail@sapphire.in.net>'

ENV GIN_MODE=release
COPY . /go/src/amaya382/go-api-server-template
RUN set -eux && \
    make

# ---------- #

FROM scratch AS example
LABEL MAINTAINER 'amaya <mail@sapphire.in.net>'

COPY --from=release-base \
    /go/src/amaya382/go-api-server-template/example /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/example"]

