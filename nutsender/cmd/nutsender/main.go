package main

import (
  "flag"
  "log"
  "os"
  "path/filepath"
  "time"
)

const (
  version = "0.9.7"
  program = "nutsender"
  listenport = "31701"
  statCheckInterval = 7
)


type Message struct {
  Name string
  Text string
  Date string
  Addr string
  Host string
}

type Session struct {
  hostname string
}

var (
  optionConfigFilepath = flag.String("conf", "config.xml", "config file path")
  optionLoggingOn      = flag.Bool("log", false, "enable message logging")

  messageRecvCh = make(chan Message)
  messageSendCh = make(chan Message)
)

func main() {
  flag.Parse()

  config, err := loadConfigFromFile(*optionConfigFilepath, &Config{})

  // create a new if the file can not be opened
  if err != nil {
    config = &initialConfig
    saveConfigToFile(*optionConfigFilepath, &initialConfig)
  }

  if *optionLoggingOn {
    logfile, err := openLogfile()
    if err != nil {
      panic("fail to open logfile")
    }
    defer logfile.Close()

    log.SetFlags(0)
    log.SetOutput(logfile)
  }

  session := &Session{}
  hostname, err := os.Hostname()
  if err != nil {
    hostname = "unknown"
  }
  session.hostname = hostname

  go waitMessageAndSend(config, session)
  go serveHttp()

  runMainWindow(config)
}

// waitMessageAndSend is waiting for a message typed in the GUI
// over the channel. waitMessageAndSend pick up messages, and
// send HTTP request to other client.
func waitMessageAndSend(config *Config, session *Session) {
  for m := range messageSendCh {
    myname := config.MyName
    host   := session.hostname
    toname := m.Name
    addr   := m.Addr
    text   := m.Text
    go sendHttpRequest(myname, host, toname, addr, text)
  }
}

// pushReceivedMessage called from receiveHttpMessage in serveHttp
func pushReceivedMessage(m Message) {
  messageRecvCh <- m
}

func openLogfile() (*os.File, error) {

  path := "."
  if *optionConfigFilepath != "config.xml" {
    path = filepath.Dir(filepath.Clean(*optionConfigFilepath))
  }

  name := time.Now().Format("nuts_20060102150405.txt") //YYYYMMDDhhmmss
  return os.OpenFile(filepath.Join(path, name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

