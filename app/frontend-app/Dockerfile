FROM node:alpine3.18 AS builder

RUN npm install -g npm@10.1.0
 
WORKDIR /app
COPY package.json package-lock.json ./
COPY public/ public/
COPY src/ src/

RUN yarn install
RUN export NODE_OPTIONS=--openssl-legacy-provider && yarn build

# For lua integration
# FROM fabiocicerchia/nginx-lua
# COPY nginx.conf /etc/nginx/conf.d/default.conf

FROM nginx:1.25.2

COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY --from=builder /app/build /usr/share/nginx/html

RUN touch /var/run/nginx.pid
RUN chown -R nginx:nginx /var/run/nginx.pid /usr/share/nginx/html /var/cache/nginx /var/log/nginx /etc/nginx/conf.d

# USER nginx
EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]