package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data := getFile("test.txt")
	teams := buildTeams(data)
	fmt.Println(teams)
}

// Team Building Exercise
type team struct {
	teamNumber int
	members    []member
}
type member struct {
	member int
	start  byte
	end    byte
}

func buildTeams(data []string) []team {
	teams := make([]team, 0)
	for k, v := range data {
		elfMembers := make([]member, 0)
		fmt.Printf("k: %v \t  v: %v \n", k, string(v))
		fmt.Println(v[0])

		elfMembers = append(elfMembers, member{
			member: 1,
			start:  v[0],
			end:    v[2],
		},
			member{
				member: 2,
				start:  v[4],
				end:    v[6],
			})
		teams = append(teams, team{
			teamNumber: k,
			members:    elfMembers,
		})
	}
	return teams
}
func getFile(filename string) []string {
	data, _ := os.ReadFile(fmt.Sprintf("./%v", filename))
	splitData := strings.Split(string(data), "\n")
	return splitData
}
