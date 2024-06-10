// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.muehlbachler.pulumi.proxmoxve.VM.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.muehlbachler.pulumi.proxmoxve.VM.inputs.VirtualMachine2CdromArgs;
import io.muehlbachler.pulumi.proxmoxve.VM.inputs.VirtualMachine2CloneArgs;
import io.muehlbachler.pulumi.proxmoxve.VM.inputs.VirtualMachine2CpuArgs;
import io.muehlbachler.pulumi.proxmoxve.VM.inputs.VirtualMachine2TimeoutsArgs;
import io.muehlbachler.pulumi.proxmoxve.VM.inputs.VirtualMachine2VgaArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachine2State extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachine2State Empty = new VirtualMachine2State();

    /**
     * The CD-ROM configuration. The key is the interface of the CD-ROM, could be one of `ideN`, `sataN`, `scsiN`, where N is the index of the interface. Note that `q35` machine type only supports `ide0` and `ide2` of IDE interfaces.
     * 
     */
    @Import(name="cdrom")
    private @Nullable Output<Map<String,VirtualMachine2CdromArgs>> cdrom;

    /**
     * @return The CD-ROM configuration. The key is the interface of the CD-ROM, could be one of `ideN`, `sataN`, `scsiN`, where N is the index of the interface. Note that `q35` machine type only supports `ide0` and `ide2` of IDE interfaces.
     * 
     */
    public Optional<Output<Map<String,VirtualMachine2CdromArgs>>> cdrom() {
        return Optional.ofNullable(this.cdrom);
    }

    /**
     * The cloning configuration.
     * 
     */
    @Import(name="clone")
    private @Nullable Output<VirtualMachine2CloneArgs> clone;

    /**
     * @return The cloning configuration.
     * 
     */
    public Optional<Output<VirtualMachine2CloneArgs>> clone_() {
        return Optional.ofNullable(this.clone);
    }

    /**
     * The CPU configuration.
     * 
     */
    @Import(name="cpu")
    private @Nullable Output<VirtualMachine2CpuArgs> cpu;

    /**
     * @return The CPU configuration.
     * 
     */
    public Optional<Output<VirtualMachine2CpuArgs>> cpu() {
        return Optional.ofNullable(this.cpu);
    }

    /**
     * The description of the VM.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the VM.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the VM. Doesn&#39;t have to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the VM. Doesn&#39;t have to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the node where the VM is provisioned.
     * 
     */
    @Import(name="nodeName")
    private @Nullable Output<String> nodeName;

    /**
     * @return The name of the node where the VM is provisioned.
     * 
     */
    public Optional<Output<String>> nodeName() {
        return Optional.ofNullable(this.nodeName);
    }

    /**
     * The tags assigned to the VM.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags assigned to the VM.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Set to true to create a VM template.
     * 
     */
    @Import(name="template")
    private @Nullable Output<Boolean> template;

    /**
     * @return Set to true to create a VM template.
     * 
     */
    public Optional<Output<Boolean>> template() {
        return Optional.ofNullable(this.template);
    }

    @Import(name="timeouts")
    private @Nullable Output<VirtualMachine2TimeoutsArgs> timeouts;

    public Optional<Output<VirtualMachine2TimeoutsArgs>> timeouts() {
        return Optional.ofNullable(this.timeouts);
    }

    /**
     * Configure the VGA Hardware. If you want to use high resolution modes (&gt;= 1280x1024x16) you may need to increase the vga memory option. Since QEMU 2.9 the default VGA display type is `std` for all OS types besides some Windows versions (XP and older) which use `cirrus`. The `qxl` option enables the SPICE display server. For win* OS you can select how many independent displays you want, Linux guests can add displays themself. You can also run without any graphic card, using a serial device as terminal. See the [Proxmox documentation](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#qm_virtual_machines_settings) section 10.2.8 for more information and available configuration parameters.
     * 
     */
    @Import(name="vga")
    private @Nullable Output<VirtualMachine2VgaArgs> vga;

    /**
     * @return Configure the VGA Hardware. If you want to use high resolution modes (&gt;= 1280x1024x16) you may need to increase the vga memory option. Since QEMU 2.9 the default VGA display type is `std` for all OS types besides some Windows versions (XP and older) which use `cirrus`. The `qxl` option enables the SPICE display server. For win* OS you can select how many independent displays you want, Linux guests can add displays themself. You can also run without any graphic card, using a serial device as terminal. See the [Proxmox documentation](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#qm_virtual_machines_settings) section 10.2.8 for more information and available configuration parameters.
     * 
     */
    public Optional<Output<VirtualMachine2VgaArgs>> vga() {
        return Optional.ofNullable(this.vga);
    }

    private VirtualMachine2State() {}

    private VirtualMachine2State(VirtualMachine2State $) {
        this.cdrom = $.cdrom;
        this.clone = $.clone;
        this.cpu = $.cpu;
        this.description = $.description;
        this.name = $.name;
        this.nodeName = $.nodeName;
        this.tags = $.tags;
        this.template = $.template;
        this.timeouts = $.timeouts;
        this.vga = $.vga;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachine2State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachine2State $;

        public Builder() {
            $ = new VirtualMachine2State();
        }

        public Builder(VirtualMachine2State defaults) {
            $ = new VirtualMachine2State(Objects.requireNonNull(defaults));
        }

        /**
         * @param cdrom The CD-ROM configuration. The key is the interface of the CD-ROM, could be one of `ideN`, `sataN`, `scsiN`, where N is the index of the interface. Note that `q35` machine type only supports `ide0` and `ide2` of IDE interfaces.
         * 
         * @return builder
         * 
         */
        public Builder cdrom(@Nullable Output<Map<String,VirtualMachine2CdromArgs>> cdrom) {
            $.cdrom = cdrom;
            return this;
        }

        /**
         * @param cdrom The CD-ROM configuration. The key is the interface of the CD-ROM, could be one of `ideN`, `sataN`, `scsiN`, where N is the index of the interface. Note that `q35` machine type only supports `ide0` and `ide2` of IDE interfaces.
         * 
         * @return builder
         * 
         */
        public Builder cdrom(Map<String,VirtualMachine2CdromArgs> cdrom) {
            return cdrom(Output.of(cdrom));
        }

        /**
         * @param clone The cloning configuration.
         * 
         * @return builder
         * 
         */
        public Builder clone_(@Nullable Output<VirtualMachine2CloneArgs> clone) {
            $.clone = clone;
            return this;
        }

        /**
         * @param clone The cloning configuration.
         * 
         * @return builder
         * 
         */
        public Builder clone_(VirtualMachine2CloneArgs clone) {
            return clone_(Output.of(clone));
        }

        /**
         * @param cpu The CPU configuration.
         * 
         * @return builder
         * 
         */
        public Builder cpu(@Nullable Output<VirtualMachine2CpuArgs> cpu) {
            $.cpu = cpu;
            return this;
        }

        /**
         * @param cpu The CPU configuration.
         * 
         * @return builder
         * 
         */
        public Builder cpu(VirtualMachine2CpuArgs cpu) {
            return cpu(Output.of(cpu));
        }

        /**
         * @param description The description of the VM.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the VM.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the VM. Doesn&#39;t have to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the VM. Doesn&#39;t have to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodeName The name of the node where the VM is provisioned.
         * 
         * @return builder
         * 
         */
        public Builder nodeName(@Nullable Output<String> nodeName) {
            $.nodeName = nodeName;
            return this;
        }

        /**
         * @param nodeName The name of the node where the VM is provisioned.
         * 
         * @return builder
         * 
         */
        public Builder nodeName(String nodeName) {
            return nodeName(Output.of(nodeName));
        }

        /**
         * @param tags The tags assigned to the VM.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags assigned to the VM.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags assigned to the VM.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param template Set to true to create a VM template.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<Boolean> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template Set to true to create a VM template.
         * 
         * @return builder
         * 
         */
        public Builder template(Boolean template) {
            return template(Output.of(template));
        }

        public Builder timeouts(@Nullable Output<VirtualMachine2TimeoutsArgs> timeouts) {
            $.timeouts = timeouts;
            return this;
        }

        public Builder timeouts(VirtualMachine2TimeoutsArgs timeouts) {
            return timeouts(Output.of(timeouts));
        }

        /**
         * @param vga Configure the VGA Hardware. If you want to use high resolution modes (&gt;= 1280x1024x16) you may need to increase the vga memory option. Since QEMU 2.9 the default VGA display type is `std` for all OS types besides some Windows versions (XP and older) which use `cirrus`. The `qxl` option enables the SPICE display server. For win* OS you can select how many independent displays you want, Linux guests can add displays themself. You can also run without any graphic card, using a serial device as terminal. See the [Proxmox documentation](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#qm_virtual_machines_settings) section 10.2.8 for more information and available configuration parameters.
         * 
         * @return builder
         * 
         */
        public Builder vga(@Nullable Output<VirtualMachine2VgaArgs> vga) {
            $.vga = vga;
            return this;
        }

        /**
         * @param vga Configure the VGA Hardware. If you want to use high resolution modes (&gt;= 1280x1024x16) you may need to increase the vga memory option. Since QEMU 2.9 the default VGA display type is `std` for all OS types besides some Windows versions (XP and older) which use `cirrus`. The `qxl` option enables the SPICE display server. For win* OS you can select how many independent displays you want, Linux guests can add displays themself. You can also run without any graphic card, using a serial device as terminal. See the [Proxmox documentation](https://pve.proxmox.com/pve-docs/pve-admin-guide.html#qm_virtual_machines_settings) section 10.2.8 for more information and available configuration parameters.
         * 
         * @return builder
         * 
         */
        public Builder vga(VirtualMachine2VgaArgs vga) {
            return vga(Output.of(vga));
        }

        public VirtualMachine2State build() {
            return $;
        }
    }

}
