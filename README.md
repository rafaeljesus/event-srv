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

## Stack
- Golang
- Elastic Search
- Kafka
- Docker

## Docker
This repository has automated image builds on hub.docker.com after successfully building and testing. See the `deployment` section of [circle.yml](circle.yml) for details on how this is done. Note that three environment variables need to be set on CircleCI for the deployment to work:

  * DOCKER_EMAIL - The email address associated with the user with push access to the Docker Hub repository
  * DOCKER_USER - Docker Hub username
  * DOCKER_PASS - Docker Hub password (these are all stored encrypted on CircleCI, and you can create a deployment user with limited permission on Docker Hub if you like)

```bash
$ sh buid-container
$ docker run -it -t -p 3000:3000 --name event-tracker rafaeljesus/event-tracker
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![CircleCI](https://circleci.com/gh/rafaeljesus/event-tracker.svg?style=svg)](https://circleci.com/gh/rafaeljesus/event-tracker)
[![](https://images.microbadger.com/badges/image/rafaeljesus/event-tracker.svg)](https://microbadger.com/images/rafaeljesus/event-tracker "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/rafaeljesus/event-tracker.svg)](https://microbadger.com/images/rafaeljesus/event-tracker "Get your own version badge on microbadger.com")

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Twitter [@rafaeljesus](https://twitter.com/_jesus_rafael)
