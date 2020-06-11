FROM node:8.16.2-alpine3.9

WORKDIR /usr/src/app

COPY package.json ./

RUN npm install --only=prod

COPY . .

RUN [ "touch", "responseTimes.log" ]

EXPOSE 80

CMD [ "npm", "run", "start" ]
