FROM golang
ADD . /app
WORKDIR /app/src
RUN go build 
ENTRYPOINT ./src
EXPOSE 8080