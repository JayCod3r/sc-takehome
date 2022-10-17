package folders

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
)

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
// - Component 2:
//   - Extend the existing code to facilitate pagination. You can copy over the existing methods into `folders_pagination.go` to get started.
//   - Pagination helps break down a large dataset into smaller chunks.
//   - The small data chunk will then be served to the client side usually accompanied a token that points to the next chunk.
//   - Write a short explanation on why you choosen the solution you implemented.
//   - The end result should look like this:
// ```
//   original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

//   The current result will look like this:
//   { data: [1, 2, 3, ..., 10] }

//   With pagination implementation, the result should look like this:
//   { data: [1, 2], token: "Mw==" }

//   The token can then be used to fetch more result:

//   { data : [3, 4], token: "NQ==" }

//   And more results util there's no data left:

// { data: [9, 10], token: null }
func pagination() {
	// The aim is to group up all the folders in twos and pair them with a token. Also Have a token that directs them towards the next set.
	// If I put a loop and have a map with a key. and it will put two folders into each segement then have a key. to access that area in the map.
	fmt.Println("hello")

}

func ModGetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) { // This is my modified version of GetAllFolder

	FolderOfFolders, err := ModFetchAllFoldersByOrgID(req.OrgID) //
	if err != nil {
		return &FetchFolderResponse{}, errors.New("Undefined for negative numbers")
	}
	var fetchResponse *FetchFolderResponse                         // var pointer to FetchFoldersResponse.
	fetchResponse = &FetchFolderResponse{Folders: FolderOfFolders} // ffr equals dereferences FetchFolderResponse for

	return fetchResponse, nil // it returns fetchFolderResponse and nil as the error.

}

func ModFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) { // the parameters are orgID from package uuid.UUID.  The Return type is a slice of pointers and an error.
	folders := GetSampleData() // this initialises a var folders

	resFolder := []*Folder{} // this initialises a slice of pointers.
	for _, folder := range folders {
		if folder.OrgId == orgID { // if folder.OrgID is equal to orgID , Add folder to resFolder
			resFolder = append(resFolder, folder) // this appends folder A POINTER TO Folder , to resfolder
		}
	}
	return resFolder, nil // it returns resFolder as the slice of pointer to Folder, and returns error as nil.
}
