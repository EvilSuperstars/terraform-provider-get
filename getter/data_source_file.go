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

			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	src := d.Get("url").(string)

	tmpDir, err := ioutil.TempDir("", "gg")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	dst := path.Join(tmpDir, "gg.dat")
	if err := gg.GetFile(dst, src); err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(dst)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("content", string(bytes))

	return nil
}
