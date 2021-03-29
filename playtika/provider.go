package playtika

import (
        "context"
        "fmt"
        "log"
        "strings"
        "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
 * Enable terraform to use playtika as a provider.  Fill out the
 * appropriate functions and information about this plugin.
 */

func Provider() *schema.Provider {
        return &schema.Provider{

                ResourcesMap: map[string]*schema.Resource{
                        "playtika": playtikaid(),
                },

                // note yet, but potentially pools, params and profiles
                DataSourcesMap: map[string]*schema.Resource{},

                Schema: map[string]*schema.Schema{
                        "username": {
                                Type:          schema.TypeString,
                                Optional:      true,
                                Description:   "The phpipam user"
                        },
                        "password": {
                                Type:          schema.TypeString,
                                Optional:      true,
                                Description:   "The phpipam password"
                        },
                        "endpoint": {
                                Type:        schema.TypeString,
                                Required:    true,
                                Description: "The phpipam server URL. ie: https://1.2.3.4:8092",
                                DefaultFunc: schema.MultiEnvDefaultFunc([]string{
                                        "RS_ENDPOINT",
                                }, nil),
                        },
                },

                ConfigureContextFunc: providerConfigure,
        }
}

/*
 * The config method that terraform uses to pass information about configuration
 * to the plugin.
 */
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
        log.Println("[DEBUG] Configuring the playtika provider")
        var diags diag.Diagnostics
		transCfg := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}
		client := &http.Client{Transport: transCfg}
		url := "https://phpipam.corp/api/phpipam/user/"
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		username :=d.Get("username").(string)
		password:= d.Get("password").(string)
		req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
		res_get_token, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res_get_token.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(res_get_token.Body)
		op3, _ := jq.Parse(".data.token")
		token_b, _ := op3.Apply(bodyBytes)
		token := strings.Trim(string(token_b), "\"")
		config := Config{
			token: token
			endpoint : d.Get("endpoint")
		}
        return &config, diags
}
