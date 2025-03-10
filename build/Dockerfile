###############################################
FROM --platform=$BUILDPLATFORM golang:1.24@sha256:c5adecdb7b3f8c5ca3c88648a861882849cc8b02fed68ece31e25de88ad13418 AS builder

WORKDIR /app

COPY . .

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH

RUN go build -o bin/api cmd/api/main.go

###############################################
FROM busybox AS busybox
RUN addgroup -S appgroup && adduser -S -G appgroup appuser

###############################################
FROM debian:stable-slim AS certs
RUN apt-get update && apt-get install -y ca-certificates

###############################################
FROM scratch

WORKDIR /app

COPY --from=builder /app/bin/api .

COPY --from=busybox /etc/passwd /etc/passwd
COPY --from=busybox /etc/group /etc/group

COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=certs /etc/ssl/private /etc/ssl/private
COPY --from=certs /usr/share/ca-certificates /usr/share/ca-certificates

USER appuser:appgroup

ENTRYPOINT ["/app/api"]