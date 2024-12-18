// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package permission

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a specific resource pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/Permission"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Permission.GetPool(ctx, &permission.GetPoolArgs{
//				PoolId: "operations",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPoolResult
	err := ctx.Invoke("proxmoxve:Permission/getPool:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPool.
type LookupPoolArgs struct {
	// The pool identifier.
	PoolId string `pulumi:"poolId"`
}

// A collection of values returned by getPool.
type LookupPoolResult struct {
	// The pool comment.
	Comment string `pulumi:"comment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The pool members.
	Members []GetPoolMember `pulumi:"members"`
	PoolId  string          `pulumi:"poolId"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPoolResultOutput, error) {
			args := v.(LookupPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("proxmoxve:Permission/getPool:getPool", args, LookupPoolResultOutput{}, options).(LookupPoolResultOutput), nil
		}).(LookupPoolResultOutput)
}

// A collection of arguments for invoking getPool.
type LookupPoolOutputArgs struct {
	// The pool identifier.
	PoolId pulumi.StringInput `pulumi:"poolId"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// A collection of values returned by getPool.
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// The pool comment.
func (o LookupPoolResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The pool members.
func (o LookupPoolResultOutput) Members() GetPoolMemberArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []GetPoolMember { return v.Members }).(GetPoolMemberArrayOutput)
}

func (o LookupPoolResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.PoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
