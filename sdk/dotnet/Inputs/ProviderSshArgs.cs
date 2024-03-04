// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.Inputs
{

    public sealed class ProviderSshArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use the SSH agent for authentication. Takes precedence over the `private_key` and `password` fields. Defaults to the value of the `PROXMOX_VE_SSH_AGENT` environment variable, or `false` if not set.
        /// </summary>
        [Input("agent")]
        public Input<bool>? Agent { get; set; }

        /// <summary>
        /// The path to the SSH agent socket. Defaults to the value of the `SSH_AUTH_SOCK` environment variable.
        /// </summary>
        [Input("agentSocket")]
        public Input<string>? AgentSocket { get; set; }

        [Input("nodes")]
        private InputList<Inputs.ProviderSshNodeArgs>? _nodes;

        /// <summary>
        /// Overrides for SSH connection configuration for a Proxmox VE node.
        /// </summary>
        public InputList<Inputs.ProviderSshNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.ProviderSshNodeArgs>());
            set => _nodes = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password used for the SSH connection. Defaults to the value of the `password` field of the `provider` block.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The unencrypted private key (in PEM format) used for the SSH connection. Defaults to the value of the `PROXMOX_VE_SSH_PRIVATE_KEY` environment variable.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("socks5Password")]
        private Input<string>? _socks5Password;

        /// <summary>
        /// The password for the SOCKS5 proxy server. Defaults to the value of the `PROXMOX_VE_SSH_SOCKS5_PASSWORD` environment variable.
        /// </summary>
        public Input<string>? Socks5Password
        {
            get => _socks5Password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _socks5Password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The address:port of the SOCKS5 proxy server. Defaults to the value of the `PROXMOX_VE_SSH_SOCKS5_SERVER` environment variable.
        /// </summary>
        [Input("socks5Server")]
        public Input<string>? Socks5Server { get; set; }

        /// <summary>
        /// The username for the SOCKS5 proxy server. Defaults to the value of the `PROXMOX_VE_SSH_SOCKS5_USERNAME` environment variable.
        /// </summary>
        [Input("socks5Username")]
        public Input<string>? Socks5Username { get; set; }

        /// <summary>
        /// The username used for the SSH connection. Defaults to the value of the `username` field of the `provider` block.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProviderSshArgs()
        {
        }
        public static new ProviderSshArgs Empty => new ProviderSshArgs();
    }
}
