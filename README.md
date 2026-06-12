```
░██████████                                   ░██     ░██            ░██        
░██                                           ░██     ░██            ░██        
░██         ░█████████ ░██░████  ░██████      ░██     ░██ ░██    ░██ ░████████  
░█████████       ░███  ░███           ░██     ░██████████ ░██    ░██ ░██    ░██ 
░██            ░███    ░██       ░███████     ░██     ░██ ░██    ░██ ░██    ░██ 
░██          ░███      ░██      ░██   ░██     ░██     ░██ ░██   ░███ ░███   ░██ 
░██████████ ░█████████ ░██       ░█████░██    ░██     ░██  ░█████░██ ░██░█████  
```

# Ezra Hub
Ezra Hub helps your community organize events and volunteering initiatives.

## Overview
The idea for Ezra Hub is to be a central place for community members to share their needs and volunteer to help others. Whether it's checking in on a loved one, giving someone a ride to the doctor, or just planning a gathering, Ezra Hub makes it easy to know what's going on in your community and where you can provide your assistance.

## Setup
Ezra Hub requires a RabbitMQ server and a Postgres database to run the server, and a connection to the Ezra Hub and RabbitMQ servers for the client. This guide assumes you will be running RabbitMQ from a docker container, and all other components directly on your machine.

This guide is also written from the perspective of setting up on Linux, so please report any issues or inconsistencies you experience during setup.

### Rabbitmq
Download a docker image of RabbitMQ version 4.3+, and run `rabbit.sh start` from the root of the project before starting the main server. If this is the first time, you may need to run this script twice. 

The `./rabbit.sh` file is set up assuming use of the `rabbitmq:4.3-management` image of RabbitMQ. If you are using a different version, please modify the `start_or_run` command on line 11 to reflect the version you are using.

Once the RabbitMQ server is running, you can use `rabbit.sh logs` to view the logs of the RabbitMQ server, and `rabbit.sh stop` to close the container. 

### Database
Ezra Hub uses Postgres as its database, and supports Postgres 18+. It should work with versions 16+, but anything earlier is untested. Before starting Ezra Hub for the first time, create a new database withing Postgres:
```sql
CREATE DATABASE ezra;
```
Then connect to the database with `\c ezra` create a password for the database:
```sql
ALTER USER <username> PASSWORD '<password>';
```
Note the username and password for creating your connection string for the `.env` file and [goose](https://github.com/pressly/goose) database migrations. Once you have your connection string, find the migrations file called [goose](https://github.com/jman2476/ezra-hub/blob/main/sql/schema/goose), and insert your connection string in the quotes for each command. Copy the first command into a terminal running from `./sql/schema` and run it to bring the database up to the current version.

### Server
Within the `./app/server/` folder, create a `.env` file to store your environment variables based on the `EXAMPLE.ENV` file in the root of the project. If you are running this outside of a development environment, be sure to set `PLATFORM=production` or anything other than `dev` in order to disable to database reset endpoint.

Start the server from the root of the project with `go run ./app/server`, then you will be able to start the client side. Make sure that the database and RabbitMQ server are running before starting the server, or it will not be able to make a connection.

Note: currently, Ezra Hub server is not optimized to run from a Docker container, and there is a chance it can crash from a container if not given enough memory. There are plans to improve this in the future, so if you find any bugs please report them in the issues with as much data as possible. If you can, please include the resource usage of your docker container.

### Client
Within the `./app/client/` folder, create a `.env` file for your environment variables like the server, again basing it on the appropriate section of the `EXAMPLE.ENV` file.

Once ready, start the client from the root of the project with `go run ./app/client`, and you will be brought to the opening menu of the Ezra Hub client.

## Usage
Interacting with Ezra Hub currently acts through the terminal client, with a web-based client planned for the future. Before running the client, verify that the infrastructure for the application (i.e. RabbitMQ server, database server, and Ezra Hub Server) are all set up and running, or the client will be useless.

Upon first opening the app, you will be prompted to log in or sign up, and once that's complete you will be taken to the Ezra Hub command prompt. You can type the desired command by name, enter 'help' for a description of each command, or enter 'menu' to select the desired command from a list. The options are detailed below:

| Command | Purpose | Description |
| -- | -- | -- |
| signup | Create new user account | Once new account is created on the server, the new account will be automatically logged in |
| login | Log into existing account | Also retrieves any saved events |
| logout | Log out of current account | Clears user data |
| create | Create new event | Event creation will be broadcast to all users subscribed to the feed of the new event's type |
| subscribe | Subscribe to event feeds | Choose between *ride, shopping, check-in, meal, gathering,* and *other* |
| events | Lists all events in the types the user follows | Scroll between events to get more information |
| respond | Respond to events with your availability | |
| update-user | Update current user information | All fields except password can be changed |
| update-event | Update one of the events you created | If event type is changed, subscribers of old and new type will be notified |
| menu | List all commands | Select an event to execute it |
| help | List all commands and what they do | |
| exit | Exit Ezra Hub | |

To get the most thorough experience, try opening multiple instances of the client with different accounts to see the notification system in action.

---
---
## API Endpoints

| URI | Method | Description | Details  |
| -- | -- | -- | -- |
| /admin/reset | <font color="cyan">POST</font> | Reset database instance | [`reset`](Endpoints.md#post-adminreset) |
| /admin/users | <font color="green">GET</font> | Get all registered users | [`get users`](Endpoints.md#get-adminusers) |
| /api/refresh | <font color="cyan">POST</font> | Create new refresh token for user | [`refresh`](Endpoints.md#post-apirefresh) |  
| /api/users | <font color="cyan">POST</font> | Sign up new user | [`sign up`](Endpoints.md#post-apiusers) |
| /api/login | <font color="cyan">POST</font> |  Log in to user account | [`log in`](Endpoints.md#post-apilogin) | 
| /api/users | <font color="orange">PATCH</font> | Update user information | [`update user`](Endpoints.md#patch-apiusers) |
| /api/users/subs | <font color="orange">PATCH</font> | Subscribe user to event types | [`subscribe`](Endpoints.md#patch-apiuserssubs) |
| /api/events | <font color="cyan">POST</font> | Create new event | [`create event`](Endpoints.md#post-apievents) |
| /api/events/{id} | <font color="orange">PATCH</font> | Respond to event with user's availability | [`respond to event`](Endpoints.md#patch-apieventsid) |
| /api/events/{id} | <font color="yellow">PUT</font> | Update event details | [`update event`](Endpoints.md#put-apieventsid) |
| /api/events? | <font color="green">GET</font> | Get events by type | [`get events`](Endpoints.md#get-apievents) |
| /api/events/users | <font color="green">GET</font> | Get events created by user | [`get user's events`](Endpoints.md#get-apievents-1) | 

---
---
## Contributing
If you come across any errors or have any suggestions, please report them to Issues page on the GitHub repository, and include as much detail as you can.

<!-- ## Licensing
Ezra Hub is free to use for communities under 50 persons under the AGPLv3 license, found in this repository. For larger communities or commercial applications, please contact [Jeremy McKeegan](https://github.com/jman2476) to discuss an enterprise contract with setup, maintenance and support included. -->

## Credits
Ezra Hub was created by [Jeremy McKeegan](https://github.com/jman2476) with the emotional support of his cats: Fionn, Ringo, and Hank.