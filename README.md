goroutines-server

# Development

## Run with

`docker-compose up`

## Send requests with curl

e.g.

`curl -X POST -H "content-type: application/json" http://localhost:8080/create-worker -d '{"granularity":"second", "frequency": 1}'`

`curl -X POST -H "content-type: application/json" http://localhost:8080/kill-worker -d '{"id": 1}'`

`curl -X POST -H "content-type: application/json" http://localhost:8080/get-entries -d '{"id": 1}'`

`curl -X POST -H "content-type: application/json" http://localhost:8080/get-routines'`
