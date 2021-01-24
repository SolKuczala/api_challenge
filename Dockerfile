FROM golang

COPY . /app

WORKDIR /app

RUN go build ./...
RUN go install ./...

CMD [ "cli" ]