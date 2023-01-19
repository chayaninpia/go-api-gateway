# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache git 

ADD . /app
WORKDIR /app

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /go-api-gateway 

EXPOSE 30001

FROM alpine

COPY conf /app/conf
WORKDIR /app
COPY --from=build /go-api-gateway /app/go-api-gateway 
RUN apk update && apk add nano 

CMD [ "/app/go-api-gateway" ]
