package main

import (
	"github.com/willroberts/minecraft-client"
	"log"
	//"regexp"
	//"strconv"
	//"strings"
)

/*func GetPlayerPos(name string, client *minecraft.Client) (float64, float64, float64) {
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
}*/

/*func MakeSchematic(blocks []string, x1 int, y1 int, z1 int, x2 int, y2 int, z2 int, client *minecraft.Client) [][][]string {
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
				for _, v := range blocks {
					time.Sleep(3)
					resp, err := client.SendCommand(fmt.Sprintf(
						"execute if block %d %d %d %s", x, y, z, v))
					if err != nil {
						log.Fatal(err)
					}
					tmp := ""
					if resp.Body == "Test passed\n" {
						tmp = v
					}
					schematic[x-x1][y-y1] = append(schematic[x-x1][y-y1], tmp)
				}
			}
		}
	}
	return schematic
}*/

func main() {
	client, err := minecraft.NewClient("localhost:25575")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err = client.Authenticate("password"); err != nil {
		log.Fatal(err)
	}

	//x, y, z := GetPlayerPos("tu_tutu_", client)
	//fmt.Printf("%f, %f, %f\n", x, y, z)

	//fmt.Println(MakeSchematic([]string{"minecraft:stone"}, -6, -60, 3, -3, -57, 6, client))
}
