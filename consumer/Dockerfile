ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk --no-cache add tzdata && apk add alpine-sdk git ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# Build the application -ldflags '-w -extldflags "-static"'
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o ./main ./main.go

# Run vulnerabilyty scan on build image
# FROM builder AS vulnscan
# COPY --from=aquasec/trivy:latest /usr/local/bin/trivy /usr/local/bin/trivy
# RUN trivy rootfs --no-progress /

# Build a small image
FROM alpine:latest
RUN apk update
RUN apk add tzdata
RUN apk add libcrypto1.1
RUN apk add libssl1.1
RUN apk upgrade --available
RUN rm -rf /var/cache/apk/*

COPY --from=builder /build/config/config.yaml /config
COPY --from=builder /build/main /
ENV TZ=Asia/Bangkok
ENV RUN_MODE=dev

# Command to run
ENTRYPOINT ["/main"]