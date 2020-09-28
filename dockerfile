FROM golang:1.12.15

WORKDIR /
COPY . .

ENV CMD=$CMD

RUN go build -o bin/parkinglot

CMD /bin/parkinglot $CMD