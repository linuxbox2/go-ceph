package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Info struct
type Info struct {
	Items []struct {
		ClusterID    string `json:"cluster_id"`
	} `json:"dummy"`
}

// GetInfo fetch an array of info elements (e.g., the cluster fsid)
func (api *API) GetInfo(ctx context.Context, info Info) (Info, error) {
	body, err := api.call(ctx, http.MethodGet, "/info", valueToURLParams(info))
	if err != nil {
		return Info{}, err
	}
	u := Info{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return Info{}, fmt.Errorf("%s. %s. %w", unmarshalError, string(body), err)
	}

	return u, nil
}
