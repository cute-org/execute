# Etap 1 – build Vite
FROM node:18-alpine AS build
WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Etap 2 – nginx
FROM nginx:alpine

# Skasuj domyślną konfigurację nginx
RUN rm /etc/nginx/conf.d/default.conf

# Wrzuć własną konfigurację
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Skopiuj zbudowany frontend
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
