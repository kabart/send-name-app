FROM golang:1.15-alpine AS builder

WORKDIR /go/src/

COPY go.mod .
RUN go mod download

COPY /app/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/send-name-app

FROM scratch

COPY --from=builder /go/bin/send-name-app /go/bin/send-name-app
COPY /templates/ ./templates/

ENTRYPOINT ["/go/bin/send-name-app"]