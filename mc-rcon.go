package mc_rcon

import (
	"encoding/json"
	"fmt"
	"github.com/melq/mc-rcon/maze"
	"github.com/willroberts/minecraft-client"
	"log"
	"math"
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

func GiveItsumono(name string, inventory Inventory, client *minecraft.Client) {
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

func BuildMaze(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int, material string, client *minecraft.Client) {
	sortPositions(&x1, &y1, &z1, &x2, &y2, &z2)

	length := int(math.Abs(float64(z2 - z1)))
	width := int(math.Abs(float64(x2 - x1)))
	height := int(math.Abs(float64(y2 - y1)))

	m, err := maze.CreateMaze(length, width)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range m {
		for j, vv := range v {
			if vv != 0 {
				for k := 0; k < height; k++ {
					time.Sleep(3)
					if vv != 0 {
						_, err = client.SendCommand(fmt.Sprintf(
							"setblock %d %d %d %s", x1+j, y1+k, z1+i, material))
					} else {
						_, err = client.SendCommand(fmt.Sprintf(
							"setblock %d %d %d %s", x1+j, y1+k, z1+i, "minecraft:air"))
					}
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}

	maze.DumpMaze(m)
}
