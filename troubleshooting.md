# Troubleshooting Guide

Here we will add some well-known troubleshooting issues that you may run into, along with possible solutions.

## Quickfeed Issues

1. If you have problems logging in into the application, cannot see your courses or lab submissions, please try these few steps:
   1. Make sure you are using the right URL: `uis.itest.run`.
   2. Make sure your browser does not have cookies blocked.
   3. Make sure you are logged in with the right GitHub account in case you have several accounts.
   4. Refresh the page.
   5. Log out of the Quickfeed application, then log out of your GitHub account.
      Clear all browser data.
      Log back in to Quickfeed here: [https://uis.itest.run](uis.itest.run).

## General Test Issues

1. What is this `TestLintAG` thing that gives me test failures?

   Most lab assignments include a `TestLintAG` checker that checks that your Go code
   - follows Go coding style as defined by the `gofmt` program,
   - follows (some of the) best practices for Go coding,
   - does not contain any TODO or FIXME items.

   If you are getting a message like: `File is not goimports-ed (goimports)`, this means that you are not using the proper formatting of your Go code.
   To fix this, use the Go plugin for VSCode and ensure that it works to format your code.
   The formatter works when you save your file.
   It is easy to check that it works, by adding a line that is incorrectly formatted, e.g. `var myName =   "John Doe"`.
   (make sure to include some extra spaces between the tokens.)
   When you save your Go file, the spaces should be removed automatically.

   Another alternative is to run the `go fmt` command in the same directory as your code, but that's a bit annoying to remember.
   You can of course also configure any other editor to run the `goimports` tool.

## GitHub and SSH Keys

If you are having issues using the `git` command to access GitHub, here are some things that you can check to identify, and hopefully solve your problem.

GitHub allows users to work with repositories using two different protocols `https` or `ssh`, each one requires their own set of [configurations](https://docs.github.com/en/github/using-git/which-remote-url-should-i-use) and uses a different URL to connect with the GitHub servers.

- HTTPS URL: `https://github.com/YOUR_USER/SOME_REPO.git`
- SSH URL: `git@github.com:YOUR_USER/SOME_REPO.git`

In this course, we use the __ssh__ protocol to access the repositories, since it allows you to connect to GitHub without supplying your username or password every time.
But for that to work, you need to configure it properly, as described [here](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh)

1. `Permission denied (publickey)` when `clone/pull/push` a repository.

   If you are getting this error is probably because you forgot to add your public key to GitHub, or you are trying to access the repository with a different key-pair.
   In either case, take a look [here](https://docs.github.com/en/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account) and see how to add a key in the GitHub and test if it is properly configured by running:

   ```console
   ssh -T git@github.com
   ```

   The command should display a message like:

   ```text
   Hi YOUR_USERNAME! You've successfully authenticated, but GitHub does not provide shell access.
   ```

   If you get an error, ensure that you are using the correct public key in your machine to connect to GitHub.
   The content of your public key file, normally located at your home folder: `~/.ssh/id_rsa.pub` should be the same as displayed in your GitHub account settings.
   We have created a [SSH tutorial video](https://youtu.be/qik3HHZW6C0) illustrating the necessary steps (and a bit more).

2. There are many reasons that can result in the error below when cloning or pulling a GitHub repository:

   ```text
   fatal: Could not read from remote repository

   Please make sure you have the correct access rights
   and the repository exists.
   ```

   One common reason is a misconfigured remote URL.
   As explained above, we use the `ssh` protocol to avoid having to type password for each interaction with GitHub.
   Hence, if the output from the command `git remote -v` displays a URL using `https` as shown below, you will need to change these entries in order to use ssh.

   ```console
   $ git remote -v
   course-assignments  https://github.com/dat320-2021/assignments.git (fetch)
   course-assignments  https://github.com/dat320-2021/assignments.git (push)
   origin  https://github.com/dat320-2021/YOUR_USERNAME-labs.git (fetch)
   origin  https://github.com/dat320-2021/YOUR_USERNAME-labs.git (push)
   ```

   If this is the case, change the remote's URL to use ssh by running (remember to replace YOUR_USERNAME with your own):

   ```console
   git remote set-url course-assignments git@github.com:dat320-2021/assignments.git
   git remote set-url origin git@github.com:dat320-2021/YOUR_USERNAME-labs.git
   ```

   The new remote's URL should be like this:

   ```console
   $ git remote -v
   course-assignments  git@github.com:dat320-2021/assignments.git (fetch)
   course-assignments  git@github.com:dat320-2021/assignments.git (push)
   origin  git@github.com:dat320-2021/YOUR_USERNAME-labs.git (fetch)
   origin  git@github.com:dat320-2021/YOUR_USERNAME-labs.git (push)
   ```

3. Multiple ssh clients or conflicting git configurations

   If experience the following problem while using git with [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10):

   ```console
   C:\Windows\System32\OpenSSH\ssh.exe" Permission denied
   ```

   Ensure that your git configuration points to the correct ssh client path.

   ```console
   $ git config --list --global
   ...
   [core]
      sshCommand = "C:\Windows\System32\OpenSSH\ssh.exe"
   ```

   If the output of the above command displays a different path from the command `which ssh` in your Linux subsystem.

   ```console
   $ which ssh
   /usr/bin/ssh
   ```

   Then you may need to edit your configuration to use a ssh command that your user has permission to execute.
   This can be done by editing your local or global git configuration.

   To edit the global configuration (applies to all repositories on your Linux subsystem):

   ```console
   $ git config --edit --global
   ...
   [core]
      sshCommand = /usr/bin/ssh
   ```

   To edit your local configuration (applies only to the current `assignments` repository):

   ```console
   $ cd dat320-2021/assignments
   $ git config --edit
   ...
   [core]
      sshCommand = /usr/bin/ssh
   ```

4. If you get an error message *refusing to merge unrelated histories*.
   This is because you have made changes to your personal repository before fetching and merging with `assignments` repository.
   But no worries, this can be solved by running the following command:

   ```console
   git pull course-assignments main —allow-unrelated-histories
   ```

   However, your local changes should not conflict with the changes in the `assignments` repository.
   Hence, you should not have any files in folders that already exists in the `assignments` repository.

## Lab 2

1. If QuickFeed complains with the following error message:

   ```text
   package dat320/lab2: C source files not allowed when not using cgo or SWIG: hello.c
   ```

   Then you most probably have created a `hello.c` file in the assignments root folder.
   The `hello.c` file should be in the `lab2/hello` folder of the assignments repository.

## Lab 3

1. If your `cmd/terminal` implementation causes the QuickFeed tests to timeout and emit the error message below, and your implementation otherwise seems to work on your local machine.

   ```text
   terminal_ag_test.go:342: The application failed during execution: signal: killed
   ```

   Then you should make sure that you are printing the path as part of the prompt, as described in the task description (Line 13).
   Here is tip:

   ```golang
   wd, err := os.Getwd()
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("%s> ", wd)
   ```

   If you already do, the next thing to check is if you are creating readers inside the for loop:

   ```golang
   reader := bufio.NewReader(os.Stdin)
   ```

   If you do this the test will fail, and the fix is to move the `reader` outside the for loop.
