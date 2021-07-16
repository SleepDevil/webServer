package model

type Question struct {
	TopicId        string `json:"topicId"`
	QuestionName   string `json:"questionName"`
	QuestionAnswer string `json:"questionAnswer"`
}

type Topic struct {
	ID               string `json:"id"`
	TopicName        string `json:"topicName"`
	Part             string `json:"part"`
	TopicDescription string `json:"topicDescription"`
}
