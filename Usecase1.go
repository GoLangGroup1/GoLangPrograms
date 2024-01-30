package main

import "fmt"

func main() {

var option int = 0;

vendorAndNodesMap := make(map[string][]string)

for{

fmt.Println("1. Add Vendor")
fmt.Println("2. Delete Vendor")
fmt.Println("3. Add Node type")
fmt.Println("4. Delete node type")
fmt.Println("5. Display all vendor and node types")
fmt.Println("6. Exit")

fmt.Println("Enter option: ");
fmt.Scanln(&option);

switch option {
case 1:
	addVendor(vendorAndNodesMap)
case 2:
	deleteVendor(vendorAndNodesMap)
case 3:
	addNodeType(vendorAndNodesMap)
case 4:
	deleteNodeType(vendorAndNodesMap)
case 5:
	displayAll(vendorAndNodesMap)
case 6:
	return
default:
	fmt.Println("Oops! Not a valid option")
}
}

}

func addVendor(vendorAndNodesMap map[string][]string){
	var newVendor string;
	fmt.Println("Enter new Vendor:")
	fmt.Scanln(&newVendor)

	_, ok := vendorAndNodesMap[newVendor]
	if !ok {
		vendorAndNodesMap[newVendor] = make([]string, 0, 20)
	}
}

func deleteVendor(vendorAndNodesMap map[string][]string){
	var newVendor string;
	fmt.Println("Enter Vendor to be deleted:")
	fmt.Scanln(&newVendor)

	_, ok := vendorAndNodesMap[newVendor]
	if ok {
		delete(vendorAndNodesMap, newVendor)
	}	
}

func addNodeType(vendorAndNodesMap map[string][]string) {
	var newVendor string;
	var newNodeType string;
	var nodeTypeSlice []string;

	fmt.Println("Enter new/existing Vendor:")
	fmt.Scanln(&newVendor)

	fmt.Println("Enter Node type");
	fmt.Scanln(&newNodeType)

	_, ok := vendorAndNodesMap[newVendor]
	if !ok {
		vendorAndNodesMap[newVendor] = make([]string, 0, 20)
		vendorAndNodesMap[newVendor] = append(vendorAndNodesMap[newVendor], newNodeType)
	}else{

		nodeTypeSlice = vendorAndNodesMap[newVendor]
		fmt.Println(len(nodeTypeSlice))

		//to check whether the combination already exist or not
		for _,nodeType := range nodeTypeSlice {
			if(nodeType == newNodeType){
				fmt.Printf("(%s,%s) is already present \n", newVendor, newNodeType)
				return
			}
		}
		vendorAndNodesMap[newVendor] = append(vendorAndNodesMap[newVendor], newNodeType)
	}
}

func deleteNodeType(vendorAndNodesMap map[string][]string){
	
	var deleteNodeVendor string;
	var deleteNodeType string;
	var deleteNodeIndex int;
	var nodeTypeSlice []string;

	fmt.Println("Enter Vendor type:")
	fmt.Scanln(&deleteNodeVendor)

	fmt.Println("Enter the node type to delete:")
	fmt.Scanln(&deleteNodeType)

	_, ok := vendorAndNodesMap[deleteNodeVendor]
	if !ok {
		fmt.Println("No such Vendor or Node")
	}else{
		nodeTypeSlice = vendorAndNodesMap[deleteNodeVendor]

		for index,nodeType := range nodeTypeSlice {
			if(nodeType == deleteNodeType){
				deleteNodeIndex = index
			}
		}
	}
	vendorAndNodesMap[deleteNodeVendor] = append(nodeTypeSlice[:deleteNodeIndex], nodeTypeSlice[deleteNodeIndex+1:]...)
}

func displayAll(vendorAndNodesMap map[string][]string){
	fmt.Println("\n-------------- Printing Map --------------------")
	for key, value := range vendorAndNodesMap {
		fmt.Printf("%s => %v\n", key, value)
	}
	fmt.Println("-------------------- End -------------------------")
	fmt.Println("")
	fmt.Println("")
}
