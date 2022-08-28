package main

import (
  "os"
  "os/exec"
  "strings"
  "time"

  "github.com/lxn/win"
  "github.com/lxn/walk"
  . "github.com/lxn/walk/declarative"
)


type MyMainWindow struct {
  *walk.MainWindow
  te1    *walk.TextEdit
  te2    *walk.TextEdit
  tv     *walk.TableView
  model  *MemberModel
}

type MemberModel struct {
  walk.TableModelBase
  items []*MemberAttr
}

type MemberAttr struct {
  Member
  stat    string
  checked bool
}

// this function is called from TableView
func (m *MemberModel) RowCount() int {
  return len(m.items)
}

// this function is called from TableView
func (m *MemberModel) Value(row, col int) interface{} {
  item := m.items[row]
  switch col {
  case 0:
    return item.Name
  case 1:
    return item.Addr
  case 2:
    return item.stat
  }
  panic("unexpected column")
}

// this function is called from TableView
func (m *MemberModel) Checked(row int) bool {
  return m.items[row].checked
}

// this function is called from TableView
func (m *MemberModel) SetChecked(row int, checked bool) error {
  m.items[row].checked = checked
  return nil
}

func (m *MemberModel) updateMember(config *Config) {
  m.items = []*MemberAttr{}
  for _, t := range config.Member {
    stat := getInitialStatus(t.Name)
    m.items = append(m.items, &MemberAttr{ Member: Member{ Name: t.Name, Addr: t.Addr }, stat: stat, checked: false})
  }
}

func getInitialStatus(name string) string {
  prefix := []string{"-","=","/","#"}
  for _, p := range prefix {
    if strings.HasPrefix(name, p) { return "" }
  }
  return "offline"
}

func (mw *MyMainWindow) updateTitle(config *Config) {
  mw.SetTitle(program + " - [" + config.MyName + "]")
}

//sendMessage called from keyevent.
func (mw *MyMainWindow) sendMessage(text string) {
  for _, item := range mw.model.items {
    if item.checked {
      go func(m Message){
         messageSendCh <- m
      }(Message{Name: item.Name, Addr: item.Addr, Host: "", Date: "", Text: text})
    }
  }
}

func (mw *MyMainWindow) menuEdit(config *Config) {
  exec.Command("notepad.exe", *optionConfigFilepath).Run()
}

func (mw *MyMainWindow) menuReload(config *Config) {
  config, err := loadConfigFromFile(*optionConfigFilepath, &Config{})
  if err != nil {
    return
  }

  mw.model.updateMember(config)
  mw.updateTitle(config)
  mw.model.PublishRowsReset()
}

func (mw *MyMainWindow) menuAbout() {
  s := program + " [" + version + "]"
  walk.MsgBox(mw, "About", s, walk.MsgBoxIconInformation)
}


//windows api call
func moveForeground(mw *MyMainWindow) {
  hwnd := mw.Handle()
  win.SetForegroundWindow(hwnd)

  //SWP_NOSIZE = 1, SWP_NOMOVE = 2, SWP_SHOWWINDOW = 0x40
  win.SetWindowPos(hwnd,0,0,0,0,0,1|2|0x40)
}

//listenMessageAndUpdate wait and pop messages to update
func listenMessageAndUpdate(mw *MyMainWindow) {
  for m := range messageRecvCh {
    mw.te2.AppendText(m.Host+" ["+m.Name+"] (" +m.Date+" - "+m.Addr+")\r\n")
    mw.te2.AppendText(m.Text)
    mw.te2.AppendText("\r\n\r\n")
    moveForeground(mw)
    walk.MsgBox(mw, "From: "+m.Host+" ["+m.Name+"]", m.Text, walk.MsgBoxOK)
  }
}

func pollingMemberStatus(mw *MyMainWindow) {
  for {
    for _, item := range mw.model.items {
      if item.stat == "" { continue }

      ok := sendHttpStatusCheck(item.Addr)
      if ok {
        item.stat = "online"
      } else {
        item.stat = "offline"
      }
    }
    mw.model.PublishRowsReset()
    time.Sleep(statCheckInterval * time.Second)
  }
}


func runMainWindow(config *Config) {
  mw := &MyMainWindow{ model: &MemberModel{} }
  mw.model.updateMember(config)

  go listenMessageAndUpdate(mw)
  go pollingMemberStatus(mw)

  if _, err := (MainWindow{
    AssignTo: &mw.MainWindow,
    Title:    program + " - [" + config.MyName + "]",
    MinSize:  Size{320, 180},
    Size:     Size{640, 480},
    Layout:   HBox{MarginsZero: true},
    MenuItems: []MenuItem{
      Menu{
        Text: "&File",
        Items: []MenuItem{
          Separator{},
          Action{
            Text: "&Exit",
            OnTriggered: func() { mw.Close() },
          },
        },
      },
      Menu{
        Text: "&Misc",
        Items: []MenuItem{
          Action{
            Text: "&Edit Config",
            OnTriggered: func(){ mw.menuEdit(config) },
          },
          Action{
            Text: "&Reload Config",
            OnTriggered: func(){ mw.menuReload(config) },
          },
          Separator{},
          Action{
            Text: "&About",
            OnTriggered: mw.menuAbout,
          },
        },
      },
    }, //MenuItems
    Children: []Widget{
      HSplitter{
        Children: []Widget{
          VSplitter{
            Children: []Widget{
              TextEdit{
                AssignTo: &mw.te1,
                Text:     "送信はShift+Enter",
                VScroll:  true,
                HScroll:  true,
                ReadOnly: false,
                OnKeyDown: func(key walk.Key) {
                  if walk.ModifiersDown() == walk.ModShift {
                    if key == walk.KeyReturn {
                      mw.sendMessage(mw.te1.Text())
                      mw.te1.SetText("")
                    }
                  }
                },
              },
              TextEdit{
                AssignTo: &mw.te2,
                VScroll:  true,
                HScroll:  true,
                ReadOnly: true,
              },
            },
          }, //VSplitter
          TableView{
            AssignTo: &mw.tv,
            Model:    mw.model,
            AlternatingRowBGColor: walk.RGB(241,241,241),
            CheckBoxes:            true,
            LastColumnStretched:   true,
            Columns: []TableViewColumn{
              {Title: "Name"},
              {Title: "Address"},
              {Title: "Status"},
            },
            StyleCell: func(style *walk.CellStyle) {
              item := mw.model.items[style.Row()]
              switch style.Col() {
              case 2:
                if item.stat == "online" {
                  style.TextColor = walk.RGB(0, 255, 0)
                } else {
                  style.TextColor = walk.RGB(255, 0, 0)
                }
              }
            },
          }, //TableView
        },
      }, //HSplitter
    },
  }.Run()); err != nil {
    os.Exit(1)
  }
}

