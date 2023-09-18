// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IAMAccessKey extends pulumi.CustomResource {
    /**
     * Get an existing IAMAccessKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMAccessKeyState, opts?: pulumi.CustomResourceOptions): IAMAccessKey {
        return new IAMAccessKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'exoscale:index/iAMAccessKey:IAMAccessKey';

    /**
     * Returns true if the given object is an instance of IAMAccessKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMAccessKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMAccessKey.__pulumiType;
    }

    /**
     * The IAM access key (identifier).
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * ❗ The IAM access key name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ❗ A list of API operations to restrict the key to.
     */
    public readonly operations!: pulumi.Output<string[] | undefined>;
    /**
     * ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
     */
    public readonly resources!: pulumi.Output<string[] | undefined>;
    /**
     * The key secret.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * ❗ A list of tags to restrict the key to.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly tagsOperations!: pulumi.Output<string[]>;

    /**
     * Create a IAMAccessKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IAMAccessKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMAccessKeyArgs | IAMAccessKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAMAccessKeyState | undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operations"] = state ? state.operations : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsOperations"] = state ? state.tagsOperations : undefined;
        } else {
            const args = argsOrState as IAMAccessKeyArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operations"] = args ? args.operations : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["secret"] = undefined /*out*/;
            resourceInputs["tagsOperations"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IAMAccessKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMAccessKey resources.
 */
export interface IAMAccessKeyState {
    /**
     * The IAM access key (identifier).
     */
    key?: pulumi.Input<string>;
    /**
     * ❗ The IAM access key name.
     */
    name?: pulumi.Input<string>;
    /**
     * ❗ A list of API operations to restrict the key to.
     */
    operations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key secret.
     */
    secret?: pulumi.Input<string>;
    /**
     * ❗ A list of tags to restrict the key to.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tagsOperations?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a IAMAccessKey resource.
 */
export interface IAMAccessKeyArgs {
    /**
     * ❗ The IAM access key name.
     */
    name?: pulumi.Input<string>;
    /**
     * ❗ A list of API operations to restrict the key to.
     */
    operations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ A list of API [resources](https://community.exoscale.com/documentation/iam/quick-start/#restricting-api-access-keys-to-resources) to restrict the key to (`<domain>/<type>:<name>`).
     */
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ❗ A list of tags to restrict the key to.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
