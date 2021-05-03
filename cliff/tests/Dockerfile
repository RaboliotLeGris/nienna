FROM golang:1.16.2-buster@sha256:5a6302e91acb152050d661c9a081a535978c629225225ed91a8b979ad24aafcd

RUN apt-get update && apt-get install make

WORKDIR /go/src/cliff
COPY . .

WORKDIR /go/src/cliff

CMD ["make", "test_exec"]