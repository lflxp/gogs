// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"strings"

	"github.com/unknwon/com"
	"gogs.io/gogs/internal/conf"
	"gogs.io/gogs/internal/context"
	"gogs.io/gogs/internal/db"
	"gogs.io/gogs/internal/email"
	"gogs.io/gogs/internal/form"
	log "unknwon.dev/clog/v2"
)

const (
	TOOLCHAIN     = "admin/toolchains/list"
	TOOLCHAIN_NEW = "admin/toolchains/new"
)

var Data []ToolChainsData

type ToolChainsData struct {
	ID       string
	Name     string
	Type     string
	Host     string
	User     string
	Password string
	Token    string
}

func ToolChains(c *context.Context) {
	Data = []ToolChainsData{
		ToolChainsData{
			ID:       "1",
			Name:     "Gitlab",
			Type:     "Git",
			Host:     "gitlab.com",
			User:     "roto",
			Password: "system",
			Token:    "123",
		},
		ToolChainsData{
			ID:       "2",
			Name:     "GitHub",
			Type:     "Git",
			Host:     "github.com",
			User:     "root",
			Password: "system",
			Token:    "123",
		},
		ToolChainsData{
			ID:       "3",
			Name:     "扫描工具",
			Type:     "Sonar",
			Host:     "sonarqube.com",
			User:     "admin",
			Password: "*",
			Token:    "DFJ19034NPAD",
		},
		ToolChainsData{
			ID:       "4",
			Name:     "Jenkins",
			Type:     "CICD",
			Host:     "jenkins.com",
			User:     "admin",
			Password: "Jenkins123",
			Token:    "*",
		},
		ToolChainsData{
			ID:       "5",
			Name:     "RabotFrameWork",
			Type:     "Test",
			Host:     "rabot.com",
			User:     "admin",
			Password: "admin",
			Token:    "*",
		},
		ToolChainsData{
			ID:       "6",
			Name:     "ArgoCD",
			Type:     "CD",
			Host:     "argocd.com",
			User:     "admin",
			Password: "admin@default",
			Token:    "*",
		},
	}

	c.Data["Title"] = c.Tr("admin.toolchains")
	c.Data["PageIsToolChains"] = true
	c.Data["Tools"] = Data
	c.Data["Total"] = 6

	// route.RenderUserSearch(c, &route.UserSearchOptions{
	// 	Type:     db.UserOrganization,
	// 	Counter:  db.CountOrganizations,
	// 	Ranger:   db.Organizations,
	// 	PageSize: conf.UI.Admin.OrgPagingNum,
	// 	OrderBy:  "id ASC",
	// 	TplName:  TOOLCHAIN,
	// })
	c.Success(TOOLCHAIN)
}

func NewToolChainsPost(c *context.Context, f form.AdminCrateUser) {
	c.Data["Title"] = c.Tr("admin.users.new_account")
	c.Data["PageIsAdmin"] = true
	c.Data["PageIsAdminUsers"] = true

	sources, err := db.LoginSources.List(db.ListLoginSourceOpts{})
	if err != nil {
		c.Error(err, "list login sources")
		return
	}
	c.Data["Sources"] = sources

	c.Data["CanSendEmail"] = conf.Email.Enabled

	if c.HasError() {
		c.Success(TOOLCHAIN_NEW)
		return
	}

	u := &db.User{
		Name:     f.UserName,
		Email:    f.Email,
		Passwd:   f.Password,
		IsActive: true,
	}

	if len(f.LoginType) > 0 {
		fields := strings.Split(f.LoginType, "-")
		if len(fields) == 2 {
			u.LoginSource = com.StrTo(fields[1]).MustInt64()
			u.LoginName = f.LoginName
		}
	}

	if err := db.CreateUser(u); err != nil {
		switch {
		case db.IsErrUserAlreadyExist(err):
			c.Data["Err_UserName"] = true
			c.RenderWithErr(c.Tr("form.username_been_taken"), USER_NEW, &f)
		case db.IsErrEmailAlreadyUsed(err):
			c.Data["Err_Email"] = true
			c.RenderWithErr(c.Tr("form.email_been_used"), USER_NEW, &f)
		case db.IsErrNameNotAllowed(err):
			c.Data["Err_UserName"] = true
			c.RenderWithErr(c.Tr("user.form.name_not_allowed", err.(db.ErrNameNotAllowed).Value()), USER_NEW, &f)
		default:
			c.Error(err, "create user")
		}
		return
	}
	log.Trace("Account created by admin (%s): %s", c.User.Name, u.Name)

	// Send email notification.
	if f.SendNotify && conf.Email.Enabled {
		email.SendRegisterNotifyMail(c.Context, db.NewMailerUser(u))
	}

	c.Flash.Success(c.Tr("admin.users.new_success", u.Name))
	c.Redirect(conf.Server.Subpath + "/admin/users/" + com.ToStr(u.ID))
}
