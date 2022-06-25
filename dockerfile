FROM golang:latest
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

RUN mkdir -p /usr/bin/Go_Scan

WORKDIR /usr/bin/Go_Scan

COPY . /usr/bin/Go_Scan

RUN cd src/ && go mod tidy && make linux && chmod +x Go_Scan

EXPOSE 666
