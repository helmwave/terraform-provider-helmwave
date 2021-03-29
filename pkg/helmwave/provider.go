package helmwave

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const PREFIX = "helmwave_"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			PREFIX + "stack": resStack(),
		},
	}
}
