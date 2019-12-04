// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"log"
	"os"

	"gopl.io/ch4/github"
)

//!+template
import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><几种类型的print href='{{.HTMLURL}}'>{{.Number}}</几种类型的print></td>
  <td>{{.State}}</td>
  <td><几种类型的print href='{{.User.HTMLURL}}'>{{.User.Login}}</几种类型的print></td>
  <td><几种类型的print href='{{.HTMLURL}}'>{{.Title}}</几种类型的print></td>
</tr>
{{end}}
</table>
`))

//!-template

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//!-
