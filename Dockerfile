FROM golang:alpine AS builder
RUN apk --no-cache add build-base
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
COPY app app/
RUN go build -o stori_challenge

FROM alpine:3.19
COPY --from=builder /app/stori_challenge stori_challenge
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
USER nonroot
COPY txns1.csv txns1.csv
COPY txns2.csv txns2.csv
ENTRYPOINT ["/stori_challenge"]