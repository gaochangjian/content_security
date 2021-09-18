package handle

import (
	"content_security/common"
	"content_security/common/model"
	"content_security/uuid"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type ImageScan struct {
	Cmd model.Cmd
}

func (scan *ImageScan) Exec() int {
	var task  model.Task //任务

	var tasks []model.Task  //任务列表

	var profile common.Profile  //配置文件

	var err error

	var status int  //状态

	if len(scan.Cmd.Images) == 0 {
		return  common.CONTENT_PASS
	}

	profile = common.Profile{AccessKeyId: viper.GetString("accessKeyId"), AccessKeySecret: viper.GetString("accessKeySecret")}

	clientInfo := model.ClinetInfo{Ip:"123.0.0.1"}

	// 构造请求数据
	scenes := strings.Split(viper.GetString("image.scenes"), `|`);

	fmt.Println("scenes:", viper.GetString("image.scenes"))
	//添加图片
	for _,image := range scan.Cmd.Images {
		task = model.Task{DataId:uuid.Rand().Hex(), Url:image}
		tasks = append(tasks, task)
	}

	bizData := model.BizData{ viper.GetString("image.biz_type"), scenes, tasks}

	fmt.Println("biz_type:", viper.GetString("image.biz_type"))

	var client common.IAliYunClient = common.DefaultClient{Profile: profile}

	reson :=  &model.SceneReponse{}

	client_res :=  client.GetResponse(viper.GetString("image.path"), clientInfo, bizData)

	fmt.Println("aliImageApi:", client_res)

	err = json.Unmarshal([]byte(client_res), reson)
	if err != nil {
		fmt.Println(err.Error())
	}
	//0人工审核 1正常  2 违规
	status = 0

	lab_str := ""

	if reson.Code == 200 {
		//遍历每一个task，获取对应数据
		for _,data := range reson.Data  {
			for _, result := range data.Results {
				if result.Suggestion == "review" {
					status = common.CONTENT_PENDING
					return status
				}
				if result.Suggestion == "block" {
					status = common.CONTENT_BLOCK
					lab_str = strings.Join([]string{lab_str, result.Label}, ",")
					fmt.Println("违规内容:", lab_str)
					return status
				}
				if result.Suggestion == "pass" {
					status = common.CONTENT_PASS
				}
			}
		}
	} else {
		fmt.Println("内部错误")
	}
	return  status
}