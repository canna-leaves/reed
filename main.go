package main

import (
    "net"
    "time"
    "log"
    "fmt"
    "os"
    // "runtime"
    // "os/exec"
    // "errors"
    // "bytes"
    // "github.com/kataras/iris"
    // "sort"
)

/*var commands = map[string]string {
    "windows": "cmd",
}*/

var devices = map[string]string {}

/*func Open(uri string) error {
    _, ok := commands[runtime.GOOS]
    if !ok {
        return errors.New("don't know how to open")
    }

    cmd := exec.Command("cmd", "/c", "start", uri)
    return cmd.Start()
}*/

func init() {
    //log.SetFlags(log.Ldate | log.Lmicroseconds/* | log.Lshortfile*/)
    log.SetOutput(os.Stdout)
}

func recvUDPMsg(){
    addr, err := net.ResolveUDPAddr("udp", ":44444")
    if  err != nil {
        log.Fatalln(err)
    }

    conn, err := net.ListenUDP("udp", addr)
    defer conn.Close()

     if  err != nil {
        log.Fatalln(err)
    }

    var buf [1024]byte
    for {
        n, raddr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            return
        }

        log.Println(string(buf[0:n]), raddr)
        if n == 13 {
           // devices[string(buf[0:n])] = raddr.IP.String()
           devices[raddr.IP.String()] = string(buf[0:n])
        }
    }
}

func sendUDPMsg(ip string) {
    raddr := net.UDPAddr{
        IP: net.IPv4(255, 255, 255, 255),
        Port: 44444,
	}
    laddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:0", ip))
    if err != nil {
        log.Fatalln(err)
    }

    conn, err := net.ListenUDP("udp", laddr)
    defer conn.Close()
    if err != nil {
        log.Fatalln(err)
    }

    data := []byte{0x66, 0, 0, 0, 0, 0, 0, 0}
    for {
        n, err := conn.WriteToUDP(data, &raddr)
        if err != nil {
            log.Println(err)
        } else {
            log.Println("send ok", n)
        }
        time.Sleep(1000 * time.Millisecond)
    }
}

/*func webServer() {
    app := iris.New()
    //app.StaticWeb("/static", "./assets")
    app.Get("/", func(ctx iris.Context) {
        var buffer bytes.Buffer
        var keys []string
        for k, _ := range devices {
            keys = append(keys, k)
        }
        sort.Strings(keys)
        for _, k := range keys {
            buffer.WriteString(fmt.Sprintf("<a href='http://%s'>%s</a> %s<br/>", k, devices[k], k))
        }
        ctx.Writef("%s", buffer.String())
    })
    app.Run(iris.Addr(":10086"))
}*/

/*func openBrowser() {
    time.Sleep(1000 * time.Millisecond)
    err := Open("http://localhost:10086")
    if err != nil {
        log.Println(err)
    }
}*/

func main() {
    ip := "0.0.0.0"
    if len(os.Args) > 1 {
        ip = os.Args[1]
	}
    go recvUDPMsg()
    go sendUDPMsg(ip)
    // go openBrowser()
    // webServer()
    for {
        time.Sleep(1000 * time.Millisecond)
    }
}
