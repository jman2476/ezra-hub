# What do I need?

I want active users to be able to create events that will be sent to other users, and those users can rsvp to the events which will send the update to all users who subscribe to that type of event. All events and updates will be also sent to the server, which will log events and updates in the db, so when a user logs back in they will be sent a list of all active events of the type they subscribe to

# What technology do I need?
Go -> client, server logic
RabbitMQ -> queue management
PostgreSQL -> database
Docker -> client/server containerization for distribution
BubbleTea -> client-side interface
Argon2id -> client side authentication
JWT -> client authentication

# What functionality do I need?
See notebook

# What do I need to set up for it to run?
## Roadmap
1. Setup server with connection to database, users table
    - create psql database  [X]
    - design table schema   [X]
    - design endpoints  [X]
    - 
2. Setup client to basic sign in, no password
    - implement bubbletea, if reasonable [maybe later]
    - make http requests to [X]
        - create user [X]
        - log into user account [X]

3. Add authentication internal package
    - JWT, refresh tokens, password hashing
    - add refresh token table to db
4. Migrate DB to take password, add password login to server
    - first implement JWT and password hashing
    - add refresh tokens after that's working
    - add refresh tokens table for tracking
5. Add client login, verify
    - find a way to safely store information on client side
    - don't want JWT/refresh token in a bad place
6. Setup RabbitMQ to broadcast user logins
    - Temporary
    - Primarily to start RabbitMQ integration
7. Create outgoing_messages table
    - all messages get stored here as gob-data
    - has field 'sent: boolean'
    - once a message is sent to RabbitMQ, field is updated
8. Create logging queue
    - save all server events to disk in files seperated by date
    - date will be based on server's local time, new day creates new file
    
9. Add needs/events table to DB
    - each event needs date, creator, benefactor, and type
    - possibly create a type table for easy lookup between user and event table
    - update users to have a field where they subscribe to different event types

10. Add event creation endpoint to Server
    - Broadcast to all clients on event creation
    - don't worry about what users are subscribed to yet
11. Add client-side event creation
    - add subscription by event type
    - add to SERVER to only broadcast on event type
12. Add event response/update endpoint to Server
    - Broadcast updates to all clients that are subscribed
13. Add event response/update to Client
    - Users should be able to interact with other's events
    - if only one person required, when someone responds it will still be visible but shown that it's taken
14. Setup client login refresh
    - on log in, client requests active events from server
    - server checks db for active events of each type, and sends to user
    - once user has received list, subscribes to queues of each feed to maintain accuracy
    - can refresh at any point, does not disturb queues