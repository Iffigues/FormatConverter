FROM golang:latest


COPY ./app /app
WORKDIR /app
RUN go build

RUN curl -L -o pp 'https://github.com/apple/pkl/releases/download/0.28.1/pkl-linux-amd64'
RUN chmod +x pp
RUN mv  ./pp /bin/pkl
CMD ["./converter"]

