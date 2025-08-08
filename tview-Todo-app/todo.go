package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TodoItem represents a single item in our list.
type TodoItem struct {
	Text      string
	Completed bool
}

func main() {
	app := tview.NewApplication()

	// This slice holds the state of our application.
	todos := []*TodoItem{
		{Text: "Buy milk", Completed: false},
		{Text: "Learn tview", Completed: true},
		{Text: "Write a todo app", Completed: false},
	}

	// The list widget to display the todos.
	list := tview.NewList()
	list.SetBorder(true).SetTitle("Todo List")

	// The input field for adding new todos.
	inputField := tview.NewInputField().
		SetLabel("New Todo: ").
		SetFieldWidth(0). // Use full width
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorWhite)

	// A helper function to redraw the list from the `todos` slice.
	// This is the key to keeping the UI in sync with our data.
	redrawList := func() {
		list.Clear()
		for i, todo := range todos {
			prefix := "[ ] "
			if todo.Completed {
				prefix = "[✔] "
			}
			// We use the index `i` as the "shortcut" rune.
			// tview doesn't use it for navigation here, but it's good practice.
			list.AddItem(prefix+todo.Text, "", rune(i), nil)
		}
	}

	// --- Event Handlers ---

	// Handle "Enter" key on the input field to add a new todo.
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			text := inputField.GetText()
			if text != "" {
				todos = append(todos, &TodoItem{Text: text, Completed: false})
				redrawList()
				inputField.SetText("") // Clear the input field
			}
		}
	})

	// Handle "Enter" key on a list item to toggle its completed status.
	list.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(todos) {
			todos[index].Completed = !todos[index].Completed
			redrawList()
		}
	})

	// Handle other key presses on the list for deleting items.
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// `KeyDEL` is Delete on some keyboards, `KeyBackspace2` is Backspace.
		if event.Key() == tcell.KeyDelete || event.Key() == tcell.KeyBackspace2 {
			index := list.GetCurrentItem()
			if index >= 0 && index < len(todos) {
				// Remove the item from the slice
				todos = append(todos[:index], todos[index+1:]...)
				redrawList()
				// If the list is now empty, we might need to adjust focus or selection.
				// For simplicity, we just redraw. tview handles the rest gracefully.
			}
			return nil // We've handled the event
		}
		return event // We haven't handled it, pass it on
	})

	// Initial drawing of the list.
	redrawList()

	// A text view for instructions at the bottom.
	instructions := tview.NewTextView().
		SetText(fmt.Sprintf("Navigate: Arrow Keys | Toggle: Enter | Delete: Del/Backspace | Add: Type and press Enter | Quit: Ctrl-C")).
		SetTextColor(tcell.ColorGray)

	// Create the main layout using a Flex box.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).       // Arrange items vertically
		AddItem(list, 0, 1, true).         // The list takes up 1 part of the space and has focus.
		AddItem(inputField, 1, 0, false).  // The input field is 1 row high.
		AddItem(instructions, 1, 0, false) // The instructions are 1 row high.

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
