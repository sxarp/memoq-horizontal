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
ENTRYPOINT ./server
