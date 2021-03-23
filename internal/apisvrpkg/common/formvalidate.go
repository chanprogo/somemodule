package common

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/chanprogo/somemodule/pkg/log/logging"
	"github.com/gin-gonic/gin"

	"github.com/chanprogo/somemodule/pkg/constant"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {

	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		return http.StatusInternalServerError, constant.RESPONSE_CODE_SYSTEM
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, constant.INVALID_PARAMS
	}

	return http.StatusOK, constant.RESPONSE_CODE_OK
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
