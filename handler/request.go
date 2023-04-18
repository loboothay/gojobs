package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Creating Opening
type CreateOpeningRequest struct {
	Role     string
	Company  string
	Location string
	Remote   *bool
	Link     string
	Salary   int64
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("malformed request body or empty")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

//update
type UpdateOpeningRequest struct {
	Role     string
	Company  string
	Location string
	Remote   *bool
	Link     string
	Salary   int64
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
