FROM node:22.3-alpine3.19

WORKDIR /app

COPY . .

RUN npm install -g pnpm
RUN pnpm i

RUN pnpm run build

CMD ["pnpm", "run", "serve", "--port=3000"]