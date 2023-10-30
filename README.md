# Real-Time Chat App Server (API + WebSocket)

This project serves as the backend for a real-time chat application with multiple rooms. Users can create rooms and join existing ones. Below are the instructions for running the app:

## Prerequisites

- **Go (Golang)**: Ensure you have Go version 1.21.3 or higher installed.
## Build:
    1. Open a terminal or command prompt.
    2. Navigate to the root directory of your project (where `main.go` is located).
    3. Run the following command to build the server binary:
        ```bash
        go build main.go
        ```
    4. This will create an executable binary (e.g., `main`) in the same directory.

## Running the Server

1. After building, run the server using:
    ```bash
    ./main
    ```
2. The server will start listening on a specified port (e.g., `localhost:8080`).

## API Methods

### 1. Sign Up (POST - `/v1/api/auth/signup`)

Request:
```json
{
    "user_name": "user",
    "email": "email@mail.com",
    "password": "pass"
}
```

Response:
```json
{
    "user": {
        "ID": 2,
        "CreatedAt": "2023-10-30T15:47:33.246704222Z",
        "UpdatedAt": "2023-10-30T15:47:33.246704222Z",
        "DeletedAt": null,
        "UserName": "user",
        "Email": "email@mail.com"
    },
    "token": "<your_access_token>"
}
```

### 2. Login (POST - `/v1/api/auth/login`)

Request:
```json
{
    "email": "email@mail.com",
    "password": "pass"
}
```

Response:
```json
{
    "user": {
        "ID": 2,
        "CreatedAt": "2023-10-30T15:47:33.246704Z",
        "UpdatedAt": "2023-10-30T15:47:33.246704Z",
        "DeletedAt": null,
        "UserName": "user",
        "Email": "email@mail.com"
    },
    "token": "<your_access_token>"
}
```

### 3. Create Chat Room (POST - `/v1/api/chat/create`)

Request:
```json
{
    "name": "Chat name"
}
```

Response:
```json
{
    "chatRoom": {
        "ID": 1,
        "CreatedAt": "2023-10-30T15:51:50.149489876Z",
        "UpdatedAt": "2023-10-30T15:51:50.149489876Z",
        "DeletedAt": null,
        "name": "Chat name"
    }
}
```

### 4. Get Chat Rooms (GET - `/v1/api/chat/rooms`)

Response:
```json
{
    "chatRooms": [
        {
            "ID": 1,
            "CreatedAt": "2023-10-30T15:51:50.149489Z",
            "UpdatedAt": "2023-10-30T15:51:50.149489Z",
            "DeletedAt": null,
            "name": "<room_name>"
        }
    ]
}
```

## WebSocket Endpoint

### `/ws`

- **Parameters**:
    - `jwt`: Token for authentication.
    - `room_id`: Integer representing the chat room ID.

- **Body**:
    ```json
    {
        "chatRoomId": 1,
        "chatMessage": "Message"
    }
    ```
