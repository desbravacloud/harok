package jenkins

import (
	"context"
	"log"

	"github.com/bndr/gojenkins"
)

var jenkinsAddress, _ = getJenkinsAddress()
var gitOrg, _ = getGitHubOrganization()

func setJenkinsCtx() context.Context {
	ctx := context.Background()

	return ctx
}

func runJenkins() *gojenkins.Jenkins {
	var jenkinsUser, _ = getJenkinsUser()
	var jenkinsToken, _ = getJenkinsToken()

	jenkins, err := gojenkins.CreateJenkins(nil, jenkinsAddress, jenkinsUser, jenkinsToken).Init(setJenkinsCtx())
	if err != nil {
		log.Fatal(err)
	}

	return jenkins
}

func CreateJob(jobName string) (*gojenkins.Job, error) {
	configString := `<?xml version='1.1' encoding='UTF-8'?>
	<org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject plugin="workflow-multibranch@latest">
	  <actions/>
	  <description></description>
	  <keepDependencies>false</keepDependencies>
	  <properties/>
	  <scm class="hudson.scm.NullSCM"/>
	  <canRoam>true</canRoam>
	  <disabled>false</disabled>
	  <sources class="jenkins.branch.MultiBranchProject$BranchSourceList" plugin="branch-api@2.7.0">
		<data>
		  <jenkins.branch.BranchSource>
			<source class="jenkins.plugins.git.GitSCMSource" plugin="git@latest">
			  <id></id>
			  <remote>https://github.com/` + gitOrg + `/` + jobName + `.git</remote>
			  <credentialsId></credentialsId>
			  <traits>
				<jenkins.plugins.git.traits.BranchDiscoveryTrait/>
			  </traits>
			</source>
		  </jenkins.branch.BranchSource>
		</data>
		<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
	  </sources>
	  <triggers>
		<com.igalg.jenkins.plugins.mswt.trigger.ComputedFolderWebHookTrigger plugin="multibranch-scan-webhook-trigger@1.0.9">
		  <spec></spec>
		  <token>` + jobName + `</token>
		</com.igalg.jenkins.plugins.mswt.trigger.ComputedFolderWebHookTrigger>
	  </triggers>
	  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
	  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
	  <triggers class="vector"/>
	  <concurrentBuild>false</concurrentBuild>
	  <builders/>
	  <publishers/>
	  <buildWrappers/>
	  </org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject>`

	job, err := runJenkins().CreateJob(setJenkinsCtx(), configString, jobName)
	if err != nil {
		return nil, err
	}
	return job, nil
}
