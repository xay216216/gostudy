1. 默认channel是不带buffer的，只能发送并接收一个元素。为了让channel发送并接收多个参数,创建channel时定义buffer
   ch := make(chan int,2)  //带有2个buffer的channel
2.channel容量为空，只有再发送和接收端都已经ready情况下才会通信成功，否则会一直阻塞。channel不为空，只有当发送者或者接收者数量
   超过channel容量限制的时候，才会阻塞程序。
3.channel关闭后不可以再结束数据，但可以从里面读取数据
4.不要在消费端关闭channel,不要在有多个并行的生产者时对channel执行关闭操作。
