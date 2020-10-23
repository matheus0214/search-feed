package search

type Result struct {
	Type    string
	Content string
}

type Values struct {
	Content string `json:"value"`
}

type Matcher interface {
	Search([]*Values) ([]*Result, error)
}

func Match(match Matcher, result []*Result) ([]*Result, error) {
	return nil, nil
}
