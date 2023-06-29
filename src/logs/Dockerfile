FROM node:14

WORKDIR /usr/app

COPY package.json package.json
COPY yarn.lock yarn.lock

RUN yarn

WORKDIR /usr/app/ui

COPY ui/package.json package.json
COPY ui/yarn.lock yarn.lock

RUN yarn

WORKDIR /usr/app

COPY . .

WORKDIR /usr/app/ui

RUN yarn run build

WORKDIR /usr/app

RUN yarn run build

CMD [ "yarn", "run", "dev" ]