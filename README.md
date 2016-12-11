## Trace Srv

* Record any actions your system perform, along with any properties that describe the action
* A minimal docker container
* Automatically pushes it to dockerhub if tests pass

## Setup
Env vars
```bash
export TRACE_SRV_PORT=3000
export TRACE_SRV_DB=http://@docker:9200
export TRACE_SRV_BUS=localhost:9093
```

mkdir -p $GOPATH/src/github.com/rafaeljesus
cd $GOPATH/src/github.com/rafaeljesus
git clone https://github.com/rafaeljesus/trace-srv.git
cd trace-srv
glide install
go build

## Running server
```
./trace-srv
# => Starting Trace Service at port 3000
```

### Trace a Event through HTTP
- Request
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"name": "order_created", "status": "success", "payload": {}}' \
localhost:3000/v1/events
```

- Response
```json
{
  "ok": true
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![CircleCI](https://circleci.com/gh/rafaeljesus/trace-srv.svg?style=svg)](https://circleci.com/gh/rafaeljesus/trace-srv)
[![](https://images.microbadger.com/badges/image/rafaeljesus/trace-srv.svg)](https://microbadger.com/images/rafaeljesus/trace-srv "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/trace-srv.svg)](https://microbadger.com/images/rafaeljesus/trace-srv "Get your own version badge on microbadger.com")

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Twitter [@rafaeljesus](https://twitter.com/_jesus_rafael)
