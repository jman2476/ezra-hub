# API Endpoints

| URI | Method | Description | Details  |
| -- | -- | -- | -- |
| /admin/reset | <font color="cyan">POST</font> | Reset database instance | [`reset`](#post-adminreset) |
| /admin/users | <font color="green">GET</font> | Get all registered users | [`get users`](#get-adminusers) |
| /api/refresh | <font color="cyan">POST</font> | Create new refresh token for user | [`refresh`](#post-apirefresh) |  
| /api/users | <font color="cyan">POST</font> | Sign up new user | [`sign up`](#post-apiusers) |
| /api/login | <font color="cyan">POST</font> |  Log in to user account | [`log in`](#post-apilogin) | 
| /api/users | <font color="orange">PATCH</font> | Update user information | [`update user`](#patch-apiusers) |
| /api/users/subs | <font color="orange">PATCH</font> | Subscribe user to event types | [`subscribe`](#patch-apiuserssubs) |
| /api/events | <font color="cyan">POST</font> | Create new event | [`create event`](#post-apievents) |
| /api/events/{id} | <font color="orange">PATCH</font> | Respond to event with user's availability | [`respond to event`](#patch-apieventsid) |
| /api/events/{id} | <font color="yellow">PUT</font> | Update event details | [`update event`](#put-apieventsid) |
| /api/events? | <font color="green">GET</font> | Get events by type | [`get events`](#get-apievents) |
| /api/events/users | <font color="green">GET</font> | Get events created by user | [`get user's events`](#get-apievents-1) | 



## Admin Endpoints
---

### <font color="cyan">POST</font> /admin/reset
Deletes *users* table in the database, which cascades to delete all entries across all tables. Only available when environment variable is set: "PLATFORM=dev".
```json
// Response body
{
    "msg": string,
}

```


### <font color="cyan">POST</font> /admin/users
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


## User Endpoints
---

### <font color="cyan">POST</font> /api/users

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

### <font color="cyan">POST</font> /api/login

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

### <font color="cyan">POST</font> /api/refresh

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

### <font color="orange">PATCH</font> /api/users

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

### <font color="orange">PATCH</font> /api/users/subs

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


## Event Endpoints
---

### <font color="cyan">POST</font> /api/events

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

### <font color="orange">PATCH</font> /api/events/{id}

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

### <font color="yellow">PUT</font> /api/events/{id}

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

### <font color="green">GET</font> /api/events?

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

### <font color="green">GET</font> /api/events

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
