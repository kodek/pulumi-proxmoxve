// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.CT.Inputs
{

    public sealed class ContainerCpuGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        [Input("cores")]
        public Input<int>? Cores { get; set; }

        [Input("units")]
        public Input<int>? Units { get; set; }

        public ContainerCpuGetArgs()
        {
        }
        public static new ContainerCpuGetArgs Empty => new ContainerCpuGetArgs();
    }
}
