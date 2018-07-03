package main

import (
    "net"
    "time"
    "log"
    "fmt"
    "os"
)

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

func main() {
    ip := "0.0.0.0"
    if len(os.Args) > 1 {
        ip = os.Args[1]
	}
    go recvUDPMsg()
    go sendUDPMsg(ip)
    for {
        time.Sleep(1000 * time.Millisecond)
    }
}
