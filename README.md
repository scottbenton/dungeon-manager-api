# Dungeon Manager API

This is the API for Dungeon Manager, an application I am working on to make it easier for game masters playing tabletop role playing games like Dungeons and Dragons to share information with their players.

## Setup

### Create a MongoDB database

I'm using a free database directly from [mongodb](www.mongodb.com/), but you can use any host you prefer.
Once you go through setup, copy your connection string before continuing.

### Environment Variables

The following environment variables must be set to run this project.
When running the project locally, you can set your environment variables into a `.env` file
The MongoDB connection string is private, but if you are developing off of `apps.scottbenton.dev`, you can use the default `JWKS_URL` value provided below.

```
DUNGEON_MANAGER_CONNECTION_STRING=*private*
JWKS_URL=https://api.auth.scottbenton.dev/auth/jwt/jwks.json
```

### Running

Run `go run .` in order to start the API
