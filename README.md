# EC-Nafiaufa MNC

## Login
Request:
```bash
POST /login
Content-Type: application/json

{
    "id": 1,
    "username": "user1",
    "password": "pass1"
}
```
Response:
```bash
{
    "id": 1,
    "username": "user1",
    "password": "pass1"
}
```
## Pembayaran
Request:
```bash
POST /payment
Content-Type: application/json

{
    "customer": 1,
    "amount": 1000
}
```
Response:
```bash
{
    "id": 1,
    "customer": 1,
    "timestamp": "2022-12-30T12:34:56Z",
    "amount": 1000
}
```
## Logout
Request:
```bash
POST /logout
{
    "id": 1,
    "username": "user1",
    "password": "pass1"
    
}
```

