package lowSocks

import (
	"log"
	"net"
	"fmt"
	"github.com/gwuhaolin/lightsocks/local"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
)

const (
	DefaultListenAddr = ":7448"
)

func StartClientSocks(){
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 解析配置
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动 local 端并监听
	lsLocal := local.New(password, listenAddr, remoteAddr)
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
远程服务地址 remote：
%s
密码 password：
%s
	`, listenAddr, remoteAddr, password))
		log.Printf("socks-local: 启动成功 监听在 %s\n", listenAddr.String())
	}))
}
