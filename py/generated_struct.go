package main

type Person struct {
	ID    int `json:"id,omitempty" bson:"id,omitempty"`
	Owner struct {
		Login               string `json:"login,omitempty" bson:"login,omitempty"`
		ID                  int    `json:"id,omitempty" bson:"id,omitempty"`
		Avatar_URL          string `json:"avatar_url,omitempty" bson:"avatar_url,omitempty"`
		Gravatar_ID         string `json:"gravatar_id,omitempty" bson:"gravatar_id,omitempty"`
		URL                 string `json:"url,omitempty" bson:"url,omitempty"`
		HTML_URL            string `json:"html_url,omitempty" bson:"html_url,omitempty"`
		Followers_URL       string `json:"followers_url,omitempty" bson:"followers_url,omitempty"`
		Following_URL       string `json:"following_url,omitempty" bson:"following_url,omitempty"`
		Gists_URL           string `json:"gists_url,omitempty" bson:"gists_url,omitempty"`
		Starred_URL         string `json:"starred_url,omitempty" bson:"starred_url,omitempty"`
		Subscriptions_URL   string `json:"subscriptions_url,omitempty" bson:"subscriptions_url,omitempty"`
		Organizations_URL   string `json:"organizations_url,omitempty" bson:"organizations_url,omitempty"`
		Repos_URL           string `json:"repos_url,omitempty" bson:"repos_url,omitempty"`
		Events_URL          string `json:"events_url,omitempty" bson:"events_url,omitempty"`
		Received_Events_URL string `json:"received_events_url,omitempty" bson:"received_events_url,omitempty"`
		Type                string `json:"type,omitempty" bson:"type,omitempty"`
		Site_Admin          bool   `json:"site_admin,omitempty" bson:"site_admin,omitempty"`
		Listdata            []int  `json:"listdata,omitempty" bson:"listdata,omitempty"`
		Hello               []any  `json:"hello,omitempty" bson:"hello,omitempty"`
		Owner               struct {
			Login               string `json:"login,omitempty" bson:"login,omitempty"`
			ID                  int    `json:"id,omitempty" bson:"id,omitempty"`
			Avatar_URL          string `json:"avatar_url,omitempty" bson:"avatar_url,omitempty"`
			Gravatar_ID         string `json:"gravatar_id,omitempty" bson:"gravatar_id,omitempty"`
			URL                 string `json:"url,omitempty" bson:"url,omitempty"`
			HTML_URL            string `json:"html_url,omitempty" bson:"html_url,omitempty"`
			Followers_URL       string `json:"followers_url,omitempty" bson:"followers_url,omitempty"`
			Following_URL       string `json:"following_url,omitempty" bson:"following_url,omitempty"`
			Gists_URL           string `json:"gists_url,omitempty" bson:"gists_url,omitempty"`
			Starred_URL         string `json:"starred_url,omitempty" bson:"starred_url,omitempty"`
			Subscriptions_URL   string `json:"subscriptions_url,omitempty" bson:"subscriptions_url,omitempty"`
			Organizations_URL   string `json:"organizations_url,omitempty" bson:"organizations_url,omitempty"`
			Repos_URL           string `json:"repos_url,omitempty" bson:"repos_url,omitempty"`
			Events_URL          string `json:"events_url,omitempty" bson:"events_url,omitempty"`
			Received_Events_URL string `json:"received_events_url,omitempty" bson:"received_events_url,omitempty"`
			Type                string `json:"type,omitempty" bson:"type,omitempty"`
			Site_Admin          bool   `json:"site_admin,omitempty" bson:"site_admin,omitempty"`
			Listdata            []int  `json:"listdata,omitempty" bson:"listdata,omitempty"`
			Hello               []any  `json:"hello,omitempty" bson:"hello,omitempty"`
		} `json:"owner,omitempty" bson:"owner,omitempty"`
	} `json:"owner,omitempty" bson:"owner,omitempty"`
	Site_Admin bool    `json:"site_admin,omitempty" bson:"site_admin,omitempty"`
	Float      float64 `json:"float,omitempty" bson:"float,omitempty"`
	Listdata   []int   `json:"listdata,omitempty" bson:"listdata,omitempty"`
	Hello      []any   `json:"hello,omitempty" bson:"hello,omitempty"`
}
