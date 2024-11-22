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

