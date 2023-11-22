// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A security group is a collection of rules, defined at cluster level, which can
// be used in all VMs' rules. For example, you can define a group named “webserver”
// with rules to open the http and https ports. Rules can be created on the cluster
// level, on VM / Container level.
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
//			_, err := Network.NewFirewallRules(ctx, "inbound", &Network.FirewallRulesArgs{
//				NodeName: pulumi.Any(proxmox_virtual_environment_vm.Example.Node_name),
//				VmId:     pulumi.Any(proxmox_virtual_environment_vm.Example.Vm_id),
//				Rules: network.FirewallRulesRuleArray{
//					&network.FirewallRulesRuleArgs{
//						Type:    pulumi.String("in"),
//						Action:  pulumi.String("ACCEPT"),
//						Comment: pulumi.String("Allow HTTP"),
//						Dest:    pulumi.String("192.168.1.5"),
//						Dport:   pulumi.String("80"),
//						Proto:   pulumi.String("tcp"),
//						Log:     pulumi.String("info"),
//					},
//					&network.FirewallRulesRuleArgs{
//						Type:    pulumi.String("in"),
//						Action:  pulumi.String("ACCEPT"),
//						Comment: pulumi.String("Allow HTTPS"),
//						Dest:    pulumi.String("192.168.1.5"),
//						Dport:   pulumi.String("443"),
//						Proto:   pulumi.String("tcp"),
//						Log:     pulumi.String("info"),
//					},
//					&network.FirewallRulesRuleArgs{
//						SecurityGroup: pulumi.Any(proxmox_virtual_environment_cluster_firewall_security_group.Example.Name),
//						Comment:       pulumi.String("From security group"),
//						Iface:         pulumi.String("net0"),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				proxmox_virtual_environment_vm.Example,
//				proxmox_virtual_environment_cluster_firewall_security_group.Example,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FirewallRules struct {
	pulumi.CustomResourceState

	// Container ID. Leave empty for cluster level
	// rules.
	ContainerId pulumi.IntPtrOutput `pulumi:"containerId"`
	// Node name. Leave empty for cluster level rules.
	NodeName pulumi.StringPtrOutput `pulumi:"nodeName"`
	// Firewall rule block (multiple blocks supported).
	// The provider supports two types of the `rule` blocks:
	// - a rule definition block, which includes the following arguments:
	Rules FirewallRulesRuleArrayOutput `pulumi:"rules"`
	// VM ID. Leave empty for cluster level rules.
	VmId pulumi.IntPtrOutput `pulumi:"vmId"`
}

// NewFirewallRules registers a new resource with the given unique name, arguments, and options.
func NewFirewallRules(ctx *pulumi.Context,
	name string, args *FirewallRulesArgs, opts ...pulumi.ResourceOption) (*FirewallRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallRules
	err := ctx.RegisterResource("proxmoxve:Network/firewallRules:FirewallRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRules gets an existing FirewallRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRulesState, opts ...pulumi.ResourceOption) (*FirewallRules, error) {
	var resource FirewallRules
	err := ctx.ReadResource("proxmoxve:Network/firewallRules:FirewallRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRules resources.
type firewallRulesState struct {
	// Container ID. Leave empty for cluster level
	// rules.
	ContainerId *int `pulumi:"containerId"`
	// Node name. Leave empty for cluster level rules.
	NodeName *string `pulumi:"nodeName"`
	// Firewall rule block (multiple blocks supported).
	// The provider supports two types of the `rule` blocks:
	// - a rule definition block, which includes the following arguments:
	Rules []FirewallRulesRule `pulumi:"rules"`
	// VM ID. Leave empty for cluster level rules.
	VmId *int `pulumi:"vmId"`
}

type FirewallRulesState struct {
	// Container ID. Leave empty for cluster level
	// rules.
	ContainerId pulumi.IntPtrInput
	// Node name. Leave empty for cluster level rules.
	NodeName pulumi.StringPtrInput
	// Firewall rule block (multiple blocks supported).
	// The provider supports two types of the `rule` blocks:
	// - a rule definition block, which includes the following arguments:
	Rules FirewallRulesRuleArrayInput
	// VM ID. Leave empty for cluster level rules.
	VmId pulumi.IntPtrInput
}

func (FirewallRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRulesState)(nil)).Elem()
}

type firewallRulesArgs struct {
	// Container ID. Leave empty for cluster level
	// rules.
	ContainerId *int `pulumi:"containerId"`
	// Node name. Leave empty for cluster level rules.
	NodeName *string `pulumi:"nodeName"`
	// Firewall rule block (multiple blocks supported).
	// The provider supports two types of the `rule` blocks:
	// - a rule definition block, which includes the following arguments:
	Rules []FirewallRulesRule `pulumi:"rules"`
	// VM ID. Leave empty for cluster level rules.
	VmId *int `pulumi:"vmId"`
}

// The set of arguments for constructing a FirewallRules resource.
type FirewallRulesArgs struct {
	// Container ID. Leave empty for cluster level
	// rules.
	ContainerId pulumi.IntPtrInput
	// Node name. Leave empty for cluster level rules.
	NodeName pulumi.StringPtrInput
	// Firewall rule block (multiple blocks supported).
	// The provider supports two types of the `rule` blocks:
	// - a rule definition block, which includes the following arguments:
	Rules FirewallRulesRuleArrayInput
	// VM ID. Leave empty for cluster level rules.
	VmId pulumi.IntPtrInput
}

func (FirewallRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRulesArgs)(nil)).Elem()
}

