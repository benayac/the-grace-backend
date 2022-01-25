# BUILD STAGE
FROM golang:1.17-alpine AS builder
WORKDIR /workspace
COPY ./ ./

RUN go get -d -v ./...

# Build the API
ENV CGO_ENABLED = 0
ENV GOOS=linux
RUN go build -o /usr/local/bin/golang-docker ./cmd/api

# FINAL STAGE
FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/golang-docker /usr/local/bin/
#COPY --from=builder /workspace/config.json /
#COPY --from=builder /workspace/pkg/db/migrations/ /db/migrations/

RUN chown -R nobody:nogroup /usr/local/bin/golang-docker
USER nobody
EXPOSE 8080
CMD golang-docker