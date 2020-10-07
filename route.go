package matcher

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RouteMatcher struct {
	Method string
	Path   string
}

func (m *RouteMatcher) Match(actual interface{}) (success bool, err error) {
	switch routeInfo := actual.(type) {
	case gin.RouteInfo:
		if m.Method != routeInfo.Method || m.Path != routeInfo.Path {
			return false, nil
		}

		return true, nil

	case gin.RoutesInfo:
		for _, r := range routeInfo {
			if m.Method == r.Method && m.Path == r.Path {
				return true, nil
			}
		}

		return false, fmt.Errorf("Route not found: %v %v", m.Method, m.Path)

	default:
		return false, fmt.Errorf("Unexpected type: %T", actual)
	}
}

func (m *RouteMatcher) FailureMessage(actual interface{}) string {
	return ""
}

func (m *RouteMatcher) NegatedFailureMessage(actual interface{}) string {
	return ""
}

func HasRoute(method string, path string) *RouteMatcher {
	return &RouteMatcher{Method: method, Path: path}
}
