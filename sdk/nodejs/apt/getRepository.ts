// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves an APT repository from a Proxmox VE cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const example = proxmoxve.Apt.getRepository({
 *     filePath: "/etc/apt/sources.list",
 *     index: 0,
 *     node: "pve",
 * });
 * export const proxmoxVirtualEnvironmentAptRepository = example;
 * ```
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("proxmoxve:Apt/getRepository:getRepository", {
        "filePath": args.filePath,
        "index": args.index,
        "node": args.node,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    /**
     * The absolute path of the source list file that contains this repository.
     */
    filePath: string;
    /**
     * The index within the defining source list file.
     */
    index: number;
    /**
     * The name of the target Proxmox VE node.
     */
    node: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    /**
     * The associated comment.
     */
    readonly comment: string;
    /**
     * The list of components.
     */
    readonly components: string[];
    /**
     * Indicates the activation status.
     */
    readonly enabled: boolean;
    /**
     * The absolute path of the source list file that contains this repository.
     */
    readonly filePath: string;
    /**
     * The format of the defining source list file.
     */
    readonly fileType: string;
    /**
     * The unique identifier of this APT repository data source.
     */
    readonly id: string;
    /**
     * The index within the defining source list file.
     */
    readonly index: number;
    /**
     * The name of the target Proxmox VE node.
     */
    readonly node: string;
    /**
     * The list of package types.
     */
    readonly packageTypes: string[];
    /**
     * The list of package distributions.
     */
    readonly suites: string[];
    /**
     * The list of repository URIs.
     */
    readonly uris: string[];
}
/**
 * Retrieves an APT repository from a Proxmox VE cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const example = proxmoxve.Apt.getRepository({
 *     filePath: "/etc/apt/sources.list",
 *     index: 0,
 *     node: "pve",
 * });
 * export const proxmoxVirtualEnvironmentAptRepository = example;
 * ```
 */
export function getRepositoryOutput(args: GetRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("proxmoxve:Apt/getRepository:getRepository", {
        "filePath": args.filePath,
        "index": args.index,
        "node": args.node,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryOutputArgs {
    /**
     * The absolute path of the source list file that contains this repository.
     */
    filePath: pulumi.Input<string>;
    /**
     * The index within the defining source list file.
     */
    index: pulumi.Input<number>;
    /**
     * The name of the target Proxmox VE node.
     */
    node: pulumi.Input<string>;
}
