FROM node:18-alpine

RUN npm install -g http-server

WORKDIR /app


COPY package*.json ./
RUN npm install


COPY . .

EXPOSE 5173
CMD [ "npm", "run", "dev"]