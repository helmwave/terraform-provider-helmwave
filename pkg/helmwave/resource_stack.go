package helmwave

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resStack() *schema.Resource {
	return &schema.Resource{
		CreateContext: resStackCreate,
		ReadContext:   resStackRead,
		UpdateContext: resStackUpdate,
		DeleteContext: resStackDelete,
		Schema:        map[string]*schema.Schema{},
	}
}

func resStackCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resStackRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resStackUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resStackRead(ctx, d, m)
}

func resStackDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
