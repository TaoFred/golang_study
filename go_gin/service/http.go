package service

import (
	"fmt"
	"go_gin/config"
	"go_gin/ctrl"
	"go_gin/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *service) http(c *config.InplantConfig) {
	err := http.ListenAndServe(c.APIPort, initWebAPI())
	if err != nil {
		fmt.Println("listen webapi port error! err: ", err.Error())
	}
}

var router InPlantIntegrityRouter

type InPlantIntegrityRouter struct {
	GroupV1 string

	v1            *gin.RouterGroup
	anyonePathSet map[string]struct{}
	pathRoleMap   map[string]string
}

func (r InPlantIntegrityRouter) Init(group string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.v1 = engine.Group(group)
	router.GroupV1 = group
	router.anyonePathSet = make(map[string]struct{})
	router.pathRoleMap = make(map[string]string)
	return engine
}

func (r InPlantIntegrityRouter) POST(authRole, path string, handler gin.HandlerFunc) {
	r.handle("POST", authRole, path, handler)
}

func (r InPlantIntegrityRouter) DELETE(authRole, path string, handler gin.HandlerFunc) {
	r.handle("DELETE", authRole, path, handler)
}
func (r InPlantIntegrityRouter) PUT(authRole, path string, handler gin.HandlerFunc) {
	r.handle("PUT", authRole, path, handler)
}

func (r InPlantIntegrityRouter) GET(authRole, path string, handler gin.HandlerFunc) {
	r.handle("GET", authRole, path, handler)
}

func (r InPlantIntegrityRouter) handle(method, authRole, path string, handler gin.HandlerFunc) {
	switch method {

	case "POST":
		r.v1.POST(path, handler)
	case "DELETE":
		r.v1.DELETE(path, handler)
	case "PUT":
		r.v1.PUT(path, handler)
	case "GET":
		r.v1.GET(path, handler)
	default:
		return
	}

	pathKey := r.pathKey(router.GroupV1+path, method)
	if authRole == entity.MODULE_ANYONE {
		r.anyonePathSet[pathKey] = struct{}{}
	} else {
		r.pathRoleMap[pathKey] = authRole
	}
}

func (r InPlantIntegrityRouter) pathKey(path, method string) string {
	return path + ":" + method
}

func (r InPlantIntegrityRouter) IsAnyonePath(path, method string) bool {
	pathKey := r.pathKey(path, method)
	_, ok := r.anyonePathSet[pathKey]
	return ok
}

func (r InPlantIntegrityRouter) GetRole(path, method string) (string, bool) {
	pathKey := r.pathKey(path, method)
	v, ok := r.pathRoleMap[pathKey]
	return v, ok
}

func initWebAPI() http.Handler {
	engine := router.Init("v1/vxintegrity")

	// 首页
	router.GET(entity.MODULE_HOME, "/assetnums", ctrl.SearchAssetNumsAPI)
	router.GET(entity.MODULE_ANYONE, "/scalenums", ctrl.SearchAssetNumsAPI)
	router.GET(entity.MODULE_HOME, "/cfglogsnums", ctrl.SearchCfgLogsAPI)
	router.GET(entity.MODULE_HOME, "/defects", ctrl.SearchCfgDefectsAPI)
	router.GET(entity.MODULE_HOME, "/backuplist", ctrl.SearchBackupListAPI)

	// 系统
	router.POST(entity.MODULE_CONFIG, "/system", ctrl.AddSystem)
	router.DELETE(entity.MODULE_CONFIG, "/system", ctrl.DeleteSystem)
	router.PUT(entity.MODULE_CONFIG, "/system", ctrl.ModifySystem)
	router.GET(entity.MODULE_CONFIG, "/system", ctrl.GetSystem)
	return engine
}
