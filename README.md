# Go Supervise
A lightweight GoLang service supervisor.

## Getting started
```shell
 git clone https://github.com/scottjr632/go-supervise.git
 cd go-supervise
 docker-compose up
```

### Registring a worker

Make an HTTP post request to supervisor that contains 
```json
{
	"workerId": "MY-APP",
	"name": "My Application",
	"checkUpUri": "http://validtesturi.com/health",
	"expectedRespone": "{\"status\": \"up\"}"
}
```

### Checking the worker's health
Make an HTTP get request to supervisor

`curl http://127.0.0.1:8080/api/workers/health?workerId=MY-APP`
