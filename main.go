package main

import (
	"fmt"
	"os/exec"
	"log"
	"strings"
	"strconv"
	"sort"
)

func main(){
	currentWorkspaceCommand := exec.Command("hyprctl","activeworkspace")
	currentWorkspace,err := currentWorkspaceCommand.Output()
	if(err != nil){
		log.Fatal(err)
	}

	currentDisplay := string(currentWorkspace[29:])
	currentMonitorBeforeTrim := strings.Split(currentDisplay,":")
	currentMonitor := strings.TrimSpace(currentMonitorBeforeTrim[0])

	currentWorkspaceIdString := strings.TrimSpace(string(currentWorkspace[12:14]))
	currentWorkspaceId,err := strconv.Atoi(currentWorkspaceIdString)
	if(err != nil){
		log.Fatal(err)
	}

	allWorkspacesCommand := exec.Command("hyprctl", "workspaces")
	allWorkspaces, errWorkspaces := allWorkspacesCommand.Output()
	allWorkspacesBeforeTrim := string(allWorkspaces)
	if(errWorkspaces != nil){
		log.Fatal(errWorkspaces)
	}
	allWorkspacesTrimmed := strings.TrimSpace(allWorkspacesBeforeTrim)
	allWorkspacesSplit := strings.Split(allWorkspacesTrimmed,"workspace")

	var workspaceIds []int

	for i := 1; i < len(allWorkspacesSplit); i++{
		workspaceIdString := string(allWorkspacesSplit[i][7:8])
		workspaceId,err := strconv.Atoi(workspaceIdString)
		if(err != nil){
			log.Fatal(err)
		}
		workspaceIds = append(workspaceIds, workspaceId)
	}

	changeDisplay(currentMonitor,currentWorkspaceId, workspaceIds)
}

func changeDisplay(display string, id int, allIds []int){
	IdSlice := allIds
	sort.Ints(IdSlice)

	nextIndex := getNextIndex(id, IdSlice)

	changeDisplayCommand := fmt.Sprintf("dispatch workspace %d",nextIndex)
	workspaceChangeCommand := exec.Command("hyprctl",changeDisplayCommand)
	_, err := workspaceChangeCommand.Output()
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(nextIndex)
}

func getNextIndex(id int, allIds []int)int{
	var currentWorkspaceIndex int

	for index,value := range allIds{
		if value == id{
			currentWorkspaceIndex = index
			break
		}
	}

	if currentWorkspaceIndex < len(allIds) - 1{
		return allIds[currentWorkspaceIndex + 1]
	}

	return 0
}
