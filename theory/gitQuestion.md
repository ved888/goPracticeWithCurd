# Top 50 Git and GitHub Interview Questions for 2 Years of Experience

## 1. What is Git and why is it used?
**Answer**: Git is a distributed version control system used to track changes in code during software development. It allows multiple developers to work on a project simultaneously while keeping track of changes and maintaining history.

## 2. What is the difference between Git and GitHub?
**Answer**: Git is a version control system, while GitHub is a platform that hosts Git repositories, allowing for collaboration, code sharing, and remote storage of Git projects.

## 3. What are the basic Git commands you use regularly?
**Answer**: The basic Git commands include `git clone`, `git init`, `git add`, `git commit`, `git push`, `git pull`, `git merge`, and `git status`.

## 4. What is a Git repository?
**Answer**: A Git repository is a storage space where your project's files and their version history are kept. It can be local or hosted on a platform like GitHub.

## 5. What is the difference between `git pull` and `git fetch`?
**Answer**: `git fetch` downloads new data from a remote repository but doesn’t integrate it with your working files. `git pull` downloads and automatically integrates the changes into your local branch.

## 6. What is a branch in Git?
**Answer**: A branch in Git represents an independent line of development. It allows you to work on different features or fixes without affecting the main codebase (usually `main` or `master` branch).

## 7. What is the purpose of the `git merge` command?
**Answer**: The `git merge` command combines the changes from one branch into another. It helps in integrating code from multiple branches.

## 8. What is the difference between `git merge` and `git rebase`?
**Answer**: `git merge` creates a merge commit, preserving the history of the merged branch. `git rebase` moves or reapplies commits from one branch onto another, resulting in a linear history.

## 9. What is a commit in Git?
**Answer**: A commit in Git is a snapshot of the changes made to the files in your repository. It includes a unique ID (SHA), commit message, and metadata like the author and timestamp.

## 10. What is the purpose of a `.gitignore` file?
**Answer**: The `.gitignore` file tells Git which files or directories to ignore when tracking changes in a repository. Common use cases include ignoring build artifacts, dependency folders, or IDE-specific files.

## 11. What is a pull request in GitHub?
**Answer**: A pull request (PR) in GitHub is a request to merge one branch into another. It allows team members to review changes before they are merged into the main branch.

## 12. How do you resolve merge conflicts in Git?
**Answer**: Merge conflicts occur when changes in different branches conflict with each other. To resolve conflicts, you need to manually edit the conflicting files and commit the resolved changes.

## 13. What is the purpose of a Git tag?
**Answer**: A Git tag is used to mark specific points in the repository’s history, often used for marking release versions.

## 14. How do you clone a Git repository?
**Answer**: You can clone a repository using the `git clone <repository_url>` command, which copies the repository to your local machine.

## 15. What is a detached HEAD state in Git?
**Answer**: A detached HEAD state occurs when you checkout a specific commit rather than a branch, meaning you are not on any branch but still working in the history of your repository.

## 16. How do you switch between branches in Git?
**Answer**: You switch branches in Git using the `git checkout <branch_name>` or `git switch <branch_name>` command.

## 17. What does the `git status` command do?
**Answer**: The `git status` command shows the current status of the working directory and staging area, including which files are staged, modified, or untracked.

## 18. What does the `git log` command show?
**Answer**: The `git log` command shows the commit history, including commit IDs, commit messages, authors, and dates.

## 19. What is GitHub Actions?
**Answer**: GitHub Actions is an automation tool that allows you to define workflows to automate tasks like building, testing, and deploying your code.

## 20. What is the difference between `git reset` and `git revert`?
**Answer**: `git reset` undoes commits and changes the commit history, whereas `git revert` creates a new commit that undoes the changes from a previous commit without altering the commit history.

## 21. What is the purpose of the `git stash` command?
**Answer**: The `git stash` command temporarily saves changes that are not ready to be committed, allowing you to switch branches without losing your work.

## 22. What is GitHub Pages?
**Answer**: GitHub Pages is a static site hosting service provided by GitHub that allows you to host HTML, CSS, and JavaScript files directly from a GitHub repository.

## 23. How do you handle large files in Git?
**Answer**: Large files can be handled using Git Large File Storage (LFS), which stores large files outside of the Git repository while still linking to them.

