package utils

import (
	"net/http"

	"github.com/sMARCHz/go-secretaria-finance/internal/core/errors"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var statusCodeMap = map[int]codes.Code{
	http.StatusBadRequest:          codes.InvalidArgument,
	http.StatusNotFound:            codes.NotFound,
	http.StatusUnprocessableEntity: codes.FailedPrecondition,
	http.StatusInternalServerError: codes.Internal,
}

func ConvertHttpErrToGRPC(appError *errors.AppError, logger logger.Logger) error {
	statusCode, present := statusCodeMap[appError.StatusCode]
	if !present {
		logger.Errorf("appError status code['%v'] isn't in the map", appError.StatusCode)
		statusCode = codes.Internal
	}
	return status.Error(statusCode, appError.Message)
}
