package model

type HealthResponse struct {
	Status int `json:"status"`
}

type CheckAnswerResponse struct {
	User            string `json:"user"`
	Question        string `json:"question"`
	SubmittedAnswer string `json:"submittedAnswer"`
	Correct         bool   `json:"correct"`
}

type CheckAnswerRequestBody struct {
	User     string `json:"user"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type AddPlayerRequestBody struct {
	User string `json:"user"`
}

type AddPlayerResponse struct {
	User  string `json:"user"`
	Added bool   `json:"added"`
}

type CheckScoresResponse struct {
	Scores map[string]int `json:"scores"`
}

type Question struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Choices  []string `json:"choices"`
	Category string   `json:"category"`
}
