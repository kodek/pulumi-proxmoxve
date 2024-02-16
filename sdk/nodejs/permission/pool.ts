// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a resource pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@muhlba91/pulumi-proxmoxve";
 *
 * const operationsPool = new proxmoxve.permission.Pool("operationsPool", {
 *     comment: "Managed by Terraform",
 *     poolId: "operations-pool",
 * });
 * ```
 *
 * ## Import
 *
 * Instances can be imported using the `pool_id`, e.g.,
 *
 *  bash
 *
 * ```sh
 * $ pulumi import proxmoxve:Permission/pool:Pool operations_pool operations-pool
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'proxmoxve:Permission/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * The pool comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The pool members.
     */
    public /*out*/ readonly members!: pulumi.Output<outputs.Permission.PoolMember[]>;
    /**
     * The pool identifier.
     */
    public readonly poolId!: pulumi.Output<string>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PoolState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["members"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * The pool comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * The pool members.
     */
    members?: pulumi.Input<pulumi.Input<inputs.Permission.PoolMember>[]>;
    /**
     * The pool identifier.
     */
    poolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The pool comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * The pool identifier.
     */
    poolId: pulumi.Input<string>;
}
