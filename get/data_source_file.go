package get

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"time"

	gg "github.com/hashicorp/go-getter"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceReadContext,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceReadContext(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	src := d.Get("url").(string)

	tmpDir, err := ioutil.TempDir("", "gg")
	if err != nil {
		return diag.Errorf("error creating temporary directory: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	dst := path.Join(tmpDir, "gg.dat")
	if err := gg.GetFile(dst, src); err != nil {
		return diag.Errorf("error creating downloading from %s to %s: %s", src, dst, err)
	}

	bytes, err := ioutil.ReadFile(dst)
	if err != nil {
		return diag.Errorf("error reading file %s: %s", dst, err)
	}

	d.SetId(time.Now().UTC().String())
	d.Set("content", string(bytes))

	return nil
}
