package core

// NewBuildContext create a new build context
func NewBuildContext(global *GlobalBuildContext) *BuildContext {
	return &BuildContext{
		components:    map[string]bool{},
		sequence:      1,
		scripts:       []string{},
		styles:        []string{},
		jitComponents: map[string]bool{},
		global:        global,
	}
}

// NewGlobalBuildContext create a new global build context
func NewGlobalBuildContext() *GlobalBuildContext {
	return &GlobalBuildContext{
		jitComponents: map[string]bool{},
	}
}

// GetJitComponents get the just in time components
func (ctx *BuildContext) GetJitComponents() []string {
	if ctx.jitComponents == nil {
		return []string{}
	}
	jitComponents := []string{}
	for name := range ctx.jitComponents {
		jitComponents = append(jitComponents, name)
	}
	return jitComponents
}

// GetJitComponents get the just in time components
func (globalCtx *GlobalBuildContext) GetJitComponents() []string {
	if globalCtx.jitComponents == nil {
		return []string{}
	}

	jitComponents := []string{}
	for name := range globalCtx.jitComponents {
		jitComponents = append(jitComponents, name)
	}
	return jitComponents
}

func (ctx *BuildContext) addJitComponent(name string) {
	name = stmtRe.ReplaceAllString(name, "*")
	name = propRe.ReplaceAllString(name, "*")
	ctx.jitComponents[name] = true
	if ctx.global != nil {
		ctx.global.jitComponents[name] = true
	}
}
