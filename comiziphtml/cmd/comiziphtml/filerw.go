package main

import (
  "io/ioutil"
  "os"
  "encoding/xml"
)


func LoadXMLFile(path string, xmlst interface{}) error {

  file, err := os.OpenFile(path, os.O_RDONLY, 0666)
  if err != nil {
    return err
  }
  defer file.Close()

  // use sjis
  //reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())

  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    return err
  }

  err = xml.Unmarshal(bytes, xmlst)
  if err != nil {
    return err
  }

  return nil
}

func SaveXMLFile(path string, xmlst interface{}) error {

  bytes, err := xml.MarshalIndent(xmlst, "", "  ")
  if err != nil {
    return err
  }
  text := string(bytes)

  // LF to CRLF
  //text = strings.Replace(text, "\r\n", "\n", -1)
  //text = strings.Replace(text, "\n", "\r\n", -1)

  file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    return err
  }
  defer file.Close()

  // use sjis
  //writer := transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())

  _, err = file.Write([]byte(text))
  if err != nil {
    return err
  }

  return nil
}

func ReadTextFile(filepath string) (string, error) {

  file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
  if err != nil {
    return "", err
  }
  defer file.Close()

  // use sjis
  //reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())

  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    return "", err
  }

  return string(bytes), nil
}

func WriteTextFile(filepath string, text string) error {

  err := ioutil.WriteFile(filepath, []byte(text), 0666)
  if err != nil {
    return err
  }

  return nil
}


