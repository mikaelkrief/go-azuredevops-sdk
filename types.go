package azuredevopssdk

type BuildDefinition struct {
	ID      int    `json:"id"`
	name 	string `json:"name"`
	path    string  `json:"path"`
}

type Project struct {
	Id				string	`json:"id"`
	Name    		string	`json:"name"`
	Description 	string	`json:"description"`
	Capabilities	Capabilities `json:"capabilities"`
}

type Capabilities struct {
	Versioncontrol Versioncontrol `json:"versioncontrol"`
	ProcessTemplate ProcessTemplate `json:"processTemplate"`
}

type Versioncontrol struct {
	SourceControlType string `json:"sourceControlType"`
}

type ProcessTemplate struct {
	TemplateTypeId	string `json:"templateTypeId"`
}

type RespProject struct {
	Id			string	`json:"id"`
	Status    	string	`json:"status"`
	Url 		string	`json:"url"`
	
}

type RespOperation struct {
	Id			string	`json:"id"`
	Status    	string	`json:"status"`
	Url 		string	`json:"url"`
	
}