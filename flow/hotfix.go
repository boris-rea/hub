package flow

func getPrefixHotfix() (prefix string, err error) {
	gitConfig, err := NewConfig()

	if err != nil {
		return
	}

	prefix, err = gitConfig.GetPrefix("hotfix")

	return
}

func getStartBranchHotfix() (branch string, err error) {
	gitConfig, err := NewConfig()

	if err != nil {
		return
	}

	branch, err = gitConfig.GetBranch("master")

	return
}

func getFinishBranchHotfix() (branchMaster string, branchDevelop string, err error) {
	gitConfig, err := NewConfig()

	if err != nil {
		return
	}

	branchDevelop, err = gitConfig.GetBranch("develop")
	branchMaster, err = gitConfig.GetBranch("master")

	return
}

func FlowHotfixStart(hotfixName string) (err error) {
	var prefixHotfix string
	prefixHotfix, err = getPrefixHotfix()

	if err != nil {
		return
	}

	branchName := prefixHotfix + hotfixName

	var startingBranch string
	startingBranch, err = getStartBranchHotfix()

	if err != nil {
		return
	}

	cmdGit := [][]string{}

	cmdGit1 := []string{"checkout", startingBranch}
	cmdGit2 := []string{"checkout", "-b", branchName}

	cmdGit = append(cmdGit, cmdGit1, cmdGit2)

	err = launchCmdGit(cmdGit)

	return
}

func FlowHotfixFinish(hotfixName string) (err error) {
	var prefixHotfix string
	prefixHotfix, err = getPrefixHotfix()

	if err != nil {
		return
	}

	branchName := prefixHotfix + hotfixName

	var (
		masterBranch,
		developBranch string
	)

	masterBranch, developBranch, err = getFinishBranchRelease()

	if err != nil {
		return
	}

	cmdGit := [][]string{}

	cmdGit1 := []string{"checkout", masterBranch}
	cmdGit2 := []string{"merge", branchName, "--no-ff"}
	cmdGit3 := []string{"tag", "-a", hotfixName}
	cmdGit4 := []string{"checkout", developBranch}
	cmdGit5 := []string{"merge", branchName, "--no-ff"}
	cmdGit6 := []string{"checkout", masterBranch}
	cmdGit7 := []string{"branch", "-d", branchName}

	cmdGit = append(cmdGit, cmdGit1, cmdGit2, cmdGit3, cmdGit4, cmdGit5, cmdGit6, cmdGit7)

	err = launchCmdGit(cmdGit)

	return
}
