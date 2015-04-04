package slack

import (
	"net/http"
	"strconv"
)

type APIRequest struct {
	Token       string
	TeamId      string
	TeamDomain  string
	ChannelId   string
	ChannelName string
	Timestamp   float64
	UserId      string
	UserName    string
	Text        string
	TriggerWord string
}

func ParseAPIRequest(httpReq *http.Request) (response APIRequest) {
	response.Token = httpReq.PostFormValue("token")
	response.TeamId = httpReq.PostFormValue("team_id")
	response.TeamDomain = httpReq.PostFormValue("team_domain")
	response.ChannelId = httpReq.PostFormValue("channel_id")
	response.ChannelName = httpReq.PostFormValue("channel_name")

	ts := httpReq.PostFormValue("timestamp")
	response.Timestamp, _ = strconv.ParseFloat(ts, 32)
	response.UserId = httpReq.PostFormValue("user_id")
	response.UserName = httpReq.PostFormValue("user_name")
	response.Text = httpReq.PostFormValue("text")
	response.TriggerWord = httpReq.PostFormValue("trigger_word")
	return response
}

type APIResponse struct {
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UserName    string       `json:"user_name,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	AuthorIcon string   `json:"author_icon,omitempty"`
	AuthorLink string   `json:"author_link,omitempty"`
	AuthorName string   `json:"author_name,omitempty"`
	Color      string   `json:"color,omitempty"`
	Fallback   string   `json:"fallback,omitempty"`
	Fields     []Fields `json:"fields,omitempty"`
	ImageUrl   string   `json:"image_url,omitempty"`
	MarkdownIn string   `json:"mrkdwn_in,omitempty"`
	Pretext    string   `json:"pretext,omitempty"`
	Title      string   `json:"title,omitempty"`
	TitleLink  string   `json:"title_link,omitempty"`
}

type Fields struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
