package sessions

import "net/http"

const (
	flashTypeSuccess string = "success"
	flashTypeInfo    string = "info"
	flashTypeWarning string = "warning"
	flashTypeDanger  string = "danger"
)

func (c *Cookie) addFlash(w http.ResponseWriter, r *http.Request, priority string, msg string) error {
	session, err := c.store.Get(r, c.name)
	if err != nil {
		return err
	}

	session.AddFlash(msg, priority)
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cookie) AddFlashInfo(w http.ResponseWriter, r *http.Request, msg string) error {
	err := c.addFlash(w, r, flashTypeInfo, msg)
	return err
}

func (c *Cookie) AddFlashSuccess(w http.ResponseWriter, r *http.Request, msg string) error {
	err := c.addFlash(w, r, flashTypeSuccess, msg)
	return err
}

func (c *Cookie) AddFlashWarning(w http.ResponseWriter, r *http.Request, msg string) error {
	err := c.addFlash(w, r, flashTypeWarning, msg)
	return err
}

func (c *Cookie) AddFlashDanger(w http.ResponseWriter, r *http.Request, msg string) error {
	err := c.addFlash(w, r, flashTypeDanger, msg)
	return err
}

func (c *Cookie) popFlashes(w http.ResponseWriter, r *http.Request, priority string) []string {
	session, _ := c.store.Get(r, c.name)
	interfs := session.Flashes(priority)

	var flashesMsgs []string
	for _, val := range interfs {
		if f, ok := val.(string); ok {
			flashesMsgs = append(flashesMsgs, f)
		}
	}
	session.Save(r, w)
	return flashesMsgs
}

func (c *Cookie) PopFlashesInfo(w http.ResponseWriter, r *http.Request) []string {
	return c.popFlashes(w, r, flashTypeInfo)
}

func (c *Cookie) PopFlashesSuccess(w http.ResponseWriter, r *http.Request) []string {
	return c.popFlashes(w, r, flashTypeSuccess)
}

func (c *Cookie) PopFlashesWarning(w http.ResponseWriter, r *http.Request) []string {
	return c.popFlashes(w, r, flashTypeWarning)
}

func (c *Cookie) PopFlashesDanger(w http.ResponseWriter, r *http.Request) []string {
	return c.popFlashes(w, r, flashTypeDanger)
}
