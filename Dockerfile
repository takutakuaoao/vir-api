FROM golang:1.23

WORKDIR /app

COPY ./app .

ENTRYPOINT ["top", "-b"]