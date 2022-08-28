package main

import (
  "io/ioutil"
  "net/http"
  "net/url"
  "log"
  "time"
)

func makeURL(addr, path string) string {
  return "http://" + addr + ":" + listenport + path
}

func makeValues(name, host, text string) url.Values {
  return url.Values{ "name": { name }, "host": { host }, "text": { text } }
}

func sendHttpRequest(myname, host, toname, addr, text string) {
  resp, err := http.PostForm(makeURL(addr, "/message/"), makeValues(myname, host, text))
  if err != nil {
    t := "Fail to Connect: " + toname + " [" + addr + "]"
    http.PostForm(makeURL("localhost","/message/"), makeValues("ERROR", "", t))
    return
  }

  ioutil.ReadAll(resp.Body) //discard
  defer resp.Body.Close()
}


func serveHttp() {
  http.HandleFunc("/status/",  receiveHttpStatus)
  http.HandleFunc("/message/", receiveHttpMessage)
  if http.ListenAndServe(":"+listenport, nil) != nil {
    panic("fail to serve http")
  }
}

func receiveHttpStatus(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
}

func receiveHttpMessage(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  n := r.FormValue("name")
  h := r.FormValue("host")
  t := r.FormValue("text")
  d := time.Now().Format("2006/01/02 15:04:05")
  a := r.RemoteAddr

  go pushReceivedMessage(Message{Name: n, Text: t, Date: d, Addr: a, Host: h})

  if *optionLoggingOn {
    log.Print(h+" ["+n+"] (" +d+" - "+a+")\r\n")
    log.Print(t)
    log.Print("\r\n\r\n")
  }
}

func sendHttpStatusCheck(addr string) bool {
  resp, err := http.Get(makeURL(addr, "/status/"))
  if err != nil {
    return false
  }
  defer resp.Body.Close()

  ioutil.ReadAll(resp.Body) //discard
  return true
}


