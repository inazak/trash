package main

// go get -u github.com/chromedp/chromedp

import (
  "context"
  "fmt"
  "log"
  "time"
  "github.com/chromedp/chromedp"
  "github.com/chromedp/chromedp/kb"
)

var example = "what is chromedp"

func main() {

  ctx, cancel := newChromedpContext( context.Background(), true )
  ctx, cancel = context.WithTimeout(ctx, 30 * time.Second)
  defer cancel()

  var res string

  err := chromedp.Run(ctx,
    chromedp.Navigate(`https://www.google.com`),
    chromedp.SendKeys(`input[name=q]`, example),
    //chromedp.Click(`input[name=btnK]`), // doesnt work
    chromedp.KeyEvent(kb.Enter),
    chromedp.WaitVisible(`#result-stats`, chromedp.ByQuery),
    chromedp.Text(`h3`, &res, chromedp.NodeVisible, chromedp.ByQuery),
  )

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf(">> %s\n", res)
}



func newChromedpContext(ctx context.Context, headless bool) (context.Context, context.CancelFunc) {

  var opts []chromedp.ExecAllocatorOption

  for _, opt := range chromedp.DefaultExecAllocatorOptions {
    opts = append(opts, opt)
  }

  if ! headless {
    opts = append(opts,
      chromedp.Flag("headless", false),
      chromedp.Flag("hide-scrollbars", false),
      chromedp.Flag("mute-audio", false),
    )
  }

  allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
  ctx, cancel := chromedp.NewContext(allocCtx)

  return ctx, func() {
    cancel()
    allocCancel()
  }
}

