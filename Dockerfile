FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>

ADD trace-srv /trace-srv

ENV TRACE_SRV_PORT="3000"
ENV TRACE_SRV_DB="http://@docker:9200"
ENV TRACE_SRV_BUS="localhost:9093"

ENTRYPOINT ["/trace-srv"]
