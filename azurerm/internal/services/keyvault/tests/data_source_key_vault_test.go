package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMKeyVault_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_key_vault", "test")
	config := testAccDataSourceAzureRMKeyVault_basic(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKeyVaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKeyVaultExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "sku_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.key_permissions.0", "create"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.secret_permissions.0", "set"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMKeyVault_basicClassic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_key_vault", "test")
	config := testAccDataSourceAzureRMKeyVault_basic(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKeyVaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKeyVaultExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "sku.0.name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.key_permissions.0", "create"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.secret_permissions.0", "set"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMKeyVault_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_key_vault", "test")
	config := testAccDataSourceAzureRMKeyVault_complete(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKeyVaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKeyVaultExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "sku.0.name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.key_permissions.0", "get"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.secret_permissions.0", "get"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.environment", "Production"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMKeyVault_networkAcls(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_key_vault", "test")
	config := testAccDataSourceAzureRMKeyVault_networkAcls(data)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMKeyVaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMKeyVaultExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "sku.0.name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.tenant_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "access_policy.0.object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.key_permissions.0", "create"),
					resource.TestCheckResourceAttr(data.ResourceName, "access_policy.0.secret_permissions.0", "set"),
					resource.TestCheckResourceAttr(data.ResourceName, "network_acls.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "network_acls.0.default_action", "Allow"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMKeyVault_basic(data acceptance.TestData) string {
	r := testAccAzureRMKeyVault_basic(data)
	return fmt.Sprintf(`
%s

data "azurerm_key_vault" "test" {
  name                = "${azurerm_key_vault.test.name}"
  resource_group_name = "${azurerm_key_vault.test.resource_group_name}"
}
`, r)
}

func testAccDataSourceAzureRMKeyVault_complete(data acceptance.TestData) string {
	r := testAccAzureRMKeyVault_complete(data)
	return fmt.Sprintf(`
%s

data "azurerm_key_vault" "test" {
  name                = "${azurerm_key_vault.test.name}"
  resource_group_name = "${azurerm_key_vault.test.resource_group_name}"
}
`, r)
}

func testAccDataSourceAzureRMKeyVault_networkAcls(data acceptance.TestData) string {
	r := testAccAzureRMKeyVault_networkAclsUpdated(data)
	return fmt.Sprintf(`
%s

data "azurerm_key_vault" "test" {
  name                = "${azurerm_key_vault.test.name}"
  resource_group_name = "${azurerm_key_vault.test.resource_group_name}"
}
`, r)
}