## 24. How do you contribute to an open-source project on GitHub?
**Answer**: To contribute, fork the repository, create a branch, make your changes, commit them, push to your fork, and then open a pull request to the original repository.

## 25. What is the purpose of the `git cherry-pick` command?
**Answer**: `git cherry-pick` applies the changes from a specific commit onto the current branch without merging the entire branch.

## 26. What is a fast-forward merge in Git?
**Answer**: A fast-forward merge occurs when the current branch has no changes that diverge from the branch being merged, so Git can simply move the pointer to the merged commit.

## 27. What is the difference between a local and remote repository?
**Answer**: A local repository is stored on your machine, while a remote repository is hosted on a server (e.g., GitHub, GitLab) and can be accessed by multiple developers.

## 28. What is the `git diff` command?
**Answer**: The `git diff` command shows the differences between the working directory and the index (staging area) or between different commits.

## 29. What is the purpose of the `git push` command?
**Answer**: The `git push` command uploads local commits to a remote repository, making the changes available to others.

## 30. What is the `git fetch` command?
**Answer**: `git fetch` downloads the latest changes from the remote repository without modifying your local branch.

## 31. How do you remove a file from the Git index without deleting it?
**Answer**: You can remove a file from the Git index using the command `git rm --cached <file>`.

## 32. What is the `git rebase` command used for?
**Answer**: `git rebase` is used to apply commits from one branch onto another, creating a cleaner, linear commit history.

## 33. How do you add a remote repository in Git?
**Answer**: You add a remote repository using the command `git remote add <remote_name> <repository_url>`.

## 34. What is the purpose of the `git config` command?
**Answer**: The `git config` command is used to set configuration options like the user name, email, editor, and more in Git.

## 35. How do you delete a branch in Git?
**Answer**: You can delete a local branch using `git branch -d <branch_name>`, and delete a remote branch using `git push origin --delete <branch_name>`.

## 36. What is GitHub Issues?
**Answer**: GitHub Issues is a feature that allows developers to track bugs, enhancements, and tasks in a repository, providing a way to collaborate on problem-solving.

## 37. How do you resolve a merge conflict in GitHub?
**Answer**: To resolve a merge conflict, manually edit the conflicting files, then stage and commit the changes.

## 38. What is a fork in GitHub?
**Answer**: A fork is a copy of a repository that you can freely make changes to without affecting the original repository. Forks are commonly used for contributing to open-source projects.

## 39. How do you tag a commit in Git?
**Answer**: You can tag a commit using the command `git tag <tag_name>`.

## 40. How do you undo a commit in Git?
**Answer**: You can undo the last commit using `git reset --soft HEAD~1` to keep the changes in the working directory or `git reset --hard HEAD~1` to discard the changes.

## 41. What is the `.gitmodules` file?
**Answer**: The `.gitmodules` file is used to track Git submodules, which are repositories embedded within a parent repository.

## 42. What is Git rebase interactive mode?
**Answer**: Git rebase interactive mode (`git rebase -i`) allows you to edit, reorder, or squash commits before applying them onto another branch.

## 43. How do you view the commit history in Git?
**Answer**: You can view the commit history using `git log`.

## 44. What is the purpose of the `git bisect` command?
**Answer**: The `git bisect` command helps you find the commit that introduced a bug by performing a binary search through the commit history.

## 45. What is the difference between `git pull --rebase` and `git pull`?
**Answer**: `git pull --rebase` fetches changes and applies your local commits on top of them, while `git pull` merges the changes.

## 46. What are Git submodules?
**Answer**: Git submodules allow you to include one Git repository as a subdirectory in another Git repository, enabling the use of external repositories as part of your project.

## 47. How do you configure Git for multiple users on the same machine?
**Answer**: You can configure Git to use different user credentials for different repositories using `git config user.name` and `git config user.email` within each repository.

## 48. What is the GitHub API used for?
**Answer**: The GitHub API allows developers to interact programmatically with GitHub, automating tasks like managing issues, repositories, and pull requests.

## 49. How do you view all remotes associated with a Git repository?
**Answer**: You can view all remotes with the command `git remote -v`.

## 50. What are Git hooks?
**Answer**: Git hooks are scripts that run automatically on certain Git events, such as `pre-commit` or `post-merge`, allowing you to automate tasks like testing or linting code.
