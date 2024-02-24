package main

import (
	"fmt"
	"os"

	"github.com/charlieparkes/helldivers-2-stratagems/internal/key"
	"github.com/charlieparkes/helldivers-2-stratagems/internal/stratagem"
)

func main() {
	s := stratagem.NewStratagems()

	// https://steamcommunity.com/sharedfiles/filedetails/?id=3160501535
	// general
	s.Add("reinforce", key.UP, key.DOWN, key.RIGHT, key.LEFT, key.UP)
	s.Add("sos", key.UP, key.DOWN, key.RIGHT, key.UP)
	s.Add("resupply", key.DOWN, key.DOWN, key.UP, key.RIGHT)

	// patriotic admin centre
	// s.Add("machine_gun")
	s.Add("anti_materiel_rifle", key.DOWN, key.LEFT, key.RIGHT, key.UP, key.DOWN)
	// s.Add("stalwart")
	// s.Add("expendable_anti_tank")
	// s.Add("recoilless_rifle")
	// s.Add("flamethrower")
	s.Add("autocannon", key.DOWN, key.LEFT, key.DOWN, key.UP, key.UP, key.RIGHT)
	s.Add("railgun", key.DOWN, key.RIGHT, key.DOWN, key.UP, key.LEFT, key.RIGHT)
	// s.Add("spear")

	// orbital cannons
	s.Add("orbital_gatling", key.RIGHT, key.DOWN, key.LEFT, key.UP, key.UP)
	s.Add("orbital_airburst", key.RIGHT, key.RIGHT, key.RIGHT)
	s.Add("orbital_120mm", key.RIGHT, key.RIGHT, key.DOWN, key.LEFT, key.RIGHT, key.DOWN)
	s.Add("orbital_380mm", key.RIGHT, key.DOWN, key.UP, key.UP, key.LEFT, key.DOWN, key.DOWN)
	s.Add("orbital_walking", key.RIGHT, key.DOWN, key.RIGHT, key.DOWN, key.RIGHT, key.DOWN)
	s.Add("orbital_laser", key.RIGHT, key.DOWN, key.UP, key.RIGHT, key.DOWN)
	s.Add("orbital_railcannon", key.RIGHT, key.UP, key.DOWN, key.DOWN, key.RIGHT)

	// hangar
	s.Add("eagle_strafing", key.UP, key.RIGHT, key.RIGHT)
	s.Add("eagle_airstrike", key.UP, key.RIGHT, key.DOWN, key.RIGHT)
	s.Add("eagle_cluster", key.UP, key.RIGHT, key.DOWN, key.DOWN, key.RIGHT)
	s.Add("eagle_napalm", key.UP, key.RIGHT, key.DOWN, key.UP)
	// s.Add("jump_pack")
	// s.Add("eagle_smoke")
	s.Add("eagle_110mm", key.UP, key.RIGHT, key.UP, key.LEFT)
	s.Add("eagle_500kg", key.UP, key.RIGHT, key.DOWN, key.DOWN, key.DOWN)

	// bridge
	// s.Add("orbital_precision")
	s.Add("orbital_gas", key.RIGHT, key.RIGHT, key.DOWN, key.RIGHT)
	// s.Add("orbital_ems")
	// s.Add("orbital_smoke")
	// s.Add("hmg")
	// s.Add("shield_gen")
	s.Add("tesla_tower", key.DOWN, key.UP, key.RIGHT, key.UP, key.LEFT, key.RIGHT)

	// engineering
	s.Add("minefield", key.DOWN, key.LEFT, key.UP, key.RIGHT)
	s.Add("supply_pack", key.DOWN, key.LEFT, key.DOWN, key.UP, key.UP, key.DOWN)
	s.Add("grenade_launcher", key.DOWN, key.LEFT, key.UP, key.LEFT, key.DOWN)
	// s.Add("laser_canon")
	// s.Add("incendiary_mines")
	s.Add("guard_dog_rover", key.DOWN, key.UP, key.LEFT, key.UP, key.RIGHT, key.RIGHT)
	// s.Add("balistic_shield")
	s.Add("arc_thrower", key.DOWN, key.RIGHT, key.DOWN, key.UP, key.LEFT, key.LEFT)
	s.Add("shield_gen_pack", key.DOWN, key.UP, key.LEFT, key.RIGHT, key.LEFT, key.RIGHT)

	// robotics
	// s.Add("machine_gun_sentry")
	s.Add("gatling_sentry", key.DOWN, key.UP, key.RIGHT, key.LEFT)
	s.Add("mortar_sentry", key.DOWN, key.UP, key.RIGHT, key.RIGHT, key.DOWN)
	// s.Add("guard_dog")
	s.Add("autocannon_sentry", key.DOWN, key.UP, key.RIGHT, key.UP, key.LEFT, key.UP)
	// s.Add("rocket_sentry")
	// s.Add("ems_mortar_sentry")

	if len(os.Args) != 2 {
		fmt.Println("Example: ./stratagems sos")
		fmt.Println("Stratagems:")
		s.PrintAll()
		return
	}

	s.Call(stratagem.Name(os.Args[1]))
}
