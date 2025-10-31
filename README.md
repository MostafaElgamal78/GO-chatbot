# 🗨️ Go RPC Chat Application

A simple command-line chat system built using Go's **net/rpc** package.
It demonstrates client-server communication, concurrency, and basic synchronization using mutexes.

---

## 📁 Project Structure

```
Chatbot/
│
├── server.go   # RPC chat server implementation
└── client.go   # RPC chat client to send and receive messages
```

---

## 🚀 Features

* Send messages to a shared chat history.
* Retrieve the full chat history after every message.
* Thread-safe message handling using `sync.Mutex`.
* Multiple clients can connect simultaneously.

---

## ⚙️ How It Works

### Server (`server.go`)

1. Starts an RPC server on TCP port `1234`.
2. Exposes two methods:

   * `SendMessage(string, *[]string)` — adds a new message and returns the chat history.
   * `GetHistory(struct{}, *[]string)` — returns the current chat history.
3. Handles multiple clients concurrently with goroutines.

### Client (`client.go`)

1. Connects to the server using `rpc.Dial("tcp", "localhost:1234")`.
2. Accepts user input from the terminal.
3. Sends each message to the server and displays the updated chat history.
4. Type `exit` to quit the chat.

---

## 🧩 How to Run

1. Open two terminals in the project folder.

2. In the first terminal, start the server:

   ```bash
   go run server.go
   ```

3. In the second terminal, start the client:

   ```bash
   go run client.go
   ```

4. Type messages and see them broadcast in the chat history.
   You can open more terminals and run more clients to simulate multiple users.

---

## 🧠 Example Output

**Server:**

```
🚀 Chat server running on port 1234...
```

**Client:**

```
✅ Connected to chat server. Type your message below.
Type 'exit' to quit.

You: Hello
💬 Chat History:
- Hello
```

---

## 💡 Future Improvements

* Real-time message updates without re-sending messages.
* Add user names for each client.
* Store chat history in a file or database.
* Add a simple web or GUI interface.

---

## 🧑‍💻 Author

Made with Go by Mostafa Elgamal — for learning RPC and concurrent programming.
