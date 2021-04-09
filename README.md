# rest-client

A sample REST service, which is reading a json database and sending it to port domain service via GRPC. The read is done in background job and do not block the REST handlers.

You can run 

```
go build -o main .
```

And after that

```
./main
```

A HTTP service will be started on port 10000

If you want to change the hostname and other parameters you can do

```
./main --httpAddress=localhost:10001 --serviceAddress=localhost:6667 --databaseFilename=./data.json
```