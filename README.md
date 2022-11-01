# Golang MyGram

Aplikasi API Instagram sederhana menggunakan Golang

## Database Configuration
1.  PostgreSQL
2.  Database Name: mygram
3.  Database User: postgres
4.  Database Password: root
5.  Database Port: 5432

[!] You can change the database configuration in the file models/setup.go

## How to run
```bash
go mod tidy
go run main.go
```

This will runing on default localhost port 8080

## Postman Collection
You can import this postman collection to test the API

[!] Collection file in folder test

[!] Remember after login, you need to copy the token and paste it in the header Authorization with Bearer Token manually

## Fitur
- [x] Register User
- [x] Login User
- [x] Update User
- [x] Delete User

- [x] Create Photo Post
- [x] Update Photo Post
- [x] Delete Photo Post

- [x] Create Comment
- [x] Update Comment
- [x] Delete Comment

- [x] Create User Social Media
- [x] Update User Social Media
- [x] Delete User Social Media

## Get data API
### User
Register User

*Note : username and email is unique
```bash
curl --location --request POST 'http://localhost:8080/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "age": 20,
    "email": "yourmail@mail.test",
    "username": "yourname",
    "password": "yourpassword"
}'
```

Login User
```bash
curl --location --request POST 'http://localhost:8080/users/login' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'Email=yourMail' \
--data-urlencode 'Password=yourPassword'
```
This will return token, save it for next request


Update User
```bash
curl --location --request PUT 'http://localhost:8080/users/update' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "updateyourmail@test.com",
    "username": "updateyourname"
}'
```

Delete User
```bash
curl --location --request DELETE 'http://localhost:8080/users/' \
--header 'Authorization: Bearer {YOURTOKEN}'
```
Token is required

### Photo Post

Get All Your Photo Post
```bash
curl --location --request GET 'http://localhost:8080/posts/' \
--header 'Authorization: Bearer {YOURTOKEN}'
```

Create Photo Post
```bash
curl --location --request POST 'http://localhost:8080/photos' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "your title",
    "caption": "your caption",
    "photo_url": "your photo url"
}'
```

Update Photo Post
```bash
curl --location --request PUT 'http://localhost:8080/photos/{PHOTOID}' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "update your title",
    "caption": "update your caption",
    "photo_url": "update your photo url"
}'
```

Delete Photo Post
```bash
curl --location --request DELETE 'http://localhost:8080/photos/{PHOTOID}' \
--header 'Authorization: Bearer {YOURTOKEN}'
```

### Comment

Get All Your Comment
```bash
curl --location --request GET 'http://localhost:8080/comments/' \
--header 'Authorization: Bearer {YOURTOKEN}'
```

Post Comment
```bash
curl --location --request POST 'http://localhost:8080/comments' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "your message",
    "photo_id": {PHOTOID}
}'
```

Update Comment
```bash
curl --location --request PUT 'http://localhost:8080/comments/{COMMENTID}' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "update your message"
}'
```


Delete Comment
```bash
curl --location --request DELETE 'http://localhost:8080/comments/{COMMENTID}' \
--header 'Authorization: Bearer {YOURTOKEN}'
```


### User Social Media

Get All Your User Social Media
```bash
curl --location --request GET 'http://localhost:8080/socialmedias' \
--header 'Authorization: Bearer {YOURTOKEN}'
```

Create User Social Media
```bash
curl --location --request POST 'http://localhost:8080/socialmedias' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "your social media name",
    "social_media_url": "your social media url"
}'
```

Update User Social Media
```bash
curl --location --request PUT 'http://localhost:8080/socialmedias/{SOCIALMEDIAID}' \
--header 'Authorization: Bearer {YOURTOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "update your social media name",
    "social_media_url": "update your social media url"
}'
```

Delete User Social Media
```bash
curl --location --request DELETE 'http://localhost:8080/socialmedias/{SOCIALMEDIAID}' \
--header 'Authorization: Bearer {YOURTOKEN}'
```

