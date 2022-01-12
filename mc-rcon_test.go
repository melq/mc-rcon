package mc_rcon_test

import (
	"fmt"
	"github.com/melq/mc-rcon"
	"testing"
)

func TestGetPlayerPos(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	x, y, z := mc_rcon.GetPlayerPos("tu_tutu_", client)
	fmt.Printf("%f, %f, %f\n", x, y, z)
}

func TestMakeSchematic(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	schematic := []string{
		"minecraft:glass",
		"minecraft:birch_door[facing=west,half=lower]",
		"minecraft:birch_door[facing=west,half=upper]",
		"minecraft:cobblestone",
		"minecraft:oak_log",
		"minecraft:spruce_stairs[facing=north]",
		"minecraft:spruce_stairs[facing=south]",
		"minecraft:spruce_stairs[facing=west]",
		"minecraft:spruce_stairs[facing=east]",
	}
	fmt.Println(mc_rcon.MakeSchematic(
		schematic, 1, -60, -21, -4, -57, -15, client))
}

func TestBuildWithSchematic(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	materials := []string{
		"minecraft:glass",
		"minecraft:birch_planks",
		"minecraft:birch_door[facing=west,half=lower]",
		"minecraft:birch_door[facing=west,half=upper]",
		"minecraft:cobblestone",
		"minecraft:oak_log",
		"minecraft:spruce_stairs[facing=north]",
		"minecraft:spruce_stairs[facing=south]",
		"minecraft:spruce_stairs[facing=west]",
		"minecraft:spruce_stairs[facing=east]",
	}
	schematic := mc_rcon.MakeSchematic(
		materials, -4, -60, -21, 1, -57, -15, client)

	mc_rcon.BuildWithSchematic(
		schematic, -4, -60, -11, client)
}

func TestGiveItsumono(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	inventory := mc_rcon.Inventory{
		Items: []mc_rcon.Item{
			{Id: "minecraft:firework_rocket", Count: "64", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment(nil)}},
			{Id: "minecraft:netherite_boots", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "depth_strider", Lvl: "3"}, {Id: "feather_falling", Lvl: "4"}}}},
			{Id: "minecraft:elytra", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}}}},
			{Id: "minecraft:netherite_helmet", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "protection", Lvl: "4"}, {Id: "respiration", Lvl: "3"}, {Id: "aqua_affinity", Lvl: "1"}}}},
			{Id: "minecraft:netherite_leggings", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:protection", Lvl: "4s"}}}},
			{Id: "minecraft:netherite_shovel", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}}}},
			{Id: "minecraft:arrow", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment(nil)}},
			{Id: "minecraft:bow", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "infinity", Lvl: "1"}, {Id: "power", Lvl: "5"}, {Id: "punch", Lvl: "2"}, {Id: "flame", Lvl: "1"}}}},
			{Id: "minecraft:netherite_pickaxe", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: "silk_touch"}}, Enchantments: []mc_rcon.Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:silk_touch", Lvl: "1s"}}}},
			{Id: "minecraft:netherite_pickaxe", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: "fortune"}}, Enchantments: []mc_rcon.Enchantment{{Id: "minecraft:unbreaking", Lvl: "3s"}, {Id: "minecraft:mending", Lvl: "1s"}, {Id: "minecraft:efficiency", Lvl: "5s"}, {Id: "minecraft:fortune", Lvl: "3s"}}}},
			{Id: "minecraft:netherite_axe", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "efficiency", Lvl: "5"}, {Id: "sharpness", Lvl: "5"}}}},
			{Id: "minecraft:netherite_sword", Count: "", Tag: mc_rcon.Tag{Display: mc_rcon.Display{Name: mc_rcon.Name{Text: ""}}, Enchantments: []mc_rcon.Enchantment{{Id: "unbreaking", Lvl: "3"}, {Id: "mending", Lvl: "1"}, {Id: "sharpness", Lvl: "5"}, {Id: "sweeping", Lvl: "3"}}}}}}

	mc_rcon.GiveItsumono("tu_tutu_", inventory, client)
}

func TestGetInventory(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	mc_rcon.GetInventory("tu_tutu_", client)
}

func TestBuildMaze(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()
	mc_rcon.BuildMaze(23, -60, -39, 54, -57, -8, "minecraft:sea_lantern", client)
}
