package azuredevopssdk

import "fmt"
import "log"
import "net/http"
import "encoding/json"
import "bytes"

func (s *Client) CreateProject(project Project) (string, error) {
	var a [1]interface{}
	a[0] = project
	
	projectBytes, _ := json.Marshal(a[0])
	projectReader := bytes.NewReader(projectBytes)
	url := fmt.Sprintf(baseURL+"%s/_apis/projects?api-version=5.0-preview.3",s.organization)
	log.Printf(url)
	req, err := http.NewRequest("POST", url, projectReader)
	if err != nil {
		return "", err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return "", err
	}

	var resp ResponseProject
	json.Unmarshal(bytes, &resp)
	return resp.Id, nil
}