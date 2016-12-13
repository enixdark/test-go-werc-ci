// package main

// import(
//   // "fmt"
//   // "github.com/enixdark/test-go-werc-ci/npcs"
//   "fmt"
//   "os"

//   "github.com/cloudfoundry-community/go-cfenv"
//   service "github.com/cloudnativego/gogo-service/service"
// )

// func main() {
//   mob := npc.NonPlayerCharacter{Name: "Alfred"}
//   fmt.Println(mob)

//   hortense := npc.NonPlayerCharacter{Name: "hortense", 
//     Loc: npc.Location{X: 10.0, Y: 10.0, Z: 10.0}}
//   fmt.Println(hortense)

//   fmt.Printf("Alfred is %f units from Hortense.\n",
//   mob.DistanceTo(hortense))
// }
// 
// func main() {
//   port := os.Getenv("PORT")
//   if len(port) == 0 {
//     port = "3000"
//   }

//   appEnv, err := cfenv.Current()
//   if err != nil {
//     fmt.Println("CF Environment not detected.")
//   }

//   server := service.NewServer(appEnv)
//   server.Run(":" + port)
// }