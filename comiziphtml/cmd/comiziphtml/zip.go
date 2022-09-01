package main

import (
  "io"
  "io/ioutil"
  "path/filepath"
  "archive/zip"
)


// zipファイルに格納されている1枚目の画像を取り出し、
// nameで指定されたファイル名に元の拡張子をつけて、dirに保存する
// 保存したファイル名を拡張子付きで返す
func ExtractCoverImageFromZip(path, name, dir string) (string, error) {

  // 最終的に保存したファイル名
  saved := ""

  r, err := zip.OpenReader(path)
  if err != nil {
    return saved, err
  }
  defer r.Close()


  // zip内のファイル巡回
  for _, f := range r.File {

    // ディレクトリであれば無視
    if f.FileInfo().IsDir() {
      // do nothing

    // ファイルであれば1枚目を取得して終了
    } else {

      // ただし拡張子が画像ファイル以外の場合は次へスキップ
      // 拡張子のパターンはかなり絞った
      _, _, ext := SplitFilepath(f.Name)
      if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
        continue
      }

      // コンテンツの取得
      rc, err := f.Open()
      if err != nil {
        return saved, err
      }
      defer rc.Close()

      // バッファに内容を読み込む
      buf := make([]byte, f.UncompressedSize)
      _, err = io.ReadFull(rc, buf)
      if err != nil {
        return saved, err
      }

      saved = name + ext
      if err = ioutil.WriteFile(filepath.Join(dir, saved), buf, f.Mode()); err != nil {
        return saved, err
      }

      // 一枚で終わり、forを抜ける
      break
    }
  }

  return saved, nil
}



