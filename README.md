# Ezra Hub
Ezra Hub gathers all your community organizing tools into one convenient app.

## Setup

### Rabbitmq

### Database

### Server

### Client

## Endpoints

| URI | Method | Description | Parameters | Notes |
| -- | -- | -- | -- | -- |
| /admin/reset | POST | Reset database instance | - | only available when environment variable PLATFORM="dev" |
| /admin/users | GET | Get all registered users | - | - |
| /api/refresh | POST | Create new refresh token for user | [^1] Auth Header | - | 
| /api/users | POST | Sign up new user | **Body:  | - |
| /api/login | POST | - | - | - |
| /api/users | PATCH | - | - | - |
| /api/users/subs | PATCH | - | - | - |
| /api/events | POST | - | - | - |
| /api/events/{id} | PATCH | - | - | - |
| /api/events/{id} | PUT | - | - | - |
| /api/events | GET | - | - | - |
| /api/events/users | GET | - | - | - | 

[^1]: Authorization headers are formatted as: "Bearer {user_id}"


## Licensing
Ezra Hub is free to use for communities under 50 persons under the AGPLv3 license, found in this repository. For larger communities or commercial applications, please contact [Jeremy McKeegan](https://github.com/jman2476) to discuss an enterprise contract with setup, maintenance and support included.