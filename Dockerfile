FROM golang:1.9.2-alpine3.6

RUN apk add --no-cache git curl \
    && rm -rf /var/cache/apk/*

RUN set -x \
    # go get revel
    && go get -v github.com/revel/revel \
    && go get -v github.com/revel/cmd/revel \
    && go get -v gopkg.in/mgo.v2 \
    && go get -v github.com/kpawlik/geojson

ADD . src/github.com/PolytechLyon/cloud-project-equipe-8

ENTRYPOINT ["revel","run","github.com/PolytechLyon/cloud-project-equipe-8", "prod.server"]
