package main 
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	TITLE string `json:"titl"`
	ARTIST string `json:"artist"`
	PRICE float64 `json:"price"`

}
var albums = []album{
    {ID: "1", TITLE: "Blue Train", ARTIST: "John Coltrane", PRICE: 56.99},
    {ID: "2", TITLE: "Jeru", ARTIST: "Gerry Mulligan", PRICE: 17.99},
    {ID: "3", TITLE: "Sarah Vaughan and Clifford Brown", ARTIST: "Sarah Vaughan", PRICE: 39.99},
}
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums,newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func main (){

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8082")
}