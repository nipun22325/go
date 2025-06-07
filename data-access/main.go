package main

import (
    "fmt"
    "log"
    "database/sql"
    "os"
    "github.com/go-sql-driver/mysql"
)

// Declare a db variable of type *sql.DB
// This is a pointer to an sql.DB struct, which represents access to 
// a specific database, also called the database handle.
var db *sql.DB 
// Note: Making db a global variable simplifies this example. In 
// production, you’d avoid the global variable, such as by passing 
// variable to functions that need it or by wrapping it in a struct.

// Use this to hold row data returned from the query.
type Album struct {
    ID int64
    Title string
    Artist string
    Price float32
}

func main(){
    // Capture connection properties.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "recordings"
    
    // Get a database handle.
    var err error
 
    // Collect connection properties and format them into a DSN for a 
    // connection string.
    // The Config struct makes for code that’s easier to read than a
    // connection string would be.
    db, err = sql.Open("mysql", cfg.FormatDSN())
 
    if err != nil {
        log.Fatal(err)
    }
    
    // Call DB.Ping to confirm that connecting to the database works. 
    // At run time, sql.Open might not immediately connect, depending 
    // on the driver. You’re using Ping here to confirm that the 
    // database/sql package can connect when it needs to.
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    
    // Executing query
    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)
    
    // Executing single row query
    alb, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found: %v\n", alb)
    
    // Adding new album
    albID, err := addAlbum(Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99,})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)
}

// Query the database for multiple rows
// albumsByArtist function queries for albums that have the specified 
// artist name.
func albumsByArtist(name string) ([]Album, error) {
    // albums slice to hold data from returned rows
    var albums []Album
    
    // Query to return rows where artist = name.
    // Use parameterized queries to prevent SQL injection.
    // SQL and data are sent separately, so user input is treated strictly 
    // as data.
    // Avoids risks from string concatenation like fmt.Sprintf which can  
    // be exploited.
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    
    //Defer closing rows so that any resources it holds will be released   
    //when the function exits.
    // A defer statement defers the execution of a function until the 
    // surrounding function returns.
    // The deferred call's arguments are evaluated immediately, but the 
    // function call is not executed until the surrounding function returns.
    defer rows.Close()
    
    // Loop through rows, using Scan to assign column data to struct fields
    for rows.Next() {
        var alb Album 
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    
    // After the loop, check for an error from the overall query, using 
    // rows.Err. Note that if the query itself fails, checking for an error here 
    // is the only way to find out that the results are incomplete.
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from returned row
    var alb Album
    
    // Use DB.QueryRow to execute a SELECT statement to query for an album with
    // the specified ID.
    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    
    // To simplify the calling code (your code!), QueryRow doesn’t return an 
    // error. Instead, it arranges to return any query error (such as 
    // sql.ErrNoRows) from Rows.Scan later.
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil{
        if err == sql.ErrNoRows{
            // sql.ErrNoRows indicates that the query returned no rows
            return alb, fmt.Errorf("albumsById %d : no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d : %v", id, err)
    }
    return alb, nil
}

// addAlbum adds the specified album to the database
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}


