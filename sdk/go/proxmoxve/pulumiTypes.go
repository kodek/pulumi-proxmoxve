// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package proxmoxve

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type HostsEntry struct {
	Address   string   `pulumi:"address"`
	Hostnames []string `pulumi:"hostnames"`
}

// HostsEntryInput is an input type that accepts HostsEntryArgs and HostsEntryOutput values.
// You can construct a concrete instance of `HostsEntryInput` via:
//
//	HostsEntryArgs{...}
type HostsEntryInput interface {
	pulumi.Input

	ToHostsEntryOutput() HostsEntryOutput
	ToHostsEntryOutputWithContext(context.Context) HostsEntryOutput
}

type HostsEntryArgs struct {
	Address   pulumi.StringInput      `pulumi:"address"`
	Hostnames pulumi.StringArrayInput `pulumi:"hostnames"`
}

func (HostsEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostsEntry)(nil)).Elem()
}

func (i HostsEntryArgs) ToHostsEntryOutput() HostsEntryOutput {
	return i.ToHostsEntryOutputWithContext(context.Background())
}

func (i HostsEntryArgs) ToHostsEntryOutputWithContext(ctx context.Context) HostsEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostsEntryOutput)
}

// HostsEntryArrayInput is an input type that accepts HostsEntryArray and HostsEntryArrayOutput values.
// You can construct a concrete instance of `HostsEntryArrayInput` via:
//
//	HostsEntryArray{ HostsEntryArgs{...} }
type HostsEntryArrayInput interface {
	pulumi.Input

	ToHostsEntryArrayOutput() HostsEntryArrayOutput
	ToHostsEntryArrayOutputWithContext(context.Context) HostsEntryArrayOutput
}

type HostsEntryArray []HostsEntryInput

func (HostsEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostsEntry)(nil)).Elem()
}

func (i HostsEntryArray) ToHostsEntryArrayOutput() HostsEntryArrayOutput {
	return i.ToHostsEntryArrayOutputWithContext(context.Background())
}

func (i HostsEntryArray) ToHostsEntryArrayOutputWithContext(ctx context.Context) HostsEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostsEntryArrayOutput)
}

type HostsEntryOutput struct{ *pulumi.OutputState }

func (HostsEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostsEntry)(nil)).Elem()
}

func (o HostsEntryOutput) ToHostsEntryOutput() HostsEntryOutput {
	return o
}

func (o HostsEntryOutput) ToHostsEntryOutputWithContext(ctx context.Context) HostsEntryOutput {
	return o
}

func (o HostsEntryOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v HostsEntry) string { return v.Address }).(pulumi.StringOutput)
}

func (o HostsEntryOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HostsEntry) []string { return v.Hostnames }).(pulumi.StringArrayOutput)
}

type HostsEntryArrayOutput struct{ *pulumi.OutputState }

func (HostsEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostsEntry)(nil)).Elem()
}

func (o HostsEntryArrayOutput) ToHostsEntryArrayOutput() HostsEntryArrayOutput {
	return o
}

func (o HostsEntryArrayOutput) ToHostsEntryArrayOutputWithContext(ctx context.Context) HostsEntryArrayOutput {
	return o
}

func (o HostsEntryArrayOutput) Index(i pulumi.IntInput) HostsEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostsEntry {
		return vs[0].([]HostsEntry)[vs[1].(int)]
	}).(HostsEntryOutput)
}

type ProviderSsh struct {
	Agent       *bool             `pulumi:"agent"`
	AgentSocket *string           `pulumi:"agentSocket"`
	Nodes       []ProviderSshNode `pulumi:"nodes"`
	Password    *string           `pulumi:"password"`
	Username    *string           `pulumi:"username"`
}

// ProviderSshInput is an input type that accepts ProviderSshArgs and ProviderSshOutput values.
// You can construct a concrete instance of `ProviderSshInput` via:
//
//	ProviderSshArgs{...}
type ProviderSshInput interface {
	pulumi.Input

	ToProviderSshOutput() ProviderSshOutput
	ToProviderSshOutputWithContext(context.Context) ProviderSshOutput
}

type ProviderSshArgs struct {
	Agent       pulumi.BoolPtrInput       `pulumi:"agent"`
	AgentSocket pulumi.StringPtrInput     `pulumi:"agentSocket"`
	Nodes       ProviderSshNodeArrayInput `pulumi:"nodes"`
	Password    pulumi.StringPtrInput     `pulumi:"password"`
	Username    pulumi.StringPtrInput     `pulumi:"username"`
}

