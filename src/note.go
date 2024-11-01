1. 默认channel是不带buffer的，只能发送并接收一个元素。为了让channel发送并接收多个参数,创建channel时定义buffer
   ch := make(chan int,2)  //带有2个buffer的channel
2.channel容量为空，只有再发送和接收端都已经ready情况下才会通信成功，否则会一直阻塞。channel不为空，只有当发送者或者接收者数量
   超过channel容量限制的时候，才会阻塞程序。
3.channel关闭后不可以再结束数据，但可以从里面读取数据
4.不要在消费端关闭channel,不要在有多个并行的生产者时对channel执行关闭操作。
5.channel底层是一个先进先出的阻塞队列，队列满时写会阻塞，队列空时读会阻塞，同时对channel的读和写操作是线程安全的
6.缓冲型通道读写经过buffer,非缓冲型通道的数据交互不经过通道，直接通过内存写传递。
7.单向的channel只能读或者写，双向的channel既能读也能写
8.chan读：若chan=nil,在阻塞模式下read操作会阻塞，chan关闭了，read不会阻塞，如果chan中存在有效数据，会返回有效数据，如果没有有效数据，
  会返回通道数据的零值。read操作的第二个返回值能判断第一个返回值是否是有效数据，某些情况还可以判断chan是否关闭。
9.chan写：若chan=nil,在阻塞模式下send操作会阻塞，chan关闭了，send会直接panic。当chan的sendq有阻塞的goroutine时，
  贸然关闭chan会导致panic,因此要注意chan关闭的时机。
10.panic的三种情况：1）向一个关闭的channel进行写操作。2）关闭一个nil的channel。3）重复关闭一个channel。
   读写一个nil的channel都会阻塞。不用特意关闭通道，没有使用的通道会在之后被gc
11.示例：对chan的读和写是成对的，不会出现goroutine泄漏
         func main() {
      	ch := make(chan int, 1)
      	go func() {
      		defer func() {  //使用defer兜底，确保一定会读channel
      			if err := recover(); err != nil {
      				t.Logf("panic recover")
      			}
      			<-ch
      		}()
      		// do something
      	}()
      	ch <- 1
      	//do something
      }


