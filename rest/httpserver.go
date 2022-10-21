package rest

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	configuration "xqledger/apirouter/configuration"
	utils "xqledger/apirouter/utils"
)

const componentMsg = "REST Server"

var config = configuration.GlobalConfiguration

/*
item in URL (separated by /): c.Param("")
item as param (separated by ? and then &): c.Query("")
signature := c.Request.Header.Get("X-IDPal-Signature")
*/
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowWildcard:    true,
		AllowFiles:       true,
		AllowMethods:     []string{"OPTIONS, POST, GET, PUT, DELETE, PATCH"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Access-Control-Allow-Origin, Portal-Hash"},
		ExposeHeaders:    []string{"Content-Length, Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routerGroup := router.Group(config.Rest.Path)
	{
		// Admin app
		routerGroup.GET("/tenants", getTenants)
		routerGroup.GET("/activesessions", GetOpenSessions)
		routerGroup.POST("/database", CreateNewDatabase)
		routerGroup.GET("/recordcount", GetCountInCollection)
		// Return list of dbs, collections and quota for a specific account
		routerGroup.GET("/admin/info", GetAdminInfo)

		// Manage sessions
		routerGroup.POST("/session", HandleNewSession)
		routerGroup.DELETE("/session", HandleEndSession)

		// Batch Operation Endpoints
		routerGroup.POST("/batch", HandlePostPutBatch)
		routerGroup.PUT("/batch", HandlePostPutBatch)
		// routerGroup.DELETE("/batch", DeleteBatch)

		// Single Record Operation Endpoints
		routerGroup.GET("/record/:id", GetRecordById)
		routerGroup.POST("/record", HandlePostPutRecord)
		routerGroup.PUT("/record", HandlePostPutRecord)
		routerGroup.DELETE("/record/:id", DeleteRecordById)

		// History of record
		routerGroup.GET("/recordhistory", GetRecordHistory)
		routerGroup.GET("/recordevent", GetContentInCommit)
		routerGroup.GET("/recorddiff", GetDiffTwoCommitsInFile)

		// Query-based Operation Endpoints
		// routerGroup.POST("/query/:dbname", GetRecordsByQuery)

		// Keep Alive
		routerGroup.GET("/keepalive", KeepAlive)
	}
	return router
}

func StartHTTPServer() {
	fmt.Println("PORT: " + strconv.Itoa(config.Rest.Port))
	fmt.Println("MODE: " + config.Rest.Mode)

	methodMessage := "startHTTPServer"
	utils.PrintLogInfo(componentMsg, methodMessage, "Starting REST server at port "+strconv.Itoa(config.Rest.Port))
	gin.SetMode(config.Rest.Mode)
	var buffer bytes.Buffer
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(config.Rest.Port))
	router := setupRouter()
	router.Run(buffer.String())
}