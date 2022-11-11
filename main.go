package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"github.com/bryan-nice/jira-issue-creation/configuration"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {

	var issues []jira.Issue
	var err error
	var config *configuration.Config
	var jiraJql string
	var result string

	// Init configuration from environment variables
	config = new(configuration.Config)
	err = config.Init()
	if err != nil {
		log.Fatalf("Exception: %v", err)
	}

	jiraApiToken := config.JiraApiToken
	jiraUsername := config.JiraUsername
	jiraAccountUrl := config.JiraAccountUrl
	jiraProject := config.JiraProject
	jiraIssueType := config.JiraIssueType
	jiraIssueSummary := config.JiraIssueSummary
	jiraIssueDescription := config.JiraIssueDescription
	jiraIssueAttachment := config.JiraIssueAttachment
	jiraIssueLinks := config.JiraIssueLinks

	log.Fatalf("Issue Links By Me: ", jiraIssueLinks)
}
