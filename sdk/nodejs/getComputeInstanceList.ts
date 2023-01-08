// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getComputeInstanceList(args: GetComputeInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeInstanceListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getComputeInstanceList:getComputeInstanceList", {
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeInstanceList.
 */
export interface GetComputeInstanceListArgs {
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: string;
}

/**
 * A collection of values returned by getComputeInstanceList.
 */
export interface GetComputeInstanceListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instances: outputs.GetComputeInstanceListInstance[];
    readonly zone: string;
}
export function getComputeInstanceListOutput(args: GetComputeInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeInstanceListResult> {
    return pulumi.output(args).apply((a: any) => getComputeInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getComputeInstanceList.
 */
export interface GetComputeInstanceListOutputArgs {
    /**
     * The Exoscale [Zone][zone] name.
     */
    zone: pulumi.Input<string>;
}
