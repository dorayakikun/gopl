package github

// POST /repos/:owner/:repo/issues 対応struct
// Assignee, Assignees, Milestone は除外する
type PostIssueParameter struct {
	Title string `json:"title"`
	Body *string `json:"body,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
}
// PATCH /repos/:owner/:repo/issues/:issue_number 対応struct
// Assignee, Assignees, Milestone は除外する
type PatchIssueParameter struct {
	Title *string `json:"title,omitempty"`
	Body *string `json:"body,omitempty"`
	State *string `json:"state,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
}

type IssueResponse struct {
	Number int
	State  string
	Title  string
	Body   string
	Labels []Label
}

type Label struct {
	Name string
	Color string
}