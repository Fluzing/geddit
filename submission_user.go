package geddit

//temp file. Package already included this.


type SubChild struct {
	Data SubData `json:"data"`
}

type SubData struct {
	Subreddit string `json:"subreddit"`
	Created float64 `json:"created"`
}

type SubmissionUser struct {
	Children []SubChild `json:"children"`
}
