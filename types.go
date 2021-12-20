package main

import "github.com/google/go-github/v41/github"

type Issue struct {
	github.Issue
}

type IssueComment struct {
	github.IssueComment
}
