package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

//New initializes a new form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "Field Required!")
		return false
	}
	return true
}
