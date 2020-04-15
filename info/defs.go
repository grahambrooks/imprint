package info

import "reflect"

var BuildNumber string
var BuildId string
var BuildUrl string
var NodeName string
var JobName string
var BuildTag string
var PackagePath string
var JenkinsUrl string
var ExecutorNumber string
var JavaHome string
var Workspace string
var SvnVersion string
var CvsVersion string
var GitCommit string
var GitUrl string
var GitBranch string
var AwsDefaultRegion string
var AwsRegion string
var BuildArn string
var BuildImage string
var BuildSucceeding string
var BuildInitiator string
var KmsKeyId string
var LogPath string
var ResolvedSourceVersion string
var SourceRepoUrl string
var SourceVersion string
var SrcDir string
var StartTime string
var WebHookActorAccountId string
var WebHookBaseRef string
var WebHookEvent string
var WebHookPrevCommit string
var WebHookHeadRef string
var WebHookTrigger string
var DefinitionMap map[string]VariableDefinition
var CodeBuildBuildNumber string
var CodeBuildBuildId string

type VariableDefinition struct {
	Name  string
	Value *string
}


func init() {
	PackagePath = reflect.TypeOf(InfoBuild{}).PkgPath()

	DefinitionMap = make(map[string]VariableDefinition)
	DefinitionMap["BUILD_NUMBER"] = VariableDefinition{
		Name:  "BuildNumber",
		Value: &BuildNumber,
	}
	DefinitionMap["BUILD_ID"] = VariableDefinition{
		Name:  "BuildId",
		Value: &BuildId,
	}
	DefinitionMap["BUILD_URL"] = VariableDefinition{
		Name:  "BuildUrl",
		Value: &BuildUrl,
	}
	DefinitionMap["NODE_NAME"] = VariableDefinition{
		Name:  "NodeName",
		Value: &NodeName,
	}
	DefinitionMap["JOB_NAME"] = VariableDefinition{
		Name:  "JobName",
		Value: &JobName,
	}
	DefinitionMap["BUILD_TAG"] = VariableDefinition{
		Name:  "BuildTag",
		Value: &BuildTag,
	}
	DefinitionMap["JENKINS_URL"] = VariableDefinition{
		Name:  "JenkinsUrl",
		Value: &JenkinsUrl,
	}
	DefinitionMap["EXECUTOR_NUMBER"] = VariableDefinition{
		Name:  "ExecutorNumber",
		Value: &ExecutorNumber,
	}
	DefinitionMap["JAVA_HOME"] = VariableDefinition{
		Name:  "JavaHome",
		Value: &JavaHome,
	}
	DefinitionMap["WORKSPACE"] = VariableDefinition{
		Name:  "Workspace",
		Value: &Workspace,
	}
	DefinitionMap["SVN_REVISION"] = VariableDefinition{
		Name:  "SvnVersion",
		Value: &SvnVersion,
	}
	DefinitionMap["CVS_BRANCH"] = VariableDefinition{
		Name:  "CvsVersion",
		Value: &CvsVersion,
	}
	DefinitionMap["GIT_COMMIT"] = VariableDefinition{
		Name:  "GitCommit",
		Value: &GitCommit,
	}
	DefinitionMap["GIT_URL"] = VariableDefinition{
		Name:  "GitUrl",
		Value: &GitUrl,
	}
	DefinitionMap["GIT_BRANCH"] = VariableDefinition{
		Name:  "GitBranch",
		Value: &GitBranch,
	}
	//
	////AWS Code Build
	DefinitionMap["AWS_DEFAULT_REGION"] = VariableDefinition{
		Name:  "AwsDefaultRegion",
		Value: &AwsDefaultRegion,
	}
	DefinitionMap["AWS_REGION"] = VariableDefinition{
		Name:  "AwsRegion",
		Value: &AwsRegion,
	}
	DefinitionMap["CODEBUILD_BUILD_ARN"] = VariableDefinition{
		Name:  "BuildArn",
		Value: &BuildArn,
	}
	DefinitionMap["CODEBUILD_BUILD_ID"] = VariableDefinition{
		Name:  "CodeBuildBuildId",
		Value: &CodeBuildBuildId,
	}
	DefinitionMap["CODEBUILD_BUILD_IMAGE"] = VariableDefinition{
		Name:  "BuildImage",
		Value: &BuildImage,
	}
	DefinitionMap["CODEBUILD_BUILD_NUMBER"] = VariableDefinition{
		Name:  "CodeBuildBuildNumber",
		Value: &CodeBuildBuildNumber,
	}
	DefinitionMap["CODEBUILD_BUILD_SUCCEEDING"] = VariableDefinition{
		Name:  "BuildSucceeding",
		Value: &BuildSucceeding,
	}
	DefinitionMap["CODEBUILD_INITIATOR"] = VariableDefinition{
		Name:  "BuildInitiator",
		Value: &BuildInitiator,
	}
	DefinitionMap["CODEBUILD_KMS_KEY_ID"] = VariableDefinition{
		Name:  "KmsKeyId",
		Value: &KmsKeyId,
	}
	DefinitionMap["CODEBUILD_LOG_PATH"] = VariableDefinition{
		Name:  "LogPath",
		Value: &LogPath,
	}
	DefinitionMap["CODEBUILD_RESOLVED_SOURCE_VERSION"] = VariableDefinition{
		Name:  "ResolvedSourceVersion",
		Value: &ResolvedSourceVersion,
	}
	DefinitionMap["CODEBUILD_SOURCE_REPO_URL"] = VariableDefinition{
		Name:  "SourceRepoUrl",
		Value: &SourceRepoUrl,
	}
	DefinitionMap["CODEBUILD_SOURCE_VERSION"] = VariableDefinition{
		Name:  "SourceVersion",
		Value: &SourceVersion,
	}
	DefinitionMap["CODEBUILD_SRC_DIR"] = VariableDefinition{
		Name:  "SrcDir",
		Value: &SrcDir,
	}
	DefinitionMap["CODEBUILD_START_TIME"] = VariableDefinition{
		Name:  "StartTime",
		Value: &StartTime,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_ACTOR_ACCOUNT_ID"] = VariableDefinition{
		Name:  "WebHookActorAccountId",
		Value: &WebHookActorAccountId,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_BASE_REF"] = VariableDefinition{
		Name:  "WebHookBaseRef",
		Value: &WebHookBaseRef,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_EVENT"] = VariableDefinition{
		Name:  "WebHookEvent",
		Value: &WebHookEvent,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_PREV_COMMIT"] = VariableDefinition{
		Name:  "WebHookPrevCommit",
		Value: &WebHookPrevCommit,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_HEAD_REF"] = VariableDefinition{
		Name:  "WebHookHeadRef",
		Value: &WebHookHeadRef,
	}
	DefinitionMap["CODEBUILD_WEBHOOK_TRIGGER"] = VariableDefinition{
		Name:  "WebHookTrigger",
		Value: &WebHookTrigger,
	}
}
