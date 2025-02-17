// Code generated by Tndr(t1) - DO NOT EDIT.

// Tndr(t1): version: v0.0.3

package scene

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

import (
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/i"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

func HTMX(c *router.Context) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var1 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var1 == nil {
			t1_7745c5c3_Var1 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		var t1_7745c5c3_Var2 = []any{c.Get("cls:container_scene").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var2...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<section class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var2).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" role=\"media\" style=\"width: 100%\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = t1_7745c5c3_Var1.Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var3 = []any{c.Get("cls:multiPayloadElement").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var3...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<input id=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("id:sceneTrigger").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var3).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" name=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("hx:ScenePayload").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" type=\"hidden\" hx-post=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("path:CurrentComponent").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" hx-trigger=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("hx:ChangeThrottleHalfSecond").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" hx-target=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("ids:scene-active").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var4 = []any{c.Get("i18n:SceneInstructions").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var4...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<div id=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(c.Get("id:scene-active").(string)))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\" class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var4).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var5 = []any{c.Get("cls:scene-active").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var5...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<article class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var5).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\"><hgroup><h2>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var6 string
		t1_7745c5c3_Var6, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:WelcomeMessage").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 23, Col: 48}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var6))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</h2><h3>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var7 string
		t1_7745c5c3_Var7, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:WelcomeSubMessage").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 24, Col: 51}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var7))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</h3></hgroup>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var8 = []any{c.Get("cls:noComponentWrapper").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var8...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<section class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var8).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var9 = []any{c.Get("cls:nocomponent").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var9...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<div class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var9).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = i.NoComponent().Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</div>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var10 = []any{c.Get("cls:message").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var10...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<div class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var10).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\"><p>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var11 string
		t1_7745c5c3_Var11, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:SceneInstructions").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 32, Col: 49}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var11))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</p></div></section><section><ul><li><span><kbd>Esc</kbd> + <kbd>Q</kbd></span><a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var12 string
		t1_7745c5c3_Var12, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:ShowAllCommands").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 38, Col: 95}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var12))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</a></li><li><span><kbd>Esc</kbd> + <kbd>A</kbd></span><a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var13 string
		t1_7745c5c3_Var13, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:GoToFile").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 39, Col: 88}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var13))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</a></li><li><span><kbd>Esc</kbd> + <kbd>F</kbd></span><a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var14 string
		t1_7745c5c3_Var14, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:FindInFile").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 40, Col: 90}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var14))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</a></li><li><span><kbd>Esc</kbd> + <kbd>S</kbd></span><a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var15 string
		t1_7745c5c3_Var15, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:ShowSettings").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 41, Col: 92}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var15))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</a></li><li><span><kbd>Esc</kbd> + <kbd>E</kbd></span><a>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		var t1_7745c5c3_Var16 string
		t1_7745c5c3_Var16, t1_7745c5c3_Err = t1.JoinStringErrs(c.Get("i18n:ExploreOntology").(string))
		if t1_7745c5c3_Err != nil {
			return t1.Error{Err: t1_7745c5c3_Err, FileName: `layout/v1/scene/scene.t1`, Line: 42, Col: 95}
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1_7745c5c3_Var16))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</a></li></ul></section></article></div></section>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

func Propagate(c meta.Config) t1.ComponentScript {
	return t1.ComponentScript{
		Name: `__t1_Propagate_44ce`,
		Function: `function __t1_Propagate_44ce(c){const triggerInputElement = document.querySelector(c.PayloadSelector);
	triggerInputElement.value = c.Payload;
	if(c.Inspect) 
		console.log("Propagating");
		console.log( c })
		const event = new Event('change');
		triggerInputElement.dispatchEvent(event);
	}
}`,
		Call:       t1.SafeScript(`__t1_Propagate_44ce`, c),
		CallInline: t1.SafeScriptInline(`__t1_Propagate_44ce`, c),
	}
}