type FirewallRulesInput interface {
	pulumi.Input

	ToFirewallRulesOutput() FirewallRulesOutput
	ToFirewallRulesOutputWithContext(ctx context.Context) FirewallRulesOutput
}

func (*FirewallRules) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRules)(nil)).Elem()
}

func (i *FirewallRules) ToFirewallRulesOutput() FirewallRulesOutput {
	return i.ToFirewallRulesOutputWithContext(context.Background())
}

func (i *FirewallRules) ToFirewallRulesOutputWithContext(ctx context.Context) FirewallRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesOutput)
}

// FirewallRulesArrayInput is an input type that accepts FirewallRulesArray and FirewallRulesArrayOutput values.
// You can construct a concrete instance of `FirewallRulesArrayInput` via:
//
//	FirewallRulesArray{ FirewallRulesArgs{...} }
type FirewallRulesArrayInput interface {
	pulumi.Input

	ToFirewallRulesArrayOutput() FirewallRulesArrayOutput
	ToFirewallRulesArrayOutputWithContext(context.Context) FirewallRulesArrayOutput
}

type FirewallRulesArray []FirewallRulesInput

func (FirewallRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRules)(nil)).Elem()
}

func (i FirewallRulesArray) ToFirewallRulesArrayOutput() FirewallRulesArrayOutput {
	return i.ToFirewallRulesArrayOutputWithContext(context.Background())
}

func (i FirewallRulesArray) ToFirewallRulesArrayOutputWithContext(ctx context.Context) FirewallRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesArrayOutput)
}

// FirewallRulesMapInput is an input type that accepts FirewallRulesMap and FirewallRulesMapOutput values.
// You can construct a concrete instance of `FirewallRulesMapInput` via:
//
//	FirewallRulesMap{ "key": FirewallRulesArgs{...} }
type FirewallRulesMapInput interface {
	pulumi.Input

	ToFirewallRulesMapOutput() FirewallRulesMapOutput
	ToFirewallRulesMapOutputWithContext(context.Context) FirewallRulesMapOutput
}

type FirewallRulesMap map[string]FirewallRulesInput

func (FirewallRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRules)(nil)).Elem()
}

func (i FirewallRulesMap) ToFirewallRulesMapOutput() FirewallRulesMapOutput {
	return i.ToFirewallRulesMapOutputWithContext(context.Background())
}

func (i FirewallRulesMap) ToFirewallRulesMapOutputWithContext(ctx context.Context) FirewallRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesMapOutput)
}

type FirewallRulesOutput struct{ *pulumi.OutputState }

func (FirewallRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRules)(nil)).Elem()
}

func (o FirewallRulesOutput) ToFirewallRulesOutput() FirewallRulesOutput {
	return o
}

func (o FirewallRulesOutput) ToFirewallRulesOutputWithContext(ctx context.Context) FirewallRulesOutput {
	return o
}

// Container ID. Leave empty for cluster level
// rules.
func (o FirewallRulesOutput) ContainerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallRules) pulumi.IntPtrOutput { return v.ContainerId }).(pulumi.IntPtrOutput)
}

// Node name. Leave empty for cluster level rules.
func (o FirewallRulesOutput) NodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRules) pulumi.StringPtrOutput { return v.NodeName }).(pulumi.StringPtrOutput)
}

// Firewall rule block (multiple blocks supported).
// The provider supports two types of the `rule` blocks:
// - a rule definition block, which includes the following arguments:
func (o FirewallRulesOutput) Rules() FirewallRulesRuleArrayOutput {
	return o.ApplyT(func(v *FirewallRules) FirewallRulesRuleArrayOutput { return v.Rules }).(FirewallRulesRuleArrayOutput)
}

// VM ID. Leave empty for cluster level rules.
func (o FirewallRulesOutput) VmId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallRules) pulumi.IntPtrOutput { return v.VmId }).(pulumi.IntPtrOutput)
}

type FirewallRulesArrayOutput struct{ *pulumi.OutputState }

func (FirewallRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRules)(nil)).Elem()
}

func (o FirewallRulesArrayOutput) ToFirewallRulesArrayOutput() FirewallRulesArrayOutput {
	return o
}

func (o FirewallRulesArrayOutput) ToFirewallRulesArrayOutputWithContext(ctx context.Context) FirewallRulesArrayOutput {
	return o
}

func (o FirewallRulesArrayOutput) Index(i pulumi.IntInput) FirewallRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRules {
		return vs[0].([]*FirewallRules)[vs[1].(int)]
	}).(FirewallRulesOutput)
}

type FirewallRulesMapOutput struct{ *pulumi.OutputState }

func (FirewallRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRules)(nil)).Elem()
}

func (o FirewallRulesMapOutput) ToFirewallRulesMapOutput() FirewallRulesMapOutput {
	return o
}

func (o FirewallRulesMapOutput) ToFirewallRulesMapOutputWithContext(ctx context.Context) FirewallRulesMapOutput {
	return o
}

func (o FirewallRulesMapOutput) MapIndex(k pulumi.StringInput) FirewallRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRules {
		return vs[0].(map[string]*FirewallRules)[vs[1].(string)]
	}).(FirewallRulesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesInput)(nil)).Elem(), &FirewallRules{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesArrayInput)(nil)).Elem(), FirewallRulesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRulesMapInput)(nil)).Elem(), FirewallRulesMap{})
	pulumi.RegisterOutputType(FirewallRulesOutput{})
	pulumi.RegisterOutputType(FirewallRulesArrayOutput{})
	pulumi.RegisterOutputType(FirewallRulesMapOutput{})
}
