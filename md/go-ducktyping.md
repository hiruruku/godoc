# ダックタイピングとGoのインターフェース

## ダックタイピング

ダックタイピングは、「もしもあるオブジェクトがアヒルのように歩き、アヒルのように鳴くなら、それはアヒルであろう」という考え方です。型の振る舞いによって型が決まるという概念で、明示的な型宣言が不要であることが特徴です。

## Go言語とダックタイピング

Go言語では、インターフェースの実装は明示的に宣言する必要がありません。ある型がインターフェースを実装しているかどうかは、その型がインターフェースで定義されているメソッドをすべて実装しているかどうかによって決まります。

### 例

以下に、Go言語でのダックタイピングを示す例を示します。

```go
// Walkerインターフェースは、Walkメソッドを持つすべての型に適用されます。
type Walker interface {
    Walk()
}

// Dog型は、Walkメソッドを持っています。
type Dog struct {
    Name string
}

func (d Dog) Walk() {
    fmt.Println(d.Name, "is walking.")
}

// Duck型もWalkメソッドを持っています。
type Duck struct {
    Name string
}

func (d Duck) Walk() {
    fmt.Println(d.Name, "is walking.")
}
// Walkerインターフェース型を引数にとる
func TakeAWalk(w Walker) {
    w.Walk()
}

func main() {
    dog := Dog{Name: "Buddy"}
    duck := Duck{Name: "Donald"}

    // Dog型とDuck型の両方がWalkerインターフェースを満たしているため、TakeAWalk関数に渡すことができます。
    TakeAWalk(dog)
    TakeAWalk(duck)
}
```

Dog型とDuck型が共通のメソッド（Walk）を持っているため、それぞれがWalkerインターフェースを実装しているとみなされます。このように、Go言語ではダックタイピングの考え方がインターフェースの実装に取り入れられています。

Goのインターフェース実装では、型がインターフェースで定義されているメソッドをすべて実装しているかどうかによって決まり、明示的な宣言はありません。暗黙的なインターフェース実装です。