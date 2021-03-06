package openstack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOpenStackImagesV2ImageDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOpenStackImagesV2ImageDataSourceCirros,
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.openstack_images_image_v2.image_1"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "name", "CirrOS-tf_1"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "container_format", "bare"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "disk_format", "qcow2"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "min_disk_gb", "0"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "min_ram_mb", "0"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "protected", "false"),
					resource.TestCheckResourceAttr(
						"data.openstack_images_image_v2.image_1", "visibility", "private"),
				),
			},
		},
	})
}

func TestAccOpenStackImagesV2ImageDataSource_testQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOpenStackImagesV2ImageDataSourceCirros,
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceQueryTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.openstack_images_image_v2.image_1"),
				),
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceQuerySizeMin,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.openstack_images_image_v2.image_1"),
				),
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceQuerySizeMax,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.openstack_images_image_v2.image_1"),
				),
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceProperty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.openstack_images_image_v2.image_1"),
				),
			},
			{
				Config: testAccOpenStackImagesV2ImageDataSourceCirros,
			},
		},
	})
}

func testAccCheckImagesV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find image data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Image data source ID not set")
		}

		return nil
	}
}

// Standard CirrOS image.
const testAccOpenStackImagesV2ImageDataSourceCirros = `
resource "openstack_images_image_v2" "image_1" {
  name = "CirrOS-tf_1"
  container_format = "bare"
  disk_format = "qcow2"
  image_source_url = "http://download.cirros-cloud.net/0.3.5/cirros-0.3.5-x86_64-disk.img"
  tags = ["cirros-tf_1"]
  properties = {
    foo = "bar"
    bar = "foo"
  }
}

resource "openstack_images_image_v2" "image_2" {
  name = "CirrOS-tf_2"
  container_format = "bare"
  disk_format = "qcow2"
  image_source_url = "http://download.cirros-cloud.net/0.3.5/cirros-0.3.5-x86_64-disk.img"
  tags = ["cirros-tf_2"]
  properties = {
    foo = "bar"
  }
}
`

var testAccOpenStackImagesV2ImageDataSourceBasic = fmt.Sprintf(`
%s

data "openstack_images_image_v2" "image_1" {
	most_recent = true
	name = "${openstack_images_image_v2.image_1.name}"
}
`, testAccOpenStackImagesV2ImageDataSourceCirros)

var testAccOpenStackImagesV2ImageDataSourceQueryTag = fmt.Sprintf(`
%s

data "openstack_images_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	tag = "cirros-tf_1"
}
`, testAccOpenStackImagesV2ImageDataSourceCirros)

var testAccOpenStackImagesV2ImageDataSourceQuerySizeMin = fmt.Sprintf(`
%s

data "openstack_images_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	size_min = "13000000"
}
`, testAccOpenStackImagesV2ImageDataSourceCirros)

var testAccOpenStackImagesV2ImageDataSourceQuerySizeMax = fmt.Sprintf(`
%s

data "openstack_images_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	size_max = "23000000"
}
`, testAccOpenStackImagesV2ImageDataSourceCirros)

var testAccOpenStackImagesV2ImageDataSourceProperty = fmt.Sprintf(`
%s

data "openstack_images_image_v2" "image_1" {
  properties = {
    foo = "bar"
    bar = "foo"
  }
}
`, testAccOpenStackImagesV2ImageDataSourceCirros)
