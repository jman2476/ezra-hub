# Ezra Hub
Ezra Hub gathers all your community organizing tools into one convenient app.

## Setup

### Rabbitmq

### Database

### Server

### Client

## Endpoints

| URI | Method | Description | Details  |
| -- | -- | -- | -- |
| /admin/reset | POST | Reset database instance | [`reset`](#post-adminreset) |
| /admin/users | GET | Get all registered users | [`get users`](#get-adminusers) |
| /api/refresh | POST | Create new refresh token for user | [`refresh`](#post-apirefresh) |  
| /api/users | POST | Sign up new user | [`sign up`](#post-apiusers) |
| /api/login | POST |  Log in to user account | [`log in`](#post-apilogin) | 
| /api/users | PATCH | Update user information | - |
| /api/users/subs | PATCH | Subscribe user to event types | - |
| /api/events | POST | Create new event | - |
| /api/events/{id} | PATCH | Respond to event with user's availability | - |
| /api/events/{id} | PUT | Update event details | - |
| /api/events | GET | Get events by type | - |
| /api/events/users | GET | Get events created by user | - | 


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


#### POST /api/refresh


#### PATCH /api/users


#### PATCH /api/users/subs


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