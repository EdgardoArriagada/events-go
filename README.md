WIP

```sh
# start
docker-compose up

# rebuild img
docker-compose build

# list events
curl localhost/events

# create events
curl -X POST localhost/events -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"name\": \"test event\", \"description\": \"a test event\", \"location\": \"a test location\", \"dateTime\": \"2025-01-01T15:30:00.000Z\"}"
```
