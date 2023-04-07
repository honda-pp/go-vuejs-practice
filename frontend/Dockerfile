FROM node:18-alpine

WORKDIR /app

RUN apk add --no-cache git
RUN npm install -g @vue/cli

COPY . .

RUN npm install

EXPOSE 8080

CMD ["npm", "run", "serve"]