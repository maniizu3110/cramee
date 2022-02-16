#TODO:余計なファイルまで本番環境にある状態なのでステージ分けて削除する
FROM node:16-alpine3.12
WORKDIR '/app'
COPY ./package.json ./
COPY ./nuxt.config.js ./
ENV HOST=0.0.0.0

RUN npm install
COPY . .
RUN npm run build
CMD ["npm","run","start"]