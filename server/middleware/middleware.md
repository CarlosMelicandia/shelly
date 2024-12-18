These functions will run before the server handle routes get executed. After implementing authorization (authorization is not the same as authentication)
on the server, there were many weird issues that occurred with protected routes like /admin. There was also another problem of how to handle the UI when a user
was not allowed to look at a specific page. This made me decide to handle any type of authorization on the client.
