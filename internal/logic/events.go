
package logic

func EventAction(eventType string) string {
	switch eventType {
	case "PushEvent":
		return "Pushed"
	case "CreateEvent":
		return "Created"
	case "DeleteEvent":
		return "Deleted"
	case "ForkEvent":
		return "Forked"
	case "WatchEvent":
		return "Starred"
	case "IssuesEvent":
		return "Issued"
	case "IssueCommentEvent":
		return "Commented"
	case "PullRequestEvent":
		return "PR"
	case "PullRequestReviewEvent":
		return "Reviewed"
	case "PullRequestReviewCommentEvent":
		return "PR-Commented"
	case "CommitCommentEvent":
		return "Commit-Commented"
	case "ReleaseEvent":
		return "Released"
	case "PublicEvent":
		return "Publicized"
	case "MemberEvent":
		return "Added-Member"
	case "GollumEvent":
		return "Wiki-Updated"
	case "SponsorshipEvent":
		return "Sponsored"
	case "DeploymentEvent":
		return "Deployed"
	case "DeploymentStatusEvent":
		return "Deploy-Status"
	case "StatusEvent":
		return "Status-Updated"
	case "CheckRunEvent":
		return "Check-Run"
	case "CheckSuiteEvent":
		return "Check-Suite"
	case "WorkflowRunEvent":
		return "CI"
	case "SecurityAdvisoryEvent":
		return "Security"
	case "SecretScanningAlertEvent":
		return "Secret-Alert"
	case "SecretScanningAlertLocationEvent":
		return "Secret-Found"
	case "DependabotAlertEvent":
		return "Dependency-Alert"
	default:
		return "Activity"
	}
}
