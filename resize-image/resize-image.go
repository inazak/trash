package main

import (
  "flag"
	"fmt"
  "os"
  "path/filepath"
  "strings"
	"github.com/disintegration/imaging"
)

var usage =`
Usage:
  resize-image [-width=1024] [-force] FILE_PATTERN

  wildcard can be used for FILE_PATTERN, like '*.jpg'.
`

var optionsWidth   int
var optionsForce   bool
var optionsPostfix string

func main() {

  flag.IntVar(&optionsWidth, "width", 1024, "width")
  flag.IntVar(&optionsWidth, "w",     1024, "width")
  flag.BoolVar(&optionsForce, "force", false, "force")
  flag.BoolVar(&optionsForce, "f",     false, "force")
  flag.StringVar(&optionsPostfix, "postfix", "resized", "postfix")
  flag.StringVar(&optionsPostfix, "p",       "resized", "postfix")
  flag.Parse()

  if len(flag.Args()) != 1 {
    fmt.Printf("%s", usage)
    os.Exit(1)
  }

  filename := flag.Args()[0]
  width    := optionsWidth
  postfix  := optionsPostfix

  list, err := getFilepathList(filename)
  if err != nil {
    fmt.Printf("%s", err.Error())
    os.Exit(2)
  }

  filecount := len(list)
  errcount  := 0

  for i, file := range list {

    fmt.Printf("%d/%d working ... %s\n", i+1, filecount, file)

    src, err := imaging.Open(file)
    if err != nil {
      errcount += 1
      fmt.Printf("[ERROR] in opening %s\n", file)
      continue
    }

    // options are as
    // imaging.CatmullRom, imaging.MitchellNetravali,
    // imaging.Linear, imaging.Box, imaging.NearestNeighbor
    dst := imaging.Resize(src, 0, width, imaging.Lanczos)

    dir, name, ext := splitFilepath(file)
    newfile := filepath.Join(dir, name + "." + postfix + ext)

    err = imaging.Save(dst, newfile)
    if err != nil {
      errcount += 1
      fmt.Printf("[ERROR] in saving %s\n", newfile)
      continue
    }

    if optionsForce {
      err = os.Rename(newfile, file)
      if err != nil {
        errcount += 1
        fmt.Printf("[ERROR] in rename from %s to %s\n", newfile, file)
        continue
      }
    }

  } //for

  fmt.Printf("done ... error count is %d\n", errcount)
  os.Exit(0)
}


func getFilepathList(filter string) (list []string, err error) {

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


func splitFilepath(fpath string) (dir, name, ext string) {

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


