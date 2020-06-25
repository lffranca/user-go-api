FROM golang:1.14.2-alpine AS build

RUN mkdir /app

ENV GO111MODULE=on
ENV GOPATH=/app

RUN mkdir -p $GOPATH/src/github.com/lffranca/suflex-api
RUN mkdir $GOPATH/bin
RUN mkdir $GOPATH/pkg

WORKDIR $GOPATH/src/github.com/lffranca/suflex-api
COPY . $GOPATH/src/github.com/lffranca/suflex-api


WORKDIR $GOPATH/src/github.com/lffranca/suflex-api
RUN go build -v -o $GOPATH/bin/suflex-api -ldflags '-extldflags "-static"' $GOPATH/src/github.com/lffranca/suflex-api/cmd/api/main.go
WORKDIR $GOPATH

RUN mkdir /build

RUN cp $GOPATH/bin/suflex-api /build

VOLUME $GOPATH/bin

# ALPINE
FROM alpine:3.10.0

RUN apk add --no-cache libc6-compat
RUN apk add --no-cache ca-certificates
RUN mkdir /app
COPY --from=build /app/bin/suflex-api /app/suflex-api
COPY ./production.env /app/.env
RUN chmod +x /app/suflex-api

WORKDIR /app

EXPOSE 8080

CMD [ "./suflex-api" ]
