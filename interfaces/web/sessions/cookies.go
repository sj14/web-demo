package sessions

import (
	"net/http"

	"errors"

	"github.com/gorilla/sessions"
)

// Important Note: If you aren't using gorilla/mux, you need to wrap your handlers with context.ClearHandler as or
// else you will leak memory! An easy way to do this is to wrap the top-level mux when calling http.ListenAndServe:

type Cookie struct {
	name  string
	store *sessions.CookieStore
}

var ErrInvalidKeyLen = errors.New("Key length not 16, 24, or 32 bytes long")

func NewCookie(name string, key []byte) (Cookie, error) {
	var err error = nil

	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	switch len(key) {
	case 16, 24, 32:
		err = nil
	default:
		err = ErrInvalidKeyLen
	}
	return Cookie{name, sessions.NewCookieStore(key)}, err
}

func (c *Cookie) IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	session, _ := c.store.Get(r, c.name)

	// Check if usercontroller is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// ErrorHandler(w, r, http.StatusForbidden)
		return false
	} else {
		return true
	}
}

func (c *Cookie) SessionGetUserId(r *http.Request) (int64, bool) {
	session, err := c.store.Get(r, c.name)
	if err != nil {
		return -1, false
	}

	id := session.Values["id"]
	id64, ok := id.(int64)
	return id64, ok
}

func (c *Cookie) GetSessionValue(key, r *http.Request) interface{} {
	session, _ := c.store.Get(r, c.name)

	val := session.Values[key]
	return val
}

func (c *Cookie) SetSessionValue(key string, value interface{}, w http.ResponseWriter, r *http.Request) error {
	session, err := c.store.Get(r, c.name)
	if err != nil {
		return err
	}

	session.Values[key] = value
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cookie) SetLoggedIn(id int64, w http.ResponseWriter, r *http.Request) error {
	session, err := c.store.Get(r, c.name)
	if err != nil {
		return err
	}

	// Set usercontroller as authenticated
	session.Values["authenticated"] = true
	session.Values["id"] = id
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cookie) Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := c.store.Get(r, c.name)
	if err != nil {
		return err
	}

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Values["id"] = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
