package pokemon

import "fmt"

type Pokemon struct {
	ID     int
	Name   string
	Type   string
	Weight int
}

func (p Pokemon) String() string {
	return fmt.Sprintf(
		"Name:\t\t%q\n"+
			"Type:\t\t%q\n"+
			"Published:\t%v\n",
		p.Name, p.Type, p.Weight)
}

var Wild = []Pokemon{
	Pokemon{
		ID:     1,
		Name:   "Bulbasaur",
		Type:   "grass",
		Weight: 69,
	},
	Pokemon{
		ID:     2,
		Name:   "Charmander",
		Type:   "fire",
		Weight: 85,
	},
	Pokemon{
		ID:     3,
		Name:   "Squirtle",
		Type:   "water",
		Weight: 90,
	},
	Pokemon{
		ID:     4,
		Name:   "Caterpie",
		Type:   "bug",
		Weight: 29,
	},
	Pokemon{
		ID:     5,
		Name:   "Pidgey",
		Type:   "flying",
		Weight: 18,
	},
	Pokemon{
		ID:     6,
		Name:   "Rattata",
		Type:   "normal",
		Weight: 35,
	},
	Pokemon{
		ID:     7,
		Name:   "Articuno",
		Type:   "ice",
		Weight: 554,
	},
	Pokemon{
		ID:     8,
		Name:   "Zapdos",
		Type:   "electric",
		Weight: 526,
	},
	Pokemon{
		ID:     9,
		Name:   "Moltres",
		Type:   "fire",
		Weight: 600,
	},
	Pokemon{
		ID:     10,
		Name:   "Mewtwo",
		Type:   "psychic",
		Weight: 1220,
	},
}
