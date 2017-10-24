package getter

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	gg "github.com/hashicorp/go-getter"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	url := d.Get("url").(string)

	tmpDir, err := ioutil.TempDir("", "gg")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	dst := path.Join(tmpDir, "temp.dat")
	if err := gg.GetFile(dst, url); err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(dst)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("body", string(bytes))

	return nil
}
