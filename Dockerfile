FROM --platform=linux/amd64 golang:1.22-alpine AS build

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH

RUN apk add --no-cache make git

WORKDIR /go/src/github.com/awayfromserver/gobot
COPY go.mod /go/src/github.com/awayfromserver/gobot
COPY go.sum /go/src/github.com/awayfromserver/gobot

RUN --mount=type=cache,id=go-build-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/root/.cache/go-build \
    --mount=type=cache,id=go-pkg-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/go/pkg \
        go mod download -x

COPY . /go/src/github.com/awayfromserver/gobot

RUN --mount=type=cache,id=go-build-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/root/.cache/go-build \
    --mount=type=cache,id=go-pkg-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/go/pkg \
        make build
RUN mv bin/gobot* /bin/

# --------------------------------------------------------------------

FROM scratch AS gobot-linux

ARG VCS_REF
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

LABEL org.opencontainers.image.revision=$VCS_REF \
	org.opencontainers.image.source="https://github.com/awayfromserver/gobot"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/gobot_${TARGETOS}-${TARGETARCH}${TARGETVARIANT} /gobot

ENTRYPOINT [ "/gobot" ]

# --------------------------------------------------------------------

FROM alpine:3.20 AS gobot-alpine

ARG VCS_REF
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

LABEL org.opencontainers.image.revision=$VCS_REF \
	org.opencontainers.image.source="https://github.com/awayfromserver/gobot"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/gobot_${TARGETOS}-${TARGETARCH}${TARGETVARIANT} /bin/gobot

ENTRYPOINT [ "/bin/gobot" ]

# --------------------------------------------------------------------

FROM gobot-$TARGETOS AS gobot

ADD config.yaml .