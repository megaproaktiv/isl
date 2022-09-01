package main

import (
	"fmt"
	isl "identitystorelister"
	"log"
)

func main(){

	id, err := isl.GetInstanceId(isl.ClientSSO)
	if err != nil {
		log.Fatal("Error, stop")
	}

	users, err := isl.ListUsers(id, isl.Client)
	if err != nil {
		log.Fatal("Error, stop")
	}
	for _, user := range *users {
		fmt.Printf("\nUser: %v ", *user.UserName)
		groups, err := isl.ListGroupMembershipsForMember(user.UserId,id, isl.Client)
		if err != nil {
			log.Print("Cant get groups")
		}
		if len(*groups) > 0 {
			fmt.Printf("Groups: ")
		}
		for _, group := range *groups {
			name, err := isl.GroupName(group.GroupId, id, isl.Client)
			if err != nil {
				log.Print("Cant get groupname")
			}
			fmt.Printf("%v ", *name)
		}
		if len(*groups) > 0 {
			fmt.Printf("\n")
		}
	}

	groups, err := isl.ListGroups(id, isl.Client)
	if err != nil {
		log.Fatal("Error, stop")
	}
	for _, group := range *groups {
		fmt.Printf("User: %v \n", *group.DisplayName)
	}


}
