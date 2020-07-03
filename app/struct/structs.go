package _struct

///任务
type TaskDash struct {
	ID				string 	`json:"id"`
	Name 			string 	`json:"name"`
	URL				string 	`json:"url"`
	User	 		string 	`json:"user"`
	GroupID 		string 	`json:"group_id"`
	Description		string 	`json:"description"`
	Status			string 	`json:"status"`
	StartTime		string 	`json:"start_time"`
	StopTime		string 	`json:"stop_time"`
}
