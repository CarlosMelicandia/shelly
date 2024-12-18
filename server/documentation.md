# How to run the project
1. Create a terminal for the server
2. Navigate to the server entrypoint `cd server/cmd/api`
3. Run `go run main.go`

# How to setup and effectively contribute
1. Create two terminals (one for the client and another for the server)
2. Use one of the terminals to navigate to the client `cd client`. This will be used to develop client features and see it in real time with hot reloading.
3. Run `npm i && npm run dev`
4. Use the other terminal to run the server. Follow the steps "How to run the project".

Following these steps will have two localhosts running: localhost:4321 (client) and localhost:8000 (server)

## Client (Astro + Preact)
Astro is chosen as the main framework because:
* It excels in building static sites by sending minimal HTML to the client, ensuring excellent performance.
* Dynamic features are implemented using Server Islands (a unique Astro feature), allowing specific components to load JavaScript only when necessary.
* Astro has an intuitive architecture and plugin ecosystem, making it easy to integrate Tailwind CSS, ORM, and SEO tools.

Preact is utilized for dynamic interactivity due to its:
* Lightweight size (~3KB) compared to React.
* Compatibility with the React ecosystem for components.
* JavaScript with JSDoc is preferred over TypeScript to reduce technical debt while still enabling autocompletion and tooling support.

## Server (Go)
Go is chosen for its:
* Simplicity and readability, making it approachable for new developers.
* Excellent tooling for web applications.
* High performance and efficient concurrency, ideal for backend tasks.
* The server follows a clean and modular structure to improve maintainability and readability.

## Database
Turso (based on SQLite) is used for data persistence because of:
* A generous [free tier](https://turso.tech/pricing).
* Ability to create two other DB replications for free.


Below, there is a visual diagram to give understanding on how the codebase is structured:

```
server/
├── api/
│   ├── api.go
│   └── cmd/
│       ├── api/
│       │   └── main.go # Main entrypoint of the server
│       └── operations/ # CRUD operations here
│           ├── connect_discord.go
│           ├── create_hacker.go
│           ├── create_user.go
│           └── operations.md
├── internal/
│   ├── auth/
│   │   ├── utils/
│   │   │   └── util.go
│   │   ├── discord.go
│   │   └── google.go
│   ├── config/ # How we run the client dist in the server & access ENV keys
│   │   ├── client_dist.go
│   │   └── env.go
│   ├── handlers/ # Define routes and what functions will run depending on page
│   │   ├── admin.go
│   │   ├── api.go
│   │   ├── createHacker.go
│   │   ├── dashboard.go
│   │   ├── getHacker.go
│   │   └── getUser.go
│   │   └── handlers.md
│   ├── helpers/
│   │   ├── hacker/
│   │   ├── token/
│   │   └── user/
│   │   └── helpers.md
│   ├── tools/ # Function used to make CRUD operations to the DB
│   │   └── db.go
│   └── utils/
│       └── util.go
├── middleware/ # What will run before every route
│   ├── cors.go
│   ├── middleware.md
│   └── trailing_slash.go
├── documentation.md
├── go.mod
└── go.sum
```


                        +-------------------------+
                        |      Client (Astro)     |
                        |-------------------------|
                        | Static JSX & Tailwind   |
                        | Dynamic Components      |
                        | (Preact islands)        |
                        | Authorization           |
                        +-----------+-------------+
                                    |
                                    v
                        +-------------------------+
                        |      Server (Go)        |
                        |-------------------------|
                        | API Endpoints           |
                        | Handlers & Middleware   |
                        | Authentication          |
                        +-----------+-------------+
                                    |
                                    v
                        +-------------------------+
                        |       Database          |
                        |-------------------------|
                        |      Turso libSQL       |
                        +-------------------------+


