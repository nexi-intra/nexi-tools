// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Web deploy to production
---
*/
package cmds

import (
	"context"

	"github.com/365admin/nexi-tools/execution"
	"github.com/365admin/nexi-tools/utils"
)

func ProvisionWebdeployproductionPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "nexi-tools", "60-provision", "10-web.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
