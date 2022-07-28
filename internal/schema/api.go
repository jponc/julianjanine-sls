package schema

import "github.com/jponc/julianjanine/internal/types"

type HealthcheckResponse struct {
	Message string `json:"message"`
}

type GetGuestsResponse *[]types.Guest
