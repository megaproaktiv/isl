package main

import (
	"fmt"
	isl "identitystorelister"
	"log"
	"github.com/xuri/excelize/v2"
)

func main() {

	id, err := isl.GetInstanceId(isl.ClientSSO)
	if err != nil {
		log.Fatal("Error, stop")
	}

	users, err := isl.ListUsers(id, isl.Client)
	if err != nil {
		log.Fatal("Error, stop")
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(sheet, "User")
	sheet = "User"
	f.SetCellValue(sheet, "A1", "User")
	f.SetCellValue(sheet, "B1", "Groups")
	err = f.SetColWidth(sheet, "A", "A",32)
	if err != nil {
		log.Print("Cant set width")
	}
	err = f.SetColWidth(sheet, "B", "B",128)
	if err != nil {
		log.Print("Cant set width")
	}
	styleWrap, err := f.NewStyle(&excelize.Style{
        Alignment: &excelize.Alignment{
            WrapText: true,
			ShrinkToFit: true,
        },
    })
    if err != nil {
        panic(err)
    }
	// User Loop
	for i, user := range *users {
		fmt.Printf(".")
		value := fmt.Sprintf("%v / %v %v ", *user.UserName, *user.Name.GivenName, *user.Name.FamilyName)
		row := i + 2
		coordinates := fmt.Sprintf("A%d", row)
		f.SetCellValue(sheet, coordinates, value)
		// Groups for User Loop
		groups, err := isl.ListGroupMembershipsForMember(user.UserId, id, isl.Client)
		if err != nil {
			log.Print("Cant get groups")
		}
		allGroups := ""
		for _, group := range *groups {
			name, err := isl.GroupName(group.GroupId, id, isl.Client)
			if err != nil {
				log.Print("Cant get groupname")
			}
			value = fmt.Sprintf("%v ", *name)
			allGroups += value+"\r\n"
		}
		if len(allGroups) > 0{
			coordinates = fmt.Sprintf("B%d", row)
			f.SetCellValue(sheet, coordinates, allGroups)
			if err := f.SetCellStyle(sheet, coordinates, coordinates, styleWrap); err != nil {
				fmt.Println(err)
			}
		}

	}
	filename := "sso-users.xlsx"
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nExcelfile %v saved\n", filename)

	// groups, err := isl.ListGroups(id, isl.Client)
	// if err != nil {
	// 	log.Fatal("Error, stop")
	// }

	// for _, group := range *groups {
	// 	fmt.Printf("Group: %v \n", *group.DisplayName)
	// }

}
