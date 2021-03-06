package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/controllers"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/persistence/repositories"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/providers/database"
	"io/ioutil"
	"net/http"
)

type Register struct {
	controllers.Controller
}

type RegisterRequest struct {
	Nickname, Email, Password string
}

func (a *Register) Index(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, err)
	}

	rr := RegisterRequest{}
	err = json.Unmarshal(body, &rr)

	if err != nil {
		fmt.Fprint(w, err)
	}

	ur := repositories.NewUserRepository(database.GetDb())
	ur.Create(
		entities.NewUserEntity(
			rr.Nickname,
			rr.Email,
			rr.Password,
		),
	)

	fmt.Fprint(w, "User saved successfully!")
}
