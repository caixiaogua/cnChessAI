# cnChessAI
go语言写的中国象棋AI接口


### 使用tinygo编译为wasm文件
```
tinygo build -target=wasm cnchess.go
```
### 棋盘初始化
```
var Map [][]int
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
```
#### wasm执行后，会在window对象上挂载两个函数：getmoves(mapstr, x, y)和getmove(mapstr, depth)
```
let mapstr='-3,-4,-5,-6,-7,-6,-5,-4,-3;0,0,0,0,0,0,0,0,0;0,-2,0,0,0,0,0,-2,0;-1,0,-1,0,-1,0,-1,0,-1;0,0,0,0,0,0,0,0,0;0,0,0,0,1,0,0,0,0;1,0,1,0,0,0,1,0,1;0,2,0,0,0,0,0,2,0;0,0,0,0,0,0,0,0,0;3,4,5,6,7,6,5,4,3';
//执行函数返回一个对象，其中属性key为起始和目标坐标（y0,x0,y,x）
let key=getmove(mapstr, 4)	//第二个参数为计算深度

//getmoves函数返回某个棋子的可走位置
let moves=getmoves(mapstr, x, y)	//x和y为要查询棋子的坐标
```
#### 测试demo：http://159.75.26.117:9471/game/cnchess/index.html
