package api

import "strings"

func getBearerToken(authenticationHeader string) string {
	splitToken := strings.Split(authenticationHeader, "Bearer ")
	if len(splitToken) < 2 {
		return ""
	}
	return splitToken[1]
}

func isValid(user User) bool {
	return len(user.ID) >= 3 && len(user.ID) <= 16
}

func isAuthenticated(token string) bool {
	return token != "" && isValid(User{ID: token})
}
