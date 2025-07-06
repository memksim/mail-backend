mail-backend is a simple server-side backend for the [Mail](https://github.com/Gram-git/Mail) mobile app.
The server is written in Go, uses the Gin framework, and stores data in a SQLite database.

#### Stack
- Go 1.24+
- Go standard library
- Gin
- SQLite3

#### Project Structure
```
mail-backend/
├── config/      # Configuration and current user
├── handler/     # HTTP handlers
├── model/       # Data models
├── queries/     # SQL scripts for table creation
```

#### Main API Endpoints
| Method | Endpoint             | Description                      |
|--------|----------------------|----------------------------------|
| GET    | /receivedMails       | Get received (inbox) emails      |
| GET    | /sentMails           | Get sent emails                  |
| GET    | /user?email=...      | Get user by email                |
| POST   | /user                | Register a new user              |
| POST   | /mail                | Send an email                    |
| PATCH  | /bookmarkMail        | Change bookmark status           |
| PATCH  | /readMail            | Mark email as read               |

#### Example: Sending an Email
```
curl -X POST http://localhost:8080/mail \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "sender_email": "alice@example.com",
    "recipient_email": "bob@example.com",
    "title": "Hello!",
    "body": "How are you?",
    "is_bookmark": false,
    "is_read": false,
    "time": 1710000000
}'

# Response
200 OK
{
  "id": 1,
  "sender_email": "alice@example.com",
  "recipient_email": "bob@example.com",
  "title": "Hello!",
  "body": "How are you?",
  "is_bookmark": false,
  "is_read": false,
  "time": 1710000000
}
```