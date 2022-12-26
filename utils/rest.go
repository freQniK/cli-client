package utils

import (
	"encoding/json"
	"net/http"

	clitypes "github.com/sentinel-official/cli-client/types"
)

func write(w http.ResponseWriter, status int, resp *clitypes.RestResponse) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(resp)
}

func WriteErrorToResponse(w http.ResponseWriter, status int, resp *clitypes.RestError) {
	_ = write(
		w,
		status,
		clitypes.NewRestResponse(resp, nil),
	)
}

func WriteResultToResponse(w http.ResponseWriter, status int, resp interface{}) {
	_ = write(
		w,
		status,
		clitypes.NewRestResponse(nil, resp),
	)
}