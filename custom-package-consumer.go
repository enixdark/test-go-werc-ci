package main

import(
  "fmt"
  "./npcs"
)

func main() {
  mob := npc.NonPlayerCharacter{Name: "Alfred"}
  fmt.Println(mob)

  hortense := npc.NonPlayerCharacter{Name: "hortense", 
    Loc: npc.Location{X: 10.0, Y: 10.0, Z: 10.0}}
  fmt.Println(hortense)

  fmt.Printf("Alfred is %f units from Hortense.\n",
  mob.DistanceTo(hortense))
}