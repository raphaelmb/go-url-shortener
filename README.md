# URL shortener
This is a URL shortener API.

## Usage

Run `docker compose up -d` to start the API and database.

### Shorten URL

POST `http://localhost:8080/api/url/shorten`

```bash
curl -X POST \
-d '{"url": "https://www.google.com"}' \
http://localhost:8080/api/url/shorten
```

Payload example:
```json
{
    "url": "https://www.google.com"
}
``` 

Response example:
```json
{
    "data": {
        "code": "{code}"
    }
        
}
```

### Get shortened URL

GET `http://localhost:8080/api/url/{code}`

```bash
curl http://localhost:8080/api/url/{code}
```

Response example:
```json
{
    "data": {
        "full_url": "https://www.google.com"
    }
        
}
```