package main

import (
	"math/rand"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

var Map [][]int
var maxDepth int
var count int
var bad []int

type Move struct {
	fromX int
	fromY int
	toX   int
	toY   int
	score int
	key   int
	eat   int
}

var scores []int
var aiMove Move
var cache = make(map[string]int)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func InitBoard() {
	Map = make([][]int, 10)
	for j := 0; j < 10; j++ {
		Map[j] = make([]int, 9)
	}
	Map[0] = []int{-3, -4, -5, -6, -7, -6, -5, -4, -3}
	Map[1] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	Map[2] = []int{0, -2, 0, 0, 0, 0, 0, -2, 0}
	Map[3] = []int{-1, 0, -1, 0, -1, 0, -1, 0, -1}
	Map[4] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	Map[5] = []int{0, 0, 0, 0, 1, 0, 0, 0, 0}
	Map[6] = []int{1, 0, 1, 0, 0, 0, 1, 0, 1}
	Map[7] = []int{0, 2, 0, 0, 2, 0, 0, 0, 0}
	Map[8] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	Map[9] = []int{3, 4, 5, 6, 7, 6, 5, 4, 3}
}

var Bing = [10][9]int{
	{9, 9, 9, 11, 13, 11, 9, 9, 9},
	{19, 24, 34, 42, 44, 42, 34, 24, 19},
	{19, 24, 32, 37, 37, 37, 32, 24, 19},
	{19, 23, 27, 29, 30, 29, 27, 23, 19},
	{14, 18, 20, 27, 29, 27, 20, 18, 14},
	{7, 0, 13, 0, 16, 0, 13, 0, 7},
	{7, 0, 7, 0, 15, 0, 7, 0, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var Pao = [10][9]int{
	{100, 100, 96, 91, 90, 91, 96, 100, 100},
	{98, 98, 96, 92, 89, 92, 96, 98, 98},
	{97, 97, 96, 91, 92, 91, 96, 97, 97},
	{96, 99, 99, 98, 100, 98, 99, 99, 96},
	{96, 96, 96, 96, 100, 96, 96, 96, 96},
	{95, 96, 99, 96, 100, 96, 99, 96, 95},
	{96, 96, 96, 96, 96, 96, 96, 96, 96},
	{97, 96, 100, 99, 101, 99, 100, 96, 97},
	{96, 97, 98, 98, 98, 98, 98, 97, 96},
	{96, 96, 97, 99, 99, 99, 97, 96, 96},
}

var Jiang = [10][9]int{
	{0, 0, 0, 12000, 12000, 12000, 0, 0, 0},
	{0, 0, 0, 12000, 12000, 12000, 0, 0, 0},
	{0, 0, 0, 12000, 12000, 12000, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 9900, 9900, 9900, 0, 0, 0},
	{0, 0, 0, 9930, 9950, 9930, 0, 0, 0},
	{0, 0, 0, 9950, 10000, 9950, 0, 0, 0},
}

var Shi = [10][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 20, 0, 20, 0, 0, 0},
	{0, 0, 0, 0, 20, 0, 0, 0, 0},
	{0, 0, 0, 20, 0, 20, 0, 0, 0},
}

var Xiang = [10][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{20, 0, 0, 0, 25, 0, 0, 0, 20},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 25, 0, 0, 0, 25, 0, 0},
}

var Ma = [10][9]int{
	{90, 90, 90, 96, 90, 96, 90, 90, 90},
	{90, 96, 103, 97, 94, 97, 103, 96, 90},
	{92, 98, 99, 103, 99, 103, 99, 98, 92},
	{93, 108, 100, 107, 100, 107, 100, 108, 93},
	{90, 100, 99, 103, 104, 103, 99, 100, 90},
	{90, 98, 101, 102, 103, 102, 101, 98, 90},
	{92, 94, 98, 95, 98, 95, 98, 94, 92},
	{90, 92, 95, 95, 92, 95, 95, 92, 90},
	{85, 90, 92, 93, 78, 93, 92, 90, 85},
	{88, 88, 90, 88, 90, 88, 90, 88, 88},
}

var Ju = [10][9]int{
	{206, 208, 207, 213, 214, 213, 207, 208, 206},
	{206, 212, 209, 216, 233, 216, 209, 212, 206},
	{206, 208, 207, 214, 216, 214, 207, 208, 206},
	{206, 213, 213, 216, 216, 216, 213, 213, 206},
	{208, 211, 211, 214, 215, 214, 211, 211, 208},
	{208, 212, 212, 214, 215, 214, 212, 212, 208},
	{204, 209, 204, 212, 214, 212, 204, 209, 204},
	{198, 208, 204, 212, 212, 212, 204, 208, 198},
	{200, 208, 206, 212, 200, 212, 206, 208, 200},
	{194, 206, 204, 212, 200, 212, 204, 206, 194},
}

func revMap() [][]int {
	var arr [][]int
	for i := len(Map) - 1; i >= 0; i-- {
		arr = append(arr, Map[i])
	}
	return arr
}

func generateMoves(player int) []Move {
	var moves []Move
	// 遍历棋盘上的每个位置
	for j := 0; j < 10; j++ {
		for i := 0; i < 9; i++ {
			man := Map[j][i]

			// 判断当前位置是否是当前玩家的棋子
			if man*player > 0 {

				// 生成当前位置的所有可能移动
				switch abs(man) {
				case 1: // 兵
					moves = append(moves, generateBingMoves(i, j, player)...)
				case 2: // 炮
					moves = append(moves, generatePaoMoves(i, j, player)...)
				case 3: // 车
					moves = append(moves, generateJuMoves(i, j, player)...)
				case 4: // 马
					moves = append(moves, generateMaMoves(i, j, player)...)
				case 5: // 象
					moves = append(moves, generateXiangMoves(i, j, player)...)
				case 6: // 士
					moves = append(moves, generateShiMoves(i, j, player)...)
				case 7: // 将
					moves = append(moves, generateJiangMoves(i, j, player)...)
				}
			}
		}
	}

	// if player < 0 {
	// 	Map = tempMap
	// 	// for k, m := range moves {
	// 	// 	// showLog("move--", m.fromX, m.fromY, m.toX, m.toY, m.key)
	// 	// 	if m.fromY <= 9 && m.fromY >= 0 && m.toY >= 0 && m.toY <= 9 {
	// 	// 		moves[k].fromY = 9 - m.fromY
	// 	// 		moves[k].toY = 9 - m.toY
	// 	// 	}
	// 	// }
	// }

	return moves
}

func checkxy(x, y int) bool {
	return x >= 0 && x < 9 && y >= 0 && y < 10
}

// 生成兵的移动
func generateBingMoves(x, y, player int) []Move {
	var moves []Move

	// 兵的移动方向
	directions := [][]int{{0, -1}}

	// 如果兵已经过河，可以向左右移动
	if y < 5 {
		directions = [][]int{{0, -1}, {-1, 0}, {1, 0}}
	}

	for _, dir := range directions {
		x2, y2 := x+dir[0], y+dir[1]

		// 判断目标位置是否在棋盘内
		if checkxy(x2, y2) {
			// 判断目标位置是否为空或者是敌方棋子
			if Map[y2][x2]*player <= 0 {
				moves = append(moves, Move{
					fromX: x,
					fromY: y,
					toX:   x2,
					toY:   y2,
					key:   Map[y][x],
					score: Bing[y2][x2],
					eat:   Map[y2][x2],
				})
			}
		}
	}

	return moves
}

// 生成炮的移动
func generatePaoMoves(x, y, player int) []Move {
	moves := []Move{}

	// 炮的移动方向
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		capture := false

		for i := 1; ; i++ {
			newX, newY := x+i*dx, y+i*dy
			if !checkxy(newX, newY) {
				break
			}

			if Map[newY][newX] != 0 {
				if !capture {
					capture = true
				} else {
					if Map[newY][newX]*player < 0 {
						moves = append(moves, Move{
							fromX: x,
							fromY: y,
							toX:   newX,
							toY:   newY,
							key:   Map[y][x],
							score: Pao[newY][newX],
							eat:   Map[newY][newX],
						})
					}
					break
				}
			} else {
				if !capture {
					moves = append(moves, Move{
						fromX: x,
						fromY: y,
						toX:   newX,
						toY:   newY,
						key:   Map[y][x],
						score: Pao[newY][newX],
					})
				}
			}
		}
	}

	return moves
}

