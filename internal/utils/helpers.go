// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package utils

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

const (
	Charset                       = CharsetWithoutDigits + "0123456789"
	CharsetLowerCaseWithoutDigits = "abcdefghijklmnopqrstuvwxyz"
	CharsetWithoutDigits          = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	OciImageIdsVariable           = `
		variable "InstanceImageOCID" {
			type = "map"
			default = {
				// See https://docs.us-phoenix-1.oraclecloud.com/images/
				// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
				us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
				us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
				eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
				uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
			}
		}
		// Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
		data "oci_core_images" "supported_shape_images" {
			compartment_id   = "${var.tenancy_ocid}"
			shape            = "VM.Standard2.1"
			operating_system = "Oracle Linux"
		}

	`
	DefinedShieldedImageOCIDs = `
      variable "InstanceImageOCIDShieldedCompatible" {
	  type = "map"
	  default = {
		// See https://docs.us-phoenix-1.oraclecloud.com/images/
		// Oracle-provided image "Oracle-Linux-8.4-2021.07.27-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaalw6v7wxgiuedh36jzy5ilbnfjezsxey2glgg3jtlodwzltxregba"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaay7qb3q2bttzhvzzacdpweqo2mvj43tfkm5b4j46xf6pzazspz6aq"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaasehd3xu75nmbxp6lbaaontfgmowlszrx5c72mw4kks5f75euj7gq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaa7fgs4dpcjkkeemyfzyo3yo5lezqfskac45dblmgnfq5az4jmgcza"
	  }
	}`
	FlexVmImageIdsVariable = `
	variable "FlexInstanceImageOCID" {
	  type = "map"
	  default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaalgvdp6hhnulo3tlxz3mtff625s7ix6ianpmv5l7chz5rcakrxbiq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa6tp7lhyrcokdtf7vrbmxyp2pctgg4uxvt4jz4vc47qoc2ec4anha"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaadvi77prh3vjijhwe5xbd6kjg3n5ndxjcpod6om6qaiqeu3csof7a"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaw5gvriwzjhzt2tnylrfnpanz5ndztyrv3zpwhlzxdbkqsjfkwxaq"
	  }
	}
	`
	MysqlConfigurationIdVariable = `
	variable "MysqlConfigurationOCID" {
	  type = "map"
	  default = {
		us-ashburn-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaalwzc2a22xqm56fwjwfymixnulmbq3v77p5v4lcbb6qhkftxf2trq"
		us-phoenix-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaalwzc2a22xqm56fwjwfymixnulmbq3v77p5v4lcbb6qhkftxf2trq"
	  }
	}
 	`
	MysqlConfigurationIdVariableE3_2_32_OCID = `
	variable "MysqlConfigurationE3_2_32_OCID" {
	  type = "map"
	  default = {
		us-ashburn-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaakremacvh2fizcznnja5rdxry2q4nyn27afjblyrimzjmrqblhfwa"
		us-phoenix-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaakremacvh2fizcznnja5rdxry2q4nyn27afjblyrimzjmrqblhfwa"
	  }
	}
   `
	MysqlHAConfigurationIdVariable = `
	variable "MysqlHAConfigurationOCID" {
		type = "map"
		default = {
			us-ashburn-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaantprksu6phqfgr5xvyut46wdfesdszonbclybfwvahgysfjbrb4q"
			us-phoenix-1 = "ocid1.mysqlconfiguration.oc1..aaaaaaaantprksu6phqfgr5xvyut46wdfesdszonbclybfwvahgysfjbrb4q"
		}
	}
	`
	OciWindowsImageIdsVariable = `
	variable "InstanceImageOCID" {
		type = "map"
		default = {
			# The below Image OCIDs are for Windows-Server-2012-R2-Standard-Edition-VM-Gen2-2018.12.12-0
			# See https://docs.cloud.oracle.com/iaas/images/image/5e34cde5-6cef-4cc3-b8f1-c8fc3a088302/
			us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaarlo3ace3wq34aompwj3u2z2xteonboapg663woz6d2iovarowhja"
			us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaabzwak2haqxh3r7h6dajgu4enp7q7hcrreql45awryd5frjsd5l6a"
			eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaourcjktoe3gprvwfksxc36r4rxgbcjs5qvtrja6w6euivci635vq"
			uk-london-1  = "ocid1.image.oc1.uk-london-1.aaaaaaaadb4mg7ii73wkrntmiunr7x7qrh7ompczvy3xbggm27pkhotpgj2q"
		}
	}

`
	VolumeBackupPolicyDependency = `
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {
	filter {
		name = "display_name"
		values = [ "silver" ]
	}
}
`
	OsManagedImageIdsVariable = `
	variable "OsManagedImageOCID" {
	  type = "map"
	  default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaavxm3s4jskx5rcoi63rekg54e3a27v2b7tiuuumnx5owzhkul6ufq"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaqhzgbezuoq5fz7haq5p7uyydfipffclz6w7fwyzge7tcxbbloz3q"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaulz7xiht632iidvdm4iezy33fofulmerq2nkllwnkjy335qkswza"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaayt6ppuyj6q4dwb4pkkyy3llrhxntywewfk4ssd365d4cn22i6yxa"
	  }
	}
	`
)

