package common

const CONTENT_PASS = 1     //审核通过

const CONTENT_PENDING = 0  //人工审核

const CONTENT_BLOCK = 2    //审核文本违规


const (

	NORMAL   = "normal"
	SPAM     = "spam"
	AD       = "ad"
	POLITICS    = "politics"
	TERRORISM = "terrorism"
	ABUSE = "abuse"
	PRON = "pron"
	FLOOD = "flood"
	CONTRABAND = "contraband"

)

//机审建议后续操作
var  SUGGESTION_STR = map[string]string {
	"pass": "正常",
	"review": "人工审核",
	"block": "文本违规",
}

//文本垃圾检测结果的分类
var LABEL_STR = map[string]string {
	"normal": "正常文本",
	"spam": "含垃圾信息",
	"ad": "广告",
	"politics": "涉政",
	"terrorism": "暴恐",
	"abuse": "辱骂",
	"porn": "色情",
	"flood": "灌水",
	"contraband": "违禁",
	"meaningless": "无意义",
	"customized": "自定义",
}