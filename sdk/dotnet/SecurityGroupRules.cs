// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    /// <summary>
    /// !&gt; **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use the exoscale.SecurityGroupRule instead (or refer to the ad-hoc migration guide).
    /// </summary>
    [ExoscaleResourceType("exoscale:index/securityGroupRules:SecurityGroupRules")]
    public partial class SecurityGroupRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        [Output("egresses")]
        public Output<ImmutableArray<Outputs.SecurityGroupRulesEgress>> Egresses { get; private set; } = null!;

        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        [Output("ingresses")]
        public Output<ImmutableArray<Outputs.SecurityGroupRulesIngress>> Ingresses { get; private set; } = null!;

        /// <summary>
        /// ❗ The security group (name) the rules apply to (conflicts with `security_group_id`).
        /// </summary>
        [Output("securityGroup")]
        public Output<string> SecurityGroup { get; private set; } = null!;

        /// <summary>
        /// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupRules(string name, SecurityGroupRulesArgs? args = null, CustomResourceOptions? options = null)
            : base("exoscale:index/securityGroupRules:SecurityGroupRules", name, args ?? new SecurityGroupRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupRules(string name, Input<string> id, SecurityGroupRulesState? state = null, CustomResourceOptions? options = null)
            : base("exoscale:index/securityGroupRules:SecurityGroupRules", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-exoscale",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityGroupRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupRules Get(string name, Input<string> id, SecurityGroupRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupRules(name, id, state, options);
        }
    }

    public sealed class SecurityGroupRulesArgs : global::Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<Inputs.SecurityGroupRulesEgressArgs>? _egresses;

        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        public InputList<Inputs.SecurityGroupRulesEgressArgs> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<Inputs.SecurityGroupRulesEgressArgs>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<Inputs.SecurityGroupRulesIngressArgs>? _ingresses;

        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        public InputList<Inputs.SecurityGroupRulesIngressArgs> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<Inputs.SecurityGroupRulesIngressArgs>());
            set => _ingresses = value;
        }

        /// <summary>
        /// ❗ The security group (name) the rules apply to (conflicts with `security_group_id`).
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public SecurityGroupRulesArgs()
        {
        }
        public static new SecurityGroupRulesArgs Empty => new SecurityGroupRulesArgs();
    }

    public sealed class SecurityGroupRulesState : global::Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<Inputs.SecurityGroupRulesEgressGetArgs>? _egresses;

        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        public InputList<Inputs.SecurityGroupRulesEgressGetArgs> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<Inputs.SecurityGroupRulesEgressGetArgs>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<Inputs.SecurityGroupRulesIngressGetArgs>? _ingresses;

        /// <summary>
        /// A security group rule definition (can be specified multiple times).
        /// </summary>
        public InputList<Inputs.SecurityGroupRulesIngressGetArgs> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<Inputs.SecurityGroupRulesIngressGetArgs>());
            set => _ingresses = value;
        }

        /// <summary>
        /// ❗ The security group (name) the rules apply to (conflicts with `security_group_id`).
        /// </summary>
        [Input("securityGroup")]
        public Input<string>? SecurityGroup { get; set; }

        /// <summary>
        /// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public SecurityGroupRulesState()
        {
        }
        public static new SecurityGroupRulesState Empty => new SecurityGroupRulesState();
    }
}
