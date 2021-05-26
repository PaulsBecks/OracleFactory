FROM golang:alpine as buildData

WORKDIR /usr/oracle-factory

RUN apk add build-base
RUN apk add docker
RUN apk add linux-headers

COPY . .

RUN go build

EXPOSE 8080

CMD ["./OracleFactory"]