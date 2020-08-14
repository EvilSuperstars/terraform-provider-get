package get

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourceConfig_basic = `
provider "get" {}

data "get_file" "foo" {
  url = "%s"
}
`

func TestDataSource_basic(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}
	filename := filepath.Join(wd, "test.txt")
	content := "Some contents..."

	if err := ioutil.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
		t.Fatal(err)
		return
	}
	defer os.Remove(filename)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceConfig_basic, filename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.get_file.foo", "content", content),
				),
			},
		},
	})
}
