package model

import "encoding/json"

type Profile struct {
	Category string
	Pubtime  string
	Title    string
	Content  string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
