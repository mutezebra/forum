// Code generated by hertz generator.

package forum

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/mutezebra/forum/pkg/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Cors()}
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _authMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JWT()}
}

func _forumMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createcommentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createpostMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createthreadMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcommentsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getpostsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getthreadsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _forum0Mw() []app.HandlerFunc {
	// your code...
	return nil
}
