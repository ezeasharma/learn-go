FROM golang
ADD . /app
WORKDIR /app
RUN go build 
ENTRYPOINT ./app
EXPOSE 8080