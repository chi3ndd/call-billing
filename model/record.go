package model

type (
	Record struct {
		Created  int64  `json:"created" bson:"created"`
		Username string `json:"username" bson:"username"`
		Duration int64  `json:"duration" bson:"duration"`
	}

	RequestUser struct {
		Username string `param:"username"`
	}

	RequestUserCall struct {
		RequestUser
		CallDuration int64 `json:"call_duration"`
	}
)
