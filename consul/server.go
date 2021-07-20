package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
)

const (
	consulAddress = "127.0.0.1:8500"
	localIP       = "127.0.0.1"
	localPort     = 3001
)

func consulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "go-example"
	registration.Name = "go-example-http"//根据这个名称来找这个服务
	registration.Port = localPort
	registration.Tags = []string{"v1.1"}//这个就是一个标签，可以根据这个来找这个服务，相当于V1.1这种
	registration.Address = localIP


	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"                         //超时
	check.Interval = "5s"                        //健康检查频率
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到 consul
	err = client.Agent().ServiceRegister(registration)
}

//Handler 3001
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you are visiting health check api:3001"))
}

//ServerLoad 启动
func ServerLoad() {
	consulRegister()
	//定义一个http接口
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}

func main()  {
	ServerLoad()
}