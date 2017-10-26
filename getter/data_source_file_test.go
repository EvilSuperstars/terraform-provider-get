package getter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const testDataSourceConfig_basic = `
provider "getter" {}

data "getter_file" "foo" {
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

	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testDataSourceConfig_basic, filename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.getter_file.foo", "content", content),
				),
			},
		},
	})
}
