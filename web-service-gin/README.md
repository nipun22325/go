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
curl http://localhost:8080/albums/4 --header "Content-Type: application/json" --request "GET"
```
