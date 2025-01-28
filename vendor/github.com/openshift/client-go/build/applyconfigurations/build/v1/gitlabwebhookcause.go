// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GitLabWebHookCauseApplyConfiguration represents a declarative configuration of the GitLabWebHookCause type for use
// with apply.
type GitLabWebHookCauseApplyConfiguration struct {
	CommonWebHookCauseApplyConfiguration `json:",inline"`
}

// GitLabWebHookCauseApplyConfiguration constructs a declarative configuration of the GitLabWebHookCause type for use with
// apply.
func GitLabWebHookCause() *GitLabWebHookCauseApplyConfiguration {
	return &GitLabWebHookCauseApplyConfiguration{}
}

// WithRevision sets the Revision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Revision field is set to the value of the last call.
func (b *GitLabWebHookCauseApplyConfiguration) WithRevision(value *SourceRevisionApplyConfiguration) *GitLabWebHookCauseApplyConfiguration {
	b.CommonWebHookCauseApplyConfiguration.Revision = value
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *GitLabWebHookCauseApplyConfiguration) WithSecret(value string) *GitLabWebHookCauseApplyConfiguration {
	b.CommonWebHookCauseApplyConfiguration.Secret = &value
	return b
}
