package main

import (
	"context"
	"creator/controllers"
	"github.com/Alena557miles/webservergo"
)

var ctx = context.Background()

func main() {
	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	galleryC := &controllers.GalleryController{}
	webservergo.StartServer(artistC, artC, galleryC)
}
