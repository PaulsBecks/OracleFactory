FROM node:alpine

RUN npm install http-server -g

COPY ./build/ ./

CMD ["http-server", "-p", "3000", "-a", "0.0.0.0"]