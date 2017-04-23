## Event Service

* An event log as a service
* Build with Go, Elastic Search and Kafka
* A minimal docker container
* Automatically pushes it to dockerhub if tests pass

## Setup
Env vars
```bash
export PORT=3000
export DATASTORE_URL=http://@docker:9200
export BROKER_URL=localhost:9093
```

Installation
```sh
mkdir -p $GOPATH/src/github.com/rafaeljesus
cd $GOPATH/src/github.com/rafaeljesus
git clone https://github.com/rafaeljesus/event-srv.git
cd event-srv
make all
```

## Running server
```
./dist/event-srv
# => Starting Event Service at port 3000
```

### Create a Event through HTTP
- Request
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"name": "order_created", "status": "success", "payload": {}}' \
localhost:3000/events
```

- Response
```
"OK"
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![CircleCI](https://circleci.com/gh/rafaeljesus/event-srv.svg?style=svg)](https://circleci.com/gh/rafaeljesus/event-srv)
[![](https://images.microbadger.com/badges/image/rafaeljesus/event-srv.svg)](https://microbadger.com/images/rafaeljesus/event-srv "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/event-srv.svg)](https://microbadger.com/images/rafaeljesus/event-srv "Get your own version badge on microbadger.com")

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Twitter [@rafaeljesus](https://twitter.com/_jesus_rafael)
