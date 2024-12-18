// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cluster

import (
	"context"
	"reflect"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about all available nodes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/muhlba91/pulumi-proxmoxve/sdk/v6/go/proxmoxve/Cluster"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cluster.GetNodes(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNodes(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNodesResult
	err := ctx.Invoke("proxmoxve:Cluster/getNodes:getNodes", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getNodes.
type GetNodesResult struct {
	// The CPU count for each node.
	CpuCounts []int `pulumi:"cpuCounts"`
	// The CPU utilization on each node.
	CpuUtilizations []float64 `pulumi:"cpuUtilizations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The memory available on each node.
	MemoryAvailables []int `pulumi:"memoryAvailables"`
	// The memory used on each node.
	MemoryUseds []int `pulumi:"memoryUseds"`
	// The node names.
	Names []string `pulumi:"names"`
	// Whether a node is online.
	Onlines []bool `pulumi:"onlines"`
	// The SSL fingerprint for each node.
	SslFingerprints []string `pulumi:"sslFingerprints"`
	// The support level for each node.
	SupportLevels []string `pulumi:"supportLevels"`
	// The uptime in seconds for each node.
	Uptimes []int `pulumi:"uptimes"`
}

func GetNodesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetNodesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetNodesResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("proxmoxve:Cluster/getNodes:getNodes", nil, GetNodesResultOutput{}, options).(GetNodesResultOutput), nil
	}).(GetNodesResultOutput)
}

// A collection of values returned by getNodes.
type GetNodesResultOutput struct{ *pulumi.OutputState }

func (GetNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesResult)(nil)).Elem()
}

func (o GetNodesResultOutput) ToGetNodesResultOutput() GetNodesResultOutput {
	return o
}

func (o GetNodesResultOutput) ToGetNodesResultOutputWithContext(ctx context.Context) GetNodesResultOutput {
	return o
}

// The CPU count for each node.
func (o GetNodesResultOutput) CpuCounts() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []int { return v.CpuCounts }).(pulumi.IntArrayOutput)
}

// The CPU utilization on each node.
func (o GetNodesResultOutput) CpuUtilizations() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []float64 { return v.CpuUtilizations }).(pulumi.Float64ArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The memory available on each node.
func (o GetNodesResultOutput) MemoryAvailables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []int { return v.MemoryAvailables }).(pulumi.IntArrayOutput)
}

// The memory used on each node.
func (o GetNodesResultOutput) MemoryUseds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []int { return v.MemoryUseds }).(pulumi.IntArrayOutput)
}

// The node names.
func (o GetNodesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Whether a node is online.
func (o GetNodesResultOutput) Onlines() pulumi.BoolArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []bool { return v.Onlines }).(pulumi.BoolArrayOutput)
}

// The SSL fingerprint for each node.
func (o GetNodesResultOutput) SslFingerprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []string { return v.SslFingerprints }).(pulumi.StringArrayOutput)
}

// The support level for each node.
func (o GetNodesResultOutput) SupportLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []string { return v.SupportLevels }).(pulumi.StringArrayOutput)
}

// The uptime in seconds for each node.
func (o GetNodesResultOutput) Uptimes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []int { return v.Uptimes }).(pulumi.IntArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNodesResultOutput{})
}
