FROM golang:1.25-bookworm

RUN useradd -ms /bin/bash vscode && \
    apt-get update && \
    apt-get install -y \
        git \
        curl \
        make \
        bash-completion \
        vim \
        less \
        build-essential && \
    rm -rf /var/lib/apt/lists/*

USER vscode
WORKDIR /workspace

ENV GOPATH=/home/vscode/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

CMD ["bash"]