// 生成车的移动
func generateJuMoves(x, y, player int) []Move {
	moves := []Move{}

	// 车的移动方向
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]

		for i := 1; ; i++ {
			nx, ny := x+i*dx, y+i*dy
			if !checkxy(nx, ny) {
				break
			}

			if Map[ny][nx]*player <= 0 {
				moves = append(moves, Move{
					fromX: x,
					fromY: y,
					toX:   nx,
					toY:   ny,
					key:   Map[y][x],
					score: Ju[ny][nx],
					eat:   Map[ny][nx],
				})
			}
			if Map[ny][nx]*player != 0 {
				break
			}
		}
	}

	return moves
}

func generateMaMoves(x, y, player int) []Move {
	moves := []Move{}

	// 马的移动方向
	directions := [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		newX, newY := x+dx, y+dy
		legX, legY := x+dx/2, y+dy/2

		// 判断马腿位置是否为空
		if checkxy(legX, legY) && Map[legY][legX] == 0 {
			// 判断目标位置是否为空或者为敌方棋子
			if checkxy(newX, newY) && Map[newY][newX]*player <= 0 {
				moves = append(moves, Move{
					fromX: x,
					fromY: y,
					toX:   newX,
					toY:   newY,
					key:   Map[y][x],
					score: Ma[newY][newX],
					eat:   Map[newY][newX],
				})
			}
		}
	}

	return moves
}

