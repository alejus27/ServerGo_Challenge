# ServerGo_Challenge

Go TCP server to manage files between clients.

The clients can subscribe to channels to send and receive files through CLI.

(Default channels ch_1, ch_2, ch_3).

# Server
Enter the server folder and follow the next steps.

  - Run: go run .
  
  - Start: server start
  
  - Stop: server stop


# Client
Enter the client folder and follow the next steps.

  - Run: go run .
  
  - Subscribe channel: subscribe channel:name
  
  - Unsubscribe channel: unsubscribe channel:name

  - Send file: send channel:name file:path
