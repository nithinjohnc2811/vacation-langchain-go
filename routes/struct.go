package routes

type GenerateVacationIdeaRequest struct {
	FavoriteSeason string   `json:"favourite_season"`
	Hobbies        []string `json:"hobbies"`
	Budget         int      `json:"budget"`
}

type GenerateVacationIdeaResponse struct {
	Id        uuid `json:"id"`
	Completed bool `json:"completed"`
}

type GenerateVacationIdeaResponse struct {
	Id        uuid   `json:"id"`
	Completed bool   `json:"completed"`
	Idea      string `json:"idea"`
}
