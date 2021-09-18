package model

type ClinetInfo struct {
	SdkVersion 	string	`json:"sdkVersion"`
	CfgVersion 	string	`json:"cfgVersion"`
	UserType 	string	`json:"userType"`
	UserId 		string	`json:"userId"`
	UserNick 	string	`json:"userNick"`
	Avatar 		string	`json:"avatar"`
	Imei		string	`json:"imei"`
	Imsi		string	`json:"imsi"`
	Umid		string	`json:"umid"`
	Ip		string	`json:"ip"`
	Os		string	`json:"os"`
	Channel		string	`json:"channel"`
	HostAppName	string	`json:"hostAppName"`
	HostPackage	string	`json:"hostPackage"`
	HostVersion	string	`json:"hostVersion"`
}

type Task struct {
	DataId string		`json:"dataId"`
	Content    string		`json:"content"`
	Url    string		`json:"url"`
}


//传输数据
type BizData struct {
	BizType	string		`json:"bizType"`
	Scenes 	[]string	`json:"scenes"`
	Tasks	[]Task	    `json:"tasks"`
}

//输入参数
type Cmd struct {
	Content               string     `json:"content"`
	Images                []string   `json:"images"`
	Callback_pass_url     string     `json:"callback_pass_url"`
	Callback_block_url    string     `json:"callback_block_url"`
	Callback_pending_url  string     `json:"callback_pending_url"`
}

type SceneReponse struct {
	Code int `json:"code"`
	Data []struct {
		Code   int    `json:"code"`
		DataID string `json:"dataId"`
		Extras struct {
		} `json:"extras"`
		Msg     string `json:"msg"`
		Results []struct {
			Label      string  `json:"label"`
			Rate       float64 `json:"rate"`
			Scene      string  `json:"scene"`
			Suggestion string  `json:"suggestion"`
		} `json:"results"`
		TaskID string `json:"taskId"`
		URL    string `json:"url"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

type TextReponse struct {
	Code int `json:"code"`
	Data []struct {
		Code            int    `json:"code"`
		Content         string `json:"content"`
		DataID          string `json:"dataId"`
		FilteredContent string `json:"filteredContent"`
		Msg             string `json:"msg"`
		Results         []struct {
			Details []struct {
				Contexts []struct {
					Context   string `json:"context"`
					Positions []struct {
						EndPos   int `json:"endPos"`
						StartPos int `json:"startPos"`
					} `json:"positions"`
				} `json:"contexts"`
				Label string `json:"label"`
			} `json:"details"`
			Label      string  `json:"label"`
			Rate       float64 `json:"rate"`
			Scene      string  `json:"scene"`
			Suggestion string  `json:"suggestion"`
		} `json:"results"`
		TaskID string `json:"taskId"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}