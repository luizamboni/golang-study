FROM alpine:3.15


RUN wget https://golang.org/dl/go1.19.1.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN go version