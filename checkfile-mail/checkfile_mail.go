package main

import (
  "fmt"
  "os"
  "log"
  "net/smtp"
  "path/filepath"
  "strings"
)

func main() {

  if len(os.Args) != 6 {
    fmt.Fprintf(os.Stderr, "USAGE: checkfile_mail smtpserver mailfrom mailto subject FILEPATTERN\n")
    fmt.Fprintf(os.Stderr, "----\n")
    fmt.Fprintf(os.Stderr, "If the file exists, the script will send an email\n")
    fmt.Fprintf(os.Stderr, "(wildcard is avaliable in FILEPATTERN)\n")
    fmt.Fprintf(os.Stderr, "----\n")
    fmt.Fprintf(os.Stderr, "EXAMPLE:\n")
    fmt.Fprintf(os.Stderr, "checkfile_mail \"server:25\" \"from@mail\" \"to@mail\" \"title\" \"C:\\Windows\\*.xyz\"\n")
    os.Exit(1)
  }

  server  := os.Args[1]
  from    := os.Args[2]
  to      := []string{ os.Args[3] }
  subject := os.Args[4]
  filter  := os.Args[5]

  filelist, err := GetFilepathList(filter)
  if err != nil {
    log.Fatal(err)
  }

  if len(filelist) == 0 {
    os.Exit(0)
  }

  msg  := []byte("Subject: " + subject + "\r\n\r\n this is checkfile_mail script. detect: " + strings.Join(filelist,","))

  err = smtp.SendMail( server, nil, from, to, msg )
  if err != nil {
    log.Fatal(err)
  }

}

func GetFilepathList(filter string) (list []string, err error) {

  entries, err := filepath.Glob(filter)
  if err != nil {
    return nil, err
  }

  for _, entry := range entries {
    isdir, err := isDir(entry)
    if err != nil {
      return nil, err
    }
    if ! isdir {
      list = append(list, entry)
    }
  }

  return list, nil
}

func isDir(p string) (bool, error) {
  f, err := os.Stat(p)
  if err != nil {
    return false, err
  }
  return f.Mode().IsDir(), nil
}


