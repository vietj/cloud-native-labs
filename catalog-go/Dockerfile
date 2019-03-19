FROM golang:1.11

USER nobody

RUN mkdir -p /go/src/github.com/mcouliba/cloud-native-labs/catalog-go
WORKDIR /go/src/github.com/mcouliba/cloud-native-labs/catalog-go

COPY . /go/src/github.com/mcouliba/cloud-native-labs/catalog-go
RUN go build

CMD ["./catalog-go"]