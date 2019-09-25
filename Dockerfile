FROM golang:latest

RUN apt update
RUN apt -y install sudo
RUN sudo apt install -y mecab libmecab-dev mecab-ipadic-utf8 git

ENV GO111MODULE on
ENV CGO_LDFLAGS -L/usr/lib/x86_64-linux-gnu -lmecab -lstdc++
ENV CGO_CFLAGS -I/usr/include

RUN git clone https://github.com/cotton392/ctn_ai.git