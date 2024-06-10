// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ProxmoxVE.VM.Inputs
{

    public sealed class VirtualMachineDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The disk AIO mode (defaults to `io_uring`).
        /// </summary>
        [Input("aio")]
        public Input<string>? Aio { get; set; }

        /// <summary>
        /// Whether the drive should be included when making backups (defaults to `true`).
        /// </summary>
        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        /// <summary>
        /// The cache type (defaults to `none`).
        /// </summary>
        [Input("cache")]
        public Input<string>? Cache { get; set; }

        /// <summary>
        /// The identifier for the datastore to create
        /// the disk in (defaults to `local-lvm`).
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        /// <summary>
        /// Whether to pass discard/trim requests to the
        /// underlying storage. Supported values are `on`/`ignore` (defaults
        /// to `ignore`).
        /// </summary>
        [Input("discard")]
        public Input<string>? Discard { get; set; }

        /// <summary>
        /// The file format (defaults to `qcow2`).
        /// </summary>
        [Input("fileFormat")]
        public Input<string>? FileFormat { get; set; }

        /// <summary>
        /// The file ID for a disk image. The ID format is
        /// `&lt;datastore_id&gt;:&lt;content_type&gt;/&lt;file_name&gt;`, for example `local:iso/centos8.img`. Can be also taken from
        /// `proxmoxve.Download.File` resource.
        /// </summary>
        [Input("fileId")]
        public Input<string>? FileId { get; set; }

        /// <summary>
        /// The disk interface for Proxmox, currently `scsi`,
        /// `sata` and `virtio` interfaces are supported. Append the disk index at
        /// the end, for example, `virtio0` for the first virtio disk, `virtio1` for
        /// the second, etc.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Whether to use iothreads for this disk (defaults
        /// to `false`).
        /// </summary>
        [Input("iothread")]
        public Input<bool>? Iothread { get; set; }

        /// <summary>
        /// The in-datastore path to the disk image.
        /// ***Experimental.***Use to attach another VM's disks,
        /// or (as root only) host's filesystem paths (`datastore_id` empty string).
        /// See "*Example: Attached disks*".
        /// </summary>
        [Input("pathInDatastore")]
        public Input<string>? PathInDatastore { get; set; }

        /// <summary>
        /// Whether the drive should be considered for replication jobs (defaults to `true`).
        /// </summary>
        [Input("replicate")]
        public Input<bool>? Replicate { get; set; }

        /// <summary>
        /// The disk size in gigabytes (defaults to `8`).
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The speed limits.
        /// </summary>
        [Input("speed")]
        public Input<Inputs.VirtualMachineDiskSpeedArgs>? Speed { get; set; }

        /// <summary>
        /// Whether to use an SSD emulation option for this disk (
        /// defaults to `false`). Note that SSD emulation is not supported on VirtIO
        /// Block drives.
        /// </summary>
        [Input("ssd")]
        public Input<bool>? Ssd { get; set; }

        public VirtualMachineDiskArgs()
        {
        }
        public static new VirtualMachineDiskArgs Empty => new VirtualMachineDiskArgs();
    }
}
