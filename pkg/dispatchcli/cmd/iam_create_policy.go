///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/client"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
)

var (
	createPolicyLong = i18n.T(`Create a dispatch policy`)

	createPolicyExample = i18n.T(`
# Create a policy by specifying a single subject, action, and resource
dispatch iam create policy example_policy --subject user1@example.com --action get --resource "*"

# Create a policy by specifying multiple subjects, actions, and resources
dispatch iam create policy example_policy --subject user1@example.com,user2@example.com --action get,create,delete,update --resource image,function,base-image,secret

dispatch iam create policy example_policy --subject user1@example.com --subject user2@example.com --action get --action create,delete,update --resource image,function --resource base-image,secret
`)

	subjects  *[]string
	actions   *[]string
	resources *[]string
	global    *bool
)

// NewCmdIamCreatePolicy creates command responsible for dispatch policy creation
func NewCmdIamCreatePolicy(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     i18n.T(`policy POLICY_NAME --subject SUBJECTS --action ACTIONS --resource RESOURCES`),
		Short:   i18n.T("Create policy"),
		Long:    createPolicyLong,
		Example: createPolicyExample,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := identityManagerClient()
			err := createPolicy(out, errOut, cmd, args, c)
			CheckErr(err)
		},
	}

	subjects = cmd.Flags().StringSliceP("subject", "s", []string{""}, "subjects of policy rule, separated by comma")
	actions = cmd.Flags().StringSliceP("action", "a", []string{""}, "actions of policy rule, separated by comma")
	resources = cmd.Flags().StringSliceP("resource", "r", []string{""}, "resources of policy rule, separated by comma")
	global = cmd.Flags().Bool("global", false, "applies the policy globally across all organizations")
	return cmd
}

// CallCreatePolicy makes the api call to create a policy
func CallCreatePolicy(c client.IdentityClient) ModelAction {
	return func(p interface{}) error {
		policyModel := p.(*v1.Policy)

		created, err := c.CreatePolicy(context.TODO(), "", policyModel)
		if err != nil {
			return err
		}
		*policyModel = *created
		return nil
	}
}

func createPolicy(out, errOut io.Writer, cmd *cobra.Command, args []string, c client.IdentityClient) error {

	policyName := args[0]
	policyRules := []*v1.Rule{
		{
			Subjects:  *subjects,
			Actions:   *actions,
			Resources: *resources,
		},
	}

	policyModel := &v1.Policy{
		Name:   &policyName,
		Rules:  policyRules,
		Global: *global,
	}

	err := CallCreatePolicy(c)(policyModel)
	if err != nil {
		return err
	}
	if w, err := formatOutput(out, false, policyModel); w {
		return err
	}
	fmt.Fprintf(out, "Created policy: %s\n", *policyModel.Name)
	return nil
}
