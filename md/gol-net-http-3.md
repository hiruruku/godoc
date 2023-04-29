# net.httpHandleメソッド

## http.Handlerインターフェース

### 定義
``` Golang
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
- net/httpパッケージ。ServerHttpを実装することを要求している。


### 実装例
```

## 実装
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


## http.StripPrefix 関数

### 定義

``` Golang
func StripPrefix(prefix string, h Handler) Handler
```
### 実装例

- http.StripPrefix` は通常、Web サーバーから静的ファイル
  またはアセットを提供するときに使用されます。
  たとえば、CSS、JavaScript、および画像ファイルを含む
  ディレクトリ `static` があり、それらを Web サーバーから提供したい場合、
  次のように `http.StripPrefix` を使用できます。

``` golang
func main() {
    // Create a new file server handler for serving static files
    fs := http.FileServer(http.Dir("static"))

    // Strip the "/static" prefix from incoming requests
    http.Handle("/static/", http.StripPrefix("/static", fs))

    // Start the server
    http.ListenAndServe(":8080", nil)
}
```

## http.FileServer

FileServer関数は、FileSystem型(FileSystem Interface)
を引数にとって＆fileHandler（root)を返す。
&fileHandlerは、http.Handler インターフェースを実装する。


### 定義
``` golang
func FileServer(root FileSystem) Handler {
	return &fileHandler{root}
}
```

``` Golang
func (f *fileHandler) ServeHTTP(w ResponseWriter, r *Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	serveFile(w, r, f.root, path.Clean(upath), true)
}
```
