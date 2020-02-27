package bugsnag

import "time"

// Event represents an event received from Bugsnag
type Event struct {
	Account *Account `json:"account"`
	Project *Project `json:"project"`
	Trigger *Trigger `json:"trigger"`
	User    *User    `json:"user"`
	Error   *Error   `json:"error"`
	Release *Release `json:"release"`
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Trigger struct {
	Type        string      `json:"type"`
	Message     string      `json:"message"`
	Rate        int         `json:"rate"`
	StateChange string      `json:"stateChange"`
	SnoozeRule  *SnoozeRule `json:"snoozeRule"`
}

type SnoozeRule struct {
	Type      string `json:"type"`
	RuleValue int    `json:"ruleValue"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Error struct {
	ID             string        `json:"id"`
	ErrorID        string        `json:"errorId"`
	ExceptionClass string        `json:"exceptionClass"`
	Message        string        `json:"message"`
	Context        string        `json:"context"`
	FirstReceived  *time.Time    `json:"firstReceived"`
	ReceivedAt     *time.Time    `json:"receivedAt"`
	RequestURL     string        `json:"requestUrl"`
	AssignedUserID string        `json:"assignedUserId"`
	URL            string        `json:"url"`
	Severity       string        `json:"severity"`
	Status         string        `json:"status"`
	Unhandled      bool          `json:"unhandled"`
	CreatedIssue   *CreatedIssue `json:"createdIssue"`
	User           *User         `json:"user"`
	App            *App          `json:"app"`
	Device         *Device       `json:"device"`
	StackTrace     *[]Stacktrace `json:"stackTrace"`
}

type CreatedIssue struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
	Type   string `json:"type"`
	URL    string `json:"url"`
}

type App struct {
	ID                   string   `json:"id"`
	Version              string   `json:"version"`
	VersionCode          string   `json:"versionCode"`
	BundleVersion        string   `json:"bundleVersion"`
	CodeBundleID         string   `json:"codeBundleId"`
	BuildUUID            string   `json:"buildUUID"`
	ReleaseStage         string   `json:"releaseStage"`
	Type                 string   `json:"type"`
	DsymUUIDs            []string `json:"dsymUUIDs"`
	Duration             int      `json:"duration"`
	DurationInForeground int      `json:"durationInForeground"`
	InForeground         bool     `json:"inForeground"`
}

type Device struct {
	HostName       string     `json:"hostname"`
	ID             string     `json:"id"`
	Manufacturer   string     `json:"manufacturer"`
	Model          string     `json:"model"`
	ModelNumber    string     `json:"modelNumber"`
	OSName         string     `json:"osName"`
	OSVersion      string     `json:"osVersion"`
	FreeMemory     int        `json:"freeMemory"`
	TotalMemory    int        `json:"totalMemory"`
	FreeDisk       int        `json:"freeDisk"`
	BrowserName    string     `json:"browserName"`
	BrowserVersion string     `json:"browserVersion"`
	Jailbroken     bool       `json:"jailbroken"`
	Orientation    string     `json:"orientation"`
	Locale         string     `json:"locale"`
	Charging       bool       `json:"charging"`
	BatteryLevel   float32    `json:"batteryLevel"`
	Time           *time.Time `json:"time"`
	Timezone       string     `json:"timezone"`
}

type Stacktrace struct {
	InProject    bool   `json:"inProject"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`
	File         string `json:"file"`
	Method       string `json:"method"`
	// TODO: type `code` field
}

type Release struct {
	ID            string         `json:"id"`
	Version       string         `json:"version"`
	VersionCode   string         `json:"versionCode"`
	BundleVersion string         `json:"bundleVersion"`
	ReleaseStage  string         `json:"releaseStage"`
	URL           string         `json:"url"`
	ReleaseTime   *time.Time     `json:"releaseTime"`
	ReleasedBy    string         `json:"releasedBy"`
	SourceControl *SourceControl `json:"sourceControl"`
}

type SourceControl struct {
	Provider    string `json:"provider"`
	Revision    string `json:"revision"`
	RevisionURL string `json:"revisionUrl"`
	DiffURL     string `json:"diffUrl"`
}
