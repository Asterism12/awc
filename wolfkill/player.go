package wolfkill

import (
	"AWC-gateway/wechat"
	"strings"
)

const helpText = `指令列表(标*会记录在日志中)：
普通指令
help ----查看帮助
list ----列出房间人员
rename --改名
msg -----查看昨晚死亡结果
sit n ---坐到n号位置
set n ---设置自己身份

角色指令
kill n --刀n号玩家*
check n -查验n号玩家*
save ----使用解药*
poi n ---对n号玩家使用毒药*
who? ----查询狼人目标(女巫)
guard n -守卫n号玩家*
learn n -以n号玩家为榜样*
lucky? --查询自己是否被毒(猎人)

特殊指令
restart -创建新的游戏*
day -----结算晚上的行动/开始游戏*
log -----查看运行日志*


身份列表：
0 平民(默认)
1 狼人
2 预言家
3 女巫
4 猎人
5 守卫
6 混子`

const defaultName = "jerry"

var nicknames = make(map[string]string)

// help 帮助信息
func help() string {
	return helpText
}

// rename 更改昵称
func rename(req wechat.ReceiveMessageRequest) string {
	args := strings.Split(req.Content, " ")
	if len(args) <= 1 {
		return "请输入新的名字"
	}

	nicknames[req.FromUserName] = args[1]
	return "你的昵称已改为 " + args[1]
}

func getNick(id string) string {
	if nick, ok := nicknames[id]; !ok {
		return defaultName
	} else {
		return nick
	}
}

// Player 房间中的玩家信息
type Player struct {
	id   string
	role string
}
