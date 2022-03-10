package wolfkill

import (
	"AWC-gateway/wechat"
	"fmt"
)

// kill 狼人杀死玩家
func kill(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "狼人" {
		return "你不是狼人"
	}

	killTarget, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if !isValidTarget(killTarget) {
		return fmt.Sprintf("不合法的目标，检查 %d 号座位是否有人", killTarget)
	}

	set.killTarget = killTarget
	set.log = append(set.log, fmt.Sprintf("%d 号改变狼刀为 %d", seatNum, killTarget))
	return fmt.Sprintf("狼刀被改变为 %d 号玩家", killTarget)
}

// check 查验玩家
func check(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "预言家" {
		return "你不是预言家"
	}

	if set.hasChecked {
		return "今晚好累"
	}

	checkTarget, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if !isValidTarget(checkTarget) {
		return fmt.Sprintf("不合法的目标，检查 %d 号座位是否有人", checkTarget)
	}

	set.hasChecked = true
	var role string
	switch set.desk[checkTarget].role {
	case "狼人":
		role = "狼人"
	default:
		role = "好人"
	}

	set.log = append(set.log, fmt.Sprintf("%d 号查验 %d 号玩家", seatNum, checkTarget))

	return fmt.Sprintf("%d 号玩家身份是 %s", checkTarget, role)
}

// save 使用解药
func save(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "女巫" {
		return "你不是女巫"
	}

	if set.hasDrug {
		return "今晚好累"
	}
	if set.hasSaved {
		return "已经使用过解药了"
	}

	set.hasDrug = true
	set.hasSaved = true
	set.saveTarget = set.killTarget
	set.log = append(set.log, fmt.Sprintf("%d 号对 %d 号玩家使用解药", seatNum, set.saveTarget))

	return fmt.Sprintf("对 %d 号玩家使用解药", set.saveTarget)
}

// poison 使用毒药
func poison(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "女巫" {
		return "你不是女巫"
	}

	if set.hasDrug {
		return "今晚好累"
	}
	if set.hasPoison {
		return "已经使用过毒药了"
	}

	poisonTarget, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if !isValidTarget(poisonTarget) {
		return fmt.Sprintf("不合法的目标，检查 %d 号座位是否有人", poisonTarget)
	}

	set.hasDrug = true
	set.hasPoison = true
	set.poisonTarget = poisonTarget
	set.log = append(set.log, fmt.Sprintf("%d 号对 %d 号玩家使用毒药", seatNum, set.poisonTarget))

	return fmt.Sprintf("对 %d 号玩家使用毒药", poisonTarget)
}

// guard 守卫
func guard(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "守卫" {
		return "你不是守卫"
	}

	guardTarget, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if !isValidTarget(guardTarget) {
		return fmt.Sprintf("不合法的目标，检查 %d 号座位是否有人", guardTarget)
	}

	set.guardTarget = guardTarget
	set.log = append(set.log, fmt.Sprintf("%d 号守卫 %d 号玩家", seatNum, set.guardTarget))

	return fmt.Sprintf("守卫 %d 号玩家", guardTarget)
}

// learn 以某人为榜样
func learn(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "混子" {
		return "你不是混子"
	}

	if set.hasLearned {
		return "你已经树立榜样了"
	}

	learnTarget, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if !isValidTarget(learnTarget) {
		return fmt.Sprintf("不合法的目标，检查 %d 号座位是否有人", learnTarget)
	}

	set.hasLearned = true
	set.log = append(set.log, fmt.Sprintf("%d 号认定 %d 号玩家为榜样", seatNum, learnTarget))

	return fmt.Sprintf("守卫 %d 号玩家", learnTarget)
}

// isLucky 猎人是否被女巫毒中
func isLucky(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "猎人" {
		return "你不是猎人"
	}

	if set.poisonTarget == seatNum {
		return "你被女巫下了毒药，开不出枪"
	} else {
		return "状态良好，随时开枪"
	}
}

// who 检查狼人的目标
func who(req wechat.EventRequest) string {
	if !isInGame(req.FromUserName) {
		return "你不在游戏中或游戏还未开始"
	}

	seatNum := set.players[req.FromUserName]
	if set.desk[seatNum].role != "女巫" {
		return "你不是女巫"
	}

	if set.killTarget == 0 {
		return "狼人还没有确定目标，或者今晚空刀"
	} else {
		return fmt.Sprintf("今晚狼人的目标是 %d 号玩家", set.killTarget)
	}
}
