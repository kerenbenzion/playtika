package playika

import (
	"encoding/base64"
	"log"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Token struct {
	Token string `json:"token"`
}

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func resourcePlaytikaid() *schema.Resource {
	r := &schema.Resource{
		Create: resourcePlaytikaidAllocate,
		Read:   resourcePlaytikaidRead,
		Update: resourcePlaytikaidUpdate,
		Delete: resourcePlaytikaidRelease,

		Schema: map[string]*schema.Schema{
			"hostname_prefix": &schema.Schema{
				Type:        schema.TypeString,
				Description: "hostname pattern",
				Optional:    true,
			},
			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Description: "cidr of netword",
				Optional:    true,
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(25 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}

	return r
}
func resourcePlaytikaidAllocate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [resourcePlaytikaidAllocate] Allocating new ID")
	log.Println("[DBUG] [resourcePlaytikaidAllocate] printing m.config: %s", m.(*Config))
}

func resourceMachineRead(d *schema.ResourceData, m interface{}) error {

}

func resourceMachineUpdate(d *schema.ResourceData, m interface{}) error {

}

func resourceMachineRelease(d *schema.ResourceData, m interface{}) error {

}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

type Host1 struct {
	Hostname string
}

