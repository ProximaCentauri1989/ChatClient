# ChatClient

### Formulation of the task: 
Write chat client. The program after launch creates a socket on port 8080. If the port was able to bind, then the application considers itself a server and waits for messages from the client. If the port binding fails, the application joins to server and considers itself a client, after which it sends the server a welcome message (any text string) and the applications can start communication. If the client or server is shutting down, the application on the other side must report this in the terminal.
The application is considered working if the text written in the client console is displayed in the server console and vice versa, transmitted over the network.

### Project structure
The project consists of four packages.

| Package name  | Files                    | Description                |
| ------------- | ------------------------ |----------------------------|
| Main          | main.go, utils.go        | Implements main routine    |
| Server        | server.go, utils.go      | Implements server logic    |
| Client        | client.go, utils.go      | Implements client logic    |
| Config        | constants.go             |                   Constants|
| Failer        | failer.go                | Error handler              |

