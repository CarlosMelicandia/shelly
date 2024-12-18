This operations folder should contain all the CRUD operations that either the client or the server wants to do.
For example, the information that will be passed to the function inside of `create_user.go` was created on the server
when the user logged in with google. The information that will be passed to the function inside of `create_hacker.go`
was created on the client (from the application form), and be called on the server to add the hacker to the DB.
