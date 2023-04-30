# マルチプレクサに標準ライブラリのhttp.ServeMuxを使う。
## http.ServeMux
マルチプレクサは、リクエストを正しいハンドラにディスパッチ(振り分け）する。
Goの標準パッケージには、http.ServeMuxというパスベースのルーティングを提供し、
リクエストを適切なハンドラにディスパッチするオブジェクトがある。
高機能なサードパーティ製のマルチプレクサもある。

``` Golang
package main

import (
  "net/http"
)
func main() {
  mux := http.NewServeMux()　// New接頭辞、ServceMuxマルチプレクサ生成し、ポインターを渡しいている
  mux.HandleFunc("/", index) //  ServeMuxにパス(/)とハンドラー(index)を登録
  // 静的FileServerを作成（下記）
  files := http.FileServer(http.Dir("/public"))
  mux.Handle("/assets/",http.StripPrefix("/assets/",files))
  server := &http.Server{
    Addr:     "0.0.0.0:8080",
    Handler:  mux,
  }
  server.ListenAndServe()
}
```
## http.FileServer

FileServer関数は、FileSystem型(FileSystem Interface)
を引数にとって＆fileHandler（root)を返す。
&fileHandlerは、http.Handler インターフェースを実装する。


### 定義
``` Golang
func FileServer(root FileSystem) Handler {
	return &fileHandler{root}
}

func (f *fileHandler) ServeHTTP(w ResponseWriter, r *Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	serveFile(w, r, f.root, path.Clean(upath), true)
}
```

｀http.FileServer(http.Dir("/public"))｀の例では、/publicディレクトリを
  ルートとする静的ファイルサーバーを作成する。


## http.StripPrefix

### 定義

- net/httpパッケージ

``` Golang
func StripPrefix(prefix string, h Handler) Handler {
	if prefix == "" {
		return h
	}
	return HandlerFunc(func(w ResponseWriter, r *Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			NotFound(w, r)
		}
	})
}
```

http.StripPrefix` は通常、Web サーバーから静的ファイル
またはアセットを提供するときに使用されます。
引数は、prefixに削除するURLパスのプレフィックス、hは、削除後のパスで呼び出されるhttp.Handlerを渡します。
つまりは、/assets/css/style.cssというリクエストが来たら、/assets/を/css/style.cssのパスがファイル
サーバーに渡される。FileServerは下記の場合、/publicをルートディレクトリとしているので、
/public/css/style.cssファイルが返却されます。


