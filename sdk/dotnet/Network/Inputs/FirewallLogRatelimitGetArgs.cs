// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.Network.Inputs
{

    public sealed class FirewallLogRatelimitGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Initial burst of packages which will always get
        /// logged before the rate is applied (defaults to `5`).
        /// </summary>
        [Input("burst")]
        public Input<int>? Burst { get; set; }

        /// <summary>
        /// Enable or disable the log rate limit.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Frequency with which the burst bucket gets refilled
        /// (defaults to `1/second`).
        /// </summary>
        [Input("rate")]
        public Input<string>? Rate { get; set; }

        public FirewallLogRatelimitGetArgs()
        {
        }
        public static new FirewallLogRatelimitGetArgs Empty => new FirewallLogRatelimitGetArgs();
    }
}
