FROM golang

COPY . /app

WORKDIR /app

RUN go build -v ./...
RUN go install -v ./...

CMD [ "cli" ]