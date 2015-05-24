package podio

import "fmt"

type App struct {
	Id              int    `json:"app_id"`
	Name            string `json:"name"`
	Status          string `json:"status"`
	DefaultViewId   int    `json:"default_view_id"`
	URLAdd          string `json:"url_add"`
	IconId          int    `json:"icon_id"`
	LinkAdd         string `json:"link_add"`
	CurrentRevision int    `json:"current_revision"`
	ItemName        string `json:"item_name"`
	Link            string `json:"link"`
	URL             string `json:"url"`
	URLLabel        string `json:"url_label"`
	SpaceId         int    `json:"space_id"`
	Icon            string `json:"icon"`
}

func (client *Client) GetApps(spaceId int) (apps []App, err error) {
	path := fmt.Sprintf("/app/space/%d?view=micro", spaceId)
	err = client.Request("GET", path, nil, nil, &apps)
	return
}

func (client *Client) GetApp(id int) (app *App, err error) {
	path := fmt.Sprintf("/app/%d?view=micro", id)
	err = client.Request("GET", path, nil, nil, &app)
	return
}

func (client *Client) GetAppBySpaceIdAndSlug(spaceId int, slug string) (app *App, err error) {
	path := fmt.Sprintf("/app/space/%d/%s", spaceId, slug)
	err = client.Request("GET", path, nil, nil, &app)
	return
}
