FROM golang:1.12

ENV GOSRC ${GOPATH}/src

RUN mkdir -p ${GOSRC}/catalog-go
WORKDIR ${GOSRC}/catalog-go
COPY . ${GOSRC}/catalog-go
RUN go build

CMD ["./catalog-go"]