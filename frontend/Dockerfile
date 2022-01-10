# build stage
FROM node:14-alpine AS builder

WORKDIR /usr/app/
COPY . /usr/app/

RUN yarn install
RUN yarn build
RUN npm install http-server -g

EXPOSE 3000

CMD ["http-server", "./build/", "-p", "3000", "-a", "0.0.0.0"]