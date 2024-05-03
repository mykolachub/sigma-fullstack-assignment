FROM node:20-alpine

WORKDIR /app

COPY . .

RUN npm ci

RUN npm run build

EXPOSE 3333

CMD [ "npx", "serve", "build", "-p", "3333" ]
