module github.com/terraform-providers/terraform-provider-oci

require (
	github.com/aws/aws-sdk-go v1.25.2 // indirect
	github.com/fatih/color v1.7.0
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce // indirect
	github.com/hashicorp/hcl2 v0.0.0-20190618163856-0b64543c968c
	github.com/hashicorp/terraform v0.12.4-0.20190628193153-a74738cd35fc
	github.com/mitchellh/cli v1.0.0
	github.com/oracle/oci-go-sdk/v25 v25.0.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.3.0
	gopkg.in/yaml.v2 v2.2.2
)

// Uncomment this line to get OCI Go SDK from local source instead of github
//replace github.com/oracle/oci-go-sdk => ../../oracle/oci-go-sdk
