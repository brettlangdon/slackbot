package sqwiggle

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/brettlangdon/slackbot/slack"
)

type SqwiggleResponse struct {
	Snapshot         string `json:"snapshot"`
	Status           string `json:"status"`
	Name             string `json:"name"`
	EMail            string `json:"email"`
	TimeZone         string `json:"time_zone"`
	Avatar           string `json:"avatar"`
	Message          string `json:"message"`
	SnapshotInterval int    `json:"snapshot_interval"`
}

type Sqwiggle struct {
	name string
}

func NewSqwiggle() Sqwiggle {
	return Sqwiggle{
		name: "sqwiggle",
	}
}

func (this Sqwiggle) GetName() string {
	return this.name
}

func (this Sqwiggle) SetName(name string) {
	this.name = name
}

func (this Sqwiggle) callSqwiggle(token string) (response []SqwiggleResponse, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.sqwiggle.com/users", nil)
	if err != nil {
		return response, err
	}
	request.SetBasicAuth(token, "slackbot")
	sqwiggleResponse, err := client.Do(request)
	if err != nil || sqwiggleResponse.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("API call to sqwiggle with token %s failed", token))
		return response, err
	}

	defer sqwiggleResponse.Body.Close()

	contents, err := ioutil.ReadAll(sqwiggleResponse.Body)
	if err != nil {
		err = errors.New("Could not read sqwiggle api response")
		return response, err
	}
	err = json.Unmarshal(contents, &response)

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not parse sqwiggle api response: %s", string(contents)))
		return response, err
	}

	return response, nil
}

func (this Sqwiggle) HandleRequest(request slack.APIRequest, args []string) (response slack.APIResponse, err error) {
	if request.Text != "#sqwiggle" {
		err := errors.New(fmt.Sprintf("Incorrect trigger text: %s", request.Text))
		return response, err
	}

	if len(args) <= 1 {
		err := errors.New("Missing sqwiggle token argument")
		return response, err
	}

	users, err := this.callSqwiggle(args[1])

	if err != nil {
		return response, err
	}

	response.Text = "Sqwiggle Time"

	for _, user := range users {
		attachment := slack.Attachment{}
		switch user.Status {
		case "offline":
			attachment.Color = "#000000"
		case "available":
			attachment.Color = "#00CC33"
		case "busy":
			attachment.Color = "#FF3333"
		}

		attachment.Title = user.Name
		attachment.TitleLink = user.Snapshot
		attachment.ImageUrl = user.Snapshot
		attachment.Fallback = user.Snapshot + "#.jpg"

		response.Attachments = append(response.Attachments, attachment)
	}

	return response, err
}
