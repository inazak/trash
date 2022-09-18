package main

/*
type error interface {
  Error() string
}

errors パッケージにある New関数

func New(text string) error {
  return &errorString{text}
}

error型の戻り値は最後になるのが通例

func divide(x, y int) (int, error) {
  if y == 0 {
    return 0, errors.New("divide by zero")
  }
  return x / y, nil
}
*/

/*
errorのラッピングと取り出し

errors パッケージには Unwrap関数がある
これは引数の値が Unwrap() error メソッドを実装している場合、
そのUnwrapメソッドを呼び出した結果を返す

fmt.Errorf に %w を含んだ場合、Unwrapメソッドを実装したerrorを返す
return fmt.Errorf("handle xxx: %w", err)

*/

/*
errorの型アサーションの代わりに errors.As を使う
errors.As は第一引数で与えられたエラーが第二引数で与えた値に
適合する場合は、第二引数に代入し、true を返す

var de *db.Error
if errors.As(err, &de) {
  ...
}

他にチェックだけ行う errors.Is 関数もある
*/




