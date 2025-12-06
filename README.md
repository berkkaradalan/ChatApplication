# Chat Application 💬

A real-time chat application built with Go to test and showcase the [CoreGo](https://github.com/berkkaradalan/corego) package capabilities.

## What is this?

This is a WebSocket-based chat app that demonstrates how CoreGo handles authentication, MongoDB operations, and real-time communication. Think of it as a playground for testing CoreGo features in a practical scenario.

> **Note:** This project was built and tested with [CoreGo v0.1.0](https://github.com/berkkaradalan/CoreGo/releases/tag/v0.1.0)

## Features ✨

- 💬 **Real-time messaging** - WebSocket connections for instant message delivery
- 🏠 **Room-based chat** - Create rooms and chat with multiple users
- 🔐 **JWT Authentication** - Secure endpoints using CoreGo's auth system
- 💾 **MongoDB integration** - Persistent storage using CoreGo's database wrapper

## Tech Stack

- Go (Gin framework)
- MongoDB
- WebSocket (Gorilla WebSocket)
- [CoreGo](https://github.com/berkkaradalan/corego) - Custom package for auth and database operations

## WebSocket Endpoints 🔌

### Listen to room creation events 🔔
```
ws://localhost:8080/api/ws/rooms?token=YOUR_JWT_TOKEN
```


### Listen to messages in a specific room 💬
```
ws://localhost:8080/api/ws/room-messages?roomId=ROOM_ID&token=YOUR_JWT_TOKEN
```
Get real-time messages sent to a specific room.

## API Endpoints

### Authentication
- `POST /signup` - Create a new account
- `POST /login` - Login and get JWT token

### Rooms
- `POST /api/chat-room` - Create a new chat room
- `GET /api/chat-room/:id` - Get room details
- `GET /api/chat-rooms` - List all rooms

### Messages
- `POST /api/message` - Send a message to a room
- `GET /api/message` - Get messages from a room

## Setup 🚀

1. Clone the repository
2. Set up your environment variables (MongoDB connection, auth secret, etc.)
3. Run the application:
```bash
go run cmd/main.go
```

## Environment Variables

You'll need to configure:
- MongoDB connection URL
- MongoDB database name
- Auth secret for JWT
- Token expiry duration
- Server port

## Why CoreGo? 🤔

This project uses [CoreGo v0.1.0](https://github.com/berkkaradalan/CoreGo/releases/tag/v0.1.0) to handle authentication and database operations. CoreGo provides a clean interface for common backend tasks, making it easier to focus on building features rather than boilerplate code.

## Note

This is primarily a testing ground for CoreGo. If you're looking for examples of how to use CoreGo in a real application, feel free to browse the code!
