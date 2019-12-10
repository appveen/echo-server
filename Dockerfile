FROM node:10-alpine

WORKDIR /app

COPY main /app/.

EXPOSE 31000

CMD [ "./main"]