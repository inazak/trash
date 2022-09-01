package main

import (
  "os"
  "path/filepath"
  "strings"
  "time"
)

// ワイルドカードパターンを利用してファイル名、ディレクトリ名リストを取得する
func GetFilepathList(filter string, onlydir bool) (list []string, err error) {

  entries, err := filepath.Glob(filter)
  if err != nil {
    return nil, err
  }

  for _, entry := range entries {
    isdir, err := IsDir(entry)
    if err != nil {
      return nil, err
    }

    if onlydir {
      if isdir {
        list = append(list, entry)
      }
    } else {
      list = append(list, entry)
    }
  }

  return list, nil
}


func IsDir(p string) (bool, error) {
  f, err := os.Stat(p)
  if err != nil {
    return false, err
  }
  return f.Mode().IsDir(), nil
}

func GetModTime(p string) (time.Time, error) {
  f, err := os.Stat(p)
  if err != nil {
    return time.Time{}, err
  }
  return f.ModTime(), nil
}

func SplitFilepath(fpath string) (dir, name, ext string) {

  dir   = filepath.Dir(fpath)
  base := filepath.Base(fpath)

  dotindex := strings.LastIndex(base, ".")

  // file has no extention
  if dotindex == -1 {
    name = base
    ext  = ""
    return
  }

  name = base[0:dotindex]
  ext  = base[dotindex:len(base)]
  return
}


