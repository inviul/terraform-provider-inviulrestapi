package inviulrestapi

import (
	"bytes"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func dataSourceRestCall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRestCallRead,
		Schema: map[string]*schema.Schema{
			"base_uri": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base endpoint of the api.",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trail path of the api.",
			},
			"http_rest_method": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "HTTP method: GET, POST, PUT, PATCH, DELETE.",
			},
			"rest_out": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response returned by the API.",
			},
			"json_payload": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Json payload in string format with quote esc character.",
			},
		},
	}
}

func dataSourceRestCallRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var req *http.Request
	var err error

	base_uri := d.Get("base_uri").(string)
	path := d.Get("path").(string)
	http_rest_method := d.Get("http_rest_method").(string)

	if http_rest_method != "GET" {
		var jsonPayload = d.Get("json_payload").(string)
		var postBody = []byte(jsonPayload)
		payloadBody := bytes.NewBuffer(postBody)
		req, err = http.NewRequest(http_rest_method, base_uri+path, payloadBody)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(http_rest_method, base_uri+path, nil)
	}

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

	if err := d.Set("rest_out", sb); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
