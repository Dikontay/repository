package delete

import (
	"encoding/json"
	"fmt"
	"net/http"
	"repository/internal/services"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Path[len("/delete/"):]

		deletedUser, err := services.Repository().DeleteUser(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to delete user %v", err), http.StatusInternalServerError)
			return
		}

		resp := Response{
			User: *deletedUser,
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}

	}
}
