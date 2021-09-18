package main

import (
	"content_security/common"
	"content_security/common/model"
	"content_security/common/redis"
	"content_security/handle"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func main(){
	//初始化配置
	viper.SetConfigName("conf.yml")   // 配置文件名
	viper.SetConfigType("yaml")       // 配置文件类型，可以是yaml、json、xml。。。
	viper.AddConfigPath("./conf")     // 配置文件路径
	err := viper.ReadInConfig()          // 读取配置文件信息
	if err != nil {
		println("配置解析失败")
	}
	fmt.Println("程序已启动")
	fmt.Println("当前redis配置", viper.GetString("redis_queue.host"))
	//调度
	schedule();
}

//调度器
func schedule()  {

	for  {
		//采用redis调度
		redisSchedule()
		continue
	}
}

//调度方式redis
func redisSchedule() {
	var cmd model.Cmd
	var textScan handle.TextScan
	var imageScan handle.ImageScan
	var url string
	cmd_str, err := redis.GetInstance().LPop(context.Background(), viper.GetString("redis_queue.key")).Result()

	if err != nil || cmd_str == "" {
		//休眠1秒
		time.Sleep(1000000000)
		return
	}
	//解析redis值
	err = json.Unmarshal([]byte(cmd_str), &cmd)
	if err != nil {
		log.Println("解析失败:",cmd_str)
		return
	}

	//文本审核
	textScan = handle.TextScan{
		Cmd: cmd,
	}
	//执行文本审核
	text_status := textScan.Exec();
	//图片审核
	imageScan = handle.ImageScan{
		Cmd: cmd,
	}
	//执行图片审核
	image_status := imageScan.Exec()

	
	if image_status == common.CONTENT_BLOCK || text_status == common.CONTENT_BLOCK {
		fmt.Println("httpurl:",cmd.Callback_block_url)
		url = cmd.Callback_block_url
	} else if image_status == common.CONTENT_PASS && text_status == common.CONTENT_PASS {
		fmt.Println("httpurl:",cmd.Callback_pass_url)
		url = cmd.Callback_pass_url
	} else {
		fmt.Println("httpurl:",cmd.Callback_pending_url)
		url = cmd.Callback_pending_url
	}
	callback(url)
	return
}

//回调业务逻辑
func callback(url string)  {
	var err error
	_, err = http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
	}
}

