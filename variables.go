package main

var EnvironmentVariables []string

var EnvironmentVariableMap map[string]string

const packagePrefix = "github.com/grahambrooks/imprint/"

func init() {
	EnvironmentVariables = []string{
		"BUILD_NUMBER",
		"BUILD_ID",
		"BUILD_URL",
		"NODE_NAME",
		"JOB_NAME",
		"BUILD_TAG",
		"JENKINS_URL",
		"EXECUTOR_NUMBER",
		"JAVA_HOME",
		"WORKSPACE",
		"SVN_REVISION",
		"CVS_BRANCH",
		"GIT_COMMIT",
		"GIT_URL",
		"GIT_BRANCH",

		//AWS Code Build
		"AWS_DEFAULT_REGION",
		"AWS_REGION",
		"CODEBUILD_BUILD_ARN",
		"CODEBUILD_BUILD_ID",
		"CODEBUILD_BUILD_IMAGE",
		"CODEBUILD_BUILD_NUMBER",
		"CODEBUILD_BUILD_SUCCEEDING",
		"CODEBUILD_INITIATOR",
		"CODEBUILD_KMS_KEY_ID",
		"CODEBUILD_LOG_PATH",
		"CODEBUILD_RESOLVED_SOURCE_VERSION",
		"CODEBUILD_SOURCE_REPO_URL",
		"CODEBUILD_SOURCE_VERSION",
		"CODEBUILD_SRC_DIR",
		"CODEBUILD_START_TIME",
		"CODEBUILD_WEBHOOK_ACTOR_ACCOUNT_ID",
		"CODEBUILD_WEBHOOK_BASE_REF",
		"CODEBUILD_WEBHOOK_EVENT",
		"CODEBUILD_WEBHOOK_PREV_COMMIT",
		"CODEBUILD_WEBHOOK_HEAD_REF",
		"CODEBUILD_WEBHOOK_TRIGGER",
	}
	EnvironmentVariableMap = make(map[string]string)
	EnvironmentVariableMap["BUILD_NUMBER"] = packagePrefix + "info.BuildNumber"
}
