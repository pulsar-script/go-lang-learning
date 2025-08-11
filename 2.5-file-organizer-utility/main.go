package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//* defining rules
// this help to tell programe in which output dir e.g. imagesDir, documentsDir choose base on files extension
// in simple way you can use switch or even use if else , but that is not good for multiple file format like 100s u have

var Rules = map[string]string{
	".jpg": "imagesDir",
	".txt": "documentsDir",
}

func main() {

	//TODO take target directory from user
	messayDir, err := os.ReadDir("./messayDir") //* it return os.DirEntry [read below]
	// fmt.Println("")
	// fmt.Println("messay dir output =>", messayDir)

	// error handling
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	// for loops and operation on each file
	for _, entry := range messayDir {

		// check if that entry is Dir ?
		if !entry.IsDir() {

			// it is a file lets print itsname
			fileName := entry.Name()
			// fmt.Println("File - ", fileName)

			// extract file extension from from file name
			fileExtension := filepath.Ext(fileName)

			// find its sorting Dir by using rules
			destinationDirName := Rules[fileExtension]

			// TODO if sorting folder not exists
			if destinationDirName == "" {
				destinationDirName = "mesclenious"
			}

			rootDir := "./sortedDir"
			// full path
			destinationDirPath := filepath.Join(rootDir, destinationDirName)
			// *TEST
			// fmt.Println(destinationDirPath)

			//if destination Dir not exists e.g. imagesDir , documentsDir
			// lets make then
			err := os.MkdirAll(destinationDirPath, 0755)

			//error handling
			if err != nil {
				fmt.Println("")
				fmt.Println(err)
			}

			// to move files , we need full source path and destination path
			sourceDirFullPath := filepath.Join("./messayDir", fileName)
			destinationDirFullPath := filepath.Join(destinationDirPath, fileName)

			// moving files
			err = os.Rename(sourceDirFullPath, destinationDirFullPath)

			if err != nil {
				fmt.Println("")
				fmt.Println(err)
			}

		}

	}

	fmt.Println("")
	fmt.Println("All files are successfully sorted on the basis on extensions (types) and stre to new location")

}

/*
* --- Notes on Reading Directory Contents in Go ---


	While you asked about `os.Dir`, it's important to know about its modern replacement,
	`os.ReadDir`, which is what you'll almost always use.

	---------------------------------------------------------------------
	RECOMMENDED: os.ReadDir(name string) ([]os.DirEntry, error)
	---------------------------------------------------------------------
	- This is the modern, preferred function to read the contents of a directory.
	- It is more efficient than older methods because it doesn't fetch full file
	  statistics (like modification time or file size) right away.
	- It returns a "slice" (a dynamic list) of os.DirEntry objects.
	- You always get two return values: the slice of entries, and a potential error.

	---------------------------------------------------------------------
	WHAT IS os.DirEntry?
	---------------------------------------------------------------------
	- An os.DirEntry is a lightweight object that represents a SINGLE item
	  (a file or a subdirectory) found inside a directory.
	- You don't create an os.DirEntry yourself; you receive it from os.ReadDir.

	Key Methods of an os.DirEntry variable (let's call it 'entry'):
	  - entry.Name()  -> Returns the name of the file/directory as a string (e.g., "notes.txt").
	  - entry.IsDir() -> Returns 'true' if the entry is a directory, 'false' if it's a file.


 --- Example Usage ---
 The following is a complete, working example.
 It demonstrates how to use os.ReadDir and os.DirEntry to list
 the contents of the current directory
 package mai
 import (
 	"fmt"
 	"log"
 	"os"

 func main() {
 	 Define the path to the directory you want to read.
 	 The string "." is a shortcut for the current directory.
 	const dirPath = ".
 	fmt.Printf("Reading contents of directory: %s\n", dirPath)
 	fmt.Println("------------------------------------"
 	 Use os.ReadDir to get a slice of os.DirEntry objects.
 	entries, err := os.ReadDir(dirPath
 	 Always handle potential errors. For example, the directory might
 	 not exist or you may not have permission to read it.
 	if err != nil {
 		log.Fatal(err) // log.Fatal prints the error and exits the program.

 	 Loop through each os.DirEntry in the 'entries' slice.
 	for _, entry := range entries
 		 Use the IsDir() method to check if it's a directory.
 		if entry.IsDir() {
 			 If it's a directory, print its name.
 			fmt.Printf("DIR : %s\n", entry.Name())
 		} else {
 			 If it's a file, print its name.
 			fmt.Printf("FILE: %s\n", entry.Name())
 		}
 	}


*/