func (ProviderSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderSsh)(nil)).Elem()
}

func (i ProviderSshArgs) ToProviderSshOutput() ProviderSshOutput {
	return i.ToProviderSshOutputWithContext(context.Background())
}

func (i ProviderSshArgs) ToProviderSshOutputWithContext(ctx context.Context) ProviderSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderSshOutput)
}

func (i ProviderSshArgs) ToProviderSshPtrOutput() ProviderSshPtrOutput {
	return i.ToProviderSshPtrOutputWithContext(context.Background())
}

func (i ProviderSshArgs) ToProviderSshPtrOutputWithContext(ctx context.Context) ProviderSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderSshOutput).ToProviderSshPtrOutputWithContext(ctx)
}

// ProviderSshPtrInput is an input type that accepts ProviderSshArgs, ProviderSshPtr and ProviderSshPtrOutput values.
// You can construct a concrete instance of `ProviderSshPtrInput` via:
//
//	        ProviderSshArgs{...}
//
//	or:
//
//	        nil
type ProviderSshPtrInput interface {
	pulumi.Input

	ToProviderSshPtrOutput() ProviderSshPtrOutput
	ToProviderSshPtrOutputWithContext(context.Context) ProviderSshPtrOutput
}

type providerSshPtrType ProviderSshArgs

func ProviderSshPtr(v *ProviderSshArgs) ProviderSshPtrInput {
	return (*providerSshPtrType)(v)
}

func (*providerSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderSsh)(nil)).Elem()
}

func (i *providerSshPtrType) ToProviderSshPtrOutput() ProviderSshPtrOutput {
	return i.ToProviderSshPtrOutputWithContext(context.Background())
}

func (i *providerSshPtrType) ToProviderSshPtrOutputWithContext(ctx context.Context) ProviderSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderSshPtrOutput)
}

type ProviderSshOutput struct{ *pulumi.OutputState }

func (ProviderSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderSsh)(nil)).Elem()
}

func (o ProviderSshOutput) ToProviderSshOutput() ProviderSshOutput {
	return o
}

func (o ProviderSshOutput) ToProviderSshOutputWithContext(ctx context.Context) ProviderSshOutput {
	return o
}

func (o ProviderSshOutput) ToProviderSshPtrOutput() ProviderSshPtrOutput {
	return o.ToProviderSshPtrOutputWithContext(context.Background())
}

func (o ProviderSshOutput) ToProviderSshPtrOutputWithContext(ctx context.Context) ProviderSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderSsh) *ProviderSsh {
		return &v
	}).(ProviderSshPtrOutput)
}

func (o ProviderSshOutput) Agent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderSsh) *bool { return v.Agent }).(pulumi.BoolPtrOutput)
}

func (o ProviderSshOutput) AgentSocket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderSsh) *string { return v.AgentSocket }).(pulumi.StringPtrOutput)
}

func (o ProviderSshOutput) Nodes() ProviderSshNodeArrayOutput {
	return o.ApplyT(func(v ProviderSsh) []ProviderSshNode { return v.Nodes }).(ProviderSshNodeArrayOutput)
}

func (o ProviderSshOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderSsh) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ProviderSshOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderSsh) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ProviderSshPtrOutput struct{ *pulumi.OutputState }

func (ProviderSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderSsh)(nil)).Elem()
}

func (o ProviderSshPtrOutput) ToProviderSshPtrOutput() ProviderSshPtrOutput {
	return o
}

func (o ProviderSshPtrOutput) ToProviderSshPtrOutputWithContext(ctx context.Context) ProviderSshPtrOutput {
	return o
}

func (o ProviderSshPtrOutput) Elem() ProviderSshOutput {
	return o.ApplyT(func(v *ProviderSsh) ProviderSsh {
		if v != nil {
			return *v
		}
		var ret ProviderSsh
		return ret
	}).(ProviderSshOutput)
}

func (o ProviderSshPtrOutput) Agent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderSsh) *bool {
		if v == nil {
			return nil
		}
		return v.Agent
	}).(pulumi.BoolPtrOutput)
}

func (o ProviderSshPtrOutput) AgentSocket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderSsh) *string {
		if v == nil {
			return nil
		}
		return v.AgentSocket
	}).(pulumi.StringPtrOutput)
}

