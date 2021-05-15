FROM golang:latest

RUN mkdir -p -m 775 /usr/local/go/src/pocket-organize-app
RUN mkdir -p -m 775 /usr/local/go/src/pocket-organize-app/dist/js

WORKDIR /usr/local/go/src/pocket-organize-app

ADD ./go/src/pocket-organize-app /usr/local/go/src/pocket-organize-app

RUN ["go", "build", "./"]
CMD ["go", "run", "./main.go"]
