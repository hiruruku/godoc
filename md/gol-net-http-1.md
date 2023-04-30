# Go HTTPハンドラー関数の解説

- hello HTTPハンドラー関数について解説します。

``` Golang
func hello(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello World, %s!",request.URL.Path[1:])
}
```

## http.ResponseWriter

http.ResponseWriterは、Go言語のnet/httpパッケージで定義されているインターフェースです。HTTPレスポンスを構築し、クライアントに送信するための一連のメソッドを提供します。また、Writeメソッドを実装しているため、io.Writerインターフェースを満たしています。

### 定義：

    type ResponseWriter interface {
        Header() Header
        Write([]byte) (int, error)
        WriteHeader(statusCode int)
    }

- http.ResponseWriterインターフェースは、以下の3つのメソッドを持っています。

1).Header() Header: このメソッドは、HTTPレスポンスのヘッダーを表すHeader型のマップを返します。このマップにヘッダーを追加したり、ヘッダーの値を変更したりすることができます。

2).Write([]byte) (int, error): このメソッドは、HTTPレスポンスのボディにバイトスライス（[]byte）を書き込みます。書き込んだバイト数とエラー（あれば）を返します。

3).WriteHeader(statusCode int): このメソッドは、HTTPレスポンスのステータスコードを設定します。ステータスコードは、整数値（例：200, 404, 500）で指定します。
HTTPハンドラ関数内で、http.ResponseWriterインターフェースを実装したオブジェクトを使って、HTTPレスポンスのヘッダー、ボディ、ステータスコードを操作できます。

## io.Writer

ファイルやネットワーク接続、HTTPレスポンスなど、さまざまな出力先に対して同じ書き込み操作を行うことができます

type Writer interface {
    Write(p []byte) (n int, err error)
}

io.Writer インターフェースは、Write メソッドを持っています。このメソッドは、バイトスライス（[]byte）を受け取り、書き込んだバイト数とエラー（あれば）を返します。このインターフェースを実装することで、様々な出力先に対して同じ書き込み操作を行うことができます。例えば、ファイルへの書き込み、ネットワーク接続への書き込み、バッファへの書き込みなど、さまざまなシチュエーションで io.Writer を使用できます。

## ストリーム
a).ストリームとは、データが連続的に流れるシーケンスのことを指します。ストリーミングデータは、一度に全体を処理するのではなく、データの一部を順番に処理することができます。ストリームは、ファイルやネットワーク接続、メモリバッファなど、さまざまなデータソースからデータを読み取ったり、データを書き込んだりすることができます。

b).Go言語では、io.Reader インターフェースと io.Writer インターフェースを使用して、ストリームを抽象化し、簡単に操作できるようにしています。これらのインターフェースを利用することで、さまざまなデータソースからデータを読み取ったり、データを書き込んだりするコードを一貫した方法で記述することができます。
## fmt.Fprintf

fmt.Fprintfは、Go言語のfmtパッケージに定義されている関数です。指定されたフォーマットに従って文字列を生成し、io.Writerインターフェースを実装したオブジェクトに書き込みができます。
例えば、ファイルやネットワーク接続、HTTPレスポンスなど、さまざまな出力先に対して
同じ書き込み操作を行うことができます

### 定義:

``` Golang
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

1).w io.Writer: io.Writer インターフェースを実装したオブジェクト。このオブジェクトにフォーマットされた文字列が書き込まれます。

2).format string: 書式指定文字列。この文字列には、プレースホルダ（例: %s, %d, %f など）が含まれており、後続の引数で指定された値に置き換えられます。

3).a ...interface{}: 可変長引数。これらの引数は、書式指定文字列のプレースホルダに対応して順番に置き換えられます。
fmt.Fprintf 関数は、書き込んだバイト数とエラー（あれば）を返します。



