// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.Network
{
    /// <summary>
    /// Manages firewall options on VM / Container level.
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
    ///     var example = new ProxmoxVE.Network.FirewallOptions("example", new()
    ///     {
    ///         NodeName = proxmox_virtual_environment_vm.Example.Node_name,
    ///         VmId = proxmox_virtual_environment_vm.Example.Vm_id,
    ///         Dhcp = true,
    ///         Enabled = false,
    ///         Ipfilter = true,
    ///         LogLevelIn = "info",
    ///         LogLevelOut = "info",
    ///         Macfilter = false,
    ///         Ndp = true,
    ///         InputPolicy = "ACCEPT",
    ///         OutputPolicy = "ACCEPT",
    ///         Radv = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             proxmox_virtual_environment_vm.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ProxmoxVEResourceType("proxmoxve:Network/firewallOptions:FirewallOptions")]
    public partial class FirewallOptions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Container ID. Leave empty for cluster level aliases.
        /// </summary>
        [Output("containerId")]
        public Output<int?> ContainerId { get; private set; } = null!;

        /// <summary>
        /// Enable DHCP.
        /// </summary>
        [Output("dhcp")]
        public Output<bool?> Dhcp { get; private set; } = null!;

        /// <summary>
        /// Enable or disable the firewall.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The default input
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Output("inputPolicy")]
        public Output<string?> InputPolicy { get; private set; } = null!;

        /// <summary>
        /// Enable default IP filters. This is equivalent to
        /// adding an empty `ipfilter-net&lt;id&gt;` ipset for every interface. Such ipsets
        /// implicitly contain sane default restrictions such as restricting IPv6 link
        /// local addresses to the one derived from the interface's MAC address. For
        /// containers the configured IP addresses will be implicitly added.
        /// </summary>
        [Output("ipfilter")]
        public Output<bool?> Ipfilter { get; private set; } = null!;

        /// <summary>
        /// Log level for incoming
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Output("logLevelIn")]
        public Output<string?> LogLevelIn { get; private set; } = null!;

        /// <summary>
        /// Log level for outgoing
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Output("logLevelOut")]
        public Output<string?> LogLevelOut { get; private set; } = null!;

        /// <summary>
        /// Enable/disable MAC address filter.
        /// </summary>
        [Output("macfilter")]
        public Output<bool?> Macfilter { get; private set; } = null!;

        /// <summary>
        /// Enable NDP (Neighbor Discovery Protocol).
        /// </summary>
        [Output("ndp")]
        public Output<bool?> Ndp { get; private set; } = null!;

        /// <summary>
        /// Node name.
        /// </summary>
        [Output("nodeName")]
        public Output<string> NodeName { get; private set; } = null!;

        /// <summary>
        /// The default output
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Output("outputPolicy")]
        public Output<string?> OutputPolicy { get; private set; } = null!;

        /// <summary>
        /// Enable Router Advertisement.
        /// </summary>
        [Output("radv")]
        public Output<bool?> Radv { get; private set; } = null!;

        /// <summary>
        /// VM ID. Leave empty for cluster level aliases.
        /// </summary>
        [Output("vmId")]
        public Output<int?> VmId { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallOptions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallOptions(string name, FirewallOptionsArgs args, CustomResourceOptions? options = null)
            : base("proxmoxve:Network/firewallOptions:FirewallOptions", name, args ?? new FirewallOptionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallOptions(string name, Input<string> id, FirewallOptionsState? state = null, CustomResourceOptions? options = null)
            : base("proxmoxve:Network/firewallOptions:FirewallOptions", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/muhlba91/pulumi-proxmoxve",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallOptions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallOptions Get(string name, Input<string> id, FirewallOptionsState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallOptions(name, id, state, options);
        }
    }

    public sealed class FirewallOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container ID. Leave empty for cluster level aliases.
        /// </summary>
        [Input("containerId")]
        public Input<int>? ContainerId { get; set; }

        /// <summary>
        /// Enable DHCP.
        /// </summary>
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        /// <summary>
        /// Enable or disable the firewall.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The default input
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Input("inputPolicy")]
        public Input<string>? InputPolicy { get; set; }

        /// <summary>
        /// Enable default IP filters. This is equivalent to
        /// adding an empty `ipfilter-net&lt;id&gt;` ipset for every interface. Such ipsets
        /// implicitly contain sane default restrictions such as restricting IPv6 link
        /// local addresses to the one derived from the interface's MAC address. For
        /// containers the configured IP addresses will be implicitly added.
        /// </summary>
        [Input("ipfilter")]
        public Input<bool>? Ipfilter { get; set; }

        /// <summary>
        /// Log level for incoming
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Input("logLevelIn")]
        public Input<string>? LogLevelIn { get; set; }

        /// <summary>
        /// Log level for outgoing
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Input("logLevelOut")]
        public Input<string>? LogLevelOut { get; set; }

        /// <summary>
        /// Enable/disable MAC address filter.
        /// </summary>
        [Input("macfilter")]
        public Input<bool>? Macfilter { get; set; }

        /// <summary>
        /// Enable NDP (Neighbor Discovery Protocol).
        /// </summary>
        [Input("ndp")]
        public Input<bool>? Ndp { get; set; }

        /// <summary>
        /// Node name.
        /// </summary>
        [Input("nodeName", required: true)]
        public Input<string> NodeName { get; set; } = null!;

        /// <summary>
        /// The default output
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Input("outputPolicy")]
        public Input<string>? OutputPolicy { get; set; }

        /// <summary>
        /// Enable Router Advertisement.
        /// </summary>
        [Input("radv")]
        public Input<bool>? Radv { get; set; }

        /// <summary>
        /// VM ID. Leave empty for cluster level aliases.
        /// </summary>
        [Input("vmId")]
        public Input<int>? VmId { get; set; }

        public FirewallOptionsArgs()
        {
        }
        public static new FirewallOptionsArgs Empty => new FirewallOptionsArgs();
    }

    public sealed class FirewallOptionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container ID. Leave empty for cluster level aliases.
        /// </summary>
        [Input("containerId")]
        public Input<int>? ContainerId { get; set; }

        /// <summary>
        /// Enable DHCP.
        /// </summary>
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        /// <summary>
        /// Enable or disable the firewall.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The default input
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Input("inputPolicy")]
        public Input<string>? InputPolicy { get; set; }

        /// <summary>
        /// Enable default IP filters. This is equivalent to
        /// adding an empty `ipfilter-net&lt;id&gt;` ipset for every interface. Such ipsets
        /// implicitly contain sane default restrictions such as restricting IPv6 link
        /// local addresses to the one derived from the interface's MAC address. For
        /// containers the configured IP addresses will be implicitly added.
        /// </summary>
        [Input("ipfilter")]
        public Input<bool>? Ipfilter { get; set; }

        /// <summary>
        /// Log level for incoming
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Input("logLevelIn")]
        public Input<string>? LogLevelIn { get; set; }

        /// <summary>
        /// Log level for outgoing
        /// packets (`emerg`, `alert`, `crit`, `err`, `warning`, `notice`, `info`,
        /// `debug`, `nolog`).
        /// </summary>
        [Input("logLevelOut")]
        public Input<string>? LogLevelOut { get; set; }

        /// <summary>
        /// Enable/disable MAC address filter.
        /// </summary>
        [Input("macfilter")]
        public Input<bool>? Macfilter { get; set; }

        /// <summary>
        /// Enable NDP (Neighbor Discovery Protocol).
        /// </summary>
        [Input("ndp")]
        public Input<bool>? Ndp { get; set; }

        /// <summary>
        /// Node name.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// The default output
        /// policy (`ACCEPT`, `DROP`, `REJECT`).
        /// </summary>
        [Input("outputPolicy")]
        public Input<string>? OutputPolicy { get; set; }

        /// <summary>
        /// Enable Router Advertisement.
        /// </summary>
        [Input("radv")]
        public Input<bool>? Radv { get; set; }

        /// <summary>
        /// VM ID. Leave empty for cluster level aliases.
        /// </summary>
        [Input("vmId")]
        public Input<int>? VmId { get; set; }

        public FirewallOptionsState()
        {
        }
        public static new FirewallOptionsState Empty => new FirewallOptionsState();
    }
}
