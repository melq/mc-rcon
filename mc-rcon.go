package mc_rcon

import (
	"encoding/json"
	"fmt"
	"github.com/willroberts/minecraft-client"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetClient(hostport string, password string) *minecraft.Client {
	client, err := minecraft.NewClient(hostport)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Authenticate(password); err != nil {
		log.Fatal(err)
	}
	return client
}

func sortPositions(x1 *int, y1 *int, z1 *int, x2 *int, y2 *int, z2 *int) {
	if *x1 > *x2 {
		tmp := *x1
		*x1 = *x2
		*x2 = tmp
	}
	if *y1 > *y2 {
		tmp := *y1
		*y1 = *y2
		*y2 = tmp
	}
	if *z1 > *z2 {
		tmp := *z1
		*z1 = *z2
		*z2 = tmp
	}
}

func GetPlayerPos(name string, client *minecraft.Client) (float64, float64, float64) {
	resp, err := client.SendCommand(fmt.Sprintf("execute at %s run tp %s ~ ~ ~", name, name))
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`[^0-9-,.]`)
	str := reg.ReplaceAllString(resp.Body, "")
	s := strings.Split(str, ",")

	x, err := strconv.ParseFloat(s[0], 64)
	y, err := strconv.ParseFloat(s[1], 64)
	z, err := strconv.ParseFloat(s[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	return x, y, z
}

func MakeSchematic(blocks []string, x1 int, y1 int, z1 int, x2 int, y2 int, z2 int, client *minecraft.Client) [][][]string {
	sortPositions(&x1, &y1, &z1, &x2, &y2, &z2)

	var schematic [][][]string

	for x := x1; x <= x2; x++ {
		schematic = append(schematic, [][]string{})
		for y := y1; y <= y2; y++ {
			schematic[x-x1] = append(schematic[x-x1], []string{})
			for z := z1; z <= z2; z++ {
				tmp := ""
				for _, v := range blocks {
					time.Sleep(3)
					resp, err := client.SendCommand(fmt.Sprintf(
						"execute if block %d %d %d %s", x, y, z, v))
					if err != nil {
						log.Fatal(err)
					}
					if resp.Body == "Test passed\n" {
						tmp = v
					}
				}
				schematic[x-x1][y-y1] = append(schematic[x-x1][y-y1], tmp)
			}
		}
	}
	return schematic
}

func BuildWithSchematic(schematic [][][]string, x int, y int, z int, client *minecraft.Client) {
	cnt := 1
	for i, v := range schematic {
		for j, vv := range v {
			for k, vvv := range vv {
				time.Sleep(3)
				_, err := client.SendCommand(fmt.Sprintf(
					"setblock %d %d %d %s", x+i, y+j, z+k, vvv))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%d: %d %d %d %s\n", cnt, x+i, y+j, z+k, schematic[i][j][k])
				cnt++
			}
		}
	}
}

func GiveItsumono(name string, client *minecraft.Client) {
	inventory := Inventory{
		Items: []Item{
			{Id: "minecraft:firework_rocket", Count: "64", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment(nil)}},
			{Id: "minecraft:netherite_boots", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "depth_strider", Lvl: "3"}, {Id: "feather_falling", Lvl: "4"}}}},
			{Id: "minecraft:elytra", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}}}},
			{Id: "minecraft:netherite_helmet", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "respiration", Lvl: "3"}, {Id: "aqua_affinity", Lvl: "1"}}}},
			{Id: "minecraft:netherite_leggings", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:protection", Lvl: "4s"}}}},
			{Id: "minecraft:netherite_shovel", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}}}},
			{Id: "minecraft:arrow", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment(nil)}},
			{Id: "minecraft:bow", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "infinity", Lvl: "1"}, {Id: "power", Lvl: "5"}, {Id: "punch", Lvl: "2"}, {Id: "flame", Lvl: "1"}}}},
			{Id: "minecraft:netherite_pickaxe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: "silk_touch"}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:silk_touch", Lvl: "1s"}}}},
			{Id: "minecraft:netherite_pickaxe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: "fortune"}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:fortune", Lvl: "3s"}}}},
			{Id: "minecraft:netherite_axe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}, {Id: "sharpness", Lvl: "5"}}}},
			{Id: "minecraft:netherite_sword", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "sharpness", Lvl: "5"}, {Id: "sweeping", Lvl: "3"}}}}}}

	for _, item := range inventory.Items {
		time.Sleep(3)
		tag := ""
		if item.Tag.Display.Name.Text != "" {
			tag += fmt.Sprintf("display:{Name:\"\\\"%s\\\"\"},", item.Tag.Display.Name.Text)
		}
		if len(item.Tag.Enchantments) != 0 {
			tag += fmt.Sprintf("Enchantments:[")
			for _, enchantment := range item.Tag.Enchantments {
				tag += fmt.Sprintf("{id:\"%s\",lvl:%s},", enchantment.Id, enchantment.Lvl)
			}
			tag += "]"
		}
		cmd := fmt.Sprintf("give %s %s{%s}", name, item.Id, tag)
		if item.Count != "" {
			cmd += " " + item.Count
		}
		fmt.Println(cmd)
		_, err := client.SendCommand(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Name struct {
	Text string `json:"text"`
}

type Display struct {
	Name Name `json:"Name"`
}

type Enchantment struct {
	Id  string `json:"id"`
	Lvl string `json:"lvl"`
}

type Tag struct {
	Display      Display       `json:"display"`
	Enchantments []Enchantment `json:"Enchantments"`
}

type Item struct {
	Id    string `json:"id"`
	Count string `json:"Count"`
	Tag   Tag    `json:"tag"`
}

type Inventory struct {
	Items []Item `json:"items"`
}

func GetInventory(name string, client *minecraft.Client) Inventory {
	resp, err := client.SendCommand(fmt.Sprintf("data get entity %s Inventory", name))
	if err != nil {
		log.Fatal(err)
	}
	str := resp.Body[strings.Index(resp.Body, ": [")+2:]

	reg := regexp.MustCompile(`([\w\d]+): `)
	str = "{items: " + str + "}"
	str = reg.ReplaceAllString(str, "\"$1\": ")
	reg = regexp.MustCompile(`: ([\w\d]+)`)
	str = reg.ReplaceAllString(str, ": \"$1\"")
	reg = regexp.MustCompile(`'({[\w\d\s:"]+})'`)
	str = reg.ReplaceAllString(str, "$1")

	inventory := Inventory{make([]Item, 0)}
	if err := json.Unmarshal([]byte(str), &inventory); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", inventory)
	return inventory
}

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

func dumpMazeWithNum(maze [][]int) {
	for _, v := range maze {
		for _, vv := range v {
			fmt.Printf("%2d ", vv)
		}
		fmt.Println()
	}
}

func dumpMaze(maze [][]int) {
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

func BuildMaze(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int /*client *minecraft.Client*/) error {
	sortPositions(&x1, &y1, &z1, &x2, &y2, &z2)

	height := int(math.Abs(float64(z2 - z1)))
	width := int(math.Abs(float64(x2 - x1)))

	if height%2 == 0 { // サイズの奇数合わせ
		height--
	}
	if width%2 == 0 {
		width--
	}
	if height < 5 || width < 5 { // サイズのチェック
		return fmt.Errorf("size is too small")
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

	for len(startCells) != 0 { // 迷路生成
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
				if maze[s.Y+d.Y][s.X+d.X] == 0 && !isCurrentWall(currentWall, Cell{s.X + 2*d.X, s.Y + 2*d.Y}) {
					break
				}
			}
			if d.isEmpty() {
				continue
			}

			currentWall = append(currentWall, Cell{s.X + d.X, s.Y + d.Y})
			if maze[s.Y+2*d.Y][s.X+2*d.X] != 0 { //壁の拡張終了
				break
			} else {
				s = Cell{s.X + 2*d.X, s.Y + 2*d.Y}
				currentWall = append(currentWall, s)
			}
		}
		for i, v := range currentWall {
			maze[v.Y][v.X] = i + 1
		}
	}

	dumpMaze(maze)
	return nil
}
