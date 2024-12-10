// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Fetch Exoscale [Private Networks](https://community.exoscale.com/documentation/compute/private-networks/) data.
 *
 * Corresponding resource: exoscale_private_network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myPrivateNetwork = exoscale.getPrivateNetwork({
 *     zone: "ch-gva-2",
 *     name: "my-private-network",
 * });
 * export const myPrivateNetworkId = myPrivateNetwork.then(myPrivateNetwork => myPrivateNetwork.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getPrivateNetwork(args: GetPrivateNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getPrivateNetwork:getPrivateNetwork", {
        "description": args.description,
        "id": args.id,
        "labels": args.labels,
        "name": args.name,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateNetwork.
 */
export interface GetPrivateNetworkArgs {
    /**
     * The private network description.
     */
    description?: string;
    /**
     * The private network ID to match (conflicts with `name`).
     */
    id?: string;
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    /**
     * The network name to match (conflicts with `id`).
     */
    name?: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getPrivateNetwork.
 */
export interface GetPrivateNetworkResult {
    /**
     * The private network description.
     */
    readonly description?: string;
    /**
     * The first/last IPv4 addresses used by the DHCP service for dynamic leases.
     */
    readonly endIp: string;
    /**
     * The private network ID to match (conflicts with `name`).
     */
    readonly id?: string;
    /**
     * A map of key/value labels.
     */
    readonly labels?: {[key: string]: string};
    /**
     * The network name to match (conflicts with `id`).
     */
    readonly name?: string;
    /**
     * The network mask defining the IPv4 network allowed for static leases.
     */
    readonly netmask: string;
    /**
     * The first/last IPv4 addresses used by the DHCP service for dynamic leases.
     */
    readonly startIp: string;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * Fetch Exoscale [Private Networks](https://community.exoscale.com/documentation/compute/private-networks/) data.
 *
 * Corresponding resource: exoscale_private_network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myPrivateNetwork = exoscale.getPrivateNetwork({
 *     zone: "ch-gva-2",
 *     name: "my-private-network",
 * });
 * export const myPrivateNetworkId = myPrivateNetwork.then(myPrivateNetwork => myPrivateNetwork.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getPrivateNetworkOutput(args: GetPrivateNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrivateNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getPrivateNetwork:getPrivateNetwork", {
        "description": args.description,
        "id": args.id,
        "labels": args.labels,
        "name": args.name,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateNetwork.
 */
export interface GetPrivateNetworkOutputArgs {
    /**
     * The private network description.
     */
    description?: pulumi.Input<string>;
    /**
     * The private network ID to match (conflicts with `name`).
     */
    id?: pulumi.Input<string>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The network name to match (conflicts with `id`).
     */
    name?: pulumi.Input<string>;
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
