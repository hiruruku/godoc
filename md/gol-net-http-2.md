# Go言語のHTTPサーバー

Go言語のhttpパッケージを使用して、簡単なHTTPサーバーを実装する際に重要な関数と概念。
主に`http.Hanlde`インターフェース,`http.HandleFunc`,`http.ListenAndServe`の２つの関数と
デフォルトマルチプレクサーについて説明します。

## http.Handle インターフェース

net/httpパッケージ。ServerHttpを実装することを要求する。

``` Golang
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
http.Handlerタイプを実装する構造体,myHandler。
``` Golang
type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func main() {
    handler := &myHandler{}
    http.Handle("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## http.HandleFunc

`http.HandleFunc`は、http.Handlerインンターフェースを実装する関数型。
func(http.ResponseWriter, *http.Request)型の関数を引数に取り、
それをラップしてServeHTTPメソッドを持つ、関数型のインスタンスに変換する。

``` Golang
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

指定されたパターンと関連付けられたハンドラ関数をデフォルトマルチプレクサーに登録するためのメソッド。
これにより、リクエストが特定のパターンに一致する場合に、対応するハンドラ関数が呼び出されます。

``` go
http.HandleFunc("/", hello)
```
- 上記の例では、ルートパス（"/"）が指定されたリクエストが来た場合、hello関数が呼び出されます。
  そして、このパスとハンドラーfunctionの対応づけは、デフォルトマルチプレクサーに登録されます。



## http.ListenAndServe

http.ListenAndServeは、指定されたアドレスとポートでHTTPサーバーを起動し、リクエストを待ち受ける関数です。
第2引数には、リクエストを処理するためのマルチプレクサーを渡すことができます。

## デフォルトマルチプレクサー

デフォルトマルチプレクサーは、Go言語のhttpパッケージで提供されるデフォルトのリクエストマルチプレクサーです。
マルチプレクサーは、リクエストを適切なハンドラ関数にルーティングするためのコンポーネントで、
http.HandleFuncを使って登録されたパターンとハンドラ関数のマッピングを管理します。

デフォルトマルチプレクサーを使用する場合、http.ListenAndServeの第2引数にnilを渡します。
これにより、登録されたハンドラ関数を自動的に呼び出してリクエストを処理できます。

``` go
http.ListenAndServe(":8080", nil)
```