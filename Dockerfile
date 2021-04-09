FROM golang:alpine
RUN mkdir /app 
ADD . /app/
VOLUME ["/app"]
WORKDIR /app 
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ./main --serviceAddress=$SERVICE_ADDRESS --httpAddress=$HTTP_ADDRESS --databaseFilename=$DATABASE_FILE
EXPOSE 10000