package main

// comiziphtml overview
//
// 1. 読み取り対象のディレクトリ構成
//
// 親ディレクトリの下に、タイトル毎のディレクトリがあり
// タイトルディレクトリの下に、それぞれの単行本のzip圧縮ファイルがある
// コメント用テキストを含む comment.txt があってもよい
// ただし、comment.txt のエンコードはUTF-8
// 
// (dir)           (dir)           (file)
// COMICDIR -- +-- TITLE AAA --+-- FILENAME 01巻.zip
//             |               |
//             +-- TITLE BBB   +-- FILENAME 02巻.zip
//             |               |
//             +-- ...         +-- comment.txt (optional)
//                             |
//                             +-- ...
//
//
// 2. コメントファイルのフォーマット
//
// キーワード、値の列挙、セミコロン区切り
// 
//   Status: 完結; Rundown: あらすじがどうとか
//
//
// 3. 出力ディレクトリの構成
//
// コマンドを実行したディレクトリに出力する
// image にはascii順で最初の単行本の
// 最初の画像データを取得して保存している
// たいていの場合、表紙イメージになるはず
//
// - index.html
// - meta.xml
// - image --+-- [base32ed titlename].jpg or png,gif,jpeg
//           |
//           +-- [base32ed titlename].jpg
//           |
//           +-- ...

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
  "time"
  "encoding/xml"
  "sort"
  "encoding/base32"
  "bytes"
)

const timeformat = "2006-01-02 15:04:05 -0700 MST"
const imagedir   = "image"


// XMLファイルフォーマット
type XMLData struct {
  XMLName xml.Name    `xml:"comiziphtml"`
  XMLComic []XMLComic `xml:"comic"`
}

type XMLComic struct {
  Title   string `xml:"title"`
  Image   string `xml:"image"`
  Comment string `xml:"comment"`
  Updated string `xml:"updated"`
  Cover   string `xml:"cover"`
  Latest  string `xml:"latest"`
}

// タイトルをキーとしたマップ
// XMLから読んだ情報をこれに変換し、
// 保存するときはこれをXMLにする
type ComicMap map[string]*ComicInfo // key is Title
type ComicInfo struct {
  Image   string
  Comment string
  Updated time.Time
  Cover   string
  Latest  string
}


const usage = `Usage: comiziphtml TARGETDIR`

func main() {

  if len(os.Args) != 2 {
    fmt.Printf("invalid argument\n\n%s\n", usage)
    os.Exit(1)
  }

  target := os.Args[1]

  meta := &XMLData{}
  err := LoadXMLFile("meta.xml", meta)
  if err != nil {
    meta = &XMLData{}
  }

  cmap := make(ComicMap)
  cmap.FromXMLData(meta)

  dirlist, err := GetFilepathList(filepath.Join(target, "*"), true)
  if err != nil {
    fmt.Printf("ERROR: %s\n", err)
  }

  for _, dir := range dirlist {
    fmt.Printf("Working: %s\n", dir)

    err = cmap.Add(dir)
    if err != nil {
      fmt.Printf("ERROR: %s\n", err)
    }

    err = cmap.ExtractCoverImage(dir, imagedir)
    if err != nil {
      fmt.Printf("ERROR: %s\n", err)
    }
  }

  err = SaveXMLFile("meta.xml", cmap.ToXMLData())
  if err != nil {
    fmt.Printf("ERROR: %s\n", err)
  }

  html := cmap.ToHTML()
  err = WriteTextFile("index.html", html)
}