/*
---------------------------------------------------------------------
--- Notes on Key Packages & Functions Used in This Project ---
---------------------------------------------------------------------

This file contains notes on the important building blocks from Go's
standard library that were used to create the File Organizer Utility.

---------------------------------------------------------------------
--- os Package: Interacting with the Operating System ---
---------------------------------------------------------------------

The `os` package provides a platform-independent interface to
operating system functionality. We used it for creating directories,
moving files, and handling program termination.

 --- os.MkdirAll(path string, perm os.FileMode) error ---

 What it does:
   Creates a directory at the specified `path`.

 Key Points & Why we use it:
   1. Creates Parent Directories: If any parent directories in the path
      do not exist, `os.MkdirAll` will create them. For example, if you
      ask to create "FolderA/FolderB" and "FolderA" doesn't exist, it
      will create "FolderA" first, then "FolderB" inside it.
   2. Idempotent: This is the most important reason we use it. If the
      directory already exists, `os.MkdirAll` does NOTHING and does NOT
      return an error. This is perfect for our loop, as we can call it
      for every single file without worrying if the destination folder
      has already been made.
   3. `perm`: The `0755` is a standard permission mode for directories.
      It means the owner can read/write/execute, and others can read/execute.

 Example:
   err := os.MkdirAll("./Images", 0755
 --- os.Rename(oldpath, newpath string) error ---

 What it does:
   Moves a file or directory from `oldpath` to `newpath`. The function
   name is `Rename`, but it's the primary way to "move" files in Go.

 Key Points & Why we use it:
   1. Requires Full Paths: Both `oldpath` and `newpath` must be complete
      paths to the file. This is why we use `filepath.Join` to build them.
   2. Move vs. Rename: If `oldpath` and `newpath` are in different directories,
      the file is moved. If they are in the same directory but have different
      names, the file is renamed.

 Example:
   source := "./messayDir/photo.jpg"
   destination := "./Images/photo.jpg"
   err := os.Rename(source, destination
 --- os.Args []string ---

 What it is:
   A built-in variable (a slice of strings) that holds all the
   command-line arguments passed to your program when it was run.

 Key Points:
   - os.Args[0] is always the name of the program itself.
   - os.Args[1] is the first argument provided by the user.
   - IMPORTANT: Always check `len(os.Args)` before trying to access an
     index like `os.Args[1]`. If the user didn't provide an argument,
     accessing it will crash your program (a "panic").

 Example:
   if len(os.Args) < 2 {  handle error  }
   targetDir := os.Args[1
 --- os.Exit(code int) ---

 What it does:
   Causes the current program to exit immediately.

 Key Points:
   - The integer `code` is an "exit code" that tells the system if the
     program finished successfully or with an error.
   - Conventionally, `os.Exit(0)` means success.
   - A non-zero code, usually `os.Exit(1)`, indicates that an error occurred.
   - A function like `log.Fatal()` is often useful because it prints a
     message and then calls `os.Exit(1)` for you.

 Example:
   fmt.Println("Error: Directory not provided.")
   os.Exit(1
 ---------------------------------------------------------------------
 --- path/filepath Package: Managing File Paths Safely ---
 ---------------------------------------------------------------------

 The `path/filepath` package provides functions to work with file paths in a
 way that is compatible with any operating system (Windows, macOS, Linux).
 This is crucial for writing portable and bug-free code
 --- filepath.Join(elem ...string) string ---

 What it does:
   Intelligently joins any number of path elements into a single path.

 Key Points & Why we use it:
   - Platform-Agnostic: This is the #1 reason. It automatically uses the
     correct path separator for the OS (`\` on Windows, `/` on Linux/macOS).
   - Cleanliness: It handles edge cases like extra slashes for you.
   - ALWAYS use this instead of manually concatenating path strings with "+".

 Example:
    On Windows, this becomes: "C:\Users\Project\Images\photo.jpg"
    On Linux, this becomes: "/home/user/Project/Images/photo.jpg"
   path := filepath.Join("C:", "Users", "Project", "Images", "photo.jpg"
 --- filepath.Ext(path string) string ---

 What it does:
   Extracts the file extension from a path.

 Key Points:
   - It returns the part of the filename starting from the LAST dot (`.`).
   - The dot is included in the returned string.
   - If there is no dot, it returns an empty string `""`.

 Example:
   filepath.Ext("image.jpeg")       // Returns ".jpeg"
   filepath.Ext("archive.tar.gz")   // Returns ".gz"
   filepath.Ext("document")         // Returns ""
   filepath.Ext(".bash_profile")    // Returns ".bash_profile"


*/
