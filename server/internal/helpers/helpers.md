This is a pretty self-explanatory folder. We want to have use functions that we frequently call to be used in other files.
I personally decided to create sub folders to avoid circular dependency when importing and a bit more organized.
Ideally, it would be good to call the file `info.go` for a generic name. For example, if we create a folder called token,
we want the `info.go` file to have functions like GetToken and GetTokenString. If there are other function that might be of use,
put it in another file.