// 親フォルダの下のタイトル毎のディレクトリ名を取る
// この作業で新規のタイトルの追加、Comment, Updated の更新は行われるが
// Image の取得は行われない
func (cm ComicMap) Add(dir string) error {

  _, title, _ := SplitFilepath(dir)

  // zip 圧縮したファイルのみ列挙する
  filelist, err := GetFilepathList(filepath.Join(dir, "*.zip"), false)
  if err != nil {
    return err
  }

  sort.StringSlice(filelist).Sort()

  // ファイルの最初巻、最終巻を取得する
  // LatestFile になっているが最新という意味で使われていない
  cover  := ""
  latest := ""
  if len(filelist) > 0 {
    _, name, ext := SplitFilepath(filelist[0])
    cover  = name + ext
    _, name, ext = SplitFilepath(filelist[len(filelist)-1])
    latest = name + ext
  }

  // comment.txtを読んで追記する
  comment, err := ReadTextFile(filepath.Join(dir, "comment.txt"))
  if err != nil {
    comment = ""
  }
  comment = strings.Replace(comment, "\r\n", "", -1)
  comment = strings.Replace(comment, "\n", "", -1)


  // そのタイトルの中で最新のファイル更新日を取得する
  // つまり、途中の巻を入れ替えた場合でも更新日を変更することになる
  updated := time.Time{}
  for _, f := range filelist {
    t, _ := GetModTime(f)
    if updated.Before(t) {
      updated = t
    }
  }

  // 既にイメージファイルを取得済で、最新巻が同じ場合は
  // 再取得の必要はないので、それを使う
  // 最新巻が変わっている場合は、改めて取得するために
  // Imageをブランクにする
  image := ""
  if _, ok := cm[title]; ok {
    if cm[title].Image != "" {
      image = cm[title].Image
    }
  }

  cm[title] = &ComicInfo{
    Image:   image,
    Comment: comment,
    Updated: updated,
    Cover:   cover,
    Latest:  latest,
  }

  return nil
}


// Imageにパス情報を持っていない場合、zipファイルから取得して
// imagedir ディレクトリに保存する
// ファイル名はタイトルをbase32エンコードしたものにする
func (cm ComicMap) ExtractCoverImage(dir, imagedir string) error {

  _, title, _ := SplitFilepath(dir)

  if _, ok := cm[title]; ok {
    if cm[title].Image == "" {
      basename := base32.StdEncoding.EncodeToString([]byte(title))
      saved, err := ExtractCoverImageFromZip(filepath.Join(dir, cm[title].Cover), basename, imagedir)
      if err != nil {
        return err
      }
      cm[title].Image = saved
    }
  }

  return nil
}



func (cm ComicMap) ToXMLData() *XMLData {
  meta := &XMLData{}

  comiclist := []XMLComic{}
  for k, v := range cm {
    comiclist = append(comiclist, XMLComic{
      Title:   k,
      Image:   v.Image,
      Comment: v.Comment,
      Updated: v.Updated.Format(timeformat),
      Cover:   v.Cover,
      Latest:  v.Latest,
    })
  }

  meta.XMLComic = comiclist
  return meta
}

func (cm ComicMap) FromXMLData(meta *XMLData) {
  for _, c := range meta.XMLComic {
    t, _ := time.Parse(timeformat, c.Updated)
    cm[c.Title] = &ComicInfo{
      Image:   c.Image,
      Comment: c.Comment,
      Updated: t,
      Cover: c.Cover,
      Latest: c.Latest,
    }
  }
}


func (cm ComicMap) ToHTML() string {
  titlelist := []string{}

  for title, _ := range cm {
    titlelist = append(titlelist, title)
  }

  sort.Slice(titlelist, func(i, j int) bool {
    return cm[titlelist[i]].Updated.After( cm[titlelist[j]].Updated )
  })

  var html bytes.Buffer

  html.WriteString("<html><head><title>comiziphtml</title>\n")
  html.WriteString("<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/>\n")
  html.WriteString("<style type=\"text/css\">\n<!--\n")
  html.WriteString("  div {margin: 5%; clear:left;}\n")
  html.WriteString("  img {margin-right: 5%; margin-bottom: 5%; float:left;}\n")
  html.WriteString("-->\n</style>\n")
  html.WriteString("</head><body>\n\n")

  for _, title := range titlelist {
    html.WriteString("<div>\n")
    html.WriteString("  <img src=\"")
    html.WriteString(imagedir)
    html.WriteString("/")
    html.WriteString(cm[title].Image)
    html.WriteString("\" />\n")

    html.WriteString("  <h3>")
    html.WriteString(title)
    html.WriteString("</h3>\n")

    html.WriteString("  <p>Latest: ")
    html.WriteString(cm[title].Latest)
    html.WriteString("</p>\n")

    html.WriteString("  <p>Update: ")
    html.WriteString(cm[title].Updated.Format(timeformat))
    html.WriteString("</p>\n")

    html.WriteString("  <p>")
    html.WriteString(cm[title].Comment)
    html.WriteString("</p>\n")

    html.WriteString("</div>\n")
  }

  html.WriteString("</body></html>\n")

  return html.String()
}


