module github.com/kerenbenzion/playtika

go 1.13

replace github.com/digitalrebar/provision => /root/go/src/github.com/digitalrebar/provision

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.1
	github.com/kerenbenzion/playtika v0.0.0-20210329071016-9d68d587d67b
	github.com/kerenbenzion/provision v3.8.5+incompatible
	github.com/savaki/jq v0.0.0-20161209013833-0e6baecebbf8
)
