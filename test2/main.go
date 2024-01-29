package main

import (
	"fmt"
	"github.com/hexa-org/policy-mapper/models/formats/gcpBind"
	"github.com/hexa-org/policy-mapper/pkg/hexapolicysupport"
)

func main() {

	input := "examples/example_idql.json"
	idqlPolicies, err := hexapolicysupport.ParsePolicyFile(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	// instantiate gcp mapper with attribute translation for username to account.userid
	gcpMapper := gcpBind.New(map[string]string{
		"username": "account.userid",
	})

	// obtain the GCP Binding assignments from IDQL
	bindAssignments := gcpMapper.MapPoliciesToBindings(idqlPolicies)

	// Convert GCP Bind Assignments back into IDQL
	idqlPoliciesAgain, err := gcpMapper.MapBindingAssignmentsToPolicy(bindAssignments)
	if err != nil {
		fmt.Println(err.Error())
	}
}
