// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.HA
{
    public static class GetHAResources
    {
        /// <summary>
        /// Retrieves the list of High Availability resources.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using ProxmoxVE = Pulumi.ProxmoxVE;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleAll = ProxmoxVE.HA.GetHAResources.Invoke();
        /// 
        ///     var exampleVm = ProxmoxVE.HA.GetHAResources.Invoke(new()
        ///     {
        ///         Type = "vm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dataProxmoxVirtualEnvironmentHaresources"] = 
        ///         {
        ///             { "all", exampleAll.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///             { "vms", exampleVm.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///         },
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetHAResourcesResult> InvokeAsync(GetHAResourcesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHAResourcesResult>("proxmoxve:HA/getHAResources:getHAResources", args ?? new GetHAResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the list of High Availability resources.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using ProxmoxVE = Pulumi.ProxmoxVE;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleAll = ProxmoxVE.HA.GetHAResources.Invoke();
        /// 
        ///     var exampleVm = ProxmoxVE.HA.GetHAResources.Invoke(new()
        ///     {
        ///         Type = "vm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dataProxmoxVirtualEnvironmentHaresources"] = 
        ///         {
        ///             { "all", exampleAll.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///             { "vms", exampleVm.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///         },
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetHAResourcesResult> Invoke(GetHAResourcesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHAResourcesResult>("proxmoxve:HA/getHAResources:getHAResources", args ?? new GetHAResourcesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the list of High Availability resources.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using ProxmoxVE = Pulumi.ProxmoxVE;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleAll = ProxmoxVE.HA.GetHAResources.Invoke();
        /// 
        ///     var exampleVm = ProxmoxVE.HA.GetHAResources.Invoke(new()
        ///     {
        ///         Type = "vm",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["dataProxmoxVirtualEnvironmentHaresources"] = 
        ///         {
        ///             { "all", exampleAll.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///             { "vms", exampleVm.Apply(getHAResourcesResult =&gt; getHAResourcesResult.ResourceIds) },
        ///         },
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetHAResourcesResult> Invoke(GetHAResourcesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHAResourcesResult>("proxmoxve:HA/getHAResources:getHAResources", args ?? new GetHAResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHAResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetHAResourcesArgs()
        {
        }
        public static new GetHAResourcesArgs Empty => new GetHAResourcesArgs();
    }

    public sealed class GetHAResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetHAResourcesInvokeArgs()
        {
        }
        public static new GetHAResourcesInvokeArgs Empty => new GetHAResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class GetHAResourcesResult
    {
        /// <summary>
        /// The unique identifier of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identifiers of the High Availability resources.
        /// </summary>
        public readonly ImmutableArray<string> ResourceIds;
        /// <summary>
        /// The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetHAResourcesResult(
            string id,

            ImmutableArray<string> resourceIds,

            string? type)
        {
            Id = id;
            ResourceIds = resourceIds;
            Type = type;
        }
    }
}
