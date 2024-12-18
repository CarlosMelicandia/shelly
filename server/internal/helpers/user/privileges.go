// We want to have privileges be handled in one place which is here. This will be a file that you would add other privileges
// like IsUserMentor or IsUserSponsor. Following the IsUserAdmin function can give you a good idea on how to write the other
// functions.

package user

import "net/http"

func IsUserAdmin(userId string, w http.ResponseWriter, r *http.Request) bool {
	getUser, err := GetUser(w, r)

	if err != nil {
		return false
	}

	if getUser.IsAdmin {
		return true
	} else {
		return false
	}
}
