// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package permission

import (
	"context"
	"reflect"

	"errors"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a resource pool.
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
//			_, err := Permission.NewPool(ctx, "operationsPool", &Permission.PoolArgs{
//				Comment: pulumi.String("Managed by Pulumi"),
//				PoolId:  pulumi.String("operations-pool"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Instances can be imported using the `pool_id`, e.g.,
//
// bash
//
// ```sh
// $ pulumi import proxmoxve:Permission/pool:Pool operations_pool operations-pool
// ```
type Pool struct {
	pulumi.CustomResourceState

	// The pool comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The pool members.
	Members PoolMemberArrayOutput `pulumi:"members"`
	// The pool identifier.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("proxmoxve:Permission/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("proxmoxve:Permission/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// The pool comment.
	Comment *string `pulumi:"comment"`
	// The pool members.
	Members []PoolMember `pulumi:"members"`
	// The pool identifier.
	PoolId *string `pulumi:"poolId"`
}

type PoolState struct {
	// The pool comment.
	Comment pulumi.StringPtrInput
	// The pool members.
	Members PoolMemberArrayInput
	// The pool identifier.
	PoolId pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The pool comment.
	Comment *string `pulumi:"comment"`
	// The pool identifier.
	PoolId string `pulumi:"poolId"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The pool comment.
	Comment pulumi.StringPtrInput
	// The pool identifier.
	PoolId pulumi.StringInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//	PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//	PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// The pool comment.
func (o PoolOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// The pool members.
func (o PoolOutput) Members() PoolMemberArrayOutput {
	return o.ApplyT(func(v *Pool) PoolMemberArrayOutput { return v.Members }).(PoolMemberArrayOutput)
}

// The pool identifier.
func (o PoolOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pool)(nil)).Elem()
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].([]*Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pool)(nil)).Elem()
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pool {
		return vs[0].(map[string]*Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PoolInput)(nil)).Elem(), &Pool{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolArrayInput)(nil)).Elem(), PoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolMapInput)(nil)).Elem(), PoolMap{})
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
