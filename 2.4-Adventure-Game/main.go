//* Improvments are Notion folder checkout

/*
================================================================================
                    THE ALCHEMIST'S ASSISTANT - GAME GUIDE
================================================================================

This guide contains the world map and a full command walkthrough to win the game.

--------------------------------------------------------------------------------
                                  WORLD MAP
--------------------------------------------------------------------------------

                      +-----------------------+
                      |   Moonpetal Grove 🌙    |
                      |  (Item: moonpetal)    |
                      +-----------+-----------+
                                  ^ (north)
                                  |
                                  v (south)
+----------------+      +-----------------------+      +---------------+
| Crystal Cave 💎  | <--> |   Alchemist's Lab 🧪    | <--> |   Storeroom 📦  |
| (Item: crystal)|(west)|   (Item: book)        |(east)|   (Empty)       |
+----------------+      +-----------+-----------+      +---------------+
                                  ^ (south)
                                  |
                                  v (north)
                      +--------------------------+
                      |  Whispering Riverbank 🏞️  |
                      |      (Item: dew)         |
                      +--------------------------+


--------------------------------------------------------------------------------
                             COMMAND WALKTHROUGH
--------------------------------------------------------------------------------

Here is the sequence of commands to successfully complete the game:

1.  read book         (To learn what ingredients are needed)

2.  go west           (To go to the Crystal Cave)
3.  pick crystal      (To get the first ingredient)
4.  go east           (To return to the Lab)

5.  go north          (To go to the Moonpetal Grove)
6.  pick moonpetal    (To get the second ingredient)
7.  go south          (To return to the Lab)

8.  go south          (To go to the Whispering Riverbank)
9.  pick dew          (To get the third ingredient)
10. go north          (To return to the Lab)

11. show inventory    (Optional: to verify you have all three items)

12. use crystal       (To add the first ingredient to the cauldron)
13. use moonpetal     (To add the second ingredient to the cauldron)
14. use dew           (To add the final ingredient and win the game)

*/

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
		Description: "\nYou are in your master's messy laboratory. \nBubbling beakers ⚗️   and strange contraptions cover every table. \nIn the center of the room is a large Cauldron 🔥. \nThe master's enormous Recipe Book 📖 sits on a pedestal. \nDoors lead north, south, east, and west.\n",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//2 cave
	Cave := Room{
		ID:          "cave",
		Name:        "The Crystal Caves",
		Description: "\nYou enter a breathtaking cave where the walls are covered in glowing crystals. \nThe air hums with energy ✨. \nA particularly bright Glimmering Crystal is embedded in the far wall. \nThe only way out is back east to the lab.\n",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//3 grove
	Grove := Room{
		ID:          "grove",
		Name:        "Moonpetal Grove",
		Description: "\nYou find a quiet, magical grove, bathed in moonlight 🌙. \nThe plants here seem to glow with a soft, silver light. \nYou see a single, perfect Silver Moonpetal 🌸 on a rare lunar orchid. \nA path leads back south.\n",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//4 riverbank
	Riverbank := Room{

		ID:          "riverbank",
		Name:        "Whispering Riverbank",
		Description: "\nYou are standing on the bank of a gently flowing river 🏞️. \nThe water is incredibly clear. \nLarge, mossy stones near the edge are coated in sparkling River Dew 💧. \nThe path back to the lab is to the north.\n",
		Exits:       make(map[string]*Room),
		Items:       make(map[string]*Item),
	}

	//5 storeroom
	Storeroom := Room{
		ID:          "storeroom",
		Name:        "Storeroom",
		Description: "\nThis is a dusty storeroom filled with empty boxes, cracked jars, and cobwebs 🕸️. \nIt doesn't look like anything useful has been here for years. \nThe only exit is west.\n",
		Exits:       make(map[string]*Room),
	}

	// player
	ChemAssistant := Player{
		CurrentRoom: nil,
		Inventory:   make(map[string]*Item),
	}

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
	Riverbank.Items["dew"] = &Dew

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

	// //* All Room
	// Gameworld := AllRooms{
	// 	roomsMap: make(map[string]*Room),
	// }

	// // Linking
	// Gameworld.roomsMap["lab"] = &Lab
	// Gameworld.roomsMap["cave"] = &Cave
	// Gameworld.roomsMap["grove"] = &Grove
	// Gameworld.roomsMap["storeroom"] = &Storeroom

	// preparing cauldron

	CauldronPot := Cauldron{
		itemsMap: make(map[string]*Item),
	}

	// Game start

	fmt.Println("")
	fmt.Println("=========================================== The Alchemist's Assistant ⚗️  Adventure Game ============================================")
	fmt.Print("\nYour master, the great alchemist, has been called away on urgent business.\n A rare celestial event, the Starlight Convergence 🌟, is happening tonight!\n Your master left a note: \n The Elixir of Starlight must be completed by midnight. \n The recipe is in the grand grimoire. Do not fail! You must create the elixir.\n")

	ChemAssistant.loadNextRoom(&Lab)
	// GameEngine:
	for {

		fmt.Println("")
		fmt.Println("\nWhat do you want to do?")
		fmt.Print(">> ")
		scanner.Scan()
		userInputCmd := scanner.Text()

		// *TEST
		// fmt.Printf("TEST 1 userInput => %#v\n", userInputCmd)

		// *TEST
		// userInputCmd = strings.TrimSpace(userInputCmd)
		// fmt.Printf("TEST 2 userInput => %#v", userInputCmd)
		// fmt.Println("\n-------------------------------------------")

		// *TEST
		parts := strings.Fields(userInputCmd)

		var Command string
		var Argument string

		// Now you can safely access the command and its argument
		if len(parts) > 0 {

			Command = parts[0]
			// fmt.Printf("The command is: %s\n", Command)

			if len(parts) > 1 {
				Argument = parts[1]
				// fmt.Printf("The argument is: %s\n", Argument)
			}

		}

		// switch cases
		switch Command {
		case "go":
			ChemAssistant.travelToRoom(Argument) // e.g. go north

		case "read":
			if Argument == "book" {
				readBook(&ChemAssistant)

			} else {
				fmt.Println("")
				fmt.Println("I don't understand that command.")
			}

		case "pick":
			ChemAssistant.pickItem(Argument)
			// fmt.Printf("Inventory => %v", ChemAssistant.Inventory)

		case "show":
			if Argument == "inventory" {
				ChemAssistant.showInventory()
			} else {
				fmt.Println("")
				fmt.Println("I don't understand that command.")
			}

		case "use":
			CauldronPot.useItems(Argument, &ChemAssistant)

		default:
			fmt.Println("")
			fmt.Println("I don't understand that command.")
		}
	}

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

// // structure for all rooms
// type AllRooms struct {
// 	roomsMap map[string]*Room
// }

// Cauldron struct for adding all 3 ingrediant

type Cauldron struct {
	itemsMap map[string]*Item
}

// * [Method] on Player struct
func (player *Player) travelToRoom(nextRoomDirection string) {

	//! also we no need to check with current room , we are traveling base on directions , not on locations name
	// // checking if player is in that room already
	// if player.CurrentRoom.ID == nextRoomDirection {   //! we cannot compare them beacuse player have *Room.ID means , lab, cave == direction e.g. noeth , east
	// 	fmt.Println("")
	// 	fmt.Printf("You are already in %s Room\n", nextRoomDirection)
	// }

	// checking requiest next room have way from this room or not
	nextRoom, ok := player.CurrentRoom.Exits[nextRoomDirection]

	if ok {
		player.loadNextRoom(nextRoom)
		//? Why not *nextRoom ?
		// actully e.g. nextRoom =  *lab , means it already passing pointer / original value , so no need to send address of pointer
	} else {
		fmt.Println("")
		fmt.Println("You can't go that way.")
		// continue GameEngine
		//TODO [improvment] we can apply condition on commands with correct direction (that are not accessible to player in perticular room ), show cant go msg
		//TODO and  for commands like e.g. go asdsw (gibrish), can show cant understand msg
	}

}

// * [Method] for loading new room
func (player *Player) loadNextRoom(nextRoom *Room) {

	// updating players current room location
	player.CurrentRoom = nextRoom

	// printing next room text
	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Printf("Current Room %v \n", player.CurrentRoom.Name)
	fmt.Println("")
	fmt.Printf("%v", player.CurrentRoom.Description)
}

// *function for Read Book
func readBook(player *Player) {

	// check player location , it should be lab
	if player.CurrentRoom.ID == "lab" {
		fmt.Println("")
		fmt.Println("\nThe page for the Elixir of Starlight ✨ reads: \nCombine three sacred ingredients in the cauldron - a Glimmering Crystal 💎, \na Silver Moonpetal 🌸, and a drop of pure River Dew 💧.")
	} else {
		fmt.Println("")
		fmt.Println("You don't see a Book here.")
	}

}

// * [Method] for Player struct to check and add item into inventory
func (player *Player) pickItem(itemStringName string) {

	// check does item exits in that room
	foundItem, ok1 := player.CurrentRoom.Items[itemStringName]

	if !ok1 {
		fmt.Println("")
		fmt.Printf("You don't see a %v here.", itemStringName)
		return
	}

	// if exists check , already taken or not
	alreadyTakenItem, ok2 := player.Inventory[foundItem.ID]

	if ok2 {
		fmt.Println("")
		fmt.Printf("You are already carrying %v.", alreadyTakenItem.Name)
		return
	}

	// if not taken , then pick
	player.Inventory[foundItem.ID] = foundItem // when we found item in room by comma, ok syntax => that foundted item is *pointer (actule item we get by that syntax)
	// we can use that , for adding into invtory if we dont have added
	fmt.Println("")
	fmt.Printf("You have taken the %v.", player.Inventory[foundItem.ID].Name)

}

//* [Method] for Player struct to show Inventory

func (player *Player) showInventory() {
	fmt.Println("")
	fmt.Println("You are carrying:")
	// check inventory
	if len(player.Inventory) > 0 {

		for key, value := range player.Inventory {
			fmt.Printf("[%v : %v]\t", key, value.Name)
		}
	} else {
		fmt.Println("[ You are not carrying anything. ]")
	}
}

//* [Method] for player to use items from inventory

func (cauldron *Cauldron) useItems(itemString string, player *Player) {

	// check layer location it should be Lab
	if player.CurrentRoom.ID != "lab" {
		fmt.Println("")
		fmt.Println("You can't use that here. You need to be at the cauldron in the lab.")
		return
	}

	// check player have that item or not
	haveItem, ok := player.Inventory[itemString]

	if !ok {
		fmt.Println("")
		fmt.Printf("\nYou don't have a %v.\n", itemString)
		return
	}

	// check for win condition and showing message
	requiredItems := [][]string{{"dew", "The dew makes the mixture shimmer and swirl."}, {"crystal", "The crystal dissolves into liquid light in the cauldron."}, {"moonpetal", "The petal melts, releasing a scent of night air."}}

	// if have item then  add to caludron
	cauldron.itemsMap[haveItem.ID] = haveItem

	//* Here i can use simply switch syntax, but use loop for suppose items incrase  become 10 ,so that time making 10 case switch is not good option
	for _, value := range requiredItems {
		if value[0] == haveItem.ID {
			fmt.Println("")
			fmt.Printf("\n%v\n", value[1])
			break
		}
	}

	for _, itemID := range requiredItems {
		_, ok := cauldron.itemsMap[itemID[0]]

		if !ok {
			fmt.Println("")
			fmt.Print("\nYou are still need more items, that are missing !\n")
			return
		}

	}

	// if player collect all three items the , win msg
	fmt.Println("")
	fmt.Print("\nThe liquid in the cauldron begins to boil and swirl,\npulling in the starlight from the window above. \nIt transforms into a swirling, liquid galaxy in a bottle ✨! \nYou have successfully created the Elixir of Starlight! \nYou've made your master proud! \nYou win! 🏆\n")

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
