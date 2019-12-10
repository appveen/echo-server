FROM scratch

WORKDIR /app

COPY main /app/.

EXPOSE 31000

CMD [ "./main"]