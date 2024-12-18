// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ha

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the list of High Availability groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/HA"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := HA.GetHAGroups(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("dataProxmoxVirtualEnvironmentHagroups", example.GroupIds)
//			return nil
//		})
//	}
//
// ```
func GetHAGroups(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetHAGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHAGroupsResult
	err := ctx.Invoke("proxmoxve:HA/getHAGroups:getHAGroups", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getHAGroups.
type GetHAGroupsResult struct {
	// The identifiers of the High Availability groups.
	GroupIds []string `pulumi:"groupIds"`
	// The unique identifier of this resource.
	Id string `pulumi:"id"`
}

func GetHAGroupsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetHAGroupsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetHAGroupsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("proxmoxve:HA/getHAGroups:getHAGroups", nil, GetHAGroupsResultOutput{}, options).(GetHAGroupsResultOutput), nil
	}).(GetHAGroupsResultOutput)
}

// A collection of values returned by getHAGroups.
type GetHAGroupsResultOutput struct{ *pulumi.OutputState }

func (GetHAGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHAGroupsResult)(nil)).Elem()
}

func (o GetHAGroupsResultOutput) ToGetHAGroupsResultOutput() GetHAGroupsResultOutput {
	return o
}

func (o GetHAGroupsResultOutput) ToGetHAGroupsResultOutputWithContext(ctx context.Context) GetHAGroupsResultOutput {
	return o
}

// The identifiers of the High Availability groups.
func (o GetHAGroupsResultOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHAGroupsResult) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The unique identifier of this resource.
func (o GetHAGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHAGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHAGroupsResultOutput{})
}
