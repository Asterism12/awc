package wolfkill

import (
	"AWC-gateway/wechat"
	"fmt"
	"sort"
	"strings"
)

var roles = map[int]string{
	0: "平民",
	1: "狼人",
	2: "预言家",
	3: "女巫",
	4: "猎人",
	5: "守卫",
	6: "混子",
}

// Set 游戏局
type Set struct {
	desk         [21]Player     // 坐在桌旁的玩家
	log          []string       // 运行日志
	players      map[string]int // 玩家ID及对应的座位号
	isStart      bool           // 是否已开始
	killTarget   int            // 狼刀
	hasChecked   bool           // 预言家晚上是否已经预言过
	hasDrug      bool           // 女巫晚上是否用过药物
	hasSaved     bool           // 女巫是否救过人
	hasPoison    bool           // 女巫是否毒过人
	saveTarget   int            // 女巫解药目标
	poisonTarget int            // 女巫毒药目标
	guardTarget  int            // 守卫目标
	hasLearned   bool           // 是否已经学习过
	lastDeathMsg string         // 上一晚的死亡讯息
}

var set Set

// restart 开始新的游戏
func restart() string {
	set = Set{
		log:     []string{"重新开始游戏"},
		players: make(map[string]int),
	}
	return "已清空房间"
}

// list 列出局内玩家
func list() string {
	var builder strings.Builder
	for i, player := range set.desk {
		if player.id == "" {
			continue
		}
		nickname := getNick(player.id)

		builder.WriteString(fmt.Sprintf("%d %s\n", i, nickname))
	}
	return builder.String()
}

// showLog 查看日志
func showLog(req wechat.EventRequest) string {
	seatNum := set.players[req.FromUserName]
	if seatNum == 0 {
		return "你不在游戏中"
	}

	var builder strings.Builder
	for _, s := range set.log {
		builder.WriteString(fmt.Sprintf("%s\n", s))
	}

	set.log = append(set.log, fmt.Sprintf("%d 号玩家查看了运行日志", seatNum))

	return builder.String()
}

// sit 玩家坐到某个位置
func sit(req wechat.EventRequest) string {
	if set.isStart {
		return "游戏已经开始了"
	}

	seatNum, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}
	if seatNum <= 0 || seatNum > 20 {
		return "合法的座位号：1-20"
	}

	if seatBefore, ok := set.players[req.FromUserName]; ok {
		set.desk[seatBefore] = Player{}
	}

	set.desk[seatNum] = Player{
		id:   req.FromUserName,
		role: roles[0],
	}
	set.players[req.FromUserName] = seatNum

	return fmt.Sprintf("你坐到了 %d 号位置", seatNum)
}

// setRole 设置自己的角色
func setRole(req wechat.EventRequest) string {
	if set.isStart {
		return "游戏已经开始了"
	}

	seatNum := set.players[req.FromUserName]
	if seatNum == 0 {
		return "你还没坐在座位上"
	}

	roleNum, err := analyseNumberCommand(req)
	if err != nil {
		return ""
	}

	role := roles[roleNum]
	set.desk[seatNum].role = role
	return fmt.Sprintf("你的角色设置为 %s", role)
}

// isInGame 玩家是否在已开始的游戏中
func isInGame(id string) bool {
	seatNum := set.players[id]
	return seatNum != 0 && set.isStart
}

// isValidTarget 目标是否合法
func isValidTarget(i int) bool {
	return i > 0 && i < 20 && set.desk[i].id != ""
}

// day 结算晚上的行动
func day() string {
	deathPlayers := whoIsDie()
	set.lastDeathMsg = fmt.Sprintf("%v 玩家死亡", deathPlayers)

	set.killTarget = 0

	set.poisonTarget = 0
	set.saveTarget = 0
	set.hasDrug = false

	set.guardTarget = 0

	set.hasChecked = false

	set.isStart = true
	set.log = append(set.log, "新的一天开始了")

	return "结算完成，使用msg命令查看死亡信息"
}

// whoIsDie 检查今晚死亡状态
func whoIsDie() []int {
	deathPlayers := make(map[int]bool)
	if set.poisonTarget != 0 {
		deathPlayers[set.poisonTarget] = true
	}
	if set.killTarget != 0 {
		if set.saveTarget != set.killTarget && set.guardTarget != set.killTarget {
			deathPlayers[set.killTarget] = true
		}

		// 同守同救
		if set.saveTarget == set.killTarget && set.guardTarget == set.killTarget {
			deathPlayers[set.killTarget] = true
		}
	}

	var result []int
	for player := range deathPlayers {
		result = append(result, player)
	}
	sort.Ints(result)
	return result
}

// getDeathMsg 获取昨晚的死亡讯息
func getDeathMsg() string {
	return set.lastDeathMsg
}

func debug() string {
	return fmt.Sprintf("%+v", set)
}

func init() {
	restart()
}
