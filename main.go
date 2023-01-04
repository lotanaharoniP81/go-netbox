package main

import (
	"crypto/tls"
	"fmt"
	"github.com/netbox-community/go-netbox/netbox/client/extras"
	//"github.com/digitalocean/go-netbox/netbox/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/netbox/client"
	//"github.com/netbox-community/go-netbox/netbox/client/ipam"

	"net/http"
)

func main() {

	fmt.Println("hello")

	// ignore expired / self signed SSL certificates
	httpClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	// The 'host', 'protocol' and the 'token' are taken from the 'model' file (locally)
	transport := httptransport.NewWithClient(host, client.DefaultBasePath, []string{protocol}, httpClient)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", token)

	c := client.New(transport, nil)

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get all the IPs
	// working!
	//

	// get all IPs
	//req2 := ipam.NewIpamIPAddressesListParams()
	//res2, err := c.Ipam.IpamIPAddressesList(req2, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", *(res2.Payload.Count))

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get all prefixes for a specific tenant
	// working!
	//

	//var prefixes []*models.Prefix
	//exampleTenant := "sx-il1"
	//req5 := ipam.NewIpamPrefixesListParams()
	//res5, err := c.Ipam.IpamPrefixesList(req5, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//for _, p := range res5.Payload.Results {
	//	if p.Tenant != nil && p.Tenant.Display == exampleTenant {
	//		prefixes = append(prefixes, p)
	//	}
	//}
	//fmt.Printf("the prefixes: %v", prefixes)

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get the amount of the available ips per prefix
	// working!
	//

	//req3 := ipam.NewIpamPrefixesAvailableIpsListParams()
	//req3.SetID(1)
	//res3, err := c.Ipam.IpamPrefixesAvailableIpsList(req3, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", len(res3.Payload))

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get the next available IP (create it) in the relevant prefix!
	// working!
	//

	//// get the available ips per prefix
	//req7 := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(1)
	//for i := 0; i < 1; i++ {
	//	go func() {
	//		res10, err := c.Ipam.IpamPrefixesAvailableIpsCreate(req7, nil)
	//		if err != nil {
	//			fmt.Printf("%v\n", err)
	//			os.Exit(1)
	//		}
	//		fmt.Println(*res10.Payload.Address)
	//	}()
	//}
	//time.Sleep(time.Second * 5)

	///////////////////////////////////////////////////////////////////////////////////

	//
	// deletion! (by IP address - including subnet)
	// working!
	//

	////get all IPs
	//req2 := ipam.NewIpamIPAddressesListParams()
	//res2, err := c.Ipam.IpamIPAddressesList(req2, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	////fmt.Printf("%v\n", *(res2.Payload.Count))
	//
	//exampleAddress := "10.0.0.4/16"
	//var exampleID int64
	//for _, ip := range res2.Payload.Results {
	//	if *ip.Address == exampleAddress {
	//		exampleID = ip.ID
	//	}
	//}
	//fmt.Println(exampleID)
	//
	//// delete IP address
	//reqDelete := ipam.NewIpamIPAddressesDeleteParams().WithID(exampleID)
	//resDelete, err := c.Ipam.IpamIPAddressesDelete(reqDelete, nil)
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("success!")
	//fmt.Println(resDelete)

	///////////////////////////////////////////////////////////////////////////////////

	// todo: add tags to IP
	// todo: change state of IP
	// todo: handle subnet?

	///////////////////////////////////////////////////////////////////////////////////

	// todo: check this section!

	//res10, err := c.Ipam.IpamPrefixesAvailableIpsCreate(req7, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(*res10.Payload.Address)

	//res7, err := c.Ipam.IpamPrefixesAvailableIpsCreate(req7, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res7.Payload[0].Address)
	//fmt.Println()

	//for i := 0; i < 3; i ++ {
	//	go func() {
	//		req3.SetID(2)
	//		res3, err := c.Ipam.IpamPrefixesAvailableIpsList(req3, nil)
	//		if err != nil {
	//			fmt.Printf("%v\n", err)
	//			os.Exit(1)
	//		}
	//		fmt.Printf("%v\n", res3.Payload[0].Address)
	//	}()
	//}

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get all tenants list
	// working!
	//

	//req4 := tenancy.NewTenancyTenantsListParams()
	//res4, err := c.Tenancy.TenancyTenantsList(req4, nil)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Printf("%v\n", res4)

	///////////////////////////////////////////////////////////////////////////////////

	//
	// update description
	// working!
	//

	//params := &models.WritableIPAddress{}
	//
	//address := "10.0.0.1/24"
	//params.Address = &address
	//
	//description := "Description updated"
	//params.Description = description
	//
	//resource := ipam.NewIpamIPAddressesUpdateParams().WithID(25094).WithData(params)
	//_, err := c.Ipam.IpamIPAddressesUpdate(resource, nil)
	//if err != nil {
	//	fmt.Println("error!")
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Good!")

	///////////////////////////////////////////////////////////////////////////////////

	//
	// get all the tags
	// working!
	//

	tags := extras.NewExtrasTagsListParams()
	list, err := c.Extras.ExtrasTagsList(tags, nil)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Printf("The tags: %v", list.Payload.Results[0].Display)

}
