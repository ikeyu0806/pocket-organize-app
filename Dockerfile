FROM golang:latest

RUN mkdir -p -m 775 /workspace/pocket_app

WORKDIR /workspace/pocket_app

ADD ./ /workspace/pocket_app

RUN ["go", "build"]
CMD ["go", "run", "./main.go"]
