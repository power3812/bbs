FROM node:16.13.2-alpine

EXPOSE 3000
ENV HOST 0.0.0.0

RUN apk update && \
  apk add git python3 make g++ && \
  npm install -g @vue/cli nuxt create-nuxt-app && \
  yarn install
