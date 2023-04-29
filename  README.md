# godoc
主に、go関連のmdファイルをnet/httpパッケージ、mdからhtmlへの変換パッケージを使用してブラウザで簡単にみられる
ようにしています。

## Overview of the Dockerfile:
Dockerfileはマルチステージビルド用です。
Localで確認するにはDockerfile.devをお使いください。
下のコマンドを打つか、下のコマンドをshにして実行するなどができます。

## Usage

1. Build the Docker image:

``` bash
docker build --no-cache -t godoc -f Dockerfile.dev .
```

2.コンテナを実行

``` bash
docker run -it --rm -v "$(pwd)":/app -p 8080:8080 godoc go run main.go
```