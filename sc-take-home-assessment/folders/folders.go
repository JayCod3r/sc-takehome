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
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
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
