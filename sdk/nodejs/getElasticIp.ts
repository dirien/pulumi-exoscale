// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Fetch Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/) data.
 *
 * Corresponding resource: exoscale_elastic_ip.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myElasticIp = exoscale.getElasticIp({
 *     zone: "ch-gva-2",
 *     ipAddress: "1.2.3.4",
 * });
 * export const myElasticIpId = myElasticIp.then(myElasticIp => myElasticIp.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getElasticIp(args: GetElasticIpArgs, opts?: pulumi.InvokeOptions): Promise<GetElasticIpResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getElasticIp:getElasticIp", {
        "id": args.id,
        "ipAddress": args.ipAddress,
        "labels": args.labels,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getElasticIp.
 */
export interface GetElasticIpArgs {
    /**
     * The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
     */
    id?: string;
    /**
     * The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
     */
    ipAddress?: string;
    /**
     * The EIP labels to match (conflicts with `ipAddress` and `id`).
     */
    labels?: {[key: string]: string};
    /**
     * The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getElasticIp.
 */
export interface GetElasticIpResult {
    /**
     * The Elastic IP (EIP) address family (`inet4` or `inet6`).
     */
    readonly addressFamily: string;
    /**
     * The Elastic IP (EIP) CIDR.
     */
    readonly cidr: string;
    /**
     * The Elastic IP (EIP) description.
     */
    readonly description: string;
    /**
     * The *managed* EIP healthcheck configuration.
     */
    readonly healthchecks: outputs.GetElasticIpHealthcheck[];
    /**
     * The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
     */
    readonly id?: string;
    /**
     * The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
     */
    readonly ipAddress?: string;
    /**
     * The EIP labels to match (conflicts with `ipAddress` and `id`).
     */
    readonly labels?: {[key: string]: string};
    /**
     * Domain name for reverse DNS record.
     */
    readonly reverseDns: string;
    /**
     * The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * Fetch Exoscale [Elastic IPs (EIP)](https://community.exoscale.com/documentation/compute/eip/) data.
 *
 * Corresponding resource: exoscale_elastic_ip.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as exoscale from "@pulumi/exoscale";
 *
 * const myElasticIp = exoscale.getElasticIp({
 *     zone: "ch-gva-2",
 *     ipAddress: "1.2.3.4",
 * });
 * export const myElasticIpId = myElasticIp.then(myElasticIp => myElasticIp.id);
 * ```
 *
 * Please refer to the examples
 * directory for complete configuration examples.
 */
export function getElasticIpOutput(args: GetElasticIpOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetElasticIpResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getElasticIp:getElasticIp", {
        "id": args.id,
        "ipAddress": args.ipAddress,
        "labels": args.labels,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getElasticIp.
 */
export interface GetElasticIpOutputArgs {
    /**
     * The Elastic IP (EIP) ID to match (conflicts with `ipAddress` and `labels`).
     */
    id?: pulumi.Input<string>;
    /**
     * The EIP IPv4 or IPv6 address to match (conflicts with `id` and `labels`).
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The EIP labels to match (conflicts with `ipAddress` and `id`).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Exocale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
