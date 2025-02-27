package srv

import (
	"strings"

	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

type Group struct {
	Name     string
	Includes []string
}

type From struct {
	Group      string       `json:"group"`
	Display    bool         `json:"display"`
	GroupNodes []GroupNodes `json:"group_nodes"`
}

type GroupNodes struct {
	Name  string
	Nodes []identityclient.UiNode `json:"nodes"`
}

type Froms struct {
	Action   string                  `json:"action"`
	Method   string                  `json:"method"`
	Messages []identityclient.UiText `json:"messages,omitempty"`
	Froms    []From                  `json:"froms"`
}

func groupLoginUi(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "oidc",
			Includes: []string{"oidc", "default"},
		},
		{
			Name:     "password",
			Includes: []string{"password", "default"},
		},
		{
			Name:     "totp",
			Includes: []string{"totp", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupRegistrationUi(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "oidc",
			Includes: []string{"oidc", "default"},
		},
		{
			Name:     "password",
			Includes: []string{"password", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupSettingsUi(ui identityclient.UiContainer) (Froms, bool) {
	groups := []Group{
		{
			Name:     "profile",
			Includes: []string{"profile", "default"},
		},
		{
			Name:     "password",
			Includes: []string{"password", "default"},
		},
		{
			Name:     "totp",
			Includes: []string{"totp", "default"},
		},
	}

	hasImage := fixedBase64ImageForTmpl(ui)
	return groupUi(ui, groups), !hasImage
}

func groupVerificationUi(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "captcha",
			Includes: []string{"captcha", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupRecovery(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "captcha",
			Includes: []string{"captcha", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupConsentUi(ui identityclient.UiContainer) Froms {
	groups := []Group{
		{
			Name:     "consent",
			Includes: []string{"consent", "default"},
		},
	}

	return groupUi(ui, groups)
}

func groupUi(ui identityclient.UiContainer, groups []Group) Froms {
	froms := Froms{
		Action:   ui.Action,
		Method:   ui.Method,
		Messages: ui.Messages,
		Froms:    make([]From, len(groups)),
	}

	for _, n := range ui.Nodes {
		for i, g := range groups {
			for j, item := range g.Includes {
				if n.Group == item {
					froms.Froms[i].Group = g.Name
					if len(froms.Froms[i].GroupNodes) == 0 {
						froms.Froms[i].GroupNodes = make([]GroupNodes, len(g.Includes))
					}

					if n.Group == g.Name {
						froms.Froms[i].Display = true
					}

					froms.Froms[i].GroupNodes[j].Name = n.Group
					froms.Froms[i].GroupNodes[j].Nodes = append(froms.Froms[i].GroupNodes[j].Nodes, n)
					break
				}
			}
		}
	}

	return froms
}

func fixedBase64ImageForTmpl(ui identityclient.UiContainer) bool {
	var has bool
	for _, n := range ui.Nodes {
		imageAttrs := n.Attributes.UiNodeImageAttributes
		if imageAttrs != nil {
			if strings.HasPrefix(imageAttrs.Src, "data:image/png;base64,") {
				imageAttrs.Src = strings.TrimPrefix(imageAttrs.Src, "data:image/png;base64,")
			}
			has = true
		}
	}
	return has
}