func RandomString(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomStringOrHttpReplayValue(length int, charset string, httpReplayValue string) string {
	if httpreplay.ModeRecordReplay() {
		return httpReplayValue
	}
	return RandomString(length, charset)
}

// Returns a slice of keys from the given map in alphabetical order
func GetSortedKeys(source map[string]interface{}) []string {
	sortedKeys := make([]string, len(source))
	cnt := 0
	for k := range source {
		sortedKeys[cnt] = k
		cnt++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// Get obo token from file
func GetTokenFromFile(path string) (string, error) {
	token, err := ioutil.ReadFile(path)
	return string(token), err
}

// multierror with \t does not show up on Team City logs,
// replacing \t with 4 blank spaces
func CustomErrorFormat(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred:\n    * %s\n\n", es[0])
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf("%d errors occurred:\n    %s\n\n", len(es), strings.Join(points, "\n    "))
}

// Added for resource discovery AUTH
func GetProviderEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(globalvar.TfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(globalvar.OciEnvPrefix + strings.ToUpper(s))
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

func GetEnvSettingWithBlankDefault(s string) string {
	return GetEnvSettingWithDefault(s, "")
}

func GetEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(globalvar.TfEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(globalvar.OciEnvPrefix + s)
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

// Deprecated: There should be only no need to panic individually
func GetRequiredEnvSetting(s string) string {
	v := GetEnvSettingWithBlankDefault(s)
	if v == "" {
		panic(fmt.Sprintf("Required env setting %s is missing", s))
	}
	return v
}

func GetHomeFolder() string {
	if os.Getenv("TF_HOME_OVERRIDE") != "" {
		return os.Getenv("TF_HOME_OVERRIDE")
	}
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

// cleans and expands the path if it contains a tilde , returns the expanded path or the input path as is if not expansion
// was performed
func ExpandPath(filepath string) string {
	if strings.HasPrefix(filepath, fmt.Sprintf("~%c", os.PathSeparator)) {
		filepath = path.Join(GetHomeFolder(), filepath[2:])
	}
	return path.Clean(filepath)
}

func CheckProfile(profile string, path string) (err error) {
	var profileRegex = regexp.MustCompile(`^\[(.*)\]`)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	splitContent := strings.Split(content, "\n")
	for _, line := range splitContent {
		if match := profileRegex.FindStringSubmatch(line); match != nil && len(match) > 1 && match[1] == profile {
			return nil
		}
	}

	return fmt.Errorf("configuration file did not contain profile: %s", profile)
}

func CheckIncompatibleAttrsForApiKeyAuth(d *schema.ResourceData, apiKeyConfigAttributes [5]string) ([]string, bool) {
	var apiKeyConfigAttributesToUnset []string
	for _, apiKeyConfigAttribute := range apiKeyConfigAttributes {
		apiKeyConfigAttributeValue, hasConfigVariable := d.GetOkExists(apiKeyConfigAttribute)
		if (hasConfigVariable && apiKeyConfigAttributeValue != "") || GetEnvSettingWithBlankDefault(apiKeyConfigAttribute) != "" {
			apiKeyConfigAttributesToUnset = append(apiKeyConfigAttributesToUnset, apiKeyConfigAttribute)
		}
	}
	return apiKeyConfigAttributesToUnset, len(apiKeyConfigAttributesToUnset) == 0
}

func GetCertificateFileBytes(certificateFileFullPath string) (pemRaw []byte, err error) {
	absFile, err := filepath.Abs(certificateFileFullPath)
	if err != nil {
		return nil, fmt.Errorf("can't form absolute path of %s: %v", certificateFileFullPath, err)
	}

	if pemRaw, err = ioutil.ReadFile(absFile); err != nil {
		return nil, fmt.Errorf("can't read %s: %v", certificateFileFullPath, err)
	}
	return
}

func RemoveFile(file string) error {
	return os.Remove(file)
}

func WriteTempFile(data string, originFileName string) (err error) {
	f, err := os.OpenFile(originFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		_, _ = f.WriteString(data)
	}
	return err
}

func ValidateSourceValue(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	info, err := os.Stat(v)
	if err != nil {
		es = append(es, fmt.Errorf("cannot get file information for the specified source: %s", v))
		return
	}
	if info.Size() > 10000*50*1024*1024*1024 {
		es = append(es, fmt.Errorf("the specified source: %s file is too large", v))
	}
	return
}

// GetStringHashcode hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func GetStringHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// GetStringsHashcode hashes a list of strings to a unique hashcode.
func GetStringsHashcode(strings []string) string {
	var buf bytes.Buffer

	for _, s := range strings {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}

	return fmt.Sprintf("%d", GetStringHashcode(buf.String()))
}

func GetVarNameFromAttributeOfResources(attribute string, resourceType string, resourceName string) string {
	// Following format resourceType--attribute-attribute-...–resourceName
	return fmt.Sprintf(globalvar.VariableResourceLevelFormat, resourceType, strings.ReplaceAll(attribute, ".", "-"), resourceName)
}

func GetSDKServiceName(clientServiceName string) string {

	if clientServiceName == "" {
		return ""
	}
	snakeCase := strings.Replace(strings.Split(clientServiceName, ".")[0], "oci_", "", 1)
	return strings.Replace(snakeCase, "_", "", -1)
}

func ListCurrentDirectory() {
	currentDir, err := os.Getwd()
	if err != nil {
		Logf("Error getting current directory: %v", err)
	}

	// Read the files and directories in the current directory
	dirEntries, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	var dirList []string
	var fileList []string
	// Iterate and print each file/directory
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fileList = append(fileList, entry.Name())
		} else {
			dirList = append(dirList, entry.Name())
		}
	}
	Debugf("[DEBUG] Files in directory: %s\nFiles: %s\n Directories: %s", currentDir, strings.Join(fileList, ", "), strings.Join(dirList, ", "))
}
