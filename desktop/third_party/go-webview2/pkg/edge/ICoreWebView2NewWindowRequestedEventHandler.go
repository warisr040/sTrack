package edge

// sTrack addition: handler for NewWindowRequested, so pop-up windows (window.open, target=_blank)
// can be refused. A pop-up is a separate WebView without sTrack's request filter.

import "unsafe"

type _ICoreWebView2NewWindowRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri          ComProc
	PutNewWindow    ComProc
	GetNewWindow    ComProc
	PutHandled      ComProc
	GetHandled      ComProc
	GetIsUserInit   ComProc
	GetDeferral     ComProc
	GetWindowFeatur ComProc
}

type ICoreWebView2NewWindowRequestedEventArgs struct {
	vtbl *_ICoreWebView2NewWindowRequestedEventArgsVtbl
}

func (i *ICoreWebView2NewWindowRequestedEventArgs) PutHandled(handled bool) {
	v := uintptr(0)
	if handled {
		v = 1
	}
	_, _, _ = i.vtbl.PutHandled.Call(uintptr(unsafe.Pointer(i)), v)
}

type _ICoreWebView2NewWindowRequestedEventHandlerVtbl struct {
	_IUnknownVtbl
	Invoke ComProc
}

type iCoreWebView2NewWindowRequestedEventHandler struct {
	vtbl *_ICoreWebView2NewWindowRequestedEventHandlerVtbl
	impl _ICoreWebView2NewWindowRequestedEventHandlerImpl
}

func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface(this *iCoreWebView2NewWindowRequestedEventHandler, refiid, object uintptr) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef(this *iCoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease(this *iCoreWebView2NewWindowRequestedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NewWindowRequestedEventHandlerInvoke(this *iCoreWebView2NewWindowRequestedEventHandler, sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	return this.impl.NewWindowRequested(sender, args)
}

type _ICoreWebView2NewWindowRequestedEventHandlerImpl interface {
	_IUnknownImpl
	NewWindowRequested(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr
}

var _ICoreWebView2NewWindowRequestedEventHandlerFn = _ICoreWebView2NewWindowRequestedEventHandlerVtbl{
	_IUnknownVtbl{
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownQueryInterface),
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownAddRef),
		NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerIUnknownRelease),
	},
	NewComProc(_ICoreWebView2NewWindowRequestedEventHandlerInvoke),
}

func newICoreWebView2NewWindowRequestedEventHandler(impl _ICoreWebView2NewWindowRequestedEventHandlerImpl) *iCoreWebView2NewWindowRequestedEventHandler {
	return &iCoreWebView2NewWindowRequestedEventHandler{
		vtbl: &_ICoreWebView2NewWindowRequestedEventHandlerFn,
		impl: impl,
	}
}
