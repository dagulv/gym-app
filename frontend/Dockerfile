FROM node:16.15.0-alpine3.15
RUN mkdir -p /app/frontend
WORKDIR /app/frontend
COPY package.json package-lock.json ./
RUN npm install -g npm@latest
RUN npm install
COPY . .
RUN npm run build
CMD npm rebuild esbuild && npm run dev -- --host