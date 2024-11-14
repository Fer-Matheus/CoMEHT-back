FROM golang:1.23.2-bullseye AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12
COPY --from=builder /go/bin/app /
COPY .env /
CMD ["/app"]