package level5

import (
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	p.calculateRating()
	return p
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	} else {
		p.Rating = (float64(p.Goals) + (float64(p.Assists) / 2)) / float64(p.Misses)
	}
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i int, j int) bool {
		if players[i].Goals != players[j].Goals {
			return players[i].Goals > players[j].Goals
		}
		return players[i].Name < players[j].Name
	})

	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i int, j int) bool {
		if players[i].Rating != players[j].Rating {
			return players[i].Rating > players[j].Rating
		}
		return players[i].Name < players[j].Name
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i int, j int) bool {
		var i_gm float64
		var j_gm float64
		if players[i].Misses == 0 {
			i_gm = float64(players[i].Goals) / 1
		} else {
			i_gm = float64(players[i].Goals / players[i].Misses)
		}
		if players[j].Misses == 0 {
			j_gm = float64(players[j].Goals) / 1
		} else {
			j_gm = float64(players[j].Goals / players[j].Misses)
		}

		if i_gm != j_gm {
			return i_gm > j_gm
		}
		return players[i].Name < players[j].Name
	})
	return players
}

func main() {

}
