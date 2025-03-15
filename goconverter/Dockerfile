FROM golang:latest

RUN mkdir -p /tmp/file/generate
RUN mkdir -p /tmp/file//generatedpkl
RUN mkdir -p /tmp/file/file
COPY ./app /app
WORKDIR /app
RUN rm -f go.* && go mod init converter
RUN go get
RUN go get github.com/apple/pkl-go/cmd/pkl-gen-go@v0.9.0
RUN go get github.com/apple/pkl-go
RUN go install github.com/apple/pkl-go/cmd/pkl-gen-go
RUN pkl-gen-go pkl/AppConfig.pkl --base-path converter
RUN go build

RUN curl -L -o pp 'https://github.com/apple/pkl/releases/download/0.28.1/pkl-linux-amd64'
RUN chmod +x pp
RUN mv  ./pp /bin/pkl

CMD ["./converter"]

