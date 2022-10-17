package folders

import (
	"errors"
	"fmt" // imported this to check what is going on overall in the ranges.

	"github.com/gofrs/uuid"
)

// THE GETALLFOLDERS METHOD (Original version)
// This method aims to Get All Folders
// It initialises variables.  An error variable, a Folder Struct Variable, a slice of pointers variable
// Then it initialises a slice of Struct Folders.
// Then it creates a slice of pointers to Folders from calling the FetchAllFoldersByOrgID method.
// It goes through the created slice and appends into the slice of Struct of folders. f.
//
// Then initialises a slice of pointers to folder called fp
// Go through the slice of folder structs names f and grab every pointer and append to fp (a slice of pointers to folders)
//
// Create a variable that is a pointer to fetchfolderresponse
//make the ffr equal to the dereference of Folders within fp
// THEN return the ffr and nil.

//QUICK VERSION OF WHAT GETALLFOLDERS DOES.
// grabs all folders by id,  adds them to a slice of folders,
// grabs the pointers of all the structs
//and adds it to a slice of pointers to folders.
// Put it all into the FetchFolderresponse variable

func OldGetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) { // Parameters are (req *FetchFolderRequest) , It returns (*FetchFolderResponse, error)
	// var (
	// 	err error // initialises a variable of error type
	// 	f1  Folder // initialises a folder struct
	// 	fs  []*Folder // fs initialised as a slice of pointers to folders
	// )
	f := []Folder{}                           // this  initialises a slice of struct folder
	r, _ := FetchAllFoldersByOrgID(req.OrgID) // r is a slice of pointers to folder, _ is an error.
	for k, v := range r {                     // for , pointer to folder, in a slice of pointers to folder.
		f = append(f, *v) // this appends *v to the slice f
		fmt.Println(k)

	}
	var fp []*Folder        // initialises a slice of pointers to folders
	for k1, v1 := range f { // This is a range loop . Its searching through all Folders in a slice of Folders.
		fp = append(fp, &v1) // this appends &v1 to the slice of fp.
		fmt.Println(k1)
	}
	var ffr *FetchFolderResponse            // var pointer to FetchFoldersResponse.
	ffr = &FetchFolderResponse{Folders: fp} // fetchfolder response struct
	return ffr, nil                         // returns the ffr and error as nil
}

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) { // THIS IS MY MODIFIED GetAllFolders()
	FolderOfFolders, err := FetchAllFoldersByOrgID(req.OrgID) // Made variables easier for me to understand as Folder Of Folders is more readable than a letter.
	if err != nil {
		return &FetchFolderResponse{}, errors.New("Error OCCURRED IN GETALLFOLDERS") // if an error is present in the getallfolders function then it will return there was an error to help fix it if need be.
	}
	var fetchResponse *FetchFolderResponse                         // var pointer to FetchFoldersResponse.
	fetchResponse = &FetchFolderResponse{Folders: FolderOfFolders} // ffr equals dereferences FetchFolderResponse for

	return fetchResponse, nil // it returns fetchFolderResponse and nil as the error. // FetchFolder response is a struct of Folders[] *Folder

}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // this initialises a var folders

	resFolder := []*Folder{}         // this initialises a slice of pointers.
	for _, folder := range folders { // range loop through folders.
		if folder.OrgId == orgID { // if folder.OrgID is equal to orgID , Add folder to resFolder
			resFolder = append(resFolder, folder) // this appends folder A POINTER TO Folder , to resfolder
		}
	}
	return resFolder, nil // it returns resFolder as the slice of pointer to Folder, and returns error as nil.
}

//Comments
// One comment I would make would be to have clearer variables. As it would help with code readability if the variable names are more intuitive.
// The GetAllFolders function has a return of an error yet only returns nil. I would add error handling of some kind to utilise this feature. Otherwise its just a waste.

func Pagination(req *FetchFolderRequest) (*FetchFolderResponse, error) { // THIS IS MY PAGINATION
	/////////The way I got this section to work was I renamed this method to GetAllFolders. And changed the name of the other GetAllFolders Method.
	/////////
	// This section puts the data into the Folders2
	Folders2 := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	for k, v := range r { // for , pointer to folder, in a slice of pointers to folder.
		Folders2 = append(Folders2, *v) // this appends *v to the slice f
		k++
	}
	/////////
	var token int
	for i := 0; i <= 50; i++ {

		fmt.Println("Choose a token between 0 and 499")
		fmt.Scanln(&token)
		if i == 49 { // this is just to keep track if you keep going through
			fmt.Println("___________________________________________________")
			fmt.Println("One more search left ")
			fmt.Println("___________________________________________________")

		}
		if token <= -2 || token >= 500 { // This is because The tokens start at 0-499 as 0 is the first one. so there are still 500 token
			// Also I don't want them to type in a negative integer as it would bring up none of the entries.

			fmt.Println("Error: Invalid Token number")

		} else if token == -1 { // Making this the exit point
			fmt.Println("___________________________________________________")
			fmt.Println("Thank you")
			fmt.Println("___________________________________________________")
			var fetchResponse *FetchFolderResponse
			return fetchResponse, nil // returns to end // This section is a WIP
		} else if token == 0 { // Exception 1 if the token is 0 it would need a different system as it is an outlier.
			fmt.Println("___________________________________________________")
			fmt.Print("Token: ")
			fmt.Println(token)
			fmt.Print("Folder entry 1: ")
			fmt.Println(Folders2[token])
			fmt.Print("Folder entry 2: ")
			fmt.Println(Folders2[token+1])
			fmt.Println("Previous Token: No Previous Token")
			fmt.Println("Next Token: 1")
			fmt.Println("Enter -1 to end")
			fmt.Println("___________________________________________________")
		} else if token == 499 { // Exception 2 as the final token has differences.
			fmt.Println("___________________________________________________")
			fmt.Print("Token: ")
			fmt.Println(token)
			fmt.Println("Folder entry 1: ")
			fmt.Println(Folders2[token*2])
			fmt.Print("Previous Token: ")
			fmt.Println(token - 1)
			fmt.Println("Next Token: No Next Token")
			fmt.Println("Enter -1 to end")
			fmt.Println("___________________________________________________")
		} else {
			fmt.Println("___________________________________________________")
			fmt.Print("Token: ")
			fmt.Println(token)
			fmt.Print("Folder entry 1:")
			fmt.Println(Folders2[token*2])
			fmt.Print("Folder entry 2:")
			fmt.Println(Folders2[token*2+1])
			fmt.Print("Previous Token: ")
			fmt.Println(token - 1)
			fmt.Print("Next Token: ")
			fmt.Println(token + 1)
			fmt.Println("Enter -1 to end")
			fmt.Println("___________________________________________________")
		}

	}

	var fetchResponse *FetchFolderResponse
	return fetchResponse, nil
}
