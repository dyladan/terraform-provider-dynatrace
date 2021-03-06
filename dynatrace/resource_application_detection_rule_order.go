package dynatrace

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceApplicationDetectionRuleOrder() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationDetectionRuleOrderCreate,
		Read:   resourceApplicationDetectionRuleOrderRead,
		Update: resourceApplicationDetectionRuleOrderUpdate,
		Delete: resourceApplicationDetectionRuleOrderDelete,

		Schema: map[string]*schema.Schema{
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceApplicationDetectionRuleOrderCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	expectedOrder := d.Get("rules").([]interface{})
	ids := make([]string, len(expectedOrder))

	for i, id := range expectedOrder {
		ids[i] = id.(string)
	}

	err := client.UpdateApplicationNameDetectionRuleOrder(ids)

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%v", expectedOrder))

	return resourceApplicationDetectionRuleOrderRead(d, meta)
}

func resourceApplicationDetectionRuleOrderRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	expectedOrder := d.Get("rules").([]interface{})
	var actualOrder []string

	rules, err := client.AllApplicationNameDetectionRules()

	if err != nil {
		return err
	}

	for _, rule := range rules {
		for _, expectedRuleId := range expectedOrder {
			if rule.Id == expectedRuleId {
				actualOrder = append(actualOrder, rule.Id)
				break
			}
		}
	}

	d.Set("rules", actualOrder)

	return nil
}

func resourceApplicationDetectionRuleOrderUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceApplicationDetectionRuleOrderCreate(d, meta)
}

func resourceApplicationDetectionRuleOrderDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