func generateXiangMoves(x, y, player int) []Move {
	moves := []Move{}

	// 象的移动方向
	directions := [][]int{{-2, -2}, {-2, 2}, {2, -2}, {2, 2}}

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]

		// 计算象眼和目标位置
		eyeX, eyeY := x+dx/2, y+dy/2
		targetX, targetY := x+dx, y+dy

		// 检查象眼和目标位置是否在边界内
		if checkxy(eyeX, eyeY) && checkxy(targetX, targetY) && targetY > 4 {
			// 判断象眼位置是否为空
			if Map[eyeY][eyeX] == 0 {
				// 判断目标位置是否为空或为敌方棋子
				if Map[targetY][targetX]*player <= 0 {
					moves = append(moves, Move{
						fromX: x,
						fromY: y,
						toX:   targetX,
						toY:   targetY,
						key:   Map[y][x],
						score: Xiang[targetY][targetX],
						eat:   Map[targetY][targetX],
					})
				}
			}
		}
	}

	return moves
}

// 生成士的移动
func generateShiMoves(x, y, player int) []Move {
	moves := []Move{}

	// 士的移动方向
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for _, dir := range directions {
		x2, y2 := x+dir[0], y+dir[1]

		// 判断目标位置是否在九宫格内
		if x2 >= 3 && x2 <= 5 && y2 >= 7 && y2 <= 9 {
			// 判断目标位置是否为空或敌方棋子
			if Map[y2][x2]*player <= 0 {
				moves = append(moves, Move{
					fromX: x,
					fromY: y,
					toX:   x2,
					toY:   y2,
					key:   Map[y][x],
					score: Shi[y2][x2],
					eat:   Map[y2][x2],
				})
			}
		}
	}

	return moves
}

