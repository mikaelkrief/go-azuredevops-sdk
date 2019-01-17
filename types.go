package azuredevopssdk

type BuildDefinition struct {
	ID      int    `json:"id"`
	name 	string `json:"name"`
	path    string  `json:"path"`
}