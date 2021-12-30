FROM golang:alpine as buildData

WORKDIR /usr/pub-sub-oracle

RUN apk add build-base
RUN apk add docker
RUN apk add linux-headers
RUN apk add nodejs npm
RUN apk add python3
RUN npm install -g truffle

COPY . .

RUN go build

EXPOSE 8080

CMD ["./OracleFactory"]