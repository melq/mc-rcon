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

	mc_rcon.GiveItsumono("tu_tutu_", client)
}

func TestGetInventory(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	mc_rcon.GetInventory("tu_tutu_", client)
}

func TestBuildMaze(t *testing.T) {
	mc_rcon.BuildMaze(0, 0, 0, 10, 10, 10)
}
