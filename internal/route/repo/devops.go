// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"gogs.io/gogs/internal/context"
)

const (
	DEVOPSES = "repo/devops/view"

	NEW_PIPELINE = "repo/devops/new_pipeline"
)

func DevOps(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.devops")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsProject"] = true
	c.Data["Name"] = "DevOps"
	c.Success(DEVOPSES)
}

func Project(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.project")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsProject"] = true
	c.Data["Name"] = "Project"
	c.Success(DEVOPSES)
}

func NewProject(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.project")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsProject"] = true
	c.Data["Name"] = "新建项目"
	c.Success(DEVOPSES)
}

func Git(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.git")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsGit"] = true
	c.Data["Name"] = "Git"
	c.Success(DEVOPSES)
}

func Scan(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.scan")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsScan"] = true
	c.Data["Name"] = "Scan"
	c.Success(DEVOPSES)
}

func Test(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.test")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsTest"] = true
	c.Data["Name"] = "Test"
	c.Success(DEVOPSES)
}

func Deploy(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.deploy")
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsDeploy"] = true
	c.Data["Name"] = "Deploy"
	c.Success(DEVOPSES)
}

func Pipeline(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.pipeline")
	c.Data["PageIsDevOps"] = true
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsPipeline"] = true
	c.Data["Name"] = "Pipeline"
	c.Success(DEVOPSES)
}

func NewPipeline(c *context.Context) {
	c.Data["Title"] = c.Tr("repo.pipeline")
	c.Data["PageIsDevOps"] = true
	c.Data["PageIsDevOpsList"] = true
	c.Data["PageIsPipeline"] = true
	c.Data["Name"] = "新建流水线"
	c.Success(NEW_PIPELINE)
}
