# 使用するGolangのイメージを指定する
FROM golang:1.8

# 必要なパッケージなどなどをインストールする
RUN apt-get update -qq 

# ワーキングディレクトリを指定する
WORKDIR /go/src/github.com/go_example_vanilla

# sample_docker_compose直下のディレクトリをコンテナ上に載せる
ADD . .

# プロジェクトをビルド
RUN go build