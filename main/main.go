package main

import (
	"fmt"
	"isl"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func main(){

	if len(os.Args) < 2 {
		log.Fatalf("Need identitystoreID as input\n")
	}

	idInput := os.Args[1]
	if ! isl.Validate(idInput){
		log.Fatalf("wrong identityStoreId format , d- is a fixed prefix, Length Constraints: Minimum length of 1. Maximum length of 36.")
	}
	id := aws.String(idInput)

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
