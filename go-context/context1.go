package main

import (
  "fmt"
  "context"
  "sync"
  "time"
)

/*
Contextインターフェイスを満たす構造体で木構造を作る
親のコンテキストをコピーし、付加情報を加えた子を生成する


                          +--- Context3
                          |
Context1  --- Context2 ---+--- Context4
                          |
                          +--- Context5

(1) Context2 でキャンセル発生
(2) Context3-5 に伝播
(3) Context1 に対して Context2 への参照を削除

Contextインターフェイス
type Context interface {

  Deadline() (deadline time.Time, ok book)
    自動でキャンセルされる時刻と、キャンセルされる時刻を設定しているかどうかのbool

  Done() <- chan struct{}
    受信専用のチャネル

  Err() error
    コンテキストがキャンセルされた理由を返す

  Value(key interface{}) interface{}
    指定されたkeyに対応する値を返す
}

コンテキストは構造体に入れず、関数の引数として使用する
関数の引数としては、第一引数にして、変数名を ctx にする
渡すべきコンテキストが判定できない場合は、nil ではなく context.TODO を渡す


コンテキスト（の木構造）を作るための関数

func Background() Context
func TODO() Context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context

type CancelFunc func()

Background() or TODO()  ---> emptyCtx
                               |
                               +--- WithValue() --> valueCtx
                               |
                               +--- WithCancel() --> cancelCtx
                               |
                               +--- WithDeadline() --> timerCtx
                                    WithTimeout()

*/


// 並行処理のうち一つでも失敗すれば残りをスキップする処理
// ただしこのサンプルでは全部成功するパターンのみ
func doSomethingWithContext(n int) error {

  // コンテキストの生成
  ctx               := context.Background()
  cancelCtx, cancel := context.WithCancel(ctx)

  // 終了時にリソースの解放
  defer cancel()

  // ゴルーチンからのエラーを集約する
  errCh := make(chan error, n)

  // 並行処理を起動する
  wg := sync.WaitGroup{}
  for i := 0; i < n; i++ {
    current := i
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      // エラーが発生したら、キャンセル処理を行って、チャネルにエラーを送信する
      if err := doSomething1(cancelCtx, id) ; err != nil {
        cancel()
        errCh<- err
      }
      return
    }(current)
  }

  wg.Wait()

  // チャネルからメッセージを取り出す
  close(errCh)
  var errs []error
  for err := range errCh {
    errs = append(errs, err)
  }

  // もしエラーが発生してたら、最初のエラーを返す
  if len(errs) > 0 {
    return errs[0]
  }

  // 正常終了の場合
  return nil
}

// 何らかの処理を行う関数
func doSomething1(ctx context.Context, id int) error {

  // 処理に入る前にコンテキストの状態を確認する
  select {
  case <-ctx.Done():
    return ctx.Err()

  // コンテキストがキャンセルされていないので、そのまま処理に進む
  default:
  }

  fmt.Println(id)
  return nil
}




// 並行処理を行うがタイムアウトを設ける
func doSomethingWithTimeout(n int) {

  // タイムアウトする秒数
  d := 5 * time.Second

  // コンテキストの生成
  ctx              := context.Background()
  timerCtx, cancel := context.WithTimeout(ctx, d)

  // リソース解放を忘れない
  defer cancel()

  // 並行処理を起動する
  wg := sync.WaitGroup{}
  for i := 0 ; i < n ; i++ {
    current := i
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      // タイムアウトありの処理を呼び出す
      doSomething2(timerCtx, id)
      return
    }(current)
  }

  wg.Wait()

  return
}

// 何らかの処理を行う関数
func doSomething2(ctx context.Context, id int) {

  // タイムアウトするまで繰り返す
  for {
    select {
    case <-ctx.Done():
      fmt.Printf("id:%d timeout\n", id)
      return

    default:
      time.Sleep(2 * time.Second)
      fmt.Printf("id:%d working\n", id)
    }
  }
  return
}




func main() {

  fmt.Printf("--- Begin doSomethingWithContext --\n")

  doSomethingWithContext(5)

  fmt.Printf("--- End   doSomethingWithContext --\n\n")


  fmt.Printf("--- Begin doSomethingWithTimeout --\n")

  doSomethingWithTimeout(5)

  fmt.Printf("--- End   doSomethingWithTimeout --\n\n")
}



