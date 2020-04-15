package info

type CodeBuildBuild struct {
	Arn                   string `json:",omitempty"`
	Id                    string `json:",omitempty"`
	Image                 string `json:",omitempty"`
	Number                string `json:",omitempty"`
	Succeeding            string `json:",omitempty"`
	Initiator             string `json:",omitempty"`
	KmsKeyId              string `json:",omitempty"`
	LogPath               string `json:",omitempty"`
	ResolvedSourceVersion string `json:",omitempty"`
	SourceRepoUrl         string `json:",omitempty"`
	SourceVersion         string `json:",omitempty"`
	SourceDir             string `json:",omitempty"`
	StartTime             string `json:",omitempty"`
	WebHook               struct {
		ActorAccountId string `json:",omitempty"`
		BaseRef        string `json:",omitempty"`
		Event          string `json:",omitempty"`
		PrevCommit     string `json:",omitempty"`
		HeadRef        string `json:",omitempty"`
		Trigger        string `json:",omitempty"`
	}
}

type InfoBuild struct {
	BuildNumber  string `json:"build_number,omitempty" imprint:"BUILD_NUMBER"`
	Id           string `imprint:"BUILD_ID"`
	Url          string `imprint:"BUILD_URL"`
	Node         string `imprint:"NODE_NAME"`
	JobName      string `imprint:"JOB_NAME"`
	BuildTag     string `imprint:"BUILD_TAG"`
	JenkinsUrl   string `imprint:"JENKINS_URL"`
	ExecutorName string `imprint:"EXECUTOR_NUMBER"`
	JavaHome     string `imprint:"JAVA_HOME"`
	Workspace    string `imprint:"WORKSPACE"`
	ScnRevision  string `imprint:"SVN_REVISION"`
	CvsBranch    string `imprint:"CVS_BRANCH"`
	GitCommit    string `imprint:"GIT_COMMIT"`
	GitUrl       string `imprint:"GIT_URL"`
	GitBranch    string `imprint:"GIT_BRANCH"`
}

type InfoBlock struct {
	Build     InfoBuild      `json:"build,omitempty"`
	CodeBuild CodeBuildBuild `json:"code_build, omitempty"`
}

func Block() InfoBlock {
	block := InfoBlock{
		Build: InfoBuild{BuildNumber: BuildNumber},
	}
	_ = Imprint(&block)
	return block
}
