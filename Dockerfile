FROM --platform=linux/amd64 awayfromserver/upx:394 AS upx

FROM --platform=linux/amd64 golang:1.22-alpine AS build

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ENV GOOS=$TARGETOS GOARCH=$TARGETARCH

RUN apk add --no-cache make git

WORKDIR /go/src/github.com/awayfromserver/go-bot
COPY go.mod /go/src/github.com/awayfromserver/go-bot
COPY go.sum /go/src/github.com/awayfromserver/go-bot

RUN --mount=type=cache,id=go-build-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/root/.cache/go-build \
    --mount=type=cache,id=go-pkg-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/go/pkg \
        go mod download -x

COPY . /go/src/github.com/awayfromserver/go-bot

RUN --mount=type=cache,id=go-build-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/root/.cache/go-build \
    --mount=type=cache,id=go-pkg-${TARGETOS}-${TARGETARCH}${TARGETVARIANT},target=/go/pkg \
        make build
RUN mv bin/go-bot* /bin/

FROM --platform=linux/amd64 alpine:3.18 AS compress

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

RUN apk add --no-cache \
    make \
        libgcc libstdc++ ucl

ENV GOOS=$TARGETOS GOARCH=$TARGETARCH
WORKDIR /go/src/github.com/awayfromserver/go-bot
COPY Makefile .
RUN mkdir bin

COPY --from=upx /usr/bin/upx /usr/bin/upx
COPY --from=build bin/* bin/

RUN make compress
RUN mv bin/go-bot* /bin/

FROM scratch AS gobot-linux

ARG VCS_REF
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

LABEL org.opencontainers.image.revision=$VCS_REF \
	org.opencontainers.image.source="https://github.com/awayfromserver/go-bot"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/go-bot_${TARGETOS}-{${TARGETARCH}{${TARGETVARIANT} /go-bot

ENTRYPOINT [ "/go-bot" ]
