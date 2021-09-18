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


type TextScan struct {
	Cmd  model.Cmd
}


func (scan *TextScan) Exec() int {

	var task  model.Task

	var tasks []model.Task

	var profile common.Profile

	var err error

	var status int

	if scan.Cmd.Content == "" {

		return  common.CONTENT_PASS
	}

	profile = common.Profile{AccessKeyId: viper.GetString("accessKeyId"), AccessKeySecret: viper.GetString("accessKeySecret")}

	clientInfo := model.ClinetInfo{Ip:"0.0.0.0"}

	task = model.Task{
		DataId:uuid.Rand().Hex(),
		Content: scan.Cmd.Content,
	}

	tasks = []model.Task{task}
	//
	scenes := strings.Split(viper.GetString("taxt.scenes"), `|`);

	fmt.Println("scenes:", viper.GetString("taxt.scenes"))

	bizData := model.BizData{ viper.GetString("taxt.biz_type"), scenes, tasks}

	fmt.Println("biz_type:", viper.GetString("taxt.biz_type"))

	var client common.IAliYunClient = common.DefaultClient{Profile: profile}

	reson :=  &model.TextReponse{}

	client_res :=  client.GetResponse(viper.GetString("taxt.path"), clientInfo, bizData)

	fmt.Println("aliTextApi:", client_res)

	err = json.Unmarshal([]byte(client_res), reson)


	if err != nil {
		fmt.Println(err.Error())
	}
	//0人工审核 1正常  2 违规
	status = 1
	lab_str := "";
	if reson.Code == 200 {
		for _,data := range reson.Data[0].Results  {

			if data.Suggestion == "review" {
				status = common.CONTENT_PENDING
				return status
			}
			if data.Suggestion == "block" {
				status = common.CONTENT_BLOCK
				lab_str = strings.Join([]string{lab_str, common.LABEL_STR[data.Label]}, ",")
				fmt.Println("违规内容:", lab_str)
				return status
			}

			if data.Suggestion == "pass" {
				status = common.CONTENT_PASS
			}
		}

		if lab_str != "" {
			fmt.Println("block_str:",lab_str)
		}
	} else {
		fmt.Println("内部错误")
	}
	return status
}