package handler

type Action int8

const (
	_ Action = iota
	Like
	Unlike
)

var actionValues = map[string]Action{
	"like":   Like,
	"unlike": Unlike,
}

func (a Action) String() string {
	for s, v := range actionValues {
		if a == v {
			return s
		}
	}
	return "invalid"
}

func ParseAction(s string) Action {
	return actionValues[s]
}

type LikeRequest struct {
	ArticleID string `form:"article_id" json:"articleID" binding:"required"`
	Action    string `form:"action" json:"action" binding:"required"`
}

type LoginForm struct {
	Login    string `form:"login" json:"login" binding:"required,min=2"`
	Password string `form:"password" json:"password"  binding:"required,min=8,max=255"`
}

type SignUpForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Username string `form:"username" json:"username"  binding:"required",min=2,max=255`
	Password string `form:"password" json:"password"  binding:"required,min=8,max=255"`
}

type CommentForm struct {
	ArticleID string `form:"article_id" json:"articleID" binding:"required,min=1"`
	Content   string `form:"content" json:"content"  binding:"required,min=1"`
}

type ArticleForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=1"`
	Content string `form:"content" json:"content"  binding:"required,min=1"`
}
