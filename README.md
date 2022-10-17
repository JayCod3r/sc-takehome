# sc-takehome
My application to sc

**Component 1**

My overall Comments and improvements
I made it so the variables were more readable and intuitive rather than just letters. 
I find this to be easier to understand overall when variables are more related to what they are doing.

I added an error handler as it was only passing back nil and wasn't utilising that return variable so I added one in to show where there error may occur.

//

**Testing**
I tried to implement tests but this section currently is a WIP
//

**Component 2**

I completed my pagination in the Folders.go folder.
Steps to get my Function working
Rename GetAllFolders to GetAllFolders1
Rename Pagination to GetAllFolders

**My method for pagination**
It prompts the user to enter a token number from 0-499
It then will print out the two folders related to that token number, Also it will print out the previous token and the next token.

Exceptions found -
0 would act differently - therefoer there are changes in particular to the previous token, Also the way it grabs the folders.
499 would act differently - therefore there are changes in particular to the next token

Added a way out by prompting the user to input -1 as an out. 
Invalid tokens a number <= -2 or number >=500, If the user was to input an invalid number there is a catch 

.
