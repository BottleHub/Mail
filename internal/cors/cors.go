package cors

import "github.com/gin-gonic/gin"

func EnableCors(w *gin.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
