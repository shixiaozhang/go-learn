package render

import (
	"bbs-go/model"
	"bbs-go/model/constants"
	"bbs-go/package/common"
	"bbs-go/package/html"
	"bbs-go/package/markdown"
	"github.com/mlogclub/simple"
)

func BuildProject(project *model.Project) *model.ProjectResponse {
	if project == nil {
		return nil
	}
	rsp := &model.ProjectResponse{}
	rsp.ProjectId = project.Id
	rsp.User = BuildUserDefaultIfNull(project.UserId)
	rsp.Name = project.Name
	rsp.Title = project.Title
	rsp.Logo = project.Logo
	rsp.Url = project.Url
	rsp.Url = project.Url
	rsp.DocUrl = project.DocUrl
	rsp.CreateTime = project.CreateTime

	if project.ContentType == constants.ContentTypeHtml {
		rsp.Content = handleHtmlContent(project.Content)
		rsp.Summary = simple.GetSummary(simple.GetHtmlText(project.Content), constants.SummaryLen)
	} else {
		content := markdown.ToHTML(project.Content)
		summary := html.GetSummary(content, constants.SummaryLen)
		rsp.Content = handleHtmlContent(content)
		rsp.Summary = summary
	}

	return rsp
}

func BuildSimpleProjects(projects []model.Project) []model.ProjectSimpleResponse {
	if projects == nil || len(projects) == 0 {
		return nil
	}
	var responses []model.ProjectSimpleResponse
	for _, project := range projects {
		responses = append(responses, *BuildSimpleProject(&project))
	}
	return responses
}

func BuildSimpleProject(project *model.Project) *model.ProjectSimpleResponse {
	if project == nil {
		return nil
	}
	rsp := &model.ProjectSimpleResponse{}
	rsp.ProjectId = project.Id
	rsp.User = BuildUserDefaultIfNull(project.UserId)
	rsp.Name = project.Name
	rsp.Title = project.Title
	rsp.Logo = project.Logo
	rsp.Url = project.Url
	rsp.DocUrl = project.DocUrl
	rsp.DownloadUrl = project.DownloadUrl
	rsp.CreateTime = project.CreateTime

	if project.ContentType == constants.ContentTypeHtml {
		rsp.Summary = simple.GetSummary(simple.GetHtmlText(project.Content), constants.SummaryLen)
	} else {
		rsp.Summary = common.GetMarkdownSummary(project.Content)
	}

	return rsp
}
