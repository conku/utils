package utils

import (
	"net/http"

	"github.com/conku/gorm"
	"github.com/conku/qor-example/config/db"
	"github.com/conku/qor/utils"
)

// GetDB get DB from request
func GetDB(req *http.Request) *gorm.DB {
	if db := utils.GetDBFromRequest(req); db != nil {
		return db
	}
	return db.DB
}
