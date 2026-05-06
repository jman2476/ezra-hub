# Dev endpoints 
POST /admin/reset -> resets user table, must be in platform="dev"
GET /admin/users -> get list of all users' names, emails, phone numbers, ids

# User endpoints
POST /api/users -> create new user w/ name, phone number, email
POST /api/login -> log in w/ name & email