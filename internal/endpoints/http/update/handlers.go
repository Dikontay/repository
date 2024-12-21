package update

import (
	"encoding/json"
	"fmt"
	"net/http"
	"repository/internal/domain/models"
	"repository/internal/services"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Path[len("/update/"):]

		var params Params

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Invalid request params", http.StatusBadRequest)
			return
		}

		var user models.User
		user.ID = id
		user.UserProperties = params.UserProperties

		updatedUser, err := services.Repository().UpdateUser(user)

		if err != nil {
			http.Error(w, fmt.Sprintf("failed to update user %v", err), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(updatedUser)
		if err != nil {
			http.Error(w, "failed to return user", http.StatusInternalServerError)
			return
		}

	}
}
