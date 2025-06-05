package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


type album struct{
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are 
// serialized into JSON. Without them, the JSON would use the struct’s capitalized field names.

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default() // Initialize a Gin router
    
    // With Gin, you can associate a handler with an HTTP method-and-path combination. In this way, you can
    // separately route requests sent to a single path based on the method the client is using.

    // Use the GET function to associate the GET HTTP method and /albums path with a handler function.
    router.GET("/albums", getAlbums) 
    // Associate the POST method at the /albums path with the postAlbums function.
    router.POST("/albums", postAlbums)
    // Associate the /albums/:id path with the getAlbumByID function. 
    // In Gin, the colon preceding an item in the path signifies that the item is a path parameter.
    router.GET("/albums/id/:id", getAlbumByID)
    
    router.GET("/albums/artist/:artist", getAlbumsByArtist)
    
    // Use the Run function to attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}


// getAlbums responds with the list of all albums as JSON.
// gin.Context carries request details, validates and serializes JSON, and more
// the parameter c has a pointer to a Gin Context as the type
func getAlbums(c *gin.Context) { 
    c.IndentedJSON(http.StatusOK, albums) // serialize the struct into JSON and add it to the response.
}

// The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the 
// StatusOK constant from the net/http package to indicate 200 OK.
// can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the 
// indented form is much easier to work with when debugging and the size difference is usually small.


// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album
    
    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    
    // Add the new album to the slice.
    albums = append(albums, newAlbum) 
    // Add a 201 status code to the response, along with JSON representing the album you added.
    c.IndentedJSON(http.StatusCreated, newAlbum)    
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context){
    // Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a 
    // placeholder for the parameter in the path.
    id := c.Param("id") 
    
    // Loop over the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbumsByArtist(c *gin.Context){
    artist := c.Param("artist")
    
    var filteredAlbums []album
    
    for _, album := range albums {
        if album.Artist == artist {
            filteredAlbums = append(filteredAlbums, album)
        }
    }   
    
    if len(filteredAlbums) > 0 {
        c.IndentedJSON(http.StatusOK, filteredAlbums)    
    } else {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no albums by the artist in the collection"})
    }
}