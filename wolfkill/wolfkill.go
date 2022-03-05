// Package wolfkill 临时的狼人杀模块，狼人杀法官的简单实现
package wolfkill

import (
	"AWC-gateway/wechat"
	"strconv"
	"strings"
	"sync"
)

var lock sync.Mutex

// HandleMessage 处理用户消息
func HandleMessage(req wechat.ReceiveMessageRequest) string {
	lock.Lock()
	defer lock.Unlock()

	args := strings.Split(req.Content, " ")
	switch args[0] {
	case "help":
		return help()
	case "restart":
		return restart()
	case "rename":
		return rename(req)
	case "list":
		return list()
	case "log":
		return showLog(req)
	case "sit":
		return sit(req)
	case "set":
		return setRole(req)
	case "kill":
		return kill(req)
	case "check":
		return check(req)
	case "save":
		return save(req)
	case "poi":
		return poison(req)
	case "guard":
		return guard(req)
	case "learn":
		return learn(req)
	case "lucky?":
		return isLucky(req)
	case "who?":
		return who(req)
	case "day":
		return day()
	case "msg":
		return getDeathMsg()
	case "debug":
		return debug()
	}
	return ""
}

// analyseNumberCommand 解析单数字命令
func analyseNumberCommand(req wechat.ReceiveMessageRequest) (int, error) {
	args := strings.Split(req.Content, " ")
	if len(args) <= 1 {
		return 0, commandInvalid
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, commandInvalid
	}

	return num, nil
}

func init() {
	wechat.Handlers = append(wechat.Handlers, HandleMessage)
}
