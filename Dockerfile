FROM golang:1.11.0-alpine as builder

RUN apk update && apk add git && apk add ca-certificates

WORKDIR /validation-admission-controllers-go

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/validation-admission-controllers-go

# Runtime image
FROM scratch AS base
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/validation-admission-controllers-go /bin/validation-admission-controllers-go
ENTRYPOINT ["/bin/validation-admission-controllers-go"]