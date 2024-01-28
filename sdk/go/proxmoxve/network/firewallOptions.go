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

// Manages firewall options on VM / Container level.
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
//			_, err := Network.NewFirewallOptions(ctx, "example", &Network.FirewallOptionsArgs{
//				NodeName:     pulumi.Any(proxmox_virtual_environment_vm.Example.Node_name),
//				VmId:         pulumi.Any(proxmox_virtual_environment_vm.Example.Vm_id),
//				Dhcp:         pulumi.Bool(true),
//				Enabled:      pulumi.Bool(false),
//				Ipfilter:     pulumi.Bool(true),
//				LogLevelIn:   pulumi.String("info"),
//				LogLevelOut:  pulumi.String("info"),
//				Macfilter:    pulumi.Bool(false),
//				Ndp:          pulumi.Bool(true),
//				InputPolicy:  pulumi.String("ACCEPT"),
//				OutputPolicy: pulumi.String("ACCEPT"),
//				Radv:         pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				proxmox_virtual_environment_vm.Example,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FirewallOptions struct {
	pulumi.CustomResourceState

	// Container ID. Leave empty for cluster level aliases.
	ContainerId pulumi.IntPtrOutput `pulumi:"containerId"`
	// Enable DHCP.
	Dhcp pulumi.BoolPtrOutput `pulumi:"dhcp"`
	// Enable or disable the firewall.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The default input
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrOutput `pulumi:"inputPolicy"`
	// Enable default IP filters. This is equivalent to
	// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
	// implicitly contain sane default restrictions such as restricting IPv6 link
	// local addresses to the one derived from the interface's MAC address. For
	// containers the configured IP addresses will be implicitly added.
	Ipfilter pulumi.BoolPtrOutput `pulumi:"ipfilter"`
	// Log level for incoming
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelIn pulumi.StringPtrOutput `pulumi:"logLevelIn"`
	// Log level for outgoing
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelOut pulumi.StringPtrOutput `pulumi:"logLevelOut"`
	// Enable/disable MAC address filter.
	Macfilter pulumi.BoolPtrOutput `pulumi:"macfilter"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp pulumi.BoolPtrOutput `pulumi:"ndp"`
	// Node name.
	NodeName pulumi.StringOutput `pulumi:"nodeName"`
	// The default output
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrOutput `pulumi:"outputPolicy"`
	// Enable Router Advertisement.
	Radv pulumi.BoolPtrOutput `pulumi:"radv"`
	// VM ID. Leave empty for cluster level aliases.
	VmId pulumi.IntPtrOutput `pulumi:"vmId"`
}

// NewFirewallOptions registers a new resource with the given unique name, arguments, and options.
func NewFirewallOptions(ctx *pulumi.Context,
	name string, args *FirewallOptionsArgs, opts ...pulumi.ResourceOption) (*FirewallOptions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeName == nil {
		return nil, errors.New("invalid value for required argument 'NodeName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallOptions
	err := ctx.RegisterResource("proxmoxve:Network/firewallOptions:FirewallOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallOptions gets an existing FirewallOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallOptionsState, opts ...pulumi.ResourceOption) (*FirewallOptions, error) {
	var resource FirewallOptions
	err := ctx.ReadResource("proxmoxve:Network/firewallOptions:FirewallOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallOptions resources.
type firewallOptionsState struct {
	// Container ID. Leave empty for cluster level aliases.
	ContainerId *int `pulumi:"containerId"`
	// Enable DHCP.
	Dhcp *bool `pulumi:"dhcp"`
	// Enable or disable the firewall.
	Enabled *bool `pulumi:"enabled"`
	// The default input
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy *string `pulumi:"inputPolicy"`
	// Enable default IP filters. This is equivalent to
	// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
	// implicitly contain sane default restrictions such as restricting IPv6 link
	// local addresses to the one derived from the interface's MAC address. For
	// containers the configured IP addresses will be implicitly added.
	Ipfilter *bool `pulumi:"ipfilter"`
	// Log level for incoming
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelIn *string `pulumi:"logLevelIn"`
	// Log level for outgoing
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelOut *string `pulumi:"logLevelOut"`
	// Enable/disable MAC address filter.
	Macfilter *bool `pulumi:"macfilter"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp *bool `pulumi:"ndp"`
	// Node name.
	NodeName *string `pulumi:"nodeName"`
	// The default output
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy *string `pulumi:"outputPolicy"`
	// Enable Router Advertisement.
	Radv *bool `pulumi:"radv"`
	// VM ID. Leave empty for cluster level aliases.
	VmId *int `pulumi:"vmId"`
}

type FirewallOptionsState struct {
	// Container ID. Leave empty for cluster level aliases.
	ContainerId pulumi.IntPtrInput
	// Enable DHCP.
	Dhcp pulumi.BoolPtrInput
	// Enable or disable the firewall.
	Enabled pulumi.BoolPtrInput
	// The default input
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrInput
	// Enable default IP filters. This is equivalent to
	// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
	// implicitly contain sane default restrictions such as restricting IPv6 link
	// local addresses to the one derived from the interface's MAC address. For
	// containers the configured IP addresses will be implicitly added.
	Ipfilter pulumi.BoolPtrInput
	// Log level for incoming
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelIn pulumi.StringPtrInput
	// Log level for outgoing
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelOut pulumi.StringPtrInput
	// Enable/disable MAC address filter.
	Macfilter pulumi.BoolPtrInput
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp pulumi.BoolPtrInput
	// Node name.
	NodeName pulumi.StringPtrInput
	// The default output
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrInput
	// Enable Router Advertisement.
	Radv pulumi.BoolPtrInput
	// VM ID. Leave empty for cluster level aliases.
	VmId pulumi.IntPtrInput
}

func (FirewallOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallOptionsState)(nil)).Elem()
}

type firewallOptionsArgs struct {
	// Container ID. Leave empty for cluster level aliases.
	ContainerId *int `pulumi:"containerId"`
	// Enable DHCP.
	Dhcp *bool `pulumi:"dhcp"`
	// Enable or disable the firewall.
	Enabled *bool `pulumi:"enabled"`
	// The default input
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy *string `pulumi:"inputPolicy"`
	// Enable default IP filters. This is equivalent to
	// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
	// implicitly contain sane default restrictions such as restricting IPv6 link
	// local addresses to the one derived from the interface's MAC address. For
	// containers the configured IP addresses will be implicitly added.
	Ipfilter *bool `pulumi:"ipfilter"`
	// Log level for incoming
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelIn *string `pulumi:"logLevelIn"`
	// Log level for outgoing
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelOut *string `pulumi:"logLevelOut"`
	// Enable/disable MAC address filter.
	Macfilter *bool `pulumi:"macfilter"`
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp *bool `pulumi:"ndp"`
	// Node name.
	NodeName string `pulumi:"nodeName"`
	// The default output
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy *string `pulumi:"outputPolicy"`
	// Enable Router Advertisement.
	Radv *bool `pulumi:"radv"`
	// VM ID. Leave empty for cluster level aliases.
	VmId *int `pulumi:"vmId"`
}

// The set of arguments for constructing a FirewallOptions resource.
type FirewallOptionsArgs struct {
	// Container ID. Leave empty for cluster level aliases.
	ContainerId pulumi.IntPtrInput
	// Enable DHCP.
	Dhcp pulumi.BoolPtrInput
	// Enable or disable the firewall.
	Enabled pulumi.BoolPtrInput
	// The default input
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	InputPolicy pulumi.StringPtrInput
	// Enable default IP filters. This is equivalent to
	// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
	// implicitly contain sane default restrictions such as restricting IPv6 link
	// local addresses to the one derived from the interface's MAC address. For
	// containers the configured IP addresses will be implicitly added.
	Ipfilter pulumi.BoolPtrInput
	// Log level for incoming
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelIn pulumi.StringPtrInput
	// Log level for outgoing
	// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
	// `debug`, `nolog`).
	LogLevelOut pulumi.StringPtrInput
	// Enable/disable MAC address filter.
	Macfilter pulumi.BoolPtrInput
	// Enable NDP (Neighbor Discovery Protocol).
	Ndp pulumi.BoolPtrInput
	// Node name.
	NodeName pulumi.StringInput
	// The default output
	// policy (`ACCEPT`, `DROP`, `REJECT`).
	OutputPolicy pulumi.StringPtrInput
	// Enable Router Advertisement.
	Radv pulumi.BoolPtrInput
	// VM ID. Leave empty for cluster level aliases.
	VmId pulumi.IntPtrInput
}

func (FirewallOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallOptionsArgs)(nil)).Elem()
}

type FirewallOptionsInput interface {
	pulumi.Input

	ToFirewallOptionsOutput() FirewallOptionsOutput
	ToFirewallOptionsOutputWithContext(ctx context.Context) FirewallOptionsOutput
}

func (*FirewallOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallOptions)(nil)).Elem()
}

func (i *FirewallOptions) ToFirewallOptionsOutput() FirewallOptionsOutput {
	return i.ToFirewallOptionsOutputWithContext(context.Background())
}

func (i *FirewallOptions) ToFirewallOptionsOutputWithContext(ctx context.Context) FirewallOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOptionsOutput)
}

// FirewallOptionsArrayInput is an input type that accepts FirewallOptionsArray and FirewallOptionsArrayOutput values.
// You can construct a concrete instance of `FirewallOptionsArrayInput` via:
//
//	FirewallOptionsArray{ FirewallOptionsArgs{...} }
type FirewallOptionsArrayInput interface {
	pulumi.Input

	ToFirewallOptionsArrayOutput() FirewallOptionsArrayOutput
	ToFirewallOptionsArrayOutputWithContext(context.Context) FirewallOptionsArrayOutput
}

type FirewallOptionsArray []FirewallOptionsInput

func (FirewallOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallOptions)(nil)).Elem()
}

func (i FirewallOptionsArray) ToFirewallOptionsArrayOutput() FirewallOptionsArrayOutput {
	return i.ToFirewallOptionsArrayOutputWithContext(context.Background())
}

func (i FirewallOptionsArray) ToFirewallOptionsArrayOutputWithContext(ctx context.Context) FirewallOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOptionsArrayOutput)
}

// FirewallOptionsMapInput is an input type that accepts FirewallOptionsMap and FirewallOptionsMapOutput values.
// You can construct a concrete instance of `FirewallOptionsMapInput` via:
//
//	FirewallOptionsMap{ "key": FirewallOptionsArgs{...} }
type FirewallOptionsMapInput interface {
	pulumi.Input

	ToFirewallOptionsMapOutput() FirewallOptionsMapOutput
	ToFirewallOptionsMapOutputWithContext(context.Context) FirewallOptionsMapOutput
}

type FirewallOptionsMap map[string]FirewallOptionsInput

func (FirewallOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallOptions)(nil)).Elem()
}

func (i FirewallOptionsMap) ToFirewallOptionsMapOutput() FirewallOptionsMapOutput {
	return i.ToFirewallOptionsMapOutputWithContext(context.Background())
}

func (i FirewallOptionsMap) ToFirewallOptionsMapOutputWithContext(ctx context.Context) FirewallOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOptionsMapOutput)
}

type FirewallOptionsOutput struct{ *pulumi.OutputState }

func (FirewallOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallOptions)(nil)).Elem()
}

func (o FirewallOptionsOutput) ToFirewallOptionsOutput() FirewallOptionsOutput {
	return o
}

func (o FirewallOptionsOutput) ToFirewallOptionsOutputWithContext(ctx context.Context) FirewallOptionsOutput {
	return o
}

// Container ID. Leave empty for cluster level aliases.
func (o FirewallOptionsOutput) ContainerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.IntPtrOutput { return v.ContainerId }).(pulumi.IntPtrOutput)
}

// Enable DHCP.
func (o FirewallOptionsOutput) Dhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Dhcp }).(pulumi.BoolPtrOutput)
}

// Enable or disable the firewall.
func (o FirewallOptionsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The default input
// policy (`ACCEPT`, `DROP`, `REJECT`).
func (o FirewallOptionsOutput) InputPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.StringPtrOutput { return v.InputPolicy }).(pulumi.StringPtrOutput)
}

// Enable default IP filters. This is equivalent to
// adding an empty `ipfilter-net<id>` ipset for every interface. Such ipsets
// implicitly contain sane default restrictions such as restricting IPv6 link
// local addresses to the one derived from the interface's MAC address. For
// containers the configured IP addresses will be implicitly added.
func (o FirewallOptionsOutput) Ipfilter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Ipfilter }).(pulumi.BoolPtrOutput)
}

// Log level for incoming
// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
// `debug`, `nolog`).
func (o FirewallOptionsOutput) LogLevelIn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.StringPtrOutput { return v.LogLevelIn }).(pulumi.StringPtrOutput)
}

// Log level for outgoing
// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
// `debug`, `nolog`).
func (o FirewallOptionsOutput) LogLevelOut() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.StringPtrOutput { return v.LogLevelOut }).(pulumi.StringPtrOutput)
}

// Enable/disable MAC address filter.
func (o FirewallOptionsOutput) Macfilter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Macfilter }).(pulumi.BoolPtrOutput)
}

// Enable NDP (Neighbor Discovery Protocol).
func (o FirewallOptionsOutput) Ndp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Ndp }).(pulumi.BoolPtrOutput)
}

// Node name.
func (o FirewallOptionsOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.StringOutput { return v.NodeName }).(pulumi.StringOutput)
}

// The default output
// policy (`ACCEPT`, `DROP`, `REJECT`).
func (o FirewallOptionsOutput) OutputPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.StringPtrOutput { return v.OutputPolicy }).(pulumi.StringPtrOutput)
}

// Enable Router Advertisement.
func (o FirewallOptionsOutput) Radv() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.BoolPtrOutput { return v.Radv }).(pulumi.BoolPtrOutput)
}

// VM ID. Leave empty for cluster level aliases.
func (o FirewallOptionsOutput) VmId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallOptions) pulumi.IntPtrOutput { return v.VmId }).(pulumi.IntPtrOutput)
}

type FirewallOptionsArrayOutput struct{ *pulumi.OutputState }

func (FirewallOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallOptions)(nil)).Elem()
}

func (o FirewallOptionsArrayOutput) ToFirewallOptionsArrayOutput() FirewallOptionsArrayOutput {
	return o
}

func (o FirewallOptionsArrayOutput) ToFirewallOptionsArrayOutputWithContext(ctx context.Context) FirewallOptionsArrayOutput {
	return o
}

func (o FirewallOptionsArrayOutput) Index(i pulumi.IntInput) FirewallOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallOptions {
		return vs[0].([]*FirewallOptions)[vs[1].(int)]
	}).(FirewallOptionsOutput)
}

type FirewallOptionsMapOutput struct{ *pulumi.OutputState }

func (FirewallOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallOptions)(nil)).Elem()
}

func (o FirewallOptionsMapOutput) ToFirewallOptionsMapOutput() FirewallOptionsMapOutput {
	return o
}

func (o FirewallOptionsMapOutput) ToFirewallOptionsMapOutputWithContext(ctx context.Context) FirewallOptionsMapOutput {
	return o
}

func (o FirewallOptionsMapOutput) MapIndex(k pulumi.StringInput) FirewallOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallOptions {
		return vs[0].(map[string]*FirewallOptions)[vs[1].(string)]
	}).(FirewallOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallOptionsInput)(nil)).Elem(), &FirewallOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallOptionsArrayInput)(nil)).Elem(), FirewallOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallOptionsMapInput)(nil)).Elem(), FirewallOptionsMap{})
	pulumi.RegisterOutputType(FirewallOptionsOutput{})
	pulumi.RegisterOutputType(FirewallOptionsArrayOutput{})
	pulumi.RegisterOutputType(FirewallOptionsMapOutput{})
}
