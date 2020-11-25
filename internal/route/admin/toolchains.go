// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"gogs.io/gogs/internal/conf"
	"gogs.io/gogs/internal/context"
	"gogs.io/gogs/internal/db"
	"gogs.io/gogs/internal/route"
)

const (
	TOOLCHAIN = "admin/toolchains/list"
)

func ToolChains(c *context.Context) {
	c.Data["Title"] = c.Tr("admin.toolchains")
	c.Data["PageIsToolChains"] = true

	route.RenderUserSearch(c, &route.UserSearchOptions{
		Type:     db.UserOrganization,
		Counter:  db.CountOrganizations,
		Ranger:   db.Organizations,
		PageSize: conf.UI.Admin.OrgPagingNum,
		OrderBy:  "id ASC",
		TplName:  TOOLCHAIN,
	})
}
