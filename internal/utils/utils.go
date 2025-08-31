package utils

import "github.com/inancgumus/screen"

func CallClearConsoleSc() {
	// Clears the screen
	screen.Clear()
	// Moves the cursor to the top left corner of the screen
	screen.MoveTopLeft()
}
