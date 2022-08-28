package main

import (
  "encoding/xml"
  "os"
  "io/ioutil"
  "strings"

  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)


type Config struct {
  XMLName xml.Name `xml:"config"`
  MyName  string   `xml:"myname"`
  Member  []Member `xml:"member"`
}

type Member struct {
  Name string `xml:"name"`
  Addr string `xml:"addr"`
}

var initialConfig = Config{
  MyName: "YourName",
  Member: []Member{
    { Name: "localhost", Addr: "localhost" },
    { Name: "---------", Addr: "---------" },
    { Name: "example01", Addr: "example01" },
    { Name: "example02", Addr: "example02" },
    { Name: "example03", Addr: "example03" },
  },
}


func loadConfigFromFile(filepath string, conf *Config) (*Config, error) {

  file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
  if err != nil {
    return conf, err
  }
  defer file.Close()

  reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())

  bytes, err := ioutil.ReadAll(reader)
  if err != nil {
    return conf, err
  }

  err = xml.Unmarshal(bytes, &conf)
  if err != nil {
    return conf, err
  }

  return conf, nil
}


func saveConfigToFile(filepath string, conf *Config) error {

  bytes, err := xml.MarshalIndent(conf, "", "  ")
  if err != nil {
    return err
  }

  //LF to CRLF
  text := string(bytes)
  text = strings.Replace(text, "\r\n", "\n", -1)
  text = strings.Replace(text, "\n", "\r\n", -1)

  file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    return err
  }
  defer file.Close()

  writer := transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())
  _, err = writer.Write([]byte(text))
  if err != nil {
    return err
  }

  return nil
}