// 生成将的移动
func generateJiangMoves(x, y, player int) []Move {
	moves := []Move{}

	//判断是否飞将
	j1x, j1y := 0, 0
	j2x, j2y := 0, 0
	for j := 0; j < 10; j++ {
		for i := 0; i < 9; i++ {
			if abs(Map[j][i]) == 7 {
				if j < 5 {
					j2x, j2y = i, j
				} else {
					j1x, j1y = i, j
				}
			}
		}
	}
	if j1x == j2x {
		fj := 1
		for i := j2y + 1; i < j1y; i++ {
			if Map[i][j1x] != 0 {
				fj = 0
			}
		}
		if fj == 1 {
			moves = append(moves, Move{
				fromX: j1x,
				fromY: j1y,
				toX:   j2x,
				toY:   j2y,
				key:   Map[j1y][j1x],
				score: Jiang[j2y][j2x],
			})
		}
	}

	// 将的移动方向
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		x2, y2 := x+dir[0], y+dir[1]

		// 判断目标位置是否在九宫格内
		if x2 >= 3 && x2 <= 5 && y2 >= 7 && y2 <= 9 {
			// 判断目标位置是否为敌方棋子
			if Map[y2][x2]*player <= 0 {
				moves = append(moves, Move{
					fromX: x,
					fromY: y,
					toX:   x2,
					toY:   y2,
					key:   Map[y][x],
					score: Jiang[y2][x2],
					eat:   Map[y2][x2],
				})
			}
		}
	}

	return moves
}

func evaluate() int {
	score := 0

	// 遍历棋盘上的每个位置
	for j := 0; j < 10; j++ {
		for i := 0; i < 9; i++ {
			man := Map[j][i]
			if man == 0 {
				continue
			}
			player := 1
			y := j
			if man < 0 {
				player = -1
				y = 9 - j
			}

			switch abs(man) {
			case 1: // 兵
				score += Bing[y][i] * player
			case 2: // 炮
				score += (Pao[y][i] + 50) * player
			case 3: // 车
				score += (Ju[y][i] + 100) * player
			case 4: // 马
				score += Ma[y][i] * player
			case 5: // 象
				score += Xiang[y][i] * player
			case 6: // 士
				score += Shi[y][i] * player
			case 7: // 将
				score += Jiang[y][i] * player
			}
		}
	}

	// showLog("score", score)
	return score
}

func boardToString() string {
	// boardBytes, _ := json.Marshal(Map)
	// return string(boardBytes)
	str := ""
	for j := 0; j < 10; j++ {
		for i := 0; i < 9; i++ {
			str += strconv.Itoa(Map[j][i]) + ","
		}
	}
	return str
}

var pass = 0

