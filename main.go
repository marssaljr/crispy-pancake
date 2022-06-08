package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os/exec"
  "os"
  "golang.org/x/term"
  "bufio"
  "strings"
)

var img = "";
var urlBase = "https://source.unsplash.com/random/2400x1600/?"

func getImg(urlBase string) {
  resp, err := http.Get(urlBase);
  if err != nil {
    log.Fatalln(err);
  }
  img = resp.Header.Get("x-imgix-id")
  //We Read the response body on the line below.
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  resp.Body.Close()
  ioutil.WriteFile(img, data, 0666)
  clear()
  fmt.Printf("🤌 New image! close feh window and type space to show more\n")
}

func showImg() {
  cmd := exec.Command("feh","--scale-down", img)
  cmd.Run()
}
func clear() {
  cmd := exec.Command("clear") //Linux example, its tested
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func exit() {
  cmd := exec.Command("killall feh && exit")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func startUp() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("🧑 Type a theme of wallpaper 🎨")
  fmt.Println("-------------------------------")
  for {
    fmt.Print(": ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("", text) == 0 {
      urlBase = urlBase + "backgrounds"
      return
    }
    urlBase = urlBase + text
    return
  }
}

func main() {

  startUp()
  fmt.Printf("Press [Space] to download a new wallpaper or any key to exit")
  // switch stdin into 'raw' mode
  for {
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
      fmt.Println(err)
      return
    }
    defer term.Restore(int(os.Stdin.Fd()), oldState)

    b := make([]byte, 1)
    _, err = os.Stdin.Read(b)
    if err != nil {
      fmt.Println(err)
      return
    }
    if string(b[0]) != " " {
      exit()
      return
    }

    clear()
    fmt.Printf("Downloading a new wallpaper\n");
    getImg(urlBase)
    showImg()
  }
}
