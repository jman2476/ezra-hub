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
| /api/users | PATCH | Update user information | [`update user`](#patch-apiusers) |
| /api/users/subs | PATCH | Subscribe user to event types | [`subscribe`](#patch-apiuserssubs) |
| /api/events | POST | Create new event | [`create event`](#post-apievents) |
| /api/events/{id} | PATCH | Respond to event with user's availability | [`respond to event`](#patch-apieventsid) |
| /api/events/{id} | PUT | Update event details | [`update event`](#put-apieventsid) |
| /api/events? | GET | Get events by type | [`get events`](#get-apievents) |
| /api/events/users | GET | Get events created by user | [`get user's events`](#get-apievents-1) | 



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

```json
// Request body
{
    "name": string,
    "phone_number": string, // must be valid phone number format
    "email": string, //must be valid email format
    "password": string,
    "address": string
}
```
```json
// Response body
{
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string ,
        "phone_number": string,
        "email": string,
        "jwt": empty string,
        "refresh_token": empty string,
        "subs": null,
        "address": string 
}
```

#### POST /api/login

```json
// Request body
{
    "email": string,
    "password": string
}
```
```json
// Response body
{
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string ,
        "phone_number": string,
        "email": string,
        "jwt": string,
        "refresh_token": string,
        "subs": []string,
        "address": string 
}
```

#### POST /api/refresh

```json
// Request header
Authorization : "Bearer [user refresh token]"
// No request body
```
```json
// Response body
{
    "jwt_token": string
}
```

#### PATCH /api/users

```json
// Request header
Authorization : "Bearer [user authorization token]"

// Request body
{
    "name": string,
    "email": string, // must be valid email format
    "phone_number": string, //must be valid phone number format
    "address": string
}
```
```json
// Response body
{
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string ,
        "phone_number": string,
        "email": string,
        "jwt": string,
        "refresh_token": string,
        "subs": []string,
        "address": string 
}
```

#### PATCH /api/users/subs

```json
// Request header
Authorization : "Bearer [user authorization token]"

// Request body
{
    "subscriptions": {
        // 'subscription type': int
        // subscription type options:
        //   ride, shopping, check-in, meal, gathering, other
        // int value: 1 for subscribe, other values ignored 
        // e.g. "ride": 1
    }
}
```
```json
// Response body
[
    // list of subscription type, only returning new subscriptions
    // same options as above
]
```


### Event Endpoints
---

#### POST /api/events

```json
// Request header
Authorization : "Bearer [user authorization token]"

// Request body
{
    "name": string,
    "category": string,
    "occurs_on": timestamp,
    "expires_at": timestamp,
    "min_volunteer": int32,
    "max_volunteer": int32,
    "description": string,
    "location": string
}
```
```json
// Response body
{
    "id": UUID string,
    "created_at": timestamp,
    "updated_at": timestamp,
    "name": string,
    "owner_id": UUID string,
    "category": Genre, // see app/server/event_models for definition
    "occurs_on": timestamp,
    "expires_at": timestamp,
    "min_volunteer": {
        "Int32": int32,
        "Valid": boolean
    },
    "max_volunteer": {
        "Int32": int32,
        "Valid": boolean
    },
    "description": string,
    "respondants": []UUID string,
    "location": string,
    "creator": string
}
```

#### PATCH /api/events/{id}

```json
// Request header
Authorization : "Bearer [user authorization token]"

// Request body
{
    "available": boolean
}
```
```json
// Response body: No content
// Status 204
```

#### PUT /api/events/{id}

```json
// Request header
Authorization : "Bearer [user authorization token]"

// Request body
{
    "name": string,
    "category": string,
    "occurs_on": timestamp,
    "expires_at": timestamp,
    "min_volunteer": int32,
    "max_volunteer": int32,
    "description": string,
    "old_type": string,
    "location": string
}
```
```json
// Response body
{
    "id": UUID string,
    "created_at": timestamp,
    "updated_at": timestamp,
    "name": string,
    "owner_id": UUID string,
    "category": Genre, // see app/server/event_models for definition
    "occurs_on": timestamp,
    "expires_at": timestamp,
    "min_volunteer": {
        "Int32": int32,
        "Valid": boolean
    },
    "max_volunteer": {
        "Int32": int32,
        "Valid": boolean
    },
    "description": string,
    "respondants": []UUID string,
    "location": string
}
```

#### GET /api/events?

```json
// Request header
Authorization : "Bearer [user authorization token]"

// No request body
// Query parameters: "type"
// Use additional query parameter for each category desired
// e.g. /api/events?type=meal&type=gathering will return all 
//  gathering and meal type events


```
```json
// Response body
[
    {
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string,
        "owner_id": UUID string,
        "category": Genre, // see app/server/event_models for definition
        "occurs_on": timestamp,
        "expires_at": timestamp,
        "min_volunteer": {
            "Int32": int32,
            "Valid": boolean
        },
        "max_volunteer": {
            "Int32": int32,
            "Valid": boolean
        },
        "description": string,
        "respondants": []UUID string,
        "location": string
    },
]
```

#### GET /api/events

```json
// Request header
Authorization : "Bearer [user authorization token]"

// No request body
```
```json
// Response body
[
    {
        "id": UUID string,
        "created_at": timestamp,
        "updated_at": timestamp,
        "name": string,
        "owner_id": UUID string,
        "category": Genre, // see app/server/event_models for definition
        "occurs_on": timestamp,
        "expires_at": timestamp,
        "min_volunteer": {
            "Int32": int32,
            "Valid": boolean
        },
        "max_volunteer": {
            "Int32": int32,
            "Valid": boolean
        },
        "description": string,
        "respondants": []UUID string,
        "location": string
    },
]
```

## Licensing
Ezra Hub is free to use for communities under 50 persons under the AGPLv3 license, found in this repository. For larger communities or commercial applications, please contact [Jeremy McKeegan](https://github.com/jman2476) to discuss an enterprise contract with setup, maintenance and support included.