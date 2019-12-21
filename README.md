# Go Supervise
A lightweight GoLang service supervisor.

## Getting started
### Starting a production server
```shell
 mkdir go-supervise && cd $_
 wget https://raw.githubusercontent.com/scottjr632/go-supervise/master/docker-compose.yml
```
#### Create server.config.yml file
```shell
 touch server.config.yml
 vi server.config.yml
```
##### example server config file
```yml
server:
  port: 11223

services:
  checkupservice:
    interval: 35 
    uniqueids: true

datastore:
  dbname: go-service
  connectionstring: mongodb://database:27017

jwt:
  tokenname: access_token
  protectedpath: /api/protected
```

#### Uncomment from docker-compose to include web client
```yml
# web:
#   image: docker.pkg.github.com/scottjr632/go-supervise-client/client:latest 
#   restart: unless-stopped
#   links:
#     - "server:server"
#   ports:
#     - "11221:80"
#   depends_on:
#     - server
 ```
### Start the services
```bash
 docker-compose up -d
```

### Starting a dev server
```shell
 git clone https://github.com/scottjr632/go-supervise.git
 cd go-supervise
 docker-compose up
```

## Registring a worker

Make an HTTP post request to supervisor that contains 
```json
{
	"workerId": "MY-APP",
	"name": "My Application",
	"checkUpUri": "http://validtesturi.com/health",
	"expectedRespone": "{\"status\": \"up\"}"
}
```

## Checking the worker's health
Make an HTTP get request to supervisor

`curl http://127.0.0.1:8080/api/workers/health?workerId=MY-APP` 
  
Response
```json
{
    "status": "Worker is cloudy",
    "workerId": "MY-APP",
    "name": "My Application",
    "checkUpUri": "https://landing-page.mcserver.staging.scottrichardson.dev/",
    "expectedRespone": ""
}
```
To include checkups
  
`curl http://127.0.0.1:8080/api/workers/health?workerId=MY-APP-PAGE&ic=true` 
  
Response
```json
{
    "status": "Worker is stormy",
    "checkUps": [
        {
            "worker": {
                "workerId": "MY-APP",
                "name": "My Application",
                "checkUpUri": "https://landing-page.mcserver.staging.scottrichardson.dev/",
                "expectedRespone": ""
            },
            "actualResponse": "",
            "responseCode": "200 OK"
        },
        {
            "worker": {
                "workerId": "MY-APP",
                "name": "My Application",
                "checkUpUri": "https://landing-page.mcserver.staging.scottrichardson.dev/",
                "expectedRespone": ""
            },
            "actualResponse": "",
            "responseCode": "200 OK"
        }
    ],
    "workerId": "MY-APP",
    "name": "My Application",
    "checkUpUri": "https://landing-page.mcserver.staging.scottrichardson.dev/",
    "expectedRespone": ""
}

