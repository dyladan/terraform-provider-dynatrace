package dynatrace

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAwsCredentialsConfigutation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsCredentialsConfigutationCreate,
		Read:   resourceAwsCredentialsConfigutationRead,
		Update: resourceAwsCredentialsConfigutationUpdate,
		Delete: resourceAwsCredentialsConfigutationtDelete,
		Exists: resourceAwsCredentialsConfigutationExists,

		Schema: map[string]*schema.Schema{
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"partition_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS_CN"}, false),
			},
			"authentication_data": {
				Type: schema.TypeList,
				Required:     true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"type": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{"KEYS"}, false),
						},
						"key_base_authentication": {
							Type: schema.TypeList,
							Optional:     true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema {
									"access_key": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},
									"secret_key": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},									
								},
							},
						},
						"role_based_authentication": {
							Type: schema.TypeList,
							Optional:     true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iam_role": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},
									"account_id": &schema.Schema {
										Type: schema.TypeString,
										Required: true,
									},									
								},
							},
						},
					},
				},
			},
			"tagged_only": &schema.Schema {
				Type: schema.TypeBool,
				Optional: true,
				Default: true,
			},
			"tags_to_monitor": &schema.Schema {
				Type: schema.TypeList,
				Optional:     true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"name": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
						},									
					},
				},
			},
			"supporting_services": &schema.Schema {
				Type: schema.TypeList,
				Optional:     true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"name": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
						},
						"metrics": &schema.Schema {
							Type: schema.TypeList,
							Required: true,
							MinItems : 1,
							Elem: &schema.Schema{
							    Type: schema.TypeString,
							},
						},									
					},
				},
			},
		},
	}
}

func resourceAwsCredentialsConfigutationCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsCredentialsConfigutationRead(d, meta)
}

func resourceAwsCredentialsConfigutationRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsCredentialsConfigutationUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsCredentialsConfigutationRead(d, meta)
}

func resourceAwsCredentialsConfigutationtDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsCredentialsConfigutationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return true,nil
}
