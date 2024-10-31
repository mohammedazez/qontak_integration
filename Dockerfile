FROM golang:1.22.0-alpine3.19 AS builder
LABEL stage-qontak_integration=builder

RUN apk update && apk add --no-cache git
WORKDIR /go/src/qontak_integration

COPY . .

#RUN go mod vendor
RUN go build -mod=readonly -o qontak_integration
RUN rm -rf vendor

FROM alpine:latest

ENV environment=local

RUN apk add --no-cache tzdata

RUN mkdir /app
WORKDIR /app

RUN apk add busybox-extras

EXPOSE 8888

COPY --from=builder /go/src/qontak_integration/qontak_integration /app
COPY --from=builder /go/src/qontak_integration/resource /app/resource

RUN mkdir logs

CMD /app/qontak_integration ${environment}