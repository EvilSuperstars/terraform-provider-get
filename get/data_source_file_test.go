package get

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSource_basic(t *testing.T) {
	dataSourceName := "data.get_file.test"

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
				Config: testDataSourceConfig(filename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "content", content),
				),
			},
		},
	})
}

func testDataSourceConfig(n string) string {
	return fmt.Sprintf(`
provider "get" {}

data "get_file" "test" {
  url = %[1]q
}
`, n)
}
