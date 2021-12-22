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

	fmt.Println(mc_rcon.MakeSchematic([]string{"minecraft:stone"}, -6, -60, 3, -3, -57, 6, client))
}

func TestBuildWithSchematic(t *testing.T) {
	client := mc_rcon.GetClient("localhost:25575", "password")
	defer client.Close()

	mc_rcon.BuildWithSchematic(
		mc_rcon.MakeSchematic([]string{"minecraft:stone"}, -6, -60, 3, -3, -57, 6, client),
		-15, -60, -7, client)
}
