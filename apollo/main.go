package main

import (
	"encoding/json"
	"fmt"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
	"github.com/zouyx/agollo/v4/storage"
	"time"
)

type Config struct {
	Name string `json:"name"`
}

var apolloClient *agollo.Client

func init() {
	var conf = &config.AppConfig{
		AppID:          "",
		Cluster:        "default",
		IP:             "",
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         "",
	}

	var err error
	apolloClient, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return conf, nil
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
}

func main() {
	cache := apolloClient.GetConfigCache("application")

	//增加变更的监听回调
	//apolloClient.AddChangeListener(change{})
	dbConf, err := cache.Get("base")
	if err != nil {
		panic(err)
	}

	return
	var cfg Config
	json.Unmarshal([]byte(dbConf.(string)), &cfg)
	fmt.Println(cfg)

	fmt.Println(apolloClient.GetValue("testkey"))

	cache.Set("testkey", "happy", 3)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	select {}
}

type change struct {
}

//OnChange 增加变更监控
func (c change) OnChange(event *storage.ChangeEvent) {
	fmt.Sprintf("notificationid %d namespace is %s ", event.NotificationID, event.Namespace)
	for k, v := range event.Changes {
		fmt.Println(k, v.ChangeType, v.OldValue, v.NewValue)
	}
}

//OnNewestChange 监控最新变更
func (c change) OnNewestChange(event *storage.FullChangeEvent) {
	fmt.Sprintf("OnNewestChange notificationid %d namespace is %s ", event.NotificationID, event.Namespace)
	for k, v := range event.Changes {
		fmt.Println("OnNewestChange", k, v)
	}
}
