// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use the exoscale.SecurityGroupRule instead (or refer to the ad-hoc migration guide).
 */
export class SecurityGroupRules extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRulesState, opts?: pulumi.CustomResourceOptions): SecurityGroupRules {
        return new SecurityGroupRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/securityGroupRules:SecurityGroupRules';

    /**
     * Returns true if the given object is an instance of SecurityGroupRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRules.__pulumiType;
    }

    public readonly egresses!: pulumi.Output<outputs.SecurityGroupRulesEgress[] | undefined>;
    public readonly ingresses!: pulumi.Output<outputs.SecurityGroupRulesIngress[] | undefined>;
    /**
     * The security group (name) the rules apply to (conflicts with `securityGroupId`).
     */
    public readonly securityGroup!: pulumi.Output<string>;
    /**
     * The security group (ID) the rules apply to (conficts with `security_group)`.
     */
    public readonly securityGroupId!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRulesArgs | SecurityGroupRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupRulesState | undefined;
            resourceInputs["egresses"] = state ? state.egresses : undefined;
            resourceInputs["ingresses"] = state ? state.ingresses : undefined;
            resourceInputs["securityGroup"] = state ? state.securityGroup : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
        } else {
            const args = argsOrState as SecurityGroupRulesArgs | undefined;
            resourceInputs["egresses"] = args ? args.egresses : undefined;
            resourceInputs["ingresses"] = args ? args.ingresses : undefined;
            resourceInputs["securityGroup"] = args ? args.securityGroup : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRules resources.
 */
export interface SecurityGroupRulesState {
    egresses?: pulumi.Input<pulumi.Input<inputs.SecurityGroupRulesEgress>[]>;
    ingresses?: pulumi.Input<pulumi.Input<inputs.SecurityGroupRulesIngress>[]>;
    /**
     * The security group (name) the rules apply to (conflicts with `securityGroupId`).
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The security group (ID) the rules apply to (conficts with `security_group)`.
     */
    securityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRules resource.
 */
export interface SecurityGroupRulesArgs {
    egresses?: pulumi.Input<pulumi.Input<inputs.SecurityGroupRulesEgress>[]>;
    ingresses?: pulumi.Input<pulumi.Input<inputs.SecurityGroupRulesIngress>[]>;
    /**
     * The security group (name) the rules apply to (conflicts with `securityGroupId`).
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The security group (ID) the rules apply to (conficts with `security_group)`.
     */
    securityGroupId?: pulumi.Input<string>;
}
