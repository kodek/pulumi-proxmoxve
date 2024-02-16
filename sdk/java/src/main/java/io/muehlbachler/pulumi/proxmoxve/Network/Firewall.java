// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.muehlbachler.pulumi.proxmoxve.Network;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.muehlbachler.pulumi.proxmoxve.Network.FirewallArgs;
import io.muehlbachler.pulumi.proxmoxve.Network.inputs.FirewallState;
import io.muehlbachler.pulumi.proxmoxve.Network.outputs.FirewallLogRatelimit;
import io.muehlbachler.pulumi.proxmoxve.Utilities;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages firewall options on the cluster level.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.proxmoxve.Network.Firewall;
 * import com.pulumi.proxmoxve.Network.FirewallArgs;
 * import com.pulumi.proxmoxve.Network.inputs.FirewallLogRatelimitArgs;
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
 *         var example = new Firewall(&#34;example&#34;, FirewallArgs.builder()        
 *             .ebtables(false)
 *             .enabled(false)
 *             .inputPolicy(&#34;DROP&#34;)
 *             .logRatelimit(FirewallLogRatelimitArgs.builder()
 *                 .burst(10)
 *                 .enabled(false)
 *                 .rate(&#34;5/second&#34;)
 *                 .build())
 *             .outputPolicy(&#34;ACCEPT&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Important Notes
 * 
 * Be careful not to use this resource multiple times for the same node.
 * 
 * ## Import
 * 
 * Instances can be imported without an ID, but you still need to pass one, e.g.,
 * 
 *  bash
 * 
 * ```sh
 * $ pulumi import proxmoxve:Network/firewall:Firewall example example
 * ```
 * 
 */
@ResourceType(type="proxmoxve:Network/firewall:Firewall")
public class Firewall extends com.pulumi.resources.CustomResource {
    /**
     * Enable ebtables rules cluster wide.
     * 
     */
    @Export(name="ebtables", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ebtables;

    /**
     * @return Enable ebtables rules cluster wide.
     * 
     */
    public Output<Optional<Boolean>> ebtables() {
        return Codegen.optional(this.ebtables);
    }
    /**
     * Enable or disable the log rate limit.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Enable or disable the log rate limit.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The default input policy (`ACCEPT`, `DROP`, `REJECT`).
     * 
     */
    @Export(name="inputPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inputPolicy;

    /**
     * @return The default input policy (`ACCEPT`, `DROP`, `REJECT`).
     * 
     */
    public Output<Optional<String>> inputPolicy() {
        return Codegen.optional(this.inputPolicy);
    }
    /**
     * The log rate limit.
     * 
     */
    @Export(name="logRatelimit", refs={FirewallLogRatelimit.class}, tree="[0]")
    private Output</* @Nullable */ FirewallLogRatelimit> logRatelimit;

    /**
     * @return The log rate limit.
     * 
     */
    public Output<Optional<FirewallLogRatelimit>> logRatelimit() {
        return Codegen.optional(this.logRatelimit);
    }
    /**
     * The default output policy (`ACCEPT`, `DROP`, `REJECT`).
     * 
     */
    @Export(name="outputPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> outputPolicy;

    /**
     * @return The default output policy (`ACCEPT`, `DROP`, `REJECT`).
     * 
     */
    public Output<Optional<String>> outputPolicy() {
        return Codegen.optional(this.outputPolicy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Firewall(String name) {
        this(name, FirewallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Firewall(String name, @Nullable FirewallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Firewall(String name, @Nullable FirewallArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("proxmoxve:Network/firewall:Firewall", name, args == null ? FirewallArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Firewall(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("proxmoxve:Network/firewall:Firewall", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Firewall get(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Firewall(name, id, state, options);
    }
}
