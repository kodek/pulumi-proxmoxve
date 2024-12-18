// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the list of ACME accounts.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const example = proxmoxve.Acme.getAccounts({});
 * export const dataProxmoxVirtualEnvironmentAcmeAccounts = example.then(example => example.accounts);
 * ```
 */
export function getAccounts(opts?: pulumi.InvokeOptions): Promise<GetAccountsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("proxmoxve:Acme/getAccounts:getAccounts", {
    }, opts);
}

/**
 * A collection of values returned by getAccounts.
 */
export interface GetAccountsResult {
    /**
     * The identifiers of the ACME accounts.
     */
    readonly accounts: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Retrieves the list of ACME accounts.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as proxmoxve from "@pulumi/proxmoxve";
 *
 * const example = proxmoxve.Acme.getAccounts({});
 * export const dataProxmoxVirtualEnvironmentAcmeAccounts = example.then(example => example.accounts);
 * ```
 */
export function getAccountsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccountsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("proxmoxve:Acme/getAccounts:getAccounts", {
    }, opts);
}
