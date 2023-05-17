package data

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed string `json:"completed"`
}

var Todos = []Todo{
	{ID: "1", Item: "Learn Go", Completed: "false"},
	{ID: "2", Item: "Learn React", Completed: "false"},
	{ID: "3", Item: "Learn Vue", Completed: "false"},
}