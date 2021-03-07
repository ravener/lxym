
package pages

const Git = "Git is a distributed version control and source code management system.\n\nIt does this through a series of snapshots of your project, and it works\nwith those snapshots to provide you with functionality to version and\nmanage your source code.\n\n## Versioning Concepts\n\n### What is version control?\n\nVersion control is a system that records changes to a file(s), over time.\n\n### Centralized Versioning vs. Distributed Versioning\n\n* Centralized version control focuses on synchronizing, tracking, and backing\nup files.\n* Distributed version control focuses on sharing changes. Every change has a\nunique id.\n* Distributed systems have no defined structure. You could easily have a SVN\nstyle, centralized system, with git.\n\n[Additional Information](http://git-scm.com/book/en/Getting-Started-About-Version-Control)\n\n### Why Use Git?\n\n* Can work offline.\n* Collaborating with others is easy!\n* Branching is easy!\n* Branching is fast!\n* Merging is easy!\n* Git is fast.\n* Git is flexible.\n\n## Git Architecture\n\n### Repository\n\nA set of files, directories, historical records, commits, and heads. Imagine it\nas a source code data structure, with the attribute that each source code\n\"element\" gives you access to its revision history, among other things.\n\nA git repository is comprised of the .git directory & working tree.\n\n### .git Directory (component of repository)\n\nThe .git directory contains all the configurations, logs, branches, HEAD, and\nmore.\n[Detailed List.](http://gitready.com/advanced/2009/03/23/whats-inside-your-git-directory.html)\n\n### Working Tree (component of repository)\n\nThis is basically the directories and files in your repository. It is often\nreferred to as your working directory.\n\n### Index (component of .git dir)\n\nThe Index is the staging area in git. It's basically a layer that separates\nyour working tree from the Git repository. This gives developers more power\nover what gets sent to the Git repository.\n\n### Commit\n\nA git commit is a snapshot of a set of changes, or manipulations to your\nWorking Tree. For example, if you added 5 files, and removed 2 others, these\nchanges will be contained in a commit (or snapshot). This commit can then be\npushed to other repositories, or not!\n\n### Branch\n\nA branch is essentially a pointer to the last commit you made. As you go on\ncommitting, this pointer will automatically update to point to the latest commit.\n\n### Tag\n\nA tag is a mark on specific point in history. Typically people use this\nfunctionality to mark release points (v1.0, and so on).\n\n### HEAD and head (component of .git dir)\n\nHEAD is a pointer that points to the current branch. A repository only has 1\n*active* HEAD.\nhead is a pointer that points to any commit. A repository can have any number\nof heads.\n\n### Stages of Git\n* Modified - Changes have been made to a file but file has not been committed\nto Git Database yet\n* Staged - Marks a modified file to go into your next commit snapshot\n* Committed - Files have been committed to the Git Database\n\n### Conceptual Resources\n\n* [Git For Computer Scientists](http://eagain.net/articles/git-for-computer-scientists/)\n* [Git For Designers](http://hoth.entp.com/output/git_for_designers.html)\n\n## Commands\n\n### init\n\nCreate an empty Git repository. The Git repository's settings, stored\ninformation, and more is stored in a directory (a folder) named \".git\".\n\n```bash\n$ git init\n```\n\n### config\n\nTo configure settings. Whether it be for the repository, the system itself,\nor global configurations ( global config file is `~/.gitconfig` ).\n\n```bash\n# Print & Set Some Basic Config Variables (Global)\n$ git config --global user.email \"MyEmail@Zoho.com\"\n$ git config --global user.name \"My Name\"\n```\n\n[Learn More About git config.](http://git-scm.com/docs/git-config)\n\n### help\n\nTo give you quick access to an extremely detailed guide of each command. Or to\njust give you a quick reminder of some semantics.\n\n```bash\n# Quickly check available commands\n$ git help\n\n# Check all available commands\n$ git help -a\n\n# Command specific help - user manual\n# git help <command_here>\n$ git help add\n$ git help commit\n$ git help init\n# or git <command_here> --help\n$ git add --help\n$ git commit --help\n$ git init --help\n```\n\n### ignore files\n\nTo intentionally untrack file(s) & folder(s) from git. Typically meant for\nprivate & temp files which would otherwise be shared in the repository.\n\n```bash\n$ echo \"temp/\" >> .gitignore\n$ echo \"private_key\" >> .gitignore\n```\n\n### status\n\nTo show differences between the index file (basically your working copy/repo)\nand the current HEAD commit.\n\n```bash\n# Will display the branch, untracked files, changes and other differences\n$ git status\n\n# To learn other \"tid bits\" about git status\n$ git help status\n```\n\n### add\n\nTo add files to the staging area/index. If you do not `git add` new files to\nthe staging area/index, they will not be included in commits!\n\n```bash\n# add a file in your current working directory\n$ git add HelloWorld.java\n\n# add a file in a nested dir\n$ git add /path/to/file/HelloWorld.c\n\n# Regular Expression support!\n$ git add ./*.java\n\n# You can also add everything in your working directory to the staging area.\n$ git add -A\n```\n\nThis only adds a file to the staging area/index, it doesn't commit it to the\nworking directory/repo.\n\n### branch\n\nManage your branches. You can view, edit, create, delete branches using this\ncommand.\n\n```bash\n# list existing branches & remotes\n$ git branch -a\n\n# create a new branch\n$ git branch myNewBranch\n\n# delete a branch\n$ git branch -d myBranch\n\n# rename a branch\n# git branch -m <oldname> <newname>\n$ git branch -m myBranchName myNewBranchName\n\n# edit a branch's description\n$ git branch myBranchName --edit-description\n```\n\n### tag\n\nManage your tags\n\n```bash\n# List tags\n$ git tag\n\n# Create a annotated tag\n# The -m specifies a tagging message, which is stored with the tag.\n# If you don’t specify a message for an annotated tag,\n# Git launches your editor so you can type it in.\n$ git tag -a v2.0 -m 'my version 2.0'\n\n# Show info about tag\n# That shows the tagger information, the date the commit was tagged,\n# and the annotation message before showing the commit information.\n$ git show v2.0\n\n# Push a single tag to remote\n$ git push origin v2.0\n\n# Push a lot of tags to remote\n$ git push origin --tags\n```\n\n### checkout\n\nUpdates all files in the working tree to match the version in the index, or\nspecified tree.\n\n```bash\n# Checkout a repo - defaults to master branch\n$ git checkout\n\n# Checkout a specified branch\n$ git checkout branchName\n\n# Create a new branch & switch to it\n# equivalent to \"git branch <name>; git checkout <name>\"\n\n$ git checkout -b newBranch\n```\n\n### clone\n\nClones, or copies, an existing repository into a new directory. It also adds\nremote-tracking branches for each branch in the cloned repo, which allows you\nto push to a remote branch.\n\n```bash\n# Clone learnxinyminutes-docs\n$ git clone https://github.com/adambard/learnxinyminutes-docs.git\n\n# shallow clone - faster cloning that pulls only latest snapshot\n$ git clone --depth 1 https://github.com/adambard/learnxinyminutes-docs.git\n\n# clone only a specific branch\n$ git clone -b master-cn https://github.com/adambard/learnxinyminutes-docs.git --single-branch\n```\n\n### commit\n\nStores the current contents of the index in a new \"commit.\" This commit\ncontains the changes made and a message created by the user.\n\n```bash\n# commit with a message\n$ git commit -m \"Added multiplyNumbers() function to HelloWorld.c\"\n\n# signed commit with a message (user.signingkey must have been set\n# with your GPG key e.g. git config --global user.signingkey 5173AAD5)\n$ git commit -S -m \"signed commit message\"\n\n# automatically stage modified or deleted files, except new files, and then commit\n$ git commit -a -m \"Modified foo.php and removed bar.php\"\n\n# change last commit (this deletes previous commit with a fresh commit)\n$ git commit --amend -m \"Correct message\"\n```\n\n### diff\n\nShows differences between a file in the working directory, index and commits.\n\n```bash\n# Show difference between your working dir and the index\n$ git diff\n\n# Show differences between the index and the most recent commit.\n$ git diff --cached\n\n# Show differences between your working dir and the most recent commit\n$ git diff HEAD\n```\n\n### grep\n\nAllows you to quickly search a repository.\n\nOptional Configurations:\n\n```bash\n# Thanks to Travis Jeffery for these\n# Set line numbers to be shown in grep search results\n$ git config --global grep.lineNumber true\n\n# Make search results more readable, including grouping\n$ git config --global alias.g \"grep --break --heading --line-number\"\n```\n\n```bash\n# Search for \"variableName\" in all java files\n$ git grep 'variableName' -- '*.java'\n\n# Search for a line that contains \"arrayListName\" and, \"add\" or \"remove\"\n$ git grep -e 'arrayListName' --and \\( -e add -e remove \\)\n```\n\nGoogle is your friend; for more examples\n[Git Grep Ninja](http://travisjeffery.com/b/2012/02/search-a-git-repo-like-a-ninja)\n\n### log\n\nDisplay commits to the repository.\n\n```bash\n# Show all commits\n$ git log\n\n# Show only commit message & ref\n$ git log --oneline\n\n# Show merge commits only\n$ git log --merges\n\n# Show all commits represented by an ASCII graph\n$ git log --graph\n```\n\n### merge\n\n\"Merge\" in changes from external commits into the current branch.\n\n```bash\n# Merge the specified branch into the current.\n$ git merge branchName\n\n# Always generate a merge commit when merging\n$ git merge --no-ff branchName\n```\n\n### mv\n\nRename or move a file\n\n```bash\n# Renaming a file\n$ git mv HelloWorld.c HelloNewWorld.c\n\n# Moving a file\n$ git mv HelloWorld.c ./new/path/HelloWorld.c\n\n# Force rename or move\n# \"existingFile\" already exists in the directory, will be overwritten\n$ git mv -f myFile existingFile\n```\n\n### pull\n\nPulls from a repository and merges it with another branch.\n\n```bash\n# Update your local repo, by merging in new changes\n# from the remote \"origin\" and \"master\" branch.\n# git pull <remote> <branch>\n$ git pull origin master\n\n# By default, git pull will update your current branch\n# by merging in new changes from its remote-tracking branch\n$ git pull\n\n# Merge in changes from remote branch and rebase\n# branch commits onto your local repo, like: \"git fetch <remote> <branch>, git\n# rebase <remote>/<branch>\"\n$ git pull origin master --rebase\n```\n\n### push\n\nPush and merge changes from a branch to a remote & branch.\n\n```bash\n# Push and merge changes from a local repo to a\n# remote named \"origin\" and \"master\" branch.\n# git push <remote> <branch>\n$ git push origin master\n\n# By default, git push will push and merge changes from\n# the current branch to its remote-tracking branch\n$ git push\n\n# To link up current local branch with a remote branch, add -u flag:\n$ git push -u origin master\n# Now, anytime you want to push from that same local branch, use shortcut:\n$ git push\n```\n\n### stash\n\nStashing takes the dirty state of your working directory and saves it on a\nstack of unfinished changes that you can reapply at any time.\n\nLet's say you've been doing some work in your git repo, but you want to pull\nfrom the remote. Since you have dirty (uncommitted) changes to some files, you\nare not able to run `git pull`. Instead, you can run `git stash` to save your\nchanges onto a stack!\n\n```bash\n$ git stash\nSaved working directory and index state \\\n  \"WIP on master: 049d078 added the index file\"\n  HEAD is now at 049d078 added the index file\n  (To restore them type \"git stash apply\")\n```\n\nNow you can pull!\n\n```bash\ngit pull\n```\n`...changes apply...`\n\nNow check that everything is OK\n\n```bash\n$ git status\n# On branch master\nnothing to commit, working directory clean\n```\n\nYou can see what \"hunks\" you've stashed so far using `git stash list`.\nSince the \"hunks\" are stored in a Last-In-First-Out stack, our most recent\nchange will be at top.\n\n```bash\n$ git stash list\nstash@{0}: WIP on master: 049d078 added the index file\nstash@{1}: WIP on master: c264051 Revert \"added file_size\"\nstash@{2}: WIP on master: 21d80a5 added number to log\n```\n\nNow let's apply our dirty changes back by popping them off the stack.\n\n```bash\n$ git stash pop\n# On branch master\n# Changes not staged for commit:\n#   (use \"git add <file>...\" to update what will be committed)\n#\n#      modified:   index.html\n#      modified:   lib/simplegit.rb\n#\n```\n\n`git stash apply` does the same thing\n\nNow you're ready to get back to work on your stuff!\n\n[Additional Reading.](http://git-scm.com/book/en/v1/Git-Tools-Stashing)\n\n### rebase (caution)\n\nTake all changes that were committed on one branch, and replay them onto\nanother branch.\n*Do not rebase commits that you have pushed to a public repo*.\n\n```bash\n# Rebase experimentBranch onto master\n# git rebase <basebranch> <topicbranch>\n$ git rebase master experimentBranch\n```\n\n[Additional Reading.](http://git-scm.com/book/en/Git-Branching-Rebasing)\n\n### reset (caution)\n\nReset the current HEAD to the specified state. This allows you to undo merges,\npulls, commits, adds, and more. It's a great command but also dangerous if you\ndon't know what you are doing.\n\n```bash\n# Reset the staging area, to match the latest commit (leaves dir unchanged)\n$ git reset\n\n# Reset the staging area, to match the latest commit, and overwrite working dir\n$ git reset --hard\n\n# Moves the current branch tip to the specified commit (leaves dir unchanged)\n# all changes still exist in the directory.\n$ git reset 31f2bb1\n\n# Moves the current branch tip backward to the specified commit\n# and makes the working dir match (deletes uncommitted changes and all commits\n# after the specified commit).\n$ git reset --hard 31f2bb1\n```\n\n### reflog (caution)\n\nReflog will list most of the git commands you have done for a given time period,\ndefault 90 days.\n\nThis give you the chance to reverse any git commands that have gone wrong\n(for instance, if a rebase has broken your application).\n\nYou can do this:\n\n1. `git reflog` to list all of the git commands for the rebase\n\n```\n38b323f HEAD@{0}: rebase -i (finish): returning to refs/heads/feature/add_git_reflog\n38b323f HEAD@{1}: rebase -i (pick): Clarify inc/dec operators\n4fff859 HEAD@{2}: rebase -i (pick): Update java.html.markdown\n34ed963 HEAD@{3}: rebase -i (pick): [yaml/en] Add more resources (#1666)\ned8ddf2 HEAD@{4}: rebase -i (pick): pythonstatcomp spanish translation (#1748)\n2e6c386 HEAD@{5}: rebase -i (start): checkout 02fb96d\n```\n2. Select where to reset to, in our case its `2e6c386`, or `HEAD@{5}`\n3. 'git reset --hard HEAD@{5}' this will reset your repo to that head\n4. You can start the rebase again or leave it alone.\n\n[Additional Reading.](https://git-scm.com/docs/git-reflog)\n\n### revert\n\nRevert can be used to undo a commit. It should not be confused with reset which\nrestores the state of a project to a previous point. Revert will add a new\ncommit which is the inverse of the specified commit, thus reverting it.\n\n```bash\n# Revert a specified commit\n$ git revert <commit>\n```\n\n### rm\n\nThe opposite of git add, git rm removes files from the current working tree.\n\n```bash\n# remove HelloWorld.c\n$ git rm HelloWorld.c\n\n# Remove a file from a nested dir\n$ git rm /pather/to/the/file/HelloWorld.c\n```\n\n## Further Information\n\n* [tryGit - A fun interactive way to learn Git.](http://try.github.io/levels/1/challenges/1)\n\n* [Learn Git Branching - the most visual and interactive way to learn Git on the web](http://learngitbranching.js.org/)\n\n* [Udemy Git Tutorial: A Comprehensive Guide](https://blog.udemy.com/git-tutorial-a-comprehensive-guide/)\n\n* [Git Immersion - A Guided tour that walks through the fundamentals of git](http://gitimmersion.com/)\n\n* [git-scm - Video Tutorials](http://git-scm.com/videos)\n\n* [git-scm - Documentation](http://git-scm.com/docs)\n\n* [Atlassian Git - Tutorials & Workflows](https://www.atlassian.com/git/)\n\n* [SalesForce Cheat Sheet](http://res.cloudinary.com/hy4kyit2a/image/upload/SF_git_cheatsheet.pdf)\n\n* [GitGuys](http://www.gitguys.com/)\n\n* [Git - the simple guide](http://rogerdudler.github.io/git-guide/index.html)\n\n* [Pro Git](http://www.git-scm.com/book/en/v2)\n\n* [An introduction to Git and GitHub for Beginners (Tutorial)](http://product.hubspot.com/blog/git-and-github-tutorial-for-beginners)\n\n* [The New Boston tutorial to Git covering basic commands and workflow](https://www.youtube.com/playlist?list=PL6gx4Cwl9DGAKWClAD_iKpNC0bGHxGhcx)"