FROM node:18

WORKDIR /app

COPY ./ ./

RUN npm install

RUN chmod +x run.sh

CMD ["./run.sh"]