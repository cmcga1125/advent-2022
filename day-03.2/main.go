package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type group struct {
	GroupName    rune
	GroupValue   int
	GroupMembers []string
}

func main() {
	data, _ := os.ReadFile("./data.txt")
	members := make([]string, 0)
	groupList := make([]group, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	lower := "abcdefghijklmnopqrstuvwxyz"
	mixed := lower + strings.ToUpper(lower)
	firstMember := 0
	var groupName rune
	groupMembers := make([]string, 0)
	groups := make([][]string, 0)

	for scanner.Scan() {
		members = append(members, scanner.Text())
	}
	fmt.Println(members)
	fmt.Println(len(members))
	for i := 0; i < (len(members) / 3); i++ {
		groupMembers := append(groupMembers, members[firstMember], members[firstMember+1], members[firstMember+2])
		fmt.Println(groupMembers)
		firstMember+=3
		groups = append(groups, groupMembers)
	}

	// fmt.Println(groups)

	for _, m := range groups {
		for _, m0 := range m[0] {
			for _, m1 := range m[1] {
				for _, m2 := range m[2] {
					if m0 == m2 && m1 == m2 {
						groupName = m0
						break
					}
				}
			}
		}
		// fmt.Printf("groupName: %v\n", string(groupName))
		groupValue := strings.Index(mixed, string(groupName)) + 1
		// fmt.Printf("gn: %v, \t gv: %v\n", string(groupName), string(groupValue))
		groupList = append(groupList, group{
			GroupName:    groupName,
			GroupValue:   groupValue,
			GroupMembers: groupMembers,
		})
	}
	var calval []int
	for _, result := range groupList {
		fmt.Println(result.GroupValue)
		calval = append(calval, result.GroupValue)
	}
	totalValue := sum(calval)
	fmt.Printf("total value: %v\n", totalValue)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
