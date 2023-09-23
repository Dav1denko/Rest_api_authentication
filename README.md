# Rest api authentication

Create .env file in Rest_api_authentication folder and add
```
MONGODB_USERNAME=example
MONGODB_PASSWORD=example
SIGNING_KEY=example
```

The first route issues a pair of Access, Refresh tokens for a user with an identifier (GUID)
  
- http://localhost:8000/auth/{GUID}

The second route performs a Refresh operation on a pair of Access, Refresh tokens
- http://localhost:8000/auth/{GUID}/refresh

Access token

- JWT
- SHA512
- Don't store in database

Refresh token

- Base64
- Bcrypt hash
- Store in MongoDB

Dependencies

- Gin
- Viper
- Mongo driver
- Godotenv
- Jwt