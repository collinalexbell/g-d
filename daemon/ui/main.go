// Import the ncurses library for the go lanuage

package main

import (
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {
	// Initialize the ncurses library
	win, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	defer goncurses.End()

	// Clear the screen
	win.Clear()

	if err != nil {
		log.Fatal("New Window failed")
	}

	win.Refresh()
	title := "Tree Cybernetics Daemon UI"
	x := getStartingXForCenteredTitle(win, title)
	// Print a message in the window
	win.MovePrint(1, x, title)

	drawMainMenu(win)


	// Refresh the screen
	win.Refresh()
}

// getXMidpoint
func getXMidpoint(win *goncurses.Window) int {
	// Get the window's width
	_ , width := win.MaxYX()
	// Return the midpoint
	return width / 2
}

func getStartingXForCenteredTitle(win *goncurses.Window, title string) int {
	// Get the midpoint of the window
	midpoint := getXMidpoint(win)
	// Get the length of the title
	titleLength := len(title)
	// Get the midpoint of the title
	titleMidpoint := titleLength / 2
	// Return the starting X position for the title
	return midpoint - titleMidpoint
}


// Draws the main menu for subsequent windows in this application
// Options include: 
//    - drive robot arm
//    - simulate robot arm
//    - take the AI agent daemon on a virtual adventure
//    - exit the application
func drawMainMenu(win *goncurses.Window) {
	win.Clear()
	// Get the midpoint of the window
	midpoint := getXMidpoint(win)
	// Get the length of the title
	titleLength := len("Robot Arm Control")
	// Get the midpoint of the title
	titleMidpoint := titleLength / 2
	// Draw the title
	win.MovePrint(1, getStartingXForCenteredTitle(win, "Robot Arm Control"), "Robot Arm Control")
	// Draw the menu options
	win.MovePrint(3, midpoint - titleMidpoint, "0) Drive Robot Arm")
	win.MovePrint(4, midpoint - titleMidpoint, "1) Simulate Robot Arm")
	win.MovePrint(5, midpoint - titleMidpoint, "2) Take the AI agent on a virtual adventure")
	win.MovePrint(6, midpoint - titleMidpoint, "3) Exit")
	handleMainMenuSelection(win)
}


func handleMainMenuSelection(win *goncurses.Window) {
	// Get the user's selection using goncurses
	selection := win.GetChar()
	// Handle the user's selection
	switch selection {
	case '0':
		// Drive the robot arm
		drawDriveRobotArmMenu(win)
	case '1':
		// Simulate the robot arm
		drawSimulateRobotArmMenu(win)
	case '2':
		// Take the AI agent on a virtual adventure
		drawVirtualAdventureMenu(win)
	case '3':
		// Exit the application
		win.MovePrint(7, getXMidpoint(win) - len("Goodbye!") / 2, "Goodbye!")
		win.GetChar()
	default:
		// Handle invalid selections
		win.MovePrint(7, getXMidpoint(win) - len("Invalid Selection") / 2, "Invalid Selection")
		win.GetChar()
		drawMainMenu(win)
	}
}

func drawDriveRobotArmMenu(win *goncurses.Window) {
	win.Clear()
	// Get the midpoint of the window
	midpoint := getXMidpoint(win)
	// Get the length of the title
	titleLength := len("Drive Robot Arm")
	// Get the midpoint of the title
	titleMidpoint := titleLength / 2
	// Draw the title
	win.MovePrint(1, getStartingXForCenteredTitle(win, "Drive Robot Arm"), "Drive Robot Arm")
	// Draw the menu options
	win.MovePrint(3, midpoint - titleMidpoint, "0) Drive the robot arm")
	win.MovePrint(4, midpoint - titleMidpoint, "1) Return to main menu")
	// Get the user's selection using goncurses
	selection := win.GetChar()
	// Handle the user's selection
	switch selection {
	case '0':
		// Drive the robot arm
		drawDriveRobotArmMenu(win)
	case '1':
		// Return to the main menu
		drawMainMenu(win)
	default:
		// Handle invalid selections
		win.MovePrint(7, getXMidpoint(win) - len("Invalid Selection") / 2, "Invalid Selection")
		win.GetChar()
		win.Delete()
		win.Refresh()
		drawDriveRobotArmMenu(win)
	}
}


func drawSimulateRobotArmMenu(win *goncurses.Window) {
	win.Clear()
	// Get the midpoint of the window
	midpoint := getXMidpoint(win)
	// Get the length of the title
	titleLength := len("Simulate Robot Arm")
	// Get the midpoint of the title
	titleMidpoint := titleLength / 2
	// Draw the title
	win.MovePrint(1, getStartingXForCenteredTitle(win, "Simulate Robot Arm"), "Simulate Robot Arm")
	// Draw the menu options
	win.MovePrint(3, midpoint - titleMidpoint, "0) Simulate the robot arm")
	win.MovePrint(4, midpoint - titleMidpoint, "1) Return to main menu")
	// Get the user's selection using goncurses
	selection := win.GetChar()
	// Handle the user's selection
	switch selection {
	case '0':
		// Simulate the robot arm
		drawSimulateRobotArmMenu(win)
	case '1':
		// Return to the main menu
		drawMainMenu(win)
	default:
		// Handle invalid selections
		win.MovePrint(7, getXMidpoint(win) - len("Invalid Selection") / 2, "Invalid Selection")
		win.GetChar()
		win.Delete()
		win.Refresh()
		drawSimulateRobotArmMenu(win)
	}
}

func drawVirtualAdventureMenu(win *goncurses.Window) {
	win.Clear()
	// Get the midpoint of the window
	midpoint := getXMidpoint(win)
	// Get the length of the title
	titleLength := len("Take the AI agent on a virtual adventure")
	// Get the midpoint of the title
	titleMidpoint := titleLength / 2
	// Draw the title
	win.MovePrint(1, getStartingXForCenteredTitle(win, "Take the AI agent on a virtual adventure"), "Take the AI agent on a virtual adventure")
	// Draw the menu options
	win.MovePrint(3, midpoint - titleMidpoint, "0) Take the AI agent on a virtual adventure")
	win.MovePrint(4, midpoint - titleMidpoint, "1) Return to main menu")
	// Get the user's selection using goncurses
	selection := win.GetChar()
	// Handle the user's selection
	switch selection {
	case '0':
		// Simulate the robot arm
		drawVirtualAdventureMenu(win)
	case '1':
		// Return to the main menu
		drawMainMenu(win)
	default:
		// Handle invalid selections
		win.MovePrint(7, getXMidpoint(win) - len("Invalid Selection") / 2, "Invalid Selection")
		win.GetChar()
		win.Delete()
		win.Refresh()
		drawVirtualAdventureMenu(win)
	}
}
