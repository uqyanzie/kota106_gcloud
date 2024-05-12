FROM golang:1.16.3-alpine3.13 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

# Path: Dockerfile

FROM alpine:3.13

COPY --from=builder /app/main /app/main

CMD ["/app/main"]

