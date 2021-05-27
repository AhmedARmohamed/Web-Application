package controllers

import (
	"fmt"
	"net/http"

	"github.com/AhmedARmohamed/web-applications/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// the initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// This is used to render the form where a user can create
// a new account.
// Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// This is to process the signup form when a user tries to
// create a new user account.
// Post /signup
func (U *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// r.PostForm = map[string][]string
	fmt.Println(w, r.PostForm["email"])
	fmt.Println(w, r.PostForm["password"])
}
