FROM golang:1.11-alpine AS build-env

ENV GOPATH=/go

WORKDIR $GOPATH/src

RUN mkdir /release

ADD app ./app
ADD vendor ./vendor

RUN go build -o /release/server ./app/server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /release/server /app/

# You need to add root certificates to establish ssl connections.
# Otherwise, you gonna get the error like "transport: authentication handshake failed: x509: certificate signed by unknown authority".
# You can find root certificates in /etc/ssl/certs.
RUN apk add --no-cache ca-certificates

ENTRYPOINT ./server
