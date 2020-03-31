FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o GoLogger . 
CMD ["/app/GoLogger"]