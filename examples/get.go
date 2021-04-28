package examples

import "fmt"

type Endpoints struct {
	CurrentUser      string `json:"current_user_url"`
	AuthorizationURL string `json:"authorizations_url"`
	RepositoryURL    string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	str := fmt.Sprintf("Status Code: %d\n", response.StatusCode())
	str += fmt.Sprintf("Status: %s\n", response.Status())
	str += fmt.Sprintf("Body: %s\n", response.String())

	fmt.Println(str)

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryURL))

	return &endpoints, nil
}
