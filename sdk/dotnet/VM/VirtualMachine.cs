// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.VM
{
    [ProxmoxVEResourceType("proxmoxve:VM/virtualMachine:VirtualMachine")]
    public partial class VirtualMachine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable ACPI
        /// </summary>
        [Output("acpi")]
        public Output<bool?> Acpi { get; private set; } = null!;

        /// <summary>
        /// The QEMU agent configuration
        /// </summary>
        [Output("agent")]
        public Output<Outputs.VirtualMachineAgent?> Agent { get; private set; } = null!;

        /// <summary>
        /// The audio devices
        /// </summary>
        [Output("audioDevice")]
        public Output<Outputs.VirtualMachineAudioDevice?> AudioDevice { get; private set; } = null!;

        /// <summary>
        /// The BIOS implementation
        /// </summary>
        [Output("bios")]
        public Output<string?> Bios { get; private set; } = null!;

        /// <summary>
        /// The CDROM drive
        /// </summary>
        [Output("cdrom")]
        public Output<Outputs.VirtualMachineCdrom?> Cdrom { get; private set; } = null!;

        /// <summary>
        /// The cloning configuration
        /// </summary>
        [Output("clone")]
        public Output<Outputs.VirtualMachineClone?> Clone { get; private set; } = null!;

        /// <summary>
        /// The CPU allocation
        /// </summary>
        [Output("cpu")]
        public Output<Outputs.VirtualMachineCpu?> Cpu { get; private set; } = null!;

        /// <summary>
        /// The description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The disk devices
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.VirtualMachineDisk>> Disks { get; private set; } = null!;

        /// <summary>
        /// The cloud-init configuration
        /// </summary>
        [Output("initialization")]
        public Output<Outputs.VirtualMachineInitialization?> Initialization { get; private set; } = null!;

        /// <summary>
        /// The IPv4 addresses published by the QEMU agent
        /// </summary>
        [Output("ipv4Addresses")]
        public Output<ImmutableArray<ImmutableArray<string>>> Ipv4Addresses { get; private set; } = null!;

        /// <summary>
        /// The IPv6 addresses published by the QEMU agent
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<ImmutableArray<string>>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// The keyboard layout
        /// </summary>
        [Output("keyboardLayout")]
        public Output<string?> KeyboardLayout { get; private set; } = null!;

        /// <summary>
        /// The MAC addresses for the network interfaces
        /// </summary>
        [Output("macAddresses")]
        public Output<ImmutableArray<string>> MacAddresses { get; private set; } = null!;

        /// <summary>
        /// The memory allocation
        /// </summary>
        [Output("memory")]
        public Output<Outputs.VirtualMachineMemory?> Memory { get; private set; } = null!;

        /// <summary>
        /// The name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network devices
        /// </summary>
        [Output("networkDevices")]
        public Output<ImmutableArray<Outputs.VirtualMachineNetworkDevice>> NetworkDevices { get; private set; } = null!;

        /// <summary>
        /// The network interface names published by the QEMU agent
        /// </summary>
        [Output("networkInterfaceNames")]
        public Output<ImmutableArray<string>> NetworkInterfaceNames { get; private set; } = null!;

        /// <summary>
        /// The node name
        /// </summary>
        [Output("nodeName")]
        public Output<string> NodeName { get; private set; } = null!;

        /// <summary>
        /// Start VM on Node boot
        /// </summary>
        [Output("onBoot")]
        public Output<bool?> OnBoot { get; private set; } = null!;

        /// <summary>
        /// The operating system configuration
        /// </summary>
        [Output("operatingSystem")]
        public Output<Outputs.VirtualMachineOperatingSystem?> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// The ID of the pool to assign the virtual machine to
        /// </summary>
        [Output("poolId")]
        public Output<string?> PoolId { get; private set; } = null!;

        /// <summary>
        /// Whether to reboot vm after creation
        /// </summary>
        [Output("reboot")]
        public Output<bool?> Reboot { get; private set; } = null!;

        /// <summary>
        /// The serial devices
        /// </summary>
        [Output("serialDevices")]
        public Output<ImmutableArray<Outputs.VirtualMachineSerialDevice>> SerialDevices { get; private set; } = null!;

        /// <summary>
        /// Whether to start the virtual machine
        /// </summary>
        [Output("started")]
        public Output<bool?> Started { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the USB tablet device
        /// </summary>
        [Output("tabletDevice")]
        public Output<bool?> TabletDevice { get; private set; } = null!;

        /// <summary>
        /// Whether to create a template
        /// </summary>
        [Output("template")]
        public Output<bool?> Template { get; private set; } = null!;

        /// <summary>
        /// Clone VM timeout
        /// </summary>
        [Output("timeoutClone")]
        public Output<int?> TimeoutClone { get; private set; } = null!;

        /// <summary>
        /// MoveDisk timeout
        /// </summary>
        [Output("timeoutMoveDisk")]
        public Output<int?> TimeoutMoveDisk { get; private set; } = null!;

        /// <summary>
        /// Reboot timeout
        /// </summary>
        [Output("timeoutReboot")]
        public Output<int?> TimeoutReboot { get; private set; } = null!;

        /// <summary>
        /// Shutdown timeout
        /// </summary>
        [Output("timeoutShutdownVm")]
        public Output<int?> TimeoutShutdownVm { get; private set; } = null!;

        /// <summary>
        /// Start VM timeout
        /// </summary>
        [Output("timeoutStartVm")]
        public Output<int?> TimeoutStartVm { get; private set; } = null!;

        /// <summary>
        /// Stop VM timeout
        /// </summary>
        [Output("timeoutStopVm")]
        public Output<int?> TimeoutStopVm { get; private set; } = null!;

        /// <summary>
        /// The VGA configuration
        /// </summary>
        [Output("vga")]
        public Output<Outputs.VirtualMachineVga?> Vga { get; private set; } = null!;

        /// <summary>
        /// The VM identifier
        /// </summary>
        [Output("vmId")]
        public Output<int?> VmId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("proxmoxve:VM/virtualMachine:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
            : base("proxmoxve:VM/virtualMachine:VirtualMachine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, VirtualMachineState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, state, options);
        }
    }

    public sealed class VirtualMachineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable ACPI
        /// </summary>
        [Input("acpi")]
        public Input<bool>? Acpi { get; set; }

        /// <summary>
        /// The QEMU agent configuration
        /// </summary>
        [Input("agent")]
        public Input<Inputs.VirtualMachineAgentArgs>? Agent { get; set; }

        /// <summary>
        /// The audio devices
        /// </summary>
        [Input("audioDevice")]
        public Input<Inputs.VirtualMachineAudioDeviceArgs>? AudioDevice { get; set; }

        /// <summary>
        /// The BIOS implementation
        /// </summary>
        [Input("bios")]
        public Input<string>? Bios { get; set; }

        /// <summary>
        /// The CDROM drive
        /// </summary>
        [Input("cdrom")]
        public Input<Inputs.VirtualMachineCdromArgs>? Cdrom { get; set; }

        /// <summary>
        /// The cloning configuration
        /// </summary>
        [Input("clone")]
        public Input<Inputs.VirtualMachineCloneArgs>? Clone { get; set; }

        /// <summary>
        /// The CPU allocation
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.VirtualMachineCpuArgs>? Cpu { get; set; }

        /// <summary>
        /// The description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.VirtualMachineDiskArgs>? _disks;

        /// <summary>
        /// The disk devices
        /// </summary>
        public InputList<Inputs.VirtualMachineDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.VirtualMachineDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// The cloud-init configuration
        /// </summary>
        [Input("initialization")]
        public Input<Inputs.VirtualMachineInitializationArgs>? Initialization { get; set; }

        /// <summary>
        /// The keyboard layout
        /// </summary>
        [Input("keyboardLayout")]
        public Input<string>? KeyboardLayout { get; set; }

        /// <summary>
        /// The memory allocation
        /// </summary>
        [Input("memory")]
        public Input<Inputs.VirtualMachineMemoryArgs>? Memory { get; set; }

        /// <summary>
        /// The name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkDevices")]
        private InputList<Inputs.VirtualMachineNetworkDeviceArgs>? _networkDevices;

        /// <summary>
        /// The network devices
        /// </summary>
        public InputList<Inputs.VirtualMachineNetworkDeviceArgs> NetworkDevices
        {
            get => _networkDevices ?? (_networkDevices = new InputList<Inputs.VirtualMachineNetworkDeviceArgs>());
            set => _networkDevices = value;
        }

        /// <summary>
        /// The node name
        /// </summary>
        [Input("nodeName", required: true)]
        public Input<string> NodeName { get; set; } = null!;

        /// <summary>
        /// Start VM on Node boot
        /// </summary>
        [Input("onBoot")]
        public Input<bool>? OnBoot { get; set; }

        /// <summary>
        /// The operating system configuration
        /// </summary>
        [Input("operatingSystem")]
        public Input<Inputs.VirtualMachineOperatingSystemArgs>? OperatingSystem { get; set; }

        /// <summary>
        /// The ID of the pool to assign the virtual machine to
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// Whether to reboot vm after creation
        /// </summary>
        [Input("reboot")]
        public Input<bool>? Reboot { get; set; }

        [Input("serialDevices")]
        private InputList<Inputs.VirtualMachineSerialDeviceArgs>? _serialDevices;

        /// <summary>
        /// The serial devices
        /// </summary>
        public InputList<Inputs.VirtualMachineSerialDeviceArgs> SerialDevices
        {
            get => _serialDevices ?? (_serialDevices = new InputList<Inputs.VirtualMachineSerialDeviceArgs>());
            set => _serialDevices = value;
        }

        /// <summary>
        /// Whether to start the virtual machine
        /// </summary>
        [Input("started")]
        public Input<bool>? Started { get; set; }

        /// <summary>
        /// Whether to enable the USB tablet device
        /// </summary>
        [Input("tabletDevice")]
        public Input<bool>? TabletDevice { get; set; }

        /// <summary>
        /// Whether to create a template
        /// </summary>
        [Input("template")]
        public Input<bool>? Template { get; set; }

        /// <summary>
        /// Clone VM timeout
        /// </summary>
        [Input("timeoutClone")]
        public Input<int>? TimeoutClone { get; set; }

        /// <summary>
        /// MoveDisk timeout
        /// </summary>
        [Input("timeoutMoveDisk")]
        public Input<int>? TimeoutMoveDisk { get; set; }

        /// <summary>
        /// Reboot timeout
        /// </summary>
        [Input("timeoutReboot")]
        public Input<int>? TimeoutReboot { get; set; }

        /// <summary>
        /// Shutdown timeout
        /// </summary>
        [Input("timeoutShutdownVm")]
        public Input<int>? TimeoutShutdownVm { get; set; }

        /// <summary>
        /// Start VM timeout
        /// </summary>
        [Input("timeoutStartVm")]
        public Input<int>? TimeoutStartVm { get; set; }

        /// <summary>
        /// Stop VM timeout
        /// </summary>
        [Input("timeoutStopVm")]
        public Input<int>? TimeoutStopVm { get; set; }

        /// <summary>
        /// The VGA configuration
        /// </summary>
        [Input("vga")]
        public Input<Inputs.VirtualMachineVgaArgs>? Vga { get; set; }

        /// <summary>
        /// The VM identifier
        /// </summary>
        [Input("vmId")]
        public Input<int>? VmId { get; set; }

        public VirtualMachineArgs()
        {
        }
        public static new VirtualMachineArgs Empty => new VirtualMachineArgs();
    }

    public sealed class VirtualMachineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable ACPI
        /// </summary>
        [Input("acpi")]
        public Input<bool>? Acpi { get; set; }

        /// <summary>
        /// The QEMU agent configuration
        /// </summary>
        [Input("agent")]
        public Input<Inputs.VirtualMachineAgentGetArgs>? Agent { get; set; }

        /// <summary>
        /// The audio devices
        /// </summary>
        [Input("audioDevice")]
        public Input<Inputs.VirtualMachineAudioDeviceGetArgs>? AudioDevice { get; set; }

        /// <summary>
        /// The BIOS implementation
        /// </summary>
        [Input("bios")]
        public Input<string>? Bios { get; set; }

        /// <summary>
        /// The CDROM drive
        /// </summary>
        [Input("cdrom")]
        public Input<Inputs.VirtualMachineCdromGetArgs>? Cdrom { get; set; }

        /// <summary>
        /// The cloning configuration
        /// </summary>
        [Input("clone")]
        public Input<Inputs.VirtualMachineCloneGetArgs>? Clone { get; set; }

        /// <summary>
        /// The CPU allocation
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.VirtualMachineCpuGetArgs>? Cpu { get; set; }

        /// <summary>
        /// The description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disks")]
        private InputList<Inputs.VirtualMachineDiskGetArgs>? _disks;

        /// <summary>
        /// The disk devices
        /// </summary>
        public InputList<Inputs.VirtualMachineDiskGetArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.VirtualMachineDiskGetArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// The cloud-init configuration
        /// </summary>
        [Input("initialization")]
        public Input<Inputs.VirtualMachineInitializationGetArgs>? Initialization { get; set; }

        [Input("ipv4Addresses")]
        private InputList<ImmutableArray<string>>? _ipv4Addresses;

        /// <summary>
        /// The IPv4 addresses published by the QEMU agent
        /// </summary>
        public InputList<ImmutableArray<string>> Ipv4Addresses
        {
            get => _ipv4Addresses ?? (_ipv4Addresses = new InputList<ImmutableArray<string>>());
            set => _ipv4Addresses = value;
        }

        [Input("ipv6Addresses")]
        private InputList<ImmutableArray<string>>? _ipv6Addresses;

        /// <summary>
        /// The IPv6 addresses published by the QEMU agent
        /// </summary>
        public InputList<ImmutableArray<string>> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<ImmutableArray<string>>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// The keyboard layout
        /// </summary>
        [Input("keyboardLayout")]
        public Input<string>? KeyboardLayout { get; set; }

        [Input("macAddresses")]
        private InputList<string>? _macAddresses;

        /// <summary>
        /// The MAC addresses for the network interfaces
        /// </summary>
        public InputList<string> MacAddresses
        {
            get => _macAddresses ?? (_macAddresses = new InputList<string>());
            set => _macAddresses = value;
        }

        /// <summary>
        /// The memory allocation
        /// </summary>
        [Input("memory")]
        public Input<Inputs.VirtualMachineMemoryGetArgs>? Memory { get; set; }

        /// <summary>
        /// The name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkDevices")]
        private InputList<Inputs.VirtualMachineNetworkDeviceGetArgs>? _networkDevices;

        /// <summary>
        /// The network devices
        /// </summary>
        public InputList<Inputs.VirtualMachineNetworkDeviceGetArgs> NetworkDevices
        {
            get => _networkDevices ?? (_networkDevices = new InputList<Inputs.VirtualMachineNetworkDeviceGetArgs>());
            set => _networkDevices = value;
        }

        [Input("networkInterfaceNames")]
        private InputList<string>? _networkInterfaceNames;

        /// <summary>
        /// The network interface names published by the QEMU agent
        /// </summary>
        public InputList<string> NetworkInterfaceNames
        {
            get => _networkInterfaceNames ?? (_networkInterfaceNames = new InputList<string>());
            set => _networkInterfaceNames = value;
        }

        /// <summary>
        /// The node name
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// Start VM on Node boot
        /// </summary>
        [Input("onBoot")]
        public Input<bool>? OnBoot { get; set; }

        /// <summary>
        /// The operating system configuration
        /// </summary>
        [Input("operatingSystem")]
        public Input<Inputs.VirtualMachineOperatingSystemGetArgs>? OperatingSystem { get; set; }

        /// <summary>
        /// The ID of the pool to assign the virtual machine to
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// Whether to reboot vm after creation
        /// </summary>
        [Input("reboot")]
        public Input<bool>? Reboot { get; set; }

        [Input("serialDevices")]
        private InputList<Inputs.VirtualMachineSerialDeviceGetArgs>? _serialDevices;

        /// <summary>
        /// The serial devices
        /// </summary>
        public InputList<Inputs.VirtualMachineSerialDeviceGetArgs> SerialDevices
        {
            get => _serialDevices ?? (_serialDevices = new InputList<Inputs.VirtualMachineSerialDeviceGetArgs>());
            set => _serialDevices = value;
        }

        /// <summary>
        /// Whether to start the virtual machine
        /// </summary>
        [Input("started")]
        public Input<bool>? Started { get; set; }

        /// <summary>
        /// Whether to enable the USB tablet device
        /// </summary>
        [Input("tabletDevice")]
        public Input<bool>? TabletDevice { get; set; }

        /// <summary>
        /// Whether to create a template
        /// </summary>
        [Input("template")]
        public Input<bool>? Template { get; set; }

        /// <summary>
        /// Clone VM timeout
        /// </summary>
        [Input("timeoutClone")]
        public Input<int>? TimeoutClone { get; set; }

        /// <summary>
        /// MoveDisk timeout
        /// </summary>
        [Input("timeoutMoveDisk")]
        public Input<int>? TimeoutMoveDisk { get; set; }

        /// <summary>
        /// Reboot timeout
        /// </summary>
        [Input("timeoutReboot")]
        public Input<int>? TimeoutReboot { get; set; }

        /// <summary>
        /// Shutdown timeout
        /// </summary>
        [Input("timeoutShutdownVm")]
        public Input<int>? TimeoutShutdownVm { get; set; }

        /// <summary>
        /// Start VM timeout
        /// </summary>
        [Input("timeoutStartVm")]
        public Input<int>? TimeoutStartVm { get; set; }

        /// <summary>
        /// Stop VM timeout
        /// </summary>
        [Input("timeoutStopVm")]
        public Input<int>? TimeoutStopVm { get; set; }

        /// <summary>
        /// The VGA configuration
        /// </summary>
        [Input("vga")]
        public Input<Inputs.VirtualMachineVgaGetArgs>? Vga { get; set; }

        /// <summary>
        /// The VM identifier
        /// </summary>
        [Input("vmId")]
        public Input<int>? VmId { get; set; }

        public VirtualMachineState()
        {
        }
        public static new VirtualMachineState Empty => new VirtualMachineState();
    }
}
