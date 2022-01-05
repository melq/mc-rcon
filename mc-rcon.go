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

//func GiveItsumono(name string, client *minecraft.Client) {
//	items := []string{
//		"/give tu_tutu_ minecraft:netherite_helmet{Enchantments:[{id:unbreaking,lvl:3},{id:mending,lvl:1},{id:protection,lvl:4},{id:respiration,lvl:3},{id:aqua_affinity,lvl:1}]}",
//	}
//	for _, item := range items {
//		_, err := client.SendCommand(fmt.Sprintf("give %s %s", name, item))
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//}

type Enchantment struct {
	Id  string `json:"id"`
	Lvl string `json:"lvl"`
}

type Tag struct {
	Damage       string `json:"Damage"`
	Enchantments []Enchantment
}

type Item struct {
	Slot  string `json:"Slot"`
	Id    string `json:"id"`
	Count string `json:"Count"`
	Tag   Tag    `json:"tag"`
}

type Inventory struct {
	Items []Item `json:"items"`
}

func GetInventory(name string, client *minecraft.Client) /*[]string*/ {
	//resp, err := client.SendCommand(fmt.Sprintf("data get entity %s Inventory", name))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//str := resp.Body[strings.Index(resp.Body, ": [")+2:]
	str := "[{Slot: 103b, id: \"minecraft:netherite_helmet\", Count: 1b, tag: {Damage: 0, Enchantments: [{lvl: 3, id: \"unbreaking\"}, {lvl: 1, id: \"mending\"}, {lvl: 4, id: \"protection\"}, {lvl: 3, id: \"respiration\"}, {lvl: 1, id: \"aqua_affinity\"}]}}]"
	inventory := Inventory{make([]Item, 0)}

	reg := regexp.MustCompile(`([\w\d]+): `)
	str = "{items: " + str + "}"
	str = reg.ReplaceAllString(str, "\"$1\": ")
	reg = regexp.MustCompile(`: ([\w\d]+),`)
	str = reg.ReplaceAllString(str, ": \"$1\", ")
	fmt.Println(str)
	if err := json.Unmarshal([]byte(str), &inventory); err != nil {
		log.Fatal(err)
	}
	fmt.Println(inventory)
}
