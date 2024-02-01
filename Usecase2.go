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
	fmt.Println("2. Delete event details")
	fmt.Println("3. Display event details")
	fmt.Println("4. Exit")

	fmt.Println("Enter option:")
	fmt.Scanln(&option)

	switch option {
	case 1:
		nspEventSlice = append(nspEventSlice, addEventDetails(nspEventSlice)...)
	case 2:
		nspEventSlice = deleteEventDetails(nspEventSlice)
	case 3:
		displayEventDetails(nspEventSlice)
	case 4:
		return
	default:
		fmt.Println("No such option")
	}
}
}

func displayEventDetails(nspEventSlice []NspEvent){
	
	fmt.Println("---------------------------------------------------------")
	
	if(len(nspEventSlice) > 0){
		fmt.Println("**** Displaying event details *** ")
	}else{
		fmt.Println("No events YET !!!")
	}

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

func addEventDetails(nspEventSliceIn []NspEvent)(nspEventSlice []NspEvent){
	var teamMemberSlice []string;
	var nspEvent NspEvent;
	var option int;
	var teamMember string;

	fmt.Println("Enter event name:")
	fmt.Scanln(&nspEvent.name)

	for _,value := range nspEventSliceIn {
		if(nspEvent.name == value.name){
			fmt.Println("Event already present... returning to MAIN menu")
			return nspEventSliceIn
		}
	}

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

func deleteEventDetails(nspEventSliceIn []NspEvent)(nspEventSliceOut []NspEvent){

	var deleteEvent string
	var deleteEventIndex int
	
	fmt.Println("Enter event name to be delted:")
	fmt.Scanln(&deleteEvent)

	for index,nspEvent := range nspEventSliceIn {
		if(deleteEvent == nspEvent.name) {
			deleteEventIndex = index
		}
	}

	nspEventSliceOut = append(nspEventSliceIn[:deleteEventIndex], nspEventSliceIn[deleteEventIndex+1:]...)
	return nspEventSliceOut
}