func alphaBeta(depth, alpha, beta, player int) int {

	// 到达指定深度或游戏结束时，计算当前局面得分并返回
	if depth == 0 {
		count++
		return -evaluate() - rand.Intn(10)
	}

	// boardStr := ""
	// if maxDepth > 4 && depth > 1 && depth < maxDepth {
	// 	// 将局面转换为字符串表示
	// 	boardStr = boardToString() + "@dp" + strconv.Itoa(depth*player)
	// 	// 检查缓存中是否存在该局面的评估结果
	// 	if score, ok := cache[boardStr]; ok {
	// 		pass++
	// 		// showLog("pass...", score, depth, pass)
	// 		return score + rand.Intn(10)
	// 	}
	// }

	// 根据当前玩家是最大化玩家还是最小化玩家来确定搜索的方向
	if player == -1 { // 最大化玩家

		tempMap := Map
		Map = revMap()
		// 生成当前玩家的所有合法移动
		moves := generateMoves(player)
		Map = tempMap

		// if depth == maxDepth {
		// 	arr := [][]int{}
		// 	for _, m := range moves {
		// 		arr = append(arr, []int{m.fromY, m.fromX, m.toY, m.toX, m.key, m.score})
		// 	}
		// 	str, _ := json.Marshal(arr)
		// 	showLog("moves", player, depth, string(str))
		// }

		// 对每个合法移动进行搜索
		for index, move := range moves {
			if maxDepth > 4 && depth < maxDepth-1 && index&1 == 0 {
				continue
			}
			move.fromY = 9 - move.fromY
			move.toY = 9 - move.toY
			if depth == maxDepth && len(bad) == 4 {
				if bad[1] == move.fromX && bad[0] == move.fromY && bad[3] == move.toX && bad[2] == move.toY {
					continue
				}
			}
			// 执行移动操作
			makeMove(&move)

			// str, _ := json.Marshal(Map)
			// showLog("makeMove", string(str))
			// 调用alphaBeta函数进行搜索，并更新alpha的值
			score := alphaBeta(depth-1, alpha, beta, -player)

			if score > alpha {
				alpha = score

				// 如果当前深度是最大深度，保存最佳移动
				// if depth == maxDepth && (rand.Intn(2) == 0 || len(aiMove.fromto) == 0) {
				if depth == maxDepth {
					aiMove = move
					aiMove.score = score
				}
			}

			// 撤销移动操作
			undoMove(&move)

			// showLog("alpha, beta, score, depth", alpha, beta, score, depth)
			if depth == maxDepth {
				scores = append(scores, score)
				// if len(scores) > 20 {
				// 	scores = scores[20:]
				// }

				// str, _ := json.Marshal([]int{move.fromY, move.fromX, move.toY, move.toX})
				// showLog("depth", depth, string(str), score, "moves", len(moves))
			}

			// 执行剪枝
			if alpha >= beta {
				break
			}
		}

		// if boardStr != "" {
		// 	cache[boardStr] = alpha
		// }

		return alpha
	} else { // 最小化玩家

		// 生成当前玩家的所有合法移动
		moves := generateMoves(player)
		// if depth == maxDepth-1 {
		// 	arr := [][]int{}
		// 	for _, m := range moves {
		// 		arr = append(arr, []int{m.fromY, m.fromX, m.toY, m.toX, m.key, m.score})
		// 	}
		// 	str, _ := json.Marshal(arr)
		// 	showLog("moves", player, depth, string(str))
		// }
		// 对每个合法移动进行搜索
		for _, move := range moves {
			if Map[move.toY][move.toX] == -7 {
				// str, _ := json.Marshal([]int{move.fromY, move.fromX, move.toY, move.toX})
				// showLog("dead", string(str), depth)
				return -4444
				// continue
			}
			// 执行移动操作
			makeMove(&move)
			// 调用alphaBeta函数进行搜索，并更新beta的值
			score := alphaBeta(depth-1, alpha, beta, -player)
			// if move.eat == -3 {
			// 	score -= 200
			// }
			// if move.eat == -4 && move.eat == -2 {
			// 	score -= 100
			// }
			if score < beta {
				beta = score
			}
			// 撤销移动操作
			undoMove(&move)

			// if depth == maxDepth-1 {
			// 	scores = append(scores, score)
			// 	// if len(scores) > 20 {
			// 	// 	scores = scores[20:]
			// 	// }

			// 	str, _ := json.Marshal([]int{move.fromY, move.fromX, move.toY, move.toX})
			// 	showLog("--depth", depth, string(str), score, "moves", len(moves))
			// }

			// 执行剪枝
			if alpha >= beta {
				break
			}
		}

		// if boardStr != "" {
		// 	cache[boardStr] = beta
		// }

		return beta
	}
}

