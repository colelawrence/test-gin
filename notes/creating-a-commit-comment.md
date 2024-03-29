
## Creating a Commit Comment
https://developer.github.com/v3/repos/comments/#create-a-commit-comment
POST /repos/:owner/:repo/commits/:sha/comments
### Input
```
Name  Type  	    Description
body	string	    Required. The contents of the comment.
path	string	    Relative path of the file to comment on.
position	integer	Line index in the diff to comment on.
line	    integer	Deprecated. Use position parameter instead. Line number in the file to comment on.
```
### Example
{
  "body": "Great stuff",
  "path": "file1.txt",
  "position": 4,
  "line": null
}
#### Response

Status: 201 Created
Location: https://api.github.com/repos/octocat/Hello-World/comments/1
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4999
```json
{
  "html_url": "https://github.com/octocat/Hello-World/commit/6dcb09b5b57875f334f61aebed695e2e4193db5e#commitcomment-1",
  "url": "https://api.github.com/repos/octocat/Hello-World/comments/1",
  "id": 1,
  "body": "Great stuff",
  "path": "file1.txt",
  "position": 4,
  "line": 14,
  "commit_id": "6dcb09b5b57875f334f61aebed695e2e4193db5e",
  "user": {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  },
  "created_at": "2011-04-14T16:00:49Z",
  "updated_at": "2011-04-14T16:00:49Z"
}
```