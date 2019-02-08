package dynatrace

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	dt "github.com/dyladan/dynatrace-go-client/api"
)

func resourceApplicationDetectionRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationDetectionRuleCreate,
		Read:   resourceApplicationDetectionRuleRead,
		Update: resourceApplicationDetectionRuleUpdate,
		Delete: resourceApplicationDetectionRuleDelete,

		Schema: map[string]*schema.Schema{
			"application_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_application_match_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_application_match_target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationDetectionRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	applicationId := d.Get("application_identifier").(string)
	filterPattern := d.Get("filter_pattern").(string)
	filterApplicationMatchType := d.Get("filter_application_match_type").(string)
	filterApplicationMatchTarget := d.Get("filter_application_match_target").(string)

	filterConfig := dt.NameDetectionRuleFilterConfiguration{
		Pattern:                filterPattern,
		ApplicationMatchType:   filterApplicationMatchType,
		ApplicationMatchTarget: filterApplicationMatchTarget,
	}

	body := dt.NameDetectionRuleDetail{
		ApplicationIdentifier: applicationId,
		FilterConfig:          filterConfig,
	}

	created, err := client.CreateApplicationNameDetectionRule(body)

	if err != nil {
		return err
	}

	d.SetId(created.Id)

	return resourceApplicationDetectionRuleRead(d, meta)
}

func resourceApplicationDetectionRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	rule, err := client.GetApplicationNameDetectionRule(d.Id())

	if err != nil {
		return err
	}

	d.SetId(rule.Id)

	log.Printf("%s", rule.Id)

	d.Set("filter_pattern", rule.FilterConfig.Pattern)
	d.Set("filter_application_match_type", rule.FilterConfig.ApplicationMatchType)
	d.Set("filter_application_match_target", rule.FilterConfig.ApplicationMatchTarget)

	return nil
}

func resourceApplicationDetectionRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	applicationIdentifier := d.Get("application_identifier").(string)
	filterPattern := d.Get("filter_pattern").(string)
	filterApplicationMatchType := d.Get("filter_application_match_type").(string)
	filterApplicationMatchTarget := d.Get("filter_application_match_target").(string)

	filterConfig := dt.NameDetectionRuleFilterConfiguration{
		Pattern:                filterPattern,
		ApplicationMatchType:   filterApplicationMatchType,
		ApplicationMatchTarget: filterApplicationMatchTarget,
	}

	body := dt.NameDetectionRuleDetail{
		ApplicationIdentifier: applicationIdentifier,
		FilterConfig:          filterConfig,
	}

	err := client.UpdateApplicationNameDetectionRule(d.Id(), body)

	if err != nil {
		return err
	}

	return resourceApplicationDetectionRuleRead(d, meta)
}

func resourceApplicationDetectionRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ProviderConfig).Client

	return client.DeleteApplicationNameDetectionRule(d.Id())
}
