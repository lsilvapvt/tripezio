package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Poi struct {
	item.Item

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Photos      []string `json:"photos"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Address1    string   `json:"address1"`
	Address2    string   `json:"address2"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	ZipCode     string   `json:"zipcode"`
	Country     string   `json:"country"`
	Rating      int      `json:"rating"`
	Articles    []string `json:"articles"`
	Videos      []string `json:"videos"`
}

// MarshalEditor writes a buffer of html to edit a Poi within the CMS
// and implements editor.Editable
func (p *Poi) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Poi field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", p, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Tags("Tags", p, map[string]string{
				"label":       "Tags",
				"placeholder": "+Tags",
			}),
		},
		editor.Field{
			View: editor.Input("Latitude", p, map[string]string{
				"label":       "Latitude",
				"type":        "text",
				"placeholder": "Enter the Latitude here",
			}),
		},
		editor.Field{
			View: editor.Input("Longitude", p, map[string]string{
				"label":       "Longitude",
				"type":        "text",
				"placeholder": "Enter the Longitude here",
			}),
		},
		editor.Field{
			View: editor.Select("Country", p, map[string]string{
				"label": "Country",
			}, map[string]string{
				"BR": "Brazil",
				"US": "United States",
				"CA": "Canada",
			}),
		},
		editor.Field{
			View: editor.Input("ZipCode", p, map[string]string{
				"label":       "Zip Code",
				"type":        "text",
				"placeholder": "Enter the Zip Code here",
			}),
		},
		editor.Field{
			View: editor.Input("State", p, map[string]string{
				"label":       "State",
				"type":        "text",
				"placeholder": "Enter the State here",
			}),
		},
		editor.Field{
			View: editor.Input("City", p, map[string]string{
				"label":       "City",
				"type":        "text",
				"placeholder": "Enter the City here",
			}),
		},
		editor.Field{
			View: editor.Input("Address1", p, map[string]string{
				"label":       "Address 1",
				"type":        "text",
				"placeholder": "Enter the Address here",
			}),
		},
		editor.Field{
			View: editor.Input("Address2", p, map[string]string{
				"label":       "Address 2",
				"type":        "text",
				"placeholder": "Enter the remaining of Address here",
			}),
		},
		editor.Field{
			View: editor.FileRepeater("Photos", p, map[string]string{
				"label":       "Photos",
				"placeholder": "Upload the Photos here",
			}),
		},
		editor.Field{
			View: editor.Input("Rating", p, map[string]string{
				"label":       "Rating",
				"type":        "text",
				"placeholder": "Enter the Rating here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Articles", p, map[string]string{
				"label":       "Related Articles",
				"placeholder": "Upload a related article URL here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Videos", p, map[string]string{
				"label":       "Related Videos",
				"placeholder": "Upload a related video URL here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Poi editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Poi"] = func() interface{} { return new(Poi) }
}

// String defines how a Poi is printed. Update it using more descriptive
// fields from the Poi struct type
func (p *Poi) String() string {
	return fmt.Sprintf("%s, %s %s - %s", p.Name, p.City, p.State, p.Country)
}
