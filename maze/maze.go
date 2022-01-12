package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	X int
	Y int
}

func isCurrentWall(s []Cell, n Cell) bool {
	for _, v := range s {
		if n == v {
			return true
		}
	}
	return false
}

func DumpMaze(maze [][]int) {
	for _, v := range maze {
		for _, vv := range v {
			tmp := "□"
			if vv != 0 {
				tmp = "■"
			}
			fmt.Printf("%s ", tmp)
		}
		fmt.Println()
	}
}

func (c *Cell) isEmpty() bool {
	empty := Cell{}
	return *c == empty
}

func CreateMaze(height int, width int) ([][]int, error) {
	if height%2 == 0 { // サイズの奇数合わせ
		height--
	}
	if width%2 == 0 {
		width--
	}
	if height < 5 || width < 5 { // サイズのチェック
		return nil, fmt.Errorf("size is too small")
	}

	maze := make([][]int, height) // 二次元スライス初期化
	for i := 0; i < height; i++ {
		maze[i] = make([]int, width)
	}

	var startCells []Cell
	for i, v := range maze { // 周囲の壁化と起点取得
		for j := range v {
			if i == 0 || j == 0 || i == height-1 || j == width-1 {
				maze[i][j] = -1
			} else {
				if i%2 == 0 && j%2 == 0 {
					startCells = append(startCells, Cell{j, i})
				}
			}
		}
	}

	for len(startCells) != 0 { // 迷路生成(起点リストを回す)
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(len(startCells))
		s := startCells[r]

		if maze[s.Y][s.X] != 0 { // その起点が既に壁の場合
			var tmp []Cell
			for i := 0; i < len(startCells); i++ {
				if i != r {
					tmp = append(tmp, startCells[i])
				}
			}
			startCells = tmp
			continue
		}

		currentWall := []Cell{s}

		for { // 起点から壁伸ばし処理
			d := Cell{0, 0}
			for { // 進む方向決め
				if maze[s.Y-1][s.X] != 0 && isCurrentWall(currentWall, Cell{s.X, s.Y - 2}) &&
					maze[s.Y][s.X+1] != 0 && isCurrentWall(currentWall, Cell{s.X + 2, s.Y}) &&
					maze[s.Y+1][s.X] != 0 && isCurrentWall(currentWall, Cell{s.X, s.Y + 2}) &&
					maze[s.Y][s.X-1] != 0 && isCurrentWall(currentWall, Cell{s.X - 2, s.Y}) { // どこにも進めないなら
					if len(currentWall) > 3 {
						s = currentWall[len(currentWall)-2]
						currentWall = currentWall[:len(currentWall)-2]
					} else {
						currentWall = []Cell{}
					}
					break
				}

				switch rand.Intn(4) {
				case 0:
					{
						d = Cell{0, -1}
					}
				case 1:
					{
						d = Cell{1, 0}
					}
				case 2:
					{
						d = Cell{0, 1}
					}
				case 3:
					{
						d = Cell{-1, 0}
					}
				}
				if maze[s.Y+d.Y][s.X+d.X] == 0 && !isCurrentWall(currentWall, Cell{s.X + 2*d.X, s.Y + 2*d.Y}) { // 進める方向なら
					break
				}
			}
			if d.isEmpty() { // どこにも進めなければdはEmpty
				continue
			}

			currentWall = append(currentWall, Cell{s.X + d.X, s.Y + d.Y}) // 壁に当たっても当たらなくても1マスは進む
			if maze[s.Y+2*d.Y][s.X+2*d.X] != 0 {                          // 壁に当たったら
				break // 壁の拡張終了
			} else {
				s = Cell{s.X + 2*d.X, s.Y + 2*d.Y} // 2マス進めて次のループ
				currentWall = append(currentWall, s)
			}
		}
		for i, v := range currentWall { // 壁を確定
			maze[v.Y][v.X] = i + 1
		}
	}
	return maze, nil
}
