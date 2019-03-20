FROM golang:1.11

RUN sudo mkdir /app
WORKDIR /app
COPY . /app
RUN go build -o catalog-go .

EXPOSE 8080 8888
USER 1001

CMD ["/app/catalog-go"]