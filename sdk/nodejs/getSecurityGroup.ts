// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Fetch Exoscale [Security Groups](https://community.exoscale.com/documentation/compute/security-groups/) data.
 *
 * Corresponding resource: exoscale_security_group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const mySecurityGroup = exoscale.getSecurityGroup({
 *     name: "my-security-group",
 * });
 * export const mySecurityGroupId = mySecurityGroup.then(mySecurityGroup => mySecurityGroup.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getSecurityGroup(args?: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSecurityGroup:getSecurityGroup", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupArgs {
    /**
     * The security group ID to match (conflicts with `name`)
     */
    id?: string;
    /**
     * The name to match (conflicts with `id`)
     */
    name?: string;
}

/**
 * A collection of values returned by getSecurityGroup.
 */
export interface GetSecurityGroupResult {
    /**
     * The list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
     */
    readonly externalSources: string[];
    /**
     * The security group ID to match (conflicts with `name`)
     */
    readonly id?: string;
    /**
     * The name to match (conflicts with `id`)
     */
    readonly name?: string;
}
/**
 * Fetch Exoscale [Security Groups](https://community.exoscale.com/documentation/compute/security-groups/) data.
 *
 * Corresponding resource: exoscale_security_group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const mySecurityGroup = exoscale.getSecurityGroup({
 *     name: "my-security-group",
 * });
 * export const mySecurityGroupId = mySecurityGroup.then(mySecurityGroup => mySecurityGroup.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getSecurityGroupOutput(args?: GetSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getSecurityGroup:getSecurityGroup", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupOutputArgs {
    /**
     * The security group ID to match (conflicts with `name`)
     */
    id?: pulumi.Input<string>;
    /**
     * The name to match (conflicts with `id`)
     */
    name?: pulumi.Input<string>;
}
