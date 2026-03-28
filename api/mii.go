package api

import (
	"net/http"
	"wwfc/database"
)

type MiiRequest struct {
	Secret    string `json:"secret"`
	ProfileID uint32 `json:"pid"`
}

type MiiResponse struct {
	Mii     string
	Success bool
	Error   string
}

var MiiRoute = MakeRouteSpec[MiiRequest, MiiResponse](
	true,
	"/api/mii",
	HandleMii,
	http.MethodPost,
)

func HandleMii(req any, v bool, r *http.Request) (any, int, error) {
	_req := req.(MiiRequest)
	res := MiiResponse{}

	if _req.ProfileID == 0 {
		return nil, http.StatusBadRequest, ErrPIDMissing
	}

	res.Mii = database.GetMKWFriendInfo(pool, ctx, _req.ProfileID)

	if res.Mii == "" {
		return nil, http.StatusInternalServerError, ErrUserQuery
	}

	return res, http.StatusOK, nil
}