func (o ProviderSshPtrOutput) Nodes() ProviderSshNodeArrayOutput {
	return o.ApplyT(func(v *ProviderSsh) []ProviderSshNode {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(ProviderSshNodeArrayOutput)
}

func (o ProviderSshPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderSsh) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ProviderSshPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderSsh) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ProviderSshNode struct {
	Address string `pulumi:"address"`
	Name    string `pulumi:"name"`
	Port    *int   `pulumi:"port"`
}

// ProviderSshNodeInput is an input type that accepts ProviderSshNodeArgs and ProviderSshNodeOutput values.
// You can construct a concrete instance of `ProviderSshNodeInput` via:
//
//	ProviderSshNodeArgs{...}
type ProviderSshNodeInput interface {
	pulumi.Input

	ToProviderSshNodeOutput() ProviderSshNodeOutput
	ToProviderSshNodeOutputWithContext(context.Context) ProviderSshNodeOutput
}

type ProviderSshNodeArgs struct {
	Address pulumi.StringInput `pulumi:"address"`
	Name    pulumi.StringInput `pulumi:"name"`
	Port    pulumi.IntPtrInput `pulumi:"port"`
}

func (ProviderSshNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderSshNode)(nil)).Elem()
}

func (i ProviderSshNodeArgs) ToProviderSshNodeOutput() ProviderSshNodeOutput {
	return i.ToProviderSshNodeOutputWithContext(context.Background())
}

func (i ProviderSshNodeArgs) ToProviderSshNodeOutputWithContext(ctx context.Context) ProviderSshNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderSshNodeOutput)
}

// ProviderSshNodeArrayInput is an input type that accepts ProviderSshNodeArray and ProviderSshNodeArrayOutput values.
// You can construct a concrete instance of `ProviderSshNodeArrayInput` via:
//
//	ProviderSshNodeArray{ ProviderSshNodeArgs{...} }
type ProviderSshNodeArrayInput interface {
	pulumi.Input

	ToProviderSshNodeArrayOutput() ProviderSshNodeArrayOutput
	ToProviderSshNodeArrayOutputWithContext(context.Context) ProviderSshNodeArrayOutput
}

type ProviderSshNodeArray []ProviderSshNodeInput

func (ProviderSshNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderSshNode)(nil)).Elem()
}

func (i ProviderSshNodeArray) ToProviderSshNodeArrayOutput() ProviderSshNodeArrayOutput {
	return i.ToProviderSshNodeArrayOutputWithContext(context.Background())
}

func (i ProviderSshNodeArray) ToProviderSshNodeArrayOutputWithContext(ctx context.Context) ProviderSshNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderSshNodeArrayOutput)
}

type ProviderSshNodeOutput struct{ *pulumi.OutputState }

func (ProviderSshNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderSshNode)(nil)).Elem()
}

func (o ProviderSshNodeOutput) ToProviderSshNodeOutput() ProviderSshNodeOutput {
	return o
}

func (o ProviderSshNodeOutput) ToProviderSshNodeOutputWithContext(ctx context.Context) ProviderSshNodeOutput {
	return o
}

func (o ProviderSshNodeOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderSshNode) string { return v.Address }).(pulumi.StringOutput)
}

func (o ProviderSshNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderSshNode) string { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderSshNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProviderSshNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ProviderSshNodeArrayOutput struct{ *pulumi.OutputState }

func (ProviderSshNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderSshNode)(nil)).Elem()
}

func (o ProviderSshNodeArrayOutput) ToProviderSshNodeArrayOutput() ProviderSshNodeArrayOutput {
	return o
}

func (o ProviderSshNodeArrayOutput) ToProviderSshNodeArrayOutputWithContext(ctx context.Context) ProviderSshNodeArrayOutput {
	return o
}

func (o ProviderSshNodeArrayOutput) Index(i pulumi.IntInput) ProviderSshNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderSshNode {
		return vs[0].([]ProviderSshNode)[vs[1].(int)]
	}).(ProviderSshNodeOutput)
}

type GetHostsEntry struct {
	Address   string   `pulumi:"address"`
	Hostnames []string `pulumi:"hostnames"`
}

