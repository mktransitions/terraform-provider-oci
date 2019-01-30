// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectLifecyclePolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createObjectLifecyclePolicy,
		Read:     readObjectLifecyclePolicy,
		Update:   updateObjectLifecyclePolicy,
		Delete:   deleteObjectLifecyclePolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      rulesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_amount": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},
						"time_unit": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitDays),
								string(oci_object_storage.ObjectLifecycleRuleTimeUnitYears),
							}, false),
						},

						// Optional
						"object_name_filter": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"inclusion_prefixes": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Set:      literalTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return CreateResource(d, sync)
}

func readObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

func updateObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return UpdateResource(d, sync)
}

func deleteObjectLifecyclePolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectLifecyclePolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ObjectLifecyclePolicyResourceCrud struct {
	BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.ObjectLifecyclePolicy
	DisableNotFoundRetries bool
}

func (s *ObjectLifecyclePolicyResourceCrud) ID() string {
	return getObjectLifecyclePolicyCompositeId(s.D.Get("bucket").(string), s.D.Get("namespace").(string))
}

func (s *ObjectLifecyclePolicyResourceCrud) Create() error {
	request := oci_object_storage.PutObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.Items = []oci_object_storage.ObjectLifecycleRule{}
	if rules, ok := s.D.GetOkExists("rules"); ok {
		set := rules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_object_storage.ObjectLifecycleRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToObjectLifecycleRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.PutObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectLifecyclePolicyResourceCrud) Get() error {
	request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	bucket, namespace, err := parseObjectLifecyclePolicyCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectLifecyclePolicyResourceCrud) Update() error {
	request := oci_object_storage.PutObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.Items = []oci_object_storage.ObjectLifecycleRule{}
	if rules, ok := s.D.GetOkExists("rules"); ok {
		set := rules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_object_storage.ObjectLifecycleRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := rulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToObjectLifecycleRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.PutObjectLifecyclePolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectLifecyclePolicy
	return nil
}

func (s *ObjectLifecyclePolicyResourceCrud) Delete() error {
	request := oci_object_storage.DeleteObjectLifecyclePolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeleteObjectLifecyclePolicy(context.Background(), request)
	return err
}

func (s *ObjectLifecyclePolicyResourceCrud) SetData() error {

	bucket, namespace, err := parseObjectLifecyclePolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bucket", &bucket)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	rules := []interface{}{}
	for _, item := range s.Res.Items {
		rules = append(rules, ObjectLifecycleRuleToMap(item))
	}
	s.D.Set("rules", schema.NewSet(rulesHashCodeForSets, rules))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func getObjectLifecyclePolicyCompositeId(bucket string, namespace string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	compositeId := "n/" + namespace + "/b/" + bucket + "/l"
	return compositeId
}

func parseObjectLifecyclePolicyCompositeId(compositeId string) (bucket string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*/l", compositeId)
	if !match || len(parts) != 5 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])

	return
}

func (s *ObjectLifecyclePolicyResourceCrud) mapToObjectLifecycleRule(fieldKeyFormat string) (oci_object_storage.ObjectLifecycleRule, error) {
	result := oci_object_storage.ObjectLifecycleRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		tmp := action.(string)
		result.Action = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectNameFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_name_filter")); ok {
		if tmpList := objectNameFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_name_filter"), 0)
			tmp, err := s.mapToObjectNameFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert object_name_filter, encountered error: %v", err)
			}
			result.ObjectNameFilter = &tmp
		}
	}

	if timeAmount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_amount")); ok {
		tmp := timeAmount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert timeAmount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.TimeAmount = &tmpInt64
	}

	if timeUnit, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_unit")); ok {
		tmp := oci_object_storage.ObjectLifecycleRuleTimeUnitEnum(timeUnit.(string))
		result.TimeUnit = tmp
	}

	return result, nil
}

func ObjectLifecycleRuleToMap(obj oci_object_storage.ObjectLifecycleRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = string(*obj.Action)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectNameFilter != nil {
		result["object_name_filter"] = []interface{}{ObjectNameFilterToMap(obj.ObjectNameFilter)}
	}

	if obj.TimeAmount != nil {
		result["time_amount"] = strconv.FormatInt(*obj.TimeAmount, 10)
	}

	result["time_unit"] = string(obj.TimeUnit)

	return result
}

func (s *ObjectLifecyclePolicyResourceCrud) mapToObjectNameFilter(fieldKeyFormat string) (oci_object_storage.ObjectNameFilter, error) {
	result := oci_object_storage.ObjectNameFilter{}

	result.InclusionPrefixes = []string{}
	if inclusionPrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inclusion_prefixes")); ok {
		set := inclusionPrefixes.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			tmp[i] = interfaces[i].(string)
		}
		result.InclusionPrefixes = tmp
	}

	return result, nil
}

func ObjectNameFilterToMap(obj *oci_object_storage.ObjectNameFilter) map[string]interface{} {
	result := map[string]interface{}{}

	inclusionPrefixes := []interface{}{}
	for _, item := range obj.InclusionPrefixes {
		inclusionPrefixes = append(inclusionPrefixes, item)
	}
	result["inclusion_prefixes"] = schema.NewSet(literalTypeHashCodeForSets, inclusionPrefixes)

	return result
}

func rulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if action, ok := m["action"]; ok && action != "" {
		buf.WriteString(fmt.Sprintf("%v-", action))
	}
	if isEnabled, ok := m["is_enabled"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isEnabled))
	}
	if name, ok := m["name"]; ok && name != "" {
		buf.WriteString(fmt.Sprintf("%v-", name))
	}
	if objectNameFilter, ok := m["object_name_filter"]; ok {
		if tmpList := objectNameFilter.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("object_name_filter-")
			objectNameFilterRaw := tmpList[0].(map[string]interface{})
			if inclusionPrefixes, ok := objectNameFilterRaw["inclusion_prefixes"]; ok {
				set := inclusionPrefixes.(*schema.Set)
				inclusionPrefixesArr := set.List()
				for inclusionPrefix := range inclusionPrefixesArr {
					buf.WriteString(fmt.Sprintf("%v-", inclusionPrefix))
				}
			}
		}
	}
	if timeAmount, ok := m["time_amount"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", timeAmount))
	}
	if timeUnit, ok := m["time_unit"]; ok && timeUnit != "" {
		buf.WriteString(fmt.Sprintf("%v-", timeUnit))
	}
	return hashcode.String(buf.String())
}
