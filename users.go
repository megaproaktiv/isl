package isl

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
)	

var Client *identitystore.Client

func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	Client = identitystore.NewFromConfig(cfg)
}

func ListUsers(identitystoreId *string, client *identitystore.Client)( *[]types.User, error){
	params := identitystore.ListUsersInput{
		IdentityStoreId: identitystoreId,
		MaxResults: 100,
	}

	resp, err := client.ListUsers(context.TODO(), &params)

	if err != nil {
		log.Printf("Error listusers %p", err)
		return nil, err
	}
	
	return &resp.Users, err
}

func ListGroups(identitystoreId *string, client *identitystore.Client)( *[]types.Group, error){
	params := identitystore.ListGroupsInput{
		IdentityStoreId: identitystoreId,
		MaxResults: 100,
	}

	resp, err := client.ListGroups(context.TODO(), &params)

	if err != nil {
		log.Printf("Error listgroups %p", err)
		return nil, err
	}
	
	return &resp.Groups, err
}

func ListGroupMembershipsForMember(userID *string, identitystoreId *string, client *identitystore.Client)( *[]types.GroupMembership, error){
	params := identitystore.ListGroupMembershipsForMemberInput{
		IdentityStoreId: identitystoreId,
		MaxResults: 100,
		MemberId: &types.MemberIdMemberUserId{
			Value: *userID,
		},
	}

	resp, err := client.ListGroupMembershipsForMember(context.TODO(), &params)

	if err != nil {
		log.Printf("Error listgroups %p", err)
		return nil, err
	}
	
	return &resp.GroupMemberships, err
}

func GroupName(groupID *string, identitystoreId *string, client *identitystore.Client)( *string, error){
	params := identitystore.DescribeGroupInput{
		IdentityStoreId: identitystoreId,
		GroupId: groupID,
	}
	
	resp, err := client.DescribeGroup(context.TODO(), &params)
	
	if err != nil {
		log.Printf("Error DescribeGroup %p", err)
		return nil, err
	}
	
	return resp.DisplayName, err

}

// Very basic id validation
func Validate(id string) bool{
	if !strings.HasPrefix(id, "d-"){
		return false
	}
	if ! (len(id) >2 && len(id) < 36){
		return false
	}
	return true
}