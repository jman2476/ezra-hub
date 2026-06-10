# Ezra Hub
Ezra Hub gathers all your community organizing tools into one convenient app.

## Setup

### Rabbitmq

### Database

### Server

### Client

## Endpoints

| URI | Method | Description | Parameters  |
| -- | -- | -- | -- |
| /admin/reset | POST | Reset database instance | - |
| /admin/users | GET | Get all registered users | - |
| /api/refresh | POST | Create new refresh token for user | Auth Header [^1] |  
| /api/users | POST | Sign up new user | Body: NewUser [^2] |
| /api/login | POST |  - | - | 
| /api/users | PATCH | - | - |
| /api/users/subs | PATCH | - | - |
| /api/events | POST | - | - |
| /api/events/{id} | PATCH |  - | - |
| /api/events/{id} | PUT | - | - |
| /api/events | GET | - | - |
| /api/events/users | GET | - | - | 

[^1]: Authorization headers are formatted as: "Bearer {user_id}"

[^2]: Request body structs are defined on the client side

```json
// Request header
Authorization : "Bearer [user * token]"

// Request body
{

}
```
```json
// Response body

```

### Admin Endpoints
---

#### POST /admin/reset
Deletes *users* table in the database, which cascades to delete all entries across all tables. Only available when environment variable is set: "PLATFORM=dev".
```json
// Response body
{
    "msg": string,
}

```


#### POST /admin/users
Returns array of all registered users in the database and their data, excluding current JWT and refresh tokens.

```json
// Response body
[
    {
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string ,
        "phone_number": string,
        "email": string,
        "jwt": empty string,
        "refresh_token": empty string,
        "subs": []string,
        "address": string 
    },
]

```


### User Endpoints
---

#### POST /api/users


#### POST /api/login


#### POST /api/users


#### POST /api/users


#### POST /api/users


#### POST /api/users


### Event Endpoints
---

#### POST /api/events


#### POST /api/events


#### POST /api/events


#### POST /api/events


#### POST /api/events


## Licensing
Ezra Hub is free to use for communities under 50 persons under the AGPLv3 license, found in this repository. For larger communities or commercial applications, please contact [Jeremy McKeegan](https://github.com/jman2476) to discuss an enterprise contract with setup, maintenance and support included.