package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// IServer 的接口实现，定义一个Server的服务类
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[start] Server Listenner at IP:%s,Port:%d,is starting\n", s.IP, s.Port)

	go func() {
		//1，获取一个TCP的ADDR
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err.Error())
		}
		//2.尝试监听服务器的地址
		tcpListenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IP, "err ", err.Error())
		}

		fmt.Println("start Zinx server succ, ", s.Name, " succ, Listenning...")
		//3.阻塞等待客户端连接，处理客户端连接业务（读写）
		for {
			conn, err := tcpListenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//已经与客户端建立连接，做业务
			go func() {
				for {
					buf := make([]byte, 51)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err: ", err.Error())
						continue
					}
					fmt.Printf("recv client buf %s,cnt %d\n", buf, cnt)
					//回显
					if _, err = conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err: ", err.Error())
						continue
					}

				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	//TODO 服务器的资源状态或者一些已经开辟的连接信息 进行停止或者回收
}

func (s *Server) Server() {
	//启动sever功能
	s.Start()

	//TODO 做一些启动服务器之后的一些额外业务。 start和Sever分离，在Server阻塞方便扩展，不要在start阻塞

	//阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IP:        "0.0.0.0",
		IPVersion: "tcp4",
		Port:      8999,
	}
}
