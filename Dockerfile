FROM golang:alpine AS builder
RUN apk add build-base
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o stori_challenge

FROM alpine
COPY --from=builder /app/stori_challenge stori_challenge
ENV EMAIL_PUBLIC_API_KEY=
ENV EMAIL_PRIVATE_API_KEY=
COPY txns1.csv txns1.csv
COPY txns2.csv txns2.csv
ENTRYPOINT ["/stori_challenge"]