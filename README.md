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
| /api/users | POST | - | - | - |
| /api/login | POST | - | - | - |
| /api/users | PATCH | - | - | - |
| /api/users/subs | PATCH | - | - | - |
| /api/events | POST | - | - | - |
| /api/events/{event_id} | PATCH | - | - | - |
| /api/events/{event_id} | PUT | - | - | - |
| /api/events | GET | - | - | - |
| /api/events/users | GET | - | - | - | 
| /api/refresh | POST | - | - | - | 



## Licensing
Ezra Hub is free to use for communities under 50 persons under the AGPLv3 license, found in this repository. For larger communities or commercial applications, please contact [Jeremy McKeegan](https://github.com/jman2476) to discuss an enterprise contract with setup, maintenance and support included.