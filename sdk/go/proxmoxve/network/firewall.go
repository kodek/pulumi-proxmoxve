// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages firewall options on the cluster level.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/Network"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Network.NewFirewall(ctx, "example", &Network.FirewallArgs{
//				Ebtables:    pulumi.Bool(false),
//				Enabled:     pulumi.Bool(false),
//				InputPolicy: pulumi.String("DROP"),
//				LogRatelimit: &network.FirewallLogRatelimitArgs{
//					Burst:   pulumi.Int(10),
//					Enabled: pulumi.Bool(false),
//					Rate:    pulumi.String("5/second"),
//				},
//				OutputPolicy: pulumi.String("ACCEPT"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Important Notes
//
// Be careful not to use this resource multiple times for the same node.
//
// ## Import
//
// # Instances can be imported without an ID, but you still need to pass one, e.g., bash
//
// ```sh
//
//	$ pulumi import proxmoxve:Network/firewall:Firewall example example
//
// ```
type Firewall struct {
	pulumi.CustomResourceState

	// Enable ebtables rules cluster wide.
	Ebtables pulumi.BoolPtrOutput `pulumi:"ebtables"`
	// Enable or disable the log rate limit.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrOutput `pulumi:"inputPolicy"`
	// The log rate limit.
	LogRatelimit FirewallLogRatelimitPtrOutput `pulumi:"logRatelimit"`
	// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrOutput `pulumi:"outputPolicy"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		args = &FirewallArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Firewall
	err := ctx.RegisterResource("proxmoxve:Network/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("proxmoxve:Network/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// Enable ebtables rules cluster wide.
	Ebtables *bool `pulumi:"ebtables"`
	// Enable or disable the log rate limit.
	Enabled *bool `pulumi:"enabled"`
	// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy *string `pulumi:"inputPolicy"`
	// The log rate limit.
	LogRatelimit *FirewallLogRatelimit `pulumi:"logRatelimit"`
	// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy *string `pulumi:"outputPolicy"`
}

type FirewallState struct {
	// Enable ebtables rules cluster wide.
	Ebtables pulumi.BoolPtrInput
	// Enable or disable the log rate limit.
	Enabled pulumi.BoolPtrInput
	// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrInput
	// The log rate limit.
	LogRatelimit FirewallLogRatelimitPtrInput
	// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// Enable ebtables rules cluster wide.
	Ebtables *bool `pulumi:"ebtables"`
	// Enable or disable the log rate limit.
	Enabled *bool `pulumi:"enabled"`
	// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy *string `pulumi:"inputPolicy"`
	// The log rate limit.
	LogRatelimit *FirewallLogRatelimit `pulumi:"logRatelimit"`
	// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy *string `pulumi:"outputPolicy"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// Enable ebtables rules cluster wide.
	Ebtables pulumi.BoolPtrInput
	// Enable or disable the log rate limit.
	Enabled pulumi.BoolPtrInput
	// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrInput
	// The log rate limit.
	LogRatelimit FirewallLogRatelimitPtrInput
	// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//	FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//	FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct{ *pulumi.OutputState }

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

// Enable ebtables rules cluster wide.
func (o FirewallOutput) Ebtables() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.BoolPtrOutput { return v.Ebtables }).(pulumi.BoolPtrOutput)
}

// Enable or disable the log rate limit.
func (o FirewallOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The default input policy (`ACCEPT`, `DROP`, `REJECT`).
func (o FirewallOutput) InputPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.InputPolicy }).(pulumi.StringPtrOutput)
}

// The log rate limit.
func (o FirewallOutput) LogRatelimit() FirewallLogRatelimitPtrOutput {
	return o.ApplyT(func(v *Firewall) FirewallLogRatelimitPtrOutput { return v.LogRatelimit }).(FirewallLogRatelimitPtrOutput)
}

// The default output policy (`ACCEPT`, `DROP`, `REJECT`).
func (o FirewallOutput) OutputPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.OutputPolicy }).(pulumi.StringPtrOutput)
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].([]*Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].(map[string]*Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInput)(nil)).Elem(), &Firewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallArrayInput)(nil)).Elem(), FirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMapInput)(nil)).Elem(), FirewallMap{})
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}
