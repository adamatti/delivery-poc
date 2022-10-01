FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app

FROM alpine as runner
WORKDIR /app
COPY --from=builder /app/app /app/
EXPOSE 80
CMD ["/app/app"]