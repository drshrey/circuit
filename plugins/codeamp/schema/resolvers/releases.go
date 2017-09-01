package codeamp_schema_resolvers

import (
	"context"

	"github.com/codeamp/circuit/plugins/codeamp/models"
	"github.com/codeamp/circuit/plugins/codeamp/utils"
)

func (r *Resolver) Releases(ctx context.Context) ([]*ReleaseResolver, error) {
	if _, err := utils.CheckAuth(ctx, []string{}); err != nil {
		return nil, err
	}

	var rows []codeamp_models.Release
	var results []*ReleaseResolver

	r.DB.Find(&rows)
	for _, release := range rows {
		results = append(results, &ReleaseResolver{DB: r.DB, Release: release})
	}

	return results, nil
}