// GetHostsEntryInput is an input type that accepts GetHostsEntryArgs and GetHostsEntryOutput values.
// You can construct a concrete instance of `GetHostsEntryInput` via:
//
//	GetHostsEntryArgs{...}
type GetHostsEntryInput interface {
	pulumi.Input

	ToGetHostsEntryOutput() GetHostsEntryOutput
	ToGetHostsEntryOutputWithContext(context.Context) GetHostsEntryOutput
}

type GetHostsEntryArgs struct {
	Address   pulumi.StringInput      `pulumi:"address"`
	Hostnames pulumi.StringArrayInput `pulumi:"hostnames"`
}

func (GetHostsEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostsEntry)(nil)).Elem()
}

func (i GetHostsEntryArgs) ToGetHostsEntryOutput() GetHostsEntryOutput {
	return i.ToGetHostsEntryOutputWithContext(context.Background())
}

func (i GetHostsEntryArgs) ToGetHostsEntryOutputWithContext(ctx context.Context) GetHostsEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetHostsEntryOutput)
}

// GetHostsEntryArrayInput is an input type that accepts GetHostsEntryArray and GetHostsEntryArrayOutput values.
// You can construct a concrete instance of `GetHostsEntryArrayInput` via:
//
//	GetHostsEntryArray{ GetHostsEntryArgs{...} }
type GetHostsEntryArrayInput interface {
	pulumi.Input

	ToGetHostsEntryArrayOutput() GetHostsEntryArrayOutput
	ToGetHostsEntryArrayOutputWithContext(context.Context) GetHostsEntryArrayOutput
}

type GetHostsEntryArray []GetHostsEntryInput

func (GetHostsEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetHostsEntry)(nil)).Elem()
}

func (i GetHostsEntryArray) ToGetHostsEntryArrayOutput() GetHostsEntryArrayOutput {
	return i.ToGetHostsEntryArrayOutputWithContext(context.Background())
}

func (i GetHostsEntryArray) ToGetHostsEntryArrayOutputWithContext(ctx context.Context) GetHostsEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetHostsEntryArrayOutput)
}

type GetHostsEntryOutput struct{ *pulumi.OutputState }

func (GetHostsEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostsEntry)(nil)).Elem()
}

func (o GetHostsEntryOutput) ToGetHostsEntryOutput() GetHostsEntryOutput {
	return o
}

func (o GetHostsEntryOutput) ToGetHostsEntryOutputWithContext(ctx context.Context) GetHostsEntryOutput {
	return o
}

func (o GetHostsEntryOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostsEntry) string { return v.Address }).(pulumi.StringOutput)
}

func (o GetHostsEntryOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHostsEntry) []string { return v.Hostnames }).(pulumi.StringArrayOutput)
}

type GetHostsEntryArrayOutput struct{ *pulumi.OutputState }

func (GetHostsEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetHostsEntry)(nil)).Elem()
}

func (o GetHostsEntryArrayOutput) ToGetHostsEntryArrayOutput() GetHostsEntryArrayOutput {
	return o
}

func (o GetHostsEntryArrayOutput) ToGetHostsEntryArrayOutputWithContext(ctx context.Context) GetHostsEntryArrayOutput {
	return o
}

func (o GetHostsEntryArrayOutput) Index(i pulumi.IntInput) GetHostsEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetHostsEntry {
		return vs[0].([]GetHostsEntry)[vs[1].(int)]
	}).(GetHostsEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostsEntryInput)(nil)).Elem(), HostsEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostsEntryArrayInput)(nil)).Elem(), HostsEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderSshInput)(nil)).Elem(), ProviderSshArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderSshPtrInput)(nil)).Elem(), ProviderSshArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderSshNodeInput)(nil)).Elem(), ProviderSshNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderSshNodeArrayInput)(nil)).Elem(), ProviderSshNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetHostsEntryInput)(nil)).Elem(), GetHostsEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetHostsEntryArrayInput)(nil)).Elem(), GetHostsEntryArray{})
	pulumi.RegisterOutputType(HostsEntryOutput{})
	pulumi.RegisterOutputType(HostsEntryArrayOutput{})
	pulumi.RegisterOutputType(ProviderSshOutput{})
	pulumi.RegisterOutputType(ProviderSshPtrOutput{})
	pulumi.RegisterOutputType(ProviderSshNodeOutput{})
	pulumi.RegisterOutputType(ProviderSshNodeArrayOutput{})
	pulumi.RegisterOutputType(GetHostsEntryOutput{})
	pulumi.RegisterOutputType(GetHostsEntryArrayOutput{})
}
