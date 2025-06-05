# Web Service with Gin

## Run the server using

```bash
go run .
```

## To run the APIs

### In another terminal, run the following commands for specific actions

#### Return all items

```bash
curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET"
```

#### Add an item

```bash
curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data "{\"id\": \"4\", \"title\": \"The Modern Sound of Betty Carter\", \"artist\": \"Betty Carter\", \"price\": 49.99}"
```

#### Return a specific item by ID

```bash
curl http://localhost:8080/albums/id/4 --header "Content-Type: application/json" --request "GET"
```

#### Return specific items by artist

```bash
curl http://localhost:8080/albums/artist/Sarah%20Vaughan --header "Content-Type: application/json" --request "GET"
```

 > **Note:** The artist's name is seperated with %20 instead of a space because %20 is the URL-encoded representation of a space character. 

