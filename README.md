WIP

```sh
# start
docker-compose up

# rebuild img
docker-compose build

# list events
curl localhost/events | jq

# create events
curl -X POST localhost/events  -d '{ "name": "test event", "description": "a test event", "location": "a test location", "dateTime": "2025-01-01T15:30:00.000Z"}' -H "accept: application/json" -H "Content-Type: application/json" | jq

# update events
curl -X PUT localhost/events/1  -d '{ "name": "cuando despiertes un dia (editado)", "description": "a test event 2", "location": "teatro caupolican", "dateTime": "2024-01-01T15:30:00.000Z"}'  -H "accept: application/json" -H "Content-Type: application/json" | jq

# create user
curl -X POST localhost/signup -d '{"email": "test@example.com", "password": "trustno1"}' -H "Content-Type: application/json" -H "accept: application/json" | jq

# delete user
curl -X DELETE localhost/events/1 | jq

# login user
curl -X POST localhost/login -d '{"email": "test@example.com", "password": "trustno1"}' -H "Content-Type: application/json" -H "accept: application/json" | jq
```
