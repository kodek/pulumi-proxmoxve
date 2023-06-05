// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.CT.Inputs
{

    public sealed class ContainerCloneGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        [Input("vmId", required: true)]
        public Input<int> VmId { get; set; } = null!;

        public ContainerCloneGetArgs()
        {
        }
        public static new ContainerCloneGetArgs Empty => new ContainerCloneGetArgs();
    }
}