func makeMove(move *Move) {
	// 保存移动前的棋子
	fromPiece := Map[move.fromY][move.fromX]
	toPiece := Map[move.toY][move.toX]

	// 执行移动操作
	Map[move.toY][move.toX] = fromPiece
	Map[move.fromY][move.fromX] = 0

	// 如果有吃子操作，将目标位置的棋子置为0
	if toPiece != 0 {
		move.eat = toPiece
	} else {
		move.eat = 0
	}
}
func undoMove(move *Move) {
	// 恢复移动操作
	Map[move.fromY][move.fromX] = Map[move.toY][move.toX]
	Map[move.toY][move.toX] = move.eat
}

func showLog(s ...interface{}) {
	js.Global().Get("console").Call("log", s...)
}

// func showInfo(s interface{}) {
// 	js.Global().Set("paceinfo", s)
// }

func getMoves(this js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return nil
	}

	m := args[0].String()
	if m != "" {
		arr := strings.Split(m, ";")
		for i, v := range arr {
			a := strings.Split(v, ",")
			for ii, vv := range a {
				Map[i][ii], _ = strconv.Atoi(vv)
			}
		}
	}
	i := int(args[1].Float())
	j := int(args[2].Float())

	var moves []Move
	var tempMap [][]int
	man := Map[j][i]
	player := 1

	// 判断当前位置是否是当前玩家的棋子
	if man < 0 {
		tempMap = Map
		Map = revMap()
		player = -1
		j = 9 - j
	}

	// 生成当前位置的所有可能移动
	switch abs(man) {
	case 1: // 兵
		moves = generateBingMoves(i, j, player)
	case 2: // 炮
		moves = generatePaoMoves(i, j, player)
	case 3: // 车
		moves = generateJuMoves(i, j, player)
	case 4: // 马
		moves = generateMaMoves(i, j, player)
	case 5: // 象
		moves = generateXiangMoves(i, j, player)
	case 6: // 士
		moves = generateShiMoves(i, j, player)
	case 7: // 将
		moves = generateJiangMoves(i, j, player)
	}

	if man < 0 {
		Map = tempMap
	}

	var movestr = ""

	for _, m := range moves {
		if man < 0 {
			m.toY = 9 - m.toY
		}
		movestr += strconv.Itoa(m.toX) + strconv.Itoa(m.toY) + ","
	}

	showLog("getmoves", movestr)

	return map[string]interface{}{
		"moves": movestr,
		"key":   man,
	}
}

func getMove(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	m := args[0].String()
	if m != "" {
		arr := strings.Split(m, ";")
		for i, v := range arr {
			a := strings.Split(v, ",")
			for ii, vv := range a {
				Map[i][ii], _ = strconv.Atoi(vv)
			}
		}
	}
	maxDepth = 3
	if len(args) == 2 {
		maxDepth = int(args[1].Float())
	}
	// bad = []int{9, 9, 9, 9}
	if len(args) == 3 {
		b := args[2].String()
		if len(b) == 4 {
			bad = []int{99, 99, 99, 99}
			for i, s := range b {
				bad[i], _ = strconv.Atoi(string(s))
			}
			showLog("go.bad", b)
		}
	}
	count = 0
	scores = []int{}
	aiMove = Move{}
	time0 := time.Now()
	rand.Seed(time.Now().UnixNano())
	alphaBeta(maxDepth, -99999, 99999, -1)
	time1 := time.Since(time0).Milliseconds()
	result := ""
	for _, v := range []int{aiMove.fromY, aiMove.fromX, aiMove.toY, aiMove.toX} {
		result += strconv.Itoa(v)
	}
	ss := ""
	for _, s := range scores {
		ss += strconv.Itoa(s) + ","
	}
	showLog("wasmkey", result, "score", aiMove.score, "depth", maxDepth, "pass", pass, "\ncount", count, "time", time1, ss)
	return map[string]interface{}{
		"key":   result,
		"score": aiMove.score,
		"count": count,
		"time":  time1,
		"depth": maxDepth,
	}
}

func main() {
	InitBoard()
	js.Global().Set("getmove", js.FuncOf(getMove))
	js.Global().Set("getmoves", js.FuncOf(getMoves))
	select {}
}
