// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.muehlbachler.pulumi.proxmoxve.HA;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.muehlbachler.pulumi.proxmoxve.HA.HAResourceArgs;
import io.muehlbachler.pulumi.proxmoxve.HA.inputs.HAResourceState;
import io.muehlbachler.pulumi.proxmoxve.Utilities;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Proxmox HA resources.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.proxmoxve.HA.HAResource;
 * import com.pulumi.proxmoxve.HA.HAResourceArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new HAResource("example", HAResourceArgs.builder()
 *             .resourceId("vm:123")
 *             .state("started")
 *             .group("example")
 *             .comment("Managed by Pulumi")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(proxmox_virtual_environment_hagroup.example())
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * #!/usr/bin/env sh
 * 
 * HA resources can be imported using their identifiers, e.g.:
 * 
 * ```sh
 * $ pulumi import proxmoxve:HA/hAResource:HAResource example vm:123
 * ```
 * 
 */
@ResourceType(type="proxmoxve:HA/hAResource:HAResource")
public class HAResource extends com.pulumi.resources.CustomResource {
    /**
     * The comment associated with this resource.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return The comment associated with this resource.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * The identifier of the High Availability group this resource is a member of.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> group;

    /**
     * @return The identifier of the High Availability group this resource is a member of.
     * 
     */
    public Output<Optional<String>> group() {
        return Codegen.optional(this.group);
    }
    /**
     * The maximal number of relocation attempts.
     * 
     */
    @Export(name="maxRelocate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxRelocate;

    /**
     * @return The maximal number of relocation attempts.
     * 
     */
    public Output<Optional<Integer>> maxRelocate() {
        return Codegen.optional(this.maxRelocate);
    }
    /**
     * The maximal number of restart attempts.
     * 
     */
    @Export(name="maxRestart", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxRestart;

    /**
     * @return The maximal number of restart attempts.
     * 
     */
    public Output<Optional<Integer>> maxRestart() {
        return Codegen.optional(this.maxRestart);
    }
    /**
     * The Proxmox HA resource identifier
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return The Proxmox HA resource identifier
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * The desired state of the resource.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The desired state of the resource.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The type of HA resources to create. If unset, it will be deduced from the `resource_id`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of HA resources to create. If unset, it will be deduced from the `resource_id`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HAResource(java.lang.String name) {
        this(name, HAResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HAResource(java.lang.String name, HAResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HAResource(java.lang.String name, HAResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("proxmoxve:HA/hAResource:HAResource", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HAResource(java.lang.String name, Output<java.lang.String> id, @Nullable HAResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("proxmoxve:HA/hAResource:HAResource", name, state, makeResourceOptions(options, id), false);
    }

    private static HAResourceArgs makeArgs(HAResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HAResourceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static HAResource get(java.lang.String name, Output<java.lang.String> id, @Nullable HAResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HAResource(name, id, state, options);
    }
}
