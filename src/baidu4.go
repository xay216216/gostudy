package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	logF = flag.String("log", "test.log", "Log file name") //将运行时参数 地址 绑定到logF 运行时没带参数默认logF为test.log
)

func main() {
	resp, err := http.Get("http://www.baidu.com")

	flag.Parse()                                                                //解析参数付给logF
	outfile, err := os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开

	log.SetOutput(outfile)                               //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("test out:%v \n", "test log") //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

	if err != nil {
		//handleerror
		fmt.Println(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	buf := make([]byte, 1024)
	//createfile
	f, err1 := os.OpenFile("baidu1.html", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		panic(err1)
		return
	}
	defer f.Close()

	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}
		f.WriteString(string(buf[:n])) //抓到的东西存到本地的一个文件
	}

}
