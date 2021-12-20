package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
	"github.com/sirupsen/logrus"
)

// WebhookHandle is struct to handle
type WebhookHandle struct {

}

// WebhookHandle init
func(m *WebhookHandle) WebhookHandle(c *gin.Context) {
	var event interface{}

	payload, err := github.ValidatePayload(c.Request, []byte(""))
	if err != nil {
		logrus.Errorf("validate payload failed: %v", err)
		return
	}

	event, err = github.ParseWebHook(github.WebHookType(c.Request), payload)
	if err != nil {
		logrus.Errorf("parse webhook failed: %v", err)
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf("request body: %v", string(body))

		c.Writer.WriteHeader(400)
		result, err := c.Writer.Write([]byte(err.Error()))
		if err != nil {
			logrus.Errorf("rw write: %v", result)
		}
		return
	}

	switch event := event.(type) {
	case *github.PingEvent:
		logrus.Infof("PingEvent: %v", event.GetHook().Events)
	case *github.IssueEvent:
		logrus.Infof("IssueEvent: %v", *event.ID)
	case *github.PullRequestEvent:
		logrus.Infof("PullRequestEvent: %v", *event.Number)
	case *github.IssueCommentEvent:
		logrus.Infof("IssueCommentEvent: %v", *event)
	case *github.PushEvent:
		logrus.Infof("PushEvent: %v", *event.PushID)
	}
}
