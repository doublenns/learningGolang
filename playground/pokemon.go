package main

import "fmt"

func main() {
    pokemon := map[string][]string {
        "Fire": []string{"Charmander", "Ponyta"},
        "Water": {},
        "Grass": []string{"Bulbasaur", "Oddish"},
        "Rock": []string{"Golem", "Onyx"},
    }

    fmt.Println("Fire Pokemon: ", pokemon["Fire"])
    fmt.Println("Water Pokemon: ", pokemon["Water"])

    fmt.Println("Adding water pokemon now...")
    pokemon["Water"] = append(pokemon["Water"], "Squirtle", "Goldeen", "Horsea")
    pokemon["Water"] = append(pokemon["Water"], "Staryu" )
    fmt.Println("Water Pokemon: ", pokemon["Water"])
}
