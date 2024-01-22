FROM golang:alpine AS builder
RUN apk --no-cache add build-base
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY go.work.docker go.work
COPY go.work.sum .
COPY app/go.mod app/go.mod
COPY app/go.sum app/go.sum
RUN go mod download
COPY main.go .
COPY app app/
RUN go build -o stori_challenge

FROM alpine:3.19
COPY --from=builder /app/stori_challenge stori_challenge
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
USER nonroot
ENTRYPOINT ["/stori_challenge"]