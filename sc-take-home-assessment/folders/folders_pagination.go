package folders

import (
	"errors"

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
	// f := []Folder{}
	// r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// for k, v := range r { // for , pointer to folder, in a slice of pointers to folder.
	// 	f = append(f, *v) // this appends *v to the slice f
	// 	k++
	// }

	// var token int
	// for i := 0; i <= 5; i++ {

	// 	fmt.Println("Choose a token between 0 and 499")
	// 	fmt.Scanln(&token)

	// 	if token <= -2 || token >= 500 { // This is because The tokens start at 0-499 as 0 is the first one. so there are still 500 token
	// 		// Also I don't want them to type in a negative integer as it would bring up none of the entries.

	// 		fmt.Println("Error: Invalid Token number")

	// 	} else if token == -1 { // Making this the exit point
	// 		fmt.Println("___________________________________________________")
	// 		fmt.Println("Thank you")
	// 		fmt.Println("___________________________________________________")
	// 		var fetchResponse *FetchFolderResponse
	// 		return fetchResponse, nil // returns to end // This section is a WIP
	// 	} else if token == 0 { // Exception 1 if the token is 0 it would need a different system as it is an outlier.
	// 		fmt.Println("___________________________________________________")
	// 		fmt.Print("Token: ")
	// 		fmt.Println(token)
	// 		fmt.Print("Folder entry 1: ")
	// 		fmt.Println(f[token])
	// 		fmt.Print("Folder entry 2: ")
	// 		fmt.Println(f[token+1])
	// 		fmt.Println("Previous Token: No Previous Token")
	// 		fmt.Println("Next Token: 1")
	// 		fmt.Println("Enter -1 to end")
	// 		fmt.Println("___________________________________________________")
	// 	} else if token == 499 { // Exception 2 as the final token has differences.
	// 		fmt.Println("___________________________________________________")
	// 		fmt.Print("Token: ")
	// 		fmt.Println(token)
	// 		fmt.Println("Folder entry 1: ")
	// 		fmt.Println(f[token*2])
	// 		fmt.Print("Previous Token: ")
	// 		fmt.Println(token - 1)
	// 		fmt.Print("Next Token: No Next Token")
	// 		fmt.Println("___________________________________________________")
	// 	} else {
	// 		fmt.Println("___________________________________________________")
	// 		fmt.Print("Token: ")
	// 		fmt.Println(token)
	// 		fmt.Print("Folder entry 1:")
	// 		fmt.Println(f[token*2])
	// 		fmt.Print("Folder entry 2:")
	// 		fmt.Println(f[token*2+1])
	// 		fmt.Print("Previous Token: ")
	// 		fmt.Println(token - 1)
	// 		fmt.Print("Next Token: ")
	// 		fmt.Println(token + 1)
	// 		fmt.Println("Enter -1 to end")
	// 		fmt.Println("___________________________________________________")
	// 	}

	// }

	// var fetchResponse *FetchFolderResponse
	// return fetchResponse, nil
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
