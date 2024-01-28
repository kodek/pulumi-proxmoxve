// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v5/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Ssh struct {
	Agent          *bool     `pulumi:"agent"`
	AgentSocket    *string   `pulumi:"agentSocket"`
	Nodes          []SshNode `pulumi:"nodes"`
	Password       *string   `pulumi:"password"`
	Socks5Password *string   `pulumi:"socks5Password"`
	Socks5Server   *string   `pulumi:"socks5Server"`
	Socks5Username *string   `pulumi:"socks5Username"`
	Username       *string   `pulumi:"username"`
}

// SshInput is an input type that accepts SshArgs and SshOutput values.
// You can construct a concrete instance of `SshInput` via:
//
//	SshArgs{...}
type SshInput interface {
	pulumi.Input

	ToSshOutput() SshOutput
	ToSshOutputWithContext(context.Context) SshOutput
}

type SshArgs struct {
	Agent          pulumi.BoolPtrInput   `pulumi:"agent"`
	AgentSocket    pulumi.StringPtrInput `pulumi:"agentSocket"`
	Nodes          SshNodeArrayInput     `pulumi:"nodes"`
	Password       pulumi.StringPtrInput `pulumi:"password"`
	Socks5Password pulumi.StringPtrInput `pulumi:"socks5Password"`
	Socks5Server   pulumi.StringPtrInput `pulumi:"socks5Server"`
	Socks5Username pulumi.StringPtrInput `pulumi:"socks5Username"`
	Username       pulumi.StringPtrInput `pulumi:"username"`
}

func (SshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ssh)(nil)).Elem()
}

func (i SshArgs) ToSshOutput() SshOutput {
	return i.ToSshOutputWithContext(context.Background())
}

func (i SshArgs) ToSshOutputWithContext(ctx context.Context) SshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshOutput)
}

type SshOutput struct{ *pulumi.OutputState }

func (SshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ssh)(nil)).Elem()
}

func (o SshOutput) ToSshOutput() SshOutput {
	return o
}

func (o SshOutput) ToSshOutputWithContext(ctx context.Context) SshOutput {
	return o
}

func (o SshOutput) Agent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ssh) *bool { return v.Agent }).(pulumi.BoolPtrOutput)
}

func (o SshOutput) AgentSocket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.AgentSocket }).(pulumi.StringPtrOutput)
}

func (o SshOutput) Nodes() SshNodeArrayOutput {
	return o.ApplyT(func(v Ssh) []SshNode { return v.Nodes }).(SshNodeArrayOutput)
}

func (o SshOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o SshOutput) Socks5Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.Socks5Password }).(pulumi.StringPtrOutput)
}

func (o SshOutput) Socks5Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.Socks5Server }).(pulumi.StringPtrOutput)
}

func (o SshOutput) Socks5Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.Socks5Username }).(pulumi.StringPtrOutput)
}

func (o SshOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ssh) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type SshNode struct {
	Address string `pulumi:"address"`
	Name    string `pulumi:"name"`
	Port    *int   `pulumi:"port"`
}

// SshNodeInput is an input type that accepts SshNodeArgs and SshNodeOutput values.
// You can construct a concrete instance of `SshNodeInput` via:
//
//	SshNodeArgs{...}
type SshNodeInput interface {
	pulumi.Input

	ToSshNodeOutput() SshNodeOutput
	ToSshNodeOutputWithContext(context.Context) SshNodeOutput
}

type SshNodeArgs struct {
	Address pulumi.StringInput `pulumi:"address"`
	Name    pulumi.StringInput `pulumi:"name"`
	Port    pulumi.IntPtrInput `pulumi:"port"`
}

func (SshNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshNode)(nil)).Elem()
}

func (i SshNodeArgs) ToSshNodeOutput() SshNodeOutput {
	return i.ToSshNodeOutputWithContext(context.Background())
}

func (i SshNodeArgs) ToSshNodeOutputWithContext(ctx context.Context) SshNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshNodeOutput)
}

// SshNodeArrayInput is an input type that accepts SshNodeArray and SshNodeArrayOutput values.
// You can construct a concrete instance of `SshNodeArrayInput` via:
//
//	SshNodeArray{ SshNodeArgs{...} }
type SshNodeArrayInput interface {
	pulumi.Input

	ToSshNodeArrayOutput() SshNodeArrayOutput
	ToSshNodeArrayOutputWithContext(context.Context) SshNodeArrayOutput
}

type SshNodeArray []SshNodeInput

func (SshNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshNode)(nil)).Elem()
}

func (i SshNodeArray) ToSshNodeArrayOutput() SshNodeArrayOutput {
	return i.ToSshNodeArrayOutputWithContext(context.Background())
}

func (i SshNodeArray) ToSshNodeArrayOutputWithContext(ctx context.Context) SshNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshNodeArrayOutput)
}

type SshNodeOutput struct{ *pulumi.OutputState }

func (SshNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshNode)(nil)).Elem()
}

func (o SshNodeOutput) ToSshNodeOutput() SshNodeOutput {
	return o
}

func (o SshNodeOutput) ToSshNodeOutputWithContext(ctx context.Context) SshNodeOutput {
	return o
}

func (o SshNodeOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v SshNode) string { return v.Address }).(pulumi.StringOutput)
}

func (o SshNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SshNode) string { return v.Name }).(pulumi.StringOutput)
}

func (o SshNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SshNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type SshNodeArrayOutput struct{ *pulumi.OutputState }

func (SshNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshNode)(nil)).Elem()
}

func (o SshNodeArrayOutput) ToSshNodeArrayOutput() SshNodeArrayOutput {
	return o
}

func (o SshNodeArrayOutput) ToSshNodeArrayOutputWithContext(ctx context.Context) SshNodeArrayOutput {
	return o
}

func (o SshNodeArrayOutput) Index(i pulumi.IntInput) SshNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshNode {
		return vs[0].([]SshNode)[vs[1].(int)]
	}).(SshNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SshInput)(nil)).Elem(), SshArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshNodeInput)(nil)).Elem(), SshNodeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SshNodeArrayInput)(nil)).Elem(), SshNodeArray{})
	pulumi.RegisterOutputType(SshOutput{})
	pulumi.RegisterOutputType(SshNodeOutput{})
	pulumi.RegisterOutputType(SshNodeArrayOutput{})
}
