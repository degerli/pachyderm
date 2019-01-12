## ./pachctl



### Synopsis


Access the Pachyderm API.

Environment variables:
  ADDRESS=<host>:<port>, the pachd server to connect to (e.g. 127.0.0.1:30650).


### Options

```
      --no-metrics   Don't report user metrics for this command
  -v, --verbose      Output verbose logs
```

### SEE ALSO
* [./pachctl auth](./pachctl_auth.html)	 - Auth commands manage access to data in a Pachyderm cluster
* [./pachctl commit](./pachctl_commit.html)	 - Docs for commits.
* [./pachctl completion](./pachctl_completion.html)	 - Print or install the bash completion code.
* [./pachctl copy-file](./pachctl_copy-file.html)	 - Copy files between pfs paths.
* [./pachctl create-branch](./pachctl_create-branch.html)	 - Create a new branch, or update an existing branch, on a repo.
* [./pachctl create-pipeline](./pachctl_create-pipeline.html)	 - Create a new pipeline.
* [./pachctl create-repo](./pachctl_create-repo.html)	 - Create a new repo.
* [./pachctl debug-dump](./pachctl_debug-dump.html)	 - Return a dump of running goroutines.
* [./pachctl delete-all](./pachctl_delete-all.html)	 - Delete everything.
* [./pachctl delete-branch](./pachctl_delete-branch.html)	 - Delete a branch
* [./pachctl delete-commit](./pachctl_delete-commit.html)	 - Delete an input commit.
* [./pachctl delete-file](./pachctl_delete-file.html)	 - Delete a file.
* [./pachctl delete-job](./pachctl_delete-job.html)	 - Delete a job.
* [./pachctl delete-pipeline](./pachctl_delete-pipeline.html)	 - Delete a pipeline.
* [./pachctl delete-repo](./pachctl_delete-repo.html)	 - Delete a repo.
* [./pachctl deploy](./pachctl_deploy.html)	 - Deploy a Pachyderm cluster.
* [./pachctl diff-file](./pachctl_diff-file.html)	 - Return a diff of two file trees.
* [./pachctl edit-pipeline](./pachctl_edit-pipeline.html)	 - Edit the manifest for a pipeline in your text editor.
* [./pachctl enterprise](./pachctl_enterprise.html)	 - Enterprise commands enable Pachyderm Enterprise features
* [./pachctl extract](./pachctl_extract.html)	 - Extract Pachyderm state to stdout or an object store bucket.
* [./pachctl extract-pipeline](./pachctl_extract-pipeline.html)	 - Return the manifest used to create a pipeline.
* [./pachctl file](./pachctl_file.html)	 - Docs for files.
* [./pachctl finish-commit](./pachctl_finish-commit.html)	 - Finish a started commit.
* [./pachctl flush-commit](./pachctl_flush-commit.html)	 - Wait for all commits caused by the specified commits to finish and return them.
* [./pachctl flush-job](./pachctl_flush-job.html)	 - Wait for all jobs caused by the specified commits to finish and return them.
* [./pachctl garbage-collect](./pachctl_garbage-collect.html)	 - Garbage collect unused data.
* [./pachctl get-file](./pachctl_get-file.html)	 - Return the contents of a file.
* [./pachctl get-logs](./pachctl_get-logs.html)	 - Return logs from a job.
* [./pachctl get-object](./pachctl_get-object.html)	 - Return the contents of an object
* [./pachctl get-tag](./pachctl_get-tag.html)	 - Return the contents of a tag
* [./pachctl glob-file](./pachctl_glob-file.html)	 - Return files that match a glob pattern in a commit.
* [./pachctl inspect-cluster](./pachctl_inspect-cluster.html)	 - Returns info about the pachyderm cluster
* [./pachctl inspect-commit](./pachctl_inspect-commit.html)	 - Return info about a commit.
* [./pachctl inspect-datum](./pachctl_inspect-datum.html)	 - Display detailed info about a single datum.
* [./pachctl inspect-file](./pachctl_inspect-file.html)	 - Return info about a file.
* [./pachctl inspect-job](./pachctl_inspect-job.html)	 - Return info about a job.
* [./pachctl inspect-pipeline](./pachctl_inspect-pipeline.html)	 - Return info about a pipeline.
* [./pachctl inspect-repo](./pachctl_inspect-repo.html)	 - Return info about a repo.
* [./pachctl job](./pachctl_job.html)	 - Docs for jobs.
* [./pachctl list-branch](./pachctl_list-branch.html)	 - Return all branches on a repo.
* [./pachctl list-commit](./pachctl_list-commit.html)	 - Return all commits on a set of repos.
* [./pachctl list-datum](./pachctl_list-datum.html)	 - Return the datums in a job.
* [./pachctl list-file](./pachctl_list-file.html)	 - Return the files in a directory.
* [./pachctl list-job](./pachctl_list-job.html)	 - Return info about jobs.
* [./pachctl list-pipeline](./pachctl_list-pipeline.html)	 - Return info about all pipelines.
* [./pachctl list-repo](./pachctl_list-repo.html)	 - Return all repos.
* [./pachctl mount](./pachctl_mount.html)	 - Mount pfs locally. This command blocks.
* [./pachctl pipeline](./pachctl_pipeline.html)	 - Docs for pipelines.
* [./pachctl port-forward](./pachctl_port-forward.html)	 - Forward a port on the local machine to pachd. This command blocks.
* [./pachctl put-file](./pachctl_put-file.html)	 - Put a file into the filesystem.
* [./pachctl repo](./pachctl_repo.html)	 - Docs for repos.
* [./pachctl restart-datum](./pachctl_restart-datum.html)	 - Restart a datum.
* [./pachctl restore](./pachctl_restore.html)	 - Restore Pachyderm state from stdin or an object store.
* [./pachctl set-branch](./pachctl_set-branch.html)	 - DEPRECATED Set a commit and its ancestors to a branch
* [./pachctl start-commit](./pachctl_start-commit.html)	 - Start a new commit.
* [./pachctl start-pipeline](./pachctl_start-pipeline.html)	 - Restart a stopped pipeline.
* [./pachctl stop-job](./pachctl_stop-job.html)	 - Stop a job.
* [./pachctl stop-pipeline](./pachctl_stop-pipeline.html)	 - Stop a running pipeline.
* [./pachctl subscribe-commit](./pachctl_subscribe-commit.html)	 - Print commits as they are created (finished).
* [./pachctl undeploy](./pachctl_undeploy.html)	 - Tear down a deployed Pachyderm cluster.
* [./pachctl unmount](./pachctl_unmount.html)	 - Unmount pfs.
* [./pachctl update-dash](./pachctl_update-dash.html)	 - Update and redeploy the Pachyderm Dashboard at the latest compatible version.
* [./pachctl update-pipeline](./pachctl_update-pipeline.html)	 - Update an existing Pachyderm pipeline.
* [./pachctl update-repo](./pachctl_update-repo.html)	 - Update a repo.
* [./pachctl version](./pachctl_version.html)	 - Return version information.

###### Auto generated by spf13/cobra on 11-Jan-2019
