package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"
import "errors"

type UserAPIResponse struct {
	Resource      string `json:"resource"`
	Size          int    `json:"size"`
	PageNumber    int    `json:"pageNumber"`
	TotalElements int    `json:"totalElements"`
	TotalPages    int    `json:"totalPages"`
	Users         []User `json:"users"`
}

type User struct {
	Id                      string `json:"id,omitempty"`
	Resource                string `json:"resource,omitempty"`
	Username                string `json:"username,omitempty"`
	Password                string `json:"password,omitempty"` //required
	Enabled                 bool   `json:"enabled,omitempty"`
	Type                    string `json:"type,omitempty"`
	FirstName               string `json:"firstName,omitempty"`
	LastName                string `json:"lastName,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	EmailAddr               string `json:"emailAddr,omitempty"` //required
	EmailVerified           bool   `json:"emailVerified,omitempty"`
	PhoneNumber             string `json:"phoneNumber,omitempty"`
	ExternalId              string `json:"externalId,omitempty"`
	AccessKeys              string `json:"accessKeys,omitempty"`
	DisableReason           string `json:"disableReason,omitempty"`
	AccountSource           string `json:"accountSource,omitempty"`
	Status                  string `json:"status,omitempty"`
	Detail                  string `json:"detail,omitempty"`
	ActivationData          string `json:"activationData,omitempty"`
	Created                 int64  `json:"created,omitempty"`
	LastUpdated             int64  `json:"lastUpdated,omitempty"`
	CoAdmin                 bool   `json:"coAdmin,omitempty"`
	TenantAdmin             bool   `json:"tenantAdmin,omitempty"`
	ActivationProfileId     string `json:"activationProfileId,omitempty"`
	HasSubscriptionPlanType bool   `json:"hasSubscriptionPlanType,omitempty"`
	TenantId                string `json:"tenantId,omitempty"` //required
}

func (s *Client) GetUsers() ([]User, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data UserAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	users := data.Users
	return users, nil
}

func (s *Client) GetUser(id int) (*User, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + strconv.Itoa(id))

	var data User

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	user := &data

	return user, nil
}

func (s *Client) GetUserFromEmail(email string) (*User, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data UserAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	users := data.Users

	for _, user := range users {
		if user.EmailAddr == email {

			url := fmt.Sprintf(s.BaseURL + "/v1/users/" + user.Id)

			var data User

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return nil, err
			}
			bytes, err := s.doRequest(req)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(bytes, &data)
			if err != nil {
				return nil, err
			}

			user := &data

			return user, nil
		}
	}

	return nil, errors.New("USER NOT FOUND")
}

func (s *Client) AddUser(user *User) (*User, error) {

	var data User

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	j, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	user = &data

	return user, nil
}

func (s *Client) UpdateUser(user *User) (*User, error) {

	var data User

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + user.Id)

	j, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	user = &data

	return user, nil
}

func (s *Client) DeleteUser(userId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + strconv.Itoa(userId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *Client) DeleteUserByEmail(email string) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return err
	}
	var data UserAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	users := data.Users

	for _, user := range users {
		if user.EmailAddr == email {

			url := fmt.Sprintf(s.BaseURL + "/v1/users/" + user.Id)

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return err
			}
			_, err = s.doRequest(req)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("USER NOT FOUND")
}
