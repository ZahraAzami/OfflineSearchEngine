package Server

import (
	"OfflineSearchEngine/Internals/SearchEngines/Interfaces"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ginEngine    gin.Engine
	SearchEngine Interfaces.ISearchEngine
}

func NewServer() {

}
