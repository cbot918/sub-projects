FROM node:19-alpine3.16 as build

WORKDIR /app

COPY package* .
RUN npm install --production

FROM node:19-alpine3.16 as run

WORKDIR /app

COPY . .
COPY --from=build /app/node_modules /app/node_modules

ENTRYPOINT ["npm","start"]