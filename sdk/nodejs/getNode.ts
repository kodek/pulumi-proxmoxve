// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves information about node.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const node = proxmoxve.getNode({});
 * ```
 */
export function getNode(args: GetNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetNodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("proxmoxve:index/getNode:getNode", {
        "nodeName": args.nodeName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNode.
 */
export interface GetNodeArgs {
    /**
     * The node name.
     */
    nodeName: string;
}

/**
 * A collection of values returned by getNode.
 */
export interface GetNodeResult {
    /**
     * The CPU count on the node.
     */
    readonly cpuCount: number;
    /**
     * The CPU model on the node.
     */
    readonly cpuModel: string;
    /**
     * The CPU utilization on the node.
     */
    readonly cpuSockets: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The memory available on the node.
     */
    readonly memoryAvailable: number;
    /**
     * The total memory on the node.
     */
    readonly memoryTotal: number;
    /**
     * The memory used on the node.
     */
    readonly memoryUsed: number;
    readonly nodeName: string;
    /**
     * The uptime in seconds on the node.
     */
    readonly uptime: number;
}
/**
 * Retrieves information about node.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const node = proxmoxve.getNode({});
 * ```
 */
export function getNodeOutput(args: GetNodeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("proxmoxve:index/getNode:getNode", {
        "nodeName": args.nodeName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNode.
 */
export interface GetNodeOutputArgs {
    /**
     * The node name.
     */
    nodeName: pulumi.Input<string>;
}
