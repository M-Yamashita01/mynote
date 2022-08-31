FROM node:14-alpine3.16

WORKDIR /app

RUN apk update

COPY . /app

RUN npm install

EXPOSE 3000

CMD ["npm", "run", "dev"]
