package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var target string
var threadNum int
var maxTime int
var MaxTime time.Duration

func main() {
	flag.StringVar(&target, "t", "", "用于测试的目标 ip:port")
	flag.IntVar(&threadNum, "n", 200, "每次测试都会以该参数递增，默认为 200")
	flag.IntVar(&maxTime, "s", 3, " 响应时间超过该时长即为最大支持线程 ，单位为秒，默认为 3")
	// 解析命令行参数
	flag.Parse()
	target = "baidu.com:443"
	if target == "" {
		fmt.Println("请输入 -t 参数")
	}
	MaxTime = time.Duration(maxTime) + time.Second

	routines := 200
	fmt.Println("开始测试")

	for {
		var wg sync.WaitGroup
		ready := make(chan interface{})
		thread := make(chan interface{})

		for i := 0; i < routines; i++ {
			wg.Add(1)
			go sendRequest(5, ready, thread, &wg)
		}
		wg.Wait()

		close(ready)
		select {
		case <-thread:
			fmt.Print("极限线程：")
			fmt.Println(routines)
			return
		case <-time.After(10 * time.Second):
			fmt.Print("当前线程：")
			routines += threadNum
			fmt.Println(routines)
			continue
		}

	}

}

func sendRequest(timeout int, ready <-chan interface{}, thread chan<- interface{}, wg *sync.WaitGroup) {
	wg.Done()
	<-ready
	start := time.Now()
	conn, err := net.DialTimeout("tcp", "gd.10086.cn:443", time.Duration(timeout)*time.Second)
	if err != nil {
		thread <- struct{}{}
	} else if err == nil {
		defer conn.Close()
	}
	elapsed := time.Since(start)
	if elapsed >= MaxTime {
		thread <- struct{}{}
	}

}
