package plugins

import "github.com/brettlangdon/slackbot/slack"

type Plugin interface {
	GetName() string
	HandleRequest(slack.APIRequest, []string) (slack.APIResponse, error)
}
