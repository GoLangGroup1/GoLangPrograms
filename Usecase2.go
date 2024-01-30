package main

import "fmt"

type NspEvent struct {
	date string
	name string
	primaryContact string
	organizingTeam OrganizingTeam
}

type OrganizingTeam struct {
	team [] string
	primaryContact string
}

func main() {
	var option int;
	var nspEventSlice []NspEvent

	for{

	fmt.Println("1. Add event details")
	fmt.Println("2. Display event details")
	fmt.Println("3. Exit")

	fmt.Println("Enter option:")
	fmt.Scanln(&option)

	switch option {
	case 1:
		nspEventSlice = append(nspEventSlice, addEventDetails()...)
	case 2:
		displayEventDetails(nspEventSlice)
	case 3:
		return
	default:
		fmt.Println("No such option")
	}
}
}

func displayEventDetails(nspEventSlice []NspEvent){

	fmt.Println("**** Displaying event details *** ")

	for _,nspEvent := range nspEventSlice {
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("Event Name: %s \n", nspEvent.name)
		fmt.Printf("Event Date: %s \n", nspEvent.date)
		fmt.Printf("Event Primary Contact: %s \n", nspEvent.primaryContact)
		fmt.Printf("Organizing team Primary Contact: %s \n", nspEvent.organizingTeam.primaryContact)
		fmt.Printf("Organizing team members: %v \n", nspEvent.organizingTeam.team)
	}
	fmt.Println("---------------------------------------------------------")
	fmt.Println()
}

func addEventDetails()(nspEventSlice []NspEvent){
	var teamMemberSlice []string;
	var nspEvent NspEvent;
	var option int;
	var teamMember string;

	fmt.Println("Enter event name:")
	fmt.Scanln(&nspEvent.name)

	fmt.Println("Enter event date:")
	fmt.Scanln(&nspEvent.date)

	fmt.Println("Enter primary contact (NSP Event):")
	fmt.Scanln(&nspEvent.primaryContact)

	fmt.Println("Enter team members")

loop:
	for{
		fmt.Println("1: Enter details of team member")
		fmt.Println("2. Exit")
		fmt.Scanln(&option)
	
		if option == 2 {
			break loop
		}

		fmt.Println("Enter team member:")
		fmt.Scanln(&teamMember)
		teamMemberSlice = append(teamMemberSlice, teamMember)
	}

	nspEvent.organizingTeam.team = teamMemberSlice

loopOne:
	for{

	fmt.Println("Enter primary contact (organizing event):")
	fmt.Scanln(&nspEvent.organizingTeam.primaryContact)

	var found bool = false;

		for _,member := range teamMemberSlice {
			if (member == nspEvent.organizingTeam.primaryContact) {
				found = true
			}
		}
		if (found){
			break loopOne
		}

		fmt.Println("The primary contact of organizing team is not part of team. Please enter correct contact details")
	}


	nspEventSlice = append(nspEventSlice, nspEvent)
	return nspEventSlice
}