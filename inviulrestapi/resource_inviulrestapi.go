package inviulrestapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRestCallCreate,
		ReadContext:   resourceRestCallRead,
		UpdateContext: resourceRestCallUpdate,
		DeleteContext: resourceRestCallDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"uri": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base endpoint of the api.",
			},
			"end_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trail path of the api.",
			},
			"http_method": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "HTTP method: GET, POST, PUT, PATCH, DELETE.",
			},
			"resp_out": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response returned by the API.",
			},
			"json_payload_data": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Json payload in string format with quote esc character.",
			},
		},
	}
}

func resourceRestCallCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceRestCallRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceRestCallUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var req *http.Request
	var err error

	uri := d.Get("uri").(string)
	end_path := d.Get("end_path").(string)
	http_method := d.Get("http_method").(string)

	var jsonPayload = d.Get("json_payload_data").(string)
	var postBody = []byte(jsonPayload)
	payloadBody := bytes.NewBuffer(postBody)
	req, err = http.NewRequest(http_method, uri+end_path, payloadBody)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	if resp == nil {
		bt := []byte("{ \"data\" : \"No output\"}")
		resp = []byte(bytes.NewBuffer(bt).String())
	}

	sb := string(resp)
	log.Printf(sb)

	if err := d.Set("resp_out", sb); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func resourceRestCallDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
