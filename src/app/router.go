package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ver "github.com/hashicorp/go-version"
)

type RoutingEntry struct {
	engines []*gin.RouterGroup
}

func addRoute(engine *gin.RouterGroup, method, path string, hf gin.HandlerFunc) {
	for _, path := range []string{path, path + "/"} {
		engine.Handle(method, path, hf)
	}
}

func (re *RoutingEntry) GET(path string, hf gin.HandlerFunc)    { re.handle("GET", path, hf) }
func (re *RoutingEntry) POST(path string, hf gin.HandlerFunc)   { re.handle("POST", path, hf) }
func (re *RoutingEntry) DELETE(path string, hf gin.HandlerFunc) { re.handle("DELETE", path, hf) }
func (re *RoutingEntry) PUT(path string, hf gin.HandlerFunc)    { re.handle("PUT", path, hf) }

func (re *RoutingEntry) Use(hf gin.HandlerFunc) {
	for _, engine := range re.engines {
		engine.Use(hf)
	}
}

func (re *RoutingEntry) handle(method, path string, hf gin.HandlerFunc) {
	for _, engine := range re.engines {
		addRoute(engine, method, path, hf)
	}
}

type Router struct {
	engine  *gin.Engine
	engines map[string]*gin.RouterGroup
}

func NewRouter(engine *gin.Engine, versions ...string) *Router {
	m := make(map[string]*gin.RouterGroup)
	for _, v := range versions {
		m[v] = engine.Group(v)
	}
	return &Router{
		engine:  engine,
		engines: m,
	}
}

func (r *Router) Group(path string) *Router {
	m := make(map[string]*gin.RouterGroup, len(r.engines))
	for v, engine := range r.engines {
		m[v] = engine.Group(path)
	}
	return &Router{engines: m}
}

func (r *Router) Version(ver string, callback func(*Router)) {
	group := r.engine.Group(ver)
	m := make(map[string]*gin.RouterGroup)
	m[ver] = group

	versioned := &Router{engines: m}
	callback(versioned)

	r.engines[ver] = versioned.engines[ver]
}

func (r *Router) Ver(versions ...string) *RoutingEntry {
	var entry RoutingEntry
	for _, v := range versions {
		engine, found := r.engines[v]
		if !found {
			panic(fmt.Errorf("unexpected version %s", v))
		}
		entry.engines = append(entry.engines, engine)
	}
	return &entry
}

func (r *Router) VerFrom(version string) *RoutingEntry {
	_, found := r.engines[version]
	if !found {
		panic(fmt.Errorf("unexpected version %s", version))
	}

	versions := []string{}
	v1 := ver.Must(ver.NewVersion(version))
	for v, _ := range r.engines {
		v2, _ := ver.NewVersion(v)
		if v1.Compare(v2) <= 0 {
			versions = append(versions, v)
		}
	}
	return r.Ver(versions...)
}

func (r *Router) VerTo(version string) *RoutingEntry {
	_, found := r.engines[version]
	if !found {
		panic(fmt.Errorf("unexpected version %s", version))
	}

	versions := []string{}
	v1 := ver.Must(ver.NewVersion(version))
	for v, _ := range r.engines {
		v2, _ := ver.NewVersion(v)
		if v1.Compare(v2) >= 0 {
			versions = append(versions, v)
		}
	}
	return r.Ver(versions...)
}

func (r *Router) VerAll() *RoutingEntry {
	var entry RoutingEntry
	for _, engine := range r.engines {
		entry.engines = append(entry.engines, engine)
	}
	return &entry
}

func (r *Router) Use(hf gin.HandlerFunc)                 { r.VerAll().Use(hf) }
func (r *Router) GET(path string, hf gin.HandlerFunc)    { r.VerAll().GET(path, hf) }
func (r *Router) POST(path string, hf gin.HandlerFunc)   { r.VerAll().POST(path, hf) }
func (r *Router) DELETE(path string, hf gin.HandlerFunc) { r.VerAll().DELETE(path, hf) }
func (r *Router) PUT(path string, hf gin.HandlerFunc)    { r.VerAll().PUT(path, hf) }
