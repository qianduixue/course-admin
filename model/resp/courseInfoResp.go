package resp

type CourseInfoResp struct {
	ID           uint   `json:"id"`            // 主键ID
	Title        string `json:"title"`         //课程标题
	LessonNumber int8   `json:"lesson_number"` //节数
	CreatedTime  int64  `json:"created_time"`  //添加时间
	LevelName    string `json:"level_name"`    //等级名称
}
