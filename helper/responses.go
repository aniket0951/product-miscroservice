package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	CurrentIdx  int   `json:"current_idx"`
	PreviousIdx int   `json:"previous_idx"`
	TotalCount  int64 `json:"total_count"`
}

type EmptyObj struct{}

func PaginationData() Pagination {
	return Pagination{
		CurrentIdx:  CURRENT_IDX,
		PreviousIdx: PREVIOUS_IDX,
		TotalCount:  TOTALCOUNT,
	}
}

func ResponseBuilder(apiStatus bool, msg string, err string, data interface{}, dataName string, isPagination bool) map[string]interface{} {
	response := map[string]interface{}{}

	response["status"] = apiStatus
	response["message"] = msg
	response["error"] = err
	response[dataName] = data
	if isPagination {
		var paginationData = PaginationData()

		response["pagination"] = paginationData
	}

	return response
}

func BuildSuccessResponse(msg string, data interface{}, dataName string) map[string]interface{} {
	response := ResponseBuilder(true, msg, "", data, dataName, false)
	return response
}

func BuildFailedResponse(msg string, err string, obj interface{}, dataName string) map[string]interface{} {
	response := ResponseBuilder(false, msg, err, obj, dataName, false)
	return response
}

func BuildResponseWithPagination(apiStatus bool, msg string, err string, obj interface{}, dataName string) map[string]interface{} {
	response := ResponseBuilder(apiStatus, msg, err, obj, dataName, true)
	return response
}

func RequestBodyEmptyResponse(ctx *gin.Context) {
	response := BuildFailedResponse(FAILED_PROCESS, REQUIRED_PARAMS, EmptyObj{}, DATA)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func BuildPaginationFailedResponse() map[string]interface{} {

	response := ResponseBuilder(false, FAILED_PROCESS, PAGINATION_INVALID, EmptyObj{}, DATA, false)
	return response

}

func BuildUnProcessableEntity(ctx *gin.Context, err error) {
	response := BuildFailedResponse(FAILED_PROCESS, err.Error(), EmptyObj{}, DATA)
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
}

func CheckError(err error, ctx *gin.Context) bool {
	if err != nil {
		BuildUnProcessableEntity(ctx, err)
		return true
	}

	return false
}

func CheckRequestParamEmpty(param string, ctx *gin.Context) bool {
	if param == "" || len(param) <= 0 {
		RequestBodyEmptyResponse(ctx)
		return true
	}

	return false
}

func GetRequestQueryParam(paramType string, ctx *gin.Context) string {
	return ctx.Request.URL.Query().Get(paramType)
}
