package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)
// Members struct which contains an array of members
type Members struct {
    Members []Member `json:"members"` // members is a key in json result
}

// Member struct which contains a user profile
type Member struct {
    Profile Profile `json:"profile"` // profile is a key in json result
}
// Profile struct which contains user details
type Profile struct {
    UserName string `json:"real_name"`
	Email  string `json:"email"`
	ProfileImage string `json:"image_original"`
}
func main() {
    
    response, err := http.Get("https://slack.com/api/users.list?token=SLACK_TOKEN&limit=1000&pretty=1")
    if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		var members Members
		total := 0
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &members)
		totalMembers := len(members.Members)
		if totalMembers > 0 {
			fmt.Println("These are the list of users who are not uploaded their profile picture in slack")
			for i := 0; i < len(members.Members); i++ {
				
				//fmt.Println(len(members.Members[i].Profile.ProfileImage))
				if len(members.Members[i].Profile.ProfileImage) == 0 {
					fmt.Println("User: " + members.Members[i].Profile.UserName, "| Email: " + members.Members[i].Profile.Email)
					total++
				}
			}
			fmt.Println("We need ", total, "more profile pictures :disappointed:")
		}
	}
}