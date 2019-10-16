package main

import (
	"fmt"
	"net"
	"strconv"

	"./command"
	"github.com/golang/protobuf/proto"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "10.1.3.182:8888")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//main调用完毕，关闭连接
	defer conn.Close()

	//接受服务器回复的数据，新任务
	go func() {
		// tmp := &command.CSLogin{Userid: 123}
		// date, _ := proto.Marshal(tmp)
		// str2 := string(date)
		// var msgid uint16
		// msgid = 1
		// msg := []byte(strconv.FormatInt(int64(msgid), 10))
		// str1 := string(msg)
		// str := str1 + str2
		// fmt.Println(str1)
		// msg = []byte(str)
		// conn.Write(msg)
		for i := 0; i < 5; i++ {
			info := []uint64{1, 2, 3, 4, 5, 6}
			tmp := &command.CSChooseHero{
				Userid: 123,
				Info:   info,
			}
			data, _ := proto.Marshal(tmp)
			str2 := string(data)
			var msgid uint16
			msgid = 3
			msg := []byte(strconv.FormatInt(int64(msgid), 10))
			str1 := string(msg)
			str := str1 + str2
			fmt.Println(str1)
			msg = []byte(str)
			conn.Write(msg)
		}
	}()

	//切片缓冲
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接受服务器的请求
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容，转换为字符串再打印
	}

}
