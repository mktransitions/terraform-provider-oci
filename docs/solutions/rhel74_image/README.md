This example provides a method to generate a RHEL 7.4 image for use by both VM and BM shapes.
There are several prerequisites for using this process:

1. You MUST have a valid RedHat account with subscriptions available.  The TF template needs a 
   RH Username and Password to allow you to temporarily subscribe the instance that is building the image 
   and get access to the various RH repos.
2. The template expects pre-configured VCNs and Subnets.  
3. You need to provide a URL that points to the RHEL 7.4 ISO.  This URL must contain the name of the ISO, 
   with an '.iso' extension.  An OCI Pre-Authenticated Request (PAR) works well for this operation.  How to create
   OCI PARs can be found here: https://docs.us-phoenix-1.oraclecloud.com/Content/Object/Tasks/managingobjects.htm#par.
4. The template uses filters that expect unique Compartment, VCN and Subnet names.
	NOTE: The root compartment CANNOT be used for this process.
5. The following must be specified in your shell environment (prefixed with TF_VAR_ of course):
    - tenancy_ocid
    - user_ocid
    - fingerprint
    - private_key_path
    - private_key_password (if required)
    - ssh_public_key (the actual public key, not the file)
    - region
6. The subnet to be used must have the following configuration:
	- Port 80 TCP must be allowed on the subnet
	- All ICMP traffic must be allowed on the subnet (ICMP All)

NOTE: A template env-vars file is provided as part of this example.  Simply complete the items inside the template and source the result into your shell by using:

. ./env-vars    

Using this template is simple:

1. Set your environment variables
2. Open the configuration.tf file and substitute the values in each of the sections appropriate to your environment
	NOTE: The AD is specified as either 'AD-x' or 'ad-x' where x is the AD number you wish to use for the process.
3. Execute 'terraform plan; terraform apply'
4. Get coffee or favorite beverage...
5. After your image is created, execute 'terraform destroy -force' (there will not be a resource to actually kill,
   so force is required).

What happens in the background:
The template generates a script that embeds all the configuration files needed to build the iPXE server, extract the ISO
boot the instance used to load RHEL, causes RHEL to load, builds the image, destroys the build instance, and finally destroys the iPXE server.  You are left with a custom image named "RHEL_74" in your environment.

NOTE: The source configuration files for the iPXE server are included here.  It is *STRONGLY* recommended that they not be 
      altered.
      
      
ALSO NOTE: *THE PRIVATE KEY USED TO ACCESS OCI WILL TEMPORARILY BE TRANSFERRED TO THE IPXE INSTANCE.  ONCE THE IPXE INSTANCE IS DESTROYED, THE COPY OF THE PRIVATE KEY IS DESTROYED ALONG WITH IT.*

Enjoy.
