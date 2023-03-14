package main

type AuthGitHub struct {
	Access_Token string `json:"access_token"`
	Scope        string `json:"scope"`
	Token_Type   string `json:"token_type"`
}

type DataReceived struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		TextMatches []struct {
			ObjectURL string `json:"object_url"`
			Fragment  string `json:"fragment"`
			Matches   []struct {
				Text    string `json:"text"`
				Indices []int  `json:"indices"`
			} `json:"matches"`
		} `json:"text_matches"`
	} `json:"items"`
}
