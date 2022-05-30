package bean

import "encoding/xml"

type Project struct {
	XMLName xml.Name `xml:"project"`
	Text    string   `xml:",chardata"`
	Actions struct {
		Text string `xml:",chardata"`
	} `xml:"actions"`
	Description struct {
		Text string `xml:",chardata"`
	} `xml:"description"`
	KeepDependencies struct {
		Text string `xml:",chardata"`
	} `xml:"keepDependencies"`
	Properties struct {
		Text                               string `xml:",chardata"`
		JenkinsModelBuildDiscarderProperty struct {
			Text     string `xml:",chardata"`
			Strategy struct {
				Text       string `xml:",chardata"`
				Class      string `xml:"class,attr"`
				DaysToKeep struct {
					Text string `xml:",chardata"`
				} `xml:"daysToKeep"`
				NumToKeep struct {
					Text string `xml:",chardata"`
				} `xml:"numToKeep"`
				ArtifactDaysToKeep struct {
					Text string `xml:",chardata"`
				} `xml:"artifactDaysToKeep"`
				ArtifactNumToKeep struct {
					Text string `xml:",chardata"`
				} `xml:"artifactNumToKeep"`
			} `xml:"strategy"`
		} `xml:"jenkins.model.BuildDiscarderProperty"`
	} `xml:"properties"`
	Scm struct {
		Text          string `xml:",chardata"`
		Class         string `xml:"class,attr"`
		Plugin        string `xml:"plugin,attr"`
		ConfigVersion struct {
			Text string `xml:",chardata"`
		} `xml:"configVersion"`
		UserRemoteConfigs struct {
			Text                             string `xml:",chardata"`
			HudsonPluginsGitUserRemoteConfig struct {
				Text string `xml:",chardata"`
				URL  struct {
					Text string `xml:",chardata"`
				} `xml:"url"`
				CredentialsId struct {
					Text string `xml:",chardata"`
				} `xml:"credentialsId"`
			} `xml:"hudson.plugins.git.UserRemoteConfig"`
		} `xml:"userRemoteConfigs"`
		Branches struct {
			Text                       string `xml:",chardata"`
			HudsonPluginsGitBranchSpec struct {
				Text string `xml:",chardata"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
			} `xml:"hudson.plugins.git.BranchSpec"`
		} `xml:"branches"`
		DoGenerateSubmoduleConfigurations struct {
			Text string `xml:",chardata"`
		} `xml:"doGenerateSubmoduleConfigurations"`
		SubmoduleCfg struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"submoduleCfg"`
		Extensions struct {
			Text string `xml:",chardata"`
		} `xml:"extensions"`
	} `xml:"scm"`
	CanRoam struct {
		Text string `xml:",chardata"`
	} `xml:"canRoam"`
	Disabled struct {
		Text string `xml:",chardata"`
	} `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding struct {
		Text string `xml:",chardata"`
	} `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding struct {
		Text string `xml:",chardata"`
	} `xml:"blockBuildWhenUpstreamBuilding"`
	Triggers struct {
		Text string `xml:",chardata"`
	} `xml:"triggers"`
	ConcurrentBuild struct {
		Text string `xml:",chardata"`
	} `xml:"concurrentBuild"`
	Builders struct {
		Text                      string `xml:",chardata"`
		HudsonPluginsGradleGradle struct {
			Text     string `xml:",chardata"`
			Plugin   string `xml:"plugin,attr"`
			Switches struct {
				Text string `xml:",chardata"`
			} `xml:"switches"`
			Tasks struct {
				Text string `xml:",chardata"`
			} `xml:"tasks"`
			RootBuildScriptDir struct {
				Text string `xml:",chardata"`
			} `xml:"rootBuildScriptDir"`
			BuildFile struct {
				Text string `xml:",chardata"`
			} `xml:"buildFile"`
			GradleName struct {
				Text string `xml:",chardata"`
			} `xml:"gradleName"`
			UseWrapper struct {
				Text string `xml:",chardata"`
			} `xml:"useWrapper"`
			MakeExecutable struct {
				Text string `xml:",chardata"`
			} `xml:"makeExecutable"`
			UseWorkspaceAsHome struct {
				Text string `xml:",chardata"`
			} `xml:"useWorkspaceAsHome"`
			WrapperLocation struct {
				Text string `xml:",chardata"`
			} `xml:"wrapperLocation"`
			PassAllAsSystemProperties struct {
				Text string `xml:",chardata"`
			} `xml:"passAllAsSystemProperties"`
			ProjectProperties struct {
				Text string `xml:",chardata"`
			} `xml:"projectProperties"`
			PassAllAsProjectProperties struct {
				Text string `xml:",chardata"`
			} `xml:"passAllAsProjectProperties"`
		} `xml:"hudson.plugins.gradle.Gradle"`
	} `xml:"builders"`
	Publishers struct {
		Text string `xml:",chardata"`
	} `xml:"publishers"`
	BuildWrappers struct {
		Text string `xml:",chardata"`
	} `xml:"buildWrappers"`
}

func NewProject() *Project {
	return &Project{}
}
