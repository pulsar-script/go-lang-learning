package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// scanner
	scanner := bufio.NewScanner(os.Stdin)

	//* Preparing Items

	//1 book
	Book := Item{
		ID:          "book",
		Name:        "Recipe Book",
		Location:    "lab",
		Description: "A huge, leather-bound book with arcane symbols on the cover.",
	}

	//2 crystal
	Crystal := Item{
		ID:          "crystal",
		Name:        "Glimmering Crystal",
		Location:    "cave",
		Description: "A shard of pure crystal that seems to pulse with its own inner light.",
	}

	//3 moonpetal
	Moonpetal := Item{
		ID:          "moonpetal",
		Name:        "Silver Moonpetal",
		Location:    "grove",
		Description: "A single, soft flower petal that shimmers like liquid moonlight.",
	}

	//4 dew
	Dew := Item{
		ID:          "dew",
		Name:        "River Dew",
		Location:    "riverbank",
		Description: "A perfect, shimmering drop of water, held together by magic.",
	}

	//* Preparing Rooms

	//1 lab
	Lab := Room{
		ID:          "lab",
		Name:        "Alchemist's Lab",
		Description: "You are in your master's messy laboratory. Bubbling beakers ⚗️ and strange contraptions cover every table. In the center of the room is a large Cauldron 🔥. The master's enormous Recipe Book 📖 sits on a pedestal. Doors lead north, south, east, and west.",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//2 cave
	Cave := Room{
		ID:          "cave",
		Name:        "The Crystal Caves",
		Description: "You enter a breathtaking cave where the walls are covered in glowing crystals. The air hums with energy ✨. A particularly bright Glimmering Crystal is embedded in the far wall. The only way out is back east to the lab.",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//3 grove
	Grove := Room{
		ID:          "grove",
		Name:        "Moonpetal Grove",
		Description: "You find a quiet, magical grove, bathed in moonlight 🌙. The plants here seem to glow with a soft, silver light. You see a single, perfect Silver Moonpetal 🌸 on a rare lunar orchid. A path leads back south.",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//4 riverbank
	Riverbank := Room{

		ID:          "riverbank",
		Name:        "Whispering Riverbank",
		Description: "You are standing on the bank of a gently flowing river 🏞️. The water is incredibly clear. Large, mossy stones near the edge are coated in sparkling River Dew 💧. The path back to the lab is to the north.",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//5 storeroom
	Storeroom := Room{
		ID:          "storeroom",
		Name:        "Storeroom",
		Description: "This is a dusty storeroom filled with empty boxes, cracked jars, and cobwebs 🕸️. It doesn't look like anything useful has been here for years. The only exit is west.",
		Exits:       make(map[string]*Room),
	}

	// player
	// ChemAssistant := Player{
	// 	CurrentRoom: nil,
	// 	Inventory:   make(map[string]*Item),
	// }

	//* Linking items to rooms
	//? why we cannot link rooms, items after each initilizalization
	// => Because we need all other items and books initialize before linking , before point that thing should exists

	// Lab
	Lab.Items["book"] = &Book

	// Cave
	Cave.Items["crystal"] = &Crystal

	// Grove
	Grove.Items["moonpetal"] = &Moonpetal

	// Riverbank
	Riverbank.Items["Dew"] = &Dew

	//* Linking Rooms to rooms

	// Lab
	Lab.Exits["north"] = &Grove
	Lab.Exits["south"] = &Riverbank
	Lab.Exits["east"] = &Storeroom
	Lab.Exits["west"] = &Cave

	// Cave
	Cave.Exits["east"] = &Lab

	// Grove
	Grove.Exits["south"] = &Lab

	// Riverbank
	Riverbank.Exits["north"] = &Lab

	// Storeroom
	Storeroom.Exits["west"] = &Lab

	//* All Room
	Gameworld := AllRooms{
		roomsMap: make(map[string]*Room),
	}

	// Linking
	Gameworld.roomsMap["lab"] = &Lab
	Gameworld.roomsMap["cave"] = &Cave
	Gameworld.roomsMap["grove"] = &Grove
	Gameworld.roomsMap["storeroom"] = &Storeroom

	// Game start

	fmt.Println("")
	fmt.Println("=========================================== The Alchemist's Assistant ⚗️Adventure Game ============================================")
	fmt.Print(" Your master, the great alchemist, has been called away on urgent business.\n A rare celestial event, the Starlight Convergence 🌟, is happening tonight!\n Your master left a note: \n The Elixir of Starlight must be completed by midnight. \n The recipe is in the grand grimoire. Do not fail! You must create the elixir.\n")

	// gameEngine:
	// 	for {

	fmt.Println("")
	fmt.Println("What do you want to do?")
	fmt.Print(">")
	scanner.Scan()
	userInputCmd := scanner.Text()

	// *TEST
	fmt.Printf("TEST 1 userInput => %#v\n", userInputCmd)

	// *TEST
	userInputCmd = strings.TrimSpace(userInputCmd)
	fmt.Printf("TEST 2 userInput => %#v", userInputCmd)
	fmt.Println("\n-------------------------------------------")

	// *TEST
	parts := strings.Fields(userInputCmd)

	var Command string
	var Argument string

	// Now you can safely access the command and its argument
	if len(parts) > 0 {

		Command = parts[0]
		fmt.Printf("The command is: %s\n", Command)

		if len(parts) > 1 {
			Argument = parts[1]
			fmt.Printf("The argument is: %s\n", Argument)
		}

	}

	// }

}

//* Game World Entitys struct

// Items struct
type Item struct {
	ID          string
	Name        string
	Location    string
	Description string
}

// Room struct
type Room struct {
	ID          string
	Name        string
	Description string
	Exits       map[string]*Room // "north" -> pointer to the Grove Room
	Items       map[string]*Item // "crystal" -> pointer to the Crystal Item
}

// Player struct
type Player struct {
	CurrentRoom *Room            // Pointer to the room the player is in
	Inventory   map[string]*Item // "crystal" -> pointer to the Crystal Item
}

// structure for all rooms
type AllRooms struct {
	roomsMap map[string]*Room
}

// ============================================================================
//  NOTE TO SELF: Data Structures for a Game Inventory (Map vs. Slice)
// ============================================================================

/*
 * SCENARIO:
 * We need to store a player's inventory in a text-based adventure game.
 * The most common operations will be:
 * - Checking if a player has an item ("use key", "drop sword").
 * - Adding an item ("take key").
 * - Removing an item ("drop key").
 * - Listing all items ("inventory").
 *
 * This note compares two approaches: using a map (dictionary/hash table)
 * vs. using a slice (list/array).
 */

// ----------------------------------------------------------------------------
//  APPROACH 1: Using a Map / Dictionary (Recommended)
// ----------------------------------------------------------------------------

// The key is the item's unique ID string, and the value is the item object.
// var inventory map[string]*Item

// --- CHECKING FOR AN ITEM ("use crystal") ---
/*
   if item, ok := inventory["crystal"]; ok {
        We have the item!
   }

   PROS:
   - EXTREMELY FAST: Finding an item is instant (O(1) complexity).
     It doesn't matter if the inventory has 5 or 5,000 items.
   - CLEAN CODE: The logic directly asks the question "Does this item exist?".
     This is the most common operation in the game logic.

   CONS:
   - UNORDERED: When listing items, their order is not guaranteed. For a
     game inventory, this is usually an acceptable trade-off.
*/

// --- REMOVING AN ITEM ("drop crystal") ---
/*
   delete(inventory, "crystal")

   PROS:
   - EXTREMELY FAST: Removal is also an instant (O(1)) operation.
   - SIMPLE CODE: The command to remove an item is direct and clean.
*/

// ----------------------------------------------------------------------------
//  APPROACH 2: Using a Slice / Array
// ----------------------------------------------------------------------------

// A simple ordered list of item objects.
// var inventory []*Item

// --- CHECKING FOR AN ITEM ("use crystal") -fmt
/*
   for _, item := range inventory {
       if item.ID == "crystal" {
            We have the item!
           break;
       }
   }

   CONS:
   - SLOW: To find an item, you must loop through the list (O(n) complexity).
     The more items the player has, the slower this check becomes.
   - VERBOSE CODE: The logic requires a full loop just to see if something exists.
*/

// --- REMOVING AN ITEM ("drop crystal") ---
/*
   To remove an item, you must first loop to FIND its index,
   then perform a complex operation to rebuild the slice without that element.

   CONS:
   - SLOW & COMPLEX: This is the biggest weakness. It's a multi-step,
     inefficient operation for something the player does often.
*/

// --- LISTING ITEMS ("inventory") ---
/*
   for _, item := range inventory {
        print(item.Name)
   }

   PROS:
   - ORDERED: This is the main advantage. Items are listed in the
     order they were picked up, which can feel more natural.
*/

// ============================================================================
//                                CONCLUSION
// ============================================================================
/*
 * For game inventories where you frequently look up items by a unique name/ID,
 * a MAP is almost always the better choice.
 *
 * The massive performance and code simplicity benefits for the most common
 * operations (finding, using, dropping) far outweigh the single benefit of
 * a slice (guaranteed order).
 *
 * >> Use a MAP for the inventory of the player, rooms, shops, etc.
 * >> Use a SLICE only if the order is critical AND you rarely need to find
 * or remove a specific item from the middle of the list.
 */
