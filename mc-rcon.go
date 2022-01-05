package mc_rcon

import (
	"encoding/json"
	"fmt"
	"github.com/willroberts/minecraft-client"
	"log"
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
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}
	if z1 > z2 {
		tmp := z1
		z1 = z2
		z2 = tmp
	}

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
	inventory := GetInventory(name, client)
	//inventory := Inventory{
	//	Items: []Item{
	//		{Id: "minecraft:firework_rocket", Count: "64", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment(nil)}},
	//		{Id: "minecraft:netherite_boots", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "depth_strider", Lvl: "3"}, {Id: "feather_falling", Lvl: "4"}}}},
	//		{Id: "minecraft:elytra", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}}}},
	//		{Id: "minecraft:netherite_helmet", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "respiration", Lvl: "3"}, {Id: "aqua_affinity", Lvl: "1"}}}},
	//		{Id: "minecraft:netherite_leggings", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:protection", Lvl: "4s"}}}},
	//		{Id: "minecraft:netherite_shovel", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}}}},
	//		{Id: "minecraft:arrow", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment(nil)}},
	//		{Id: "minecraft:bow", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "infinity", Lvl: "1"}, {Id: "power", Lvl: "5"}, {Id: "punch", Lvl: "2"}, {Id: "flame", Lvl: "1"}}}},
	//		{Id: "minecraft:netherite_pickaxe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: "silk_touch"}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:silk_touch", Lvl: "1s"}}}},
	//		{Id: "minecraft:netherite_pickaxe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: "fortune"}}, Enchantments: []Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:fortune", Lvl: "3s"}}}},
	//		{Id: "minecraft:netherite_axe", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}, {Id: "sharpness", Lvl: "5"}}}},
	//		{Id: "minecraft:netherite_sword", Count: "", Tag: Tag{Display: Display{Name: Name{Text: ""}}, Enchantments: []Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "sharpness", Lvl: "5"}, {Id: "sweeping", Lvl: "3"}}}}}}

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
