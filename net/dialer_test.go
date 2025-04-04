package net

import (
    "context"
    "reflect"
    "syscall"
    "testing"
    "time"
)

func TestNewDialer(t *testing.T) {
    var net  = NewNet()
    var intf = net.NewDialer()

    impl, ok := intf.(dialerFacade)
    if !ok {
        t.Fatalf(`unexpected type: %T`, intf)
    }
    
    if impl.dialer == nil {
        t.Fatalf(`expected non-nil dialer`)
    }

    var dia = impl.dialer

    if dia.Timeout != time.Duration(0) {
        t.Errorf(`unexpected timeout: %v`, dia.Timeout)
    }
    if !dia.Deadline.Equal(time.Time{}) {
        t.Errorf(`unexpected deadline: %v`, dia.Deadline)
    }
    if dia.LocalAddr != nil {
        t.Errorf(`unexpected local-addr: %v`, dia.LocalAddr)
    }
    if dia.DualStack {
        t.Errorf(`expected dual-stack to be false`)
    }
    if dia.FallbackDelay != time.Duration(0) {
        t.Errorf(`unexpected fallback-delay: %v`, dia.FallbackDelay)
    }
    if dia.KeepAlive != time.Duration(0) {
        t.Errorf(`unexpected keep-alive: %v`, dia.KeepAlive)
    }
    if !reflect.DeepEqual(dia.KeepAliveConfig, KeepAliveConfig{}) {
        t.Errorf(`unexpected keep-alive: %+v`, dia.KeepAliveConfig)
    }
    if dia.Resolver != nil {
        t.Errorf(`unexpected resolver: %+v`, dia.Resolver)
    }
    if dia.Cancel != nil {
        t.Errorf(`unexpected cancel: %+v`, dia.Cancel)
    }
    if dia.Control != nil {
        t.Errorf(`unexpected control: %p`, dia.Control)
    }
    if dia.ControlContext != nil {
        t.Errorf(`unexpected control-context: %p`, dia.ControlContext)
    }
}

func TestNewDialer_WithTimeout(t *testing.T) {
    var net  = NewNet()
    var timeout = time.Minute
    var intf = net.NewDialer(WithTimeout(timeout))

    impl, ok := intf.(dialerFacade)
    if !ok {
        t.Fatalf(`unexpected type: %T`, intf)
    }
    
    if impl.dialer == nil {
        t.Fatalf(`expected non-nil dialer`)
    }

    var dia = impl.dialer

    if dia.Timeout != timeout {
        t.Errorf(`unexpected timeout: %v`, dia.Timeout)
    }
    if !dia.Deadline.Equal(time.Time{}) {
        t.Errorf(`unexpected deadline: %v`, dia.Deadline)
    }
    if dia.LocalAddr != nil {
        t.Errorf(`unexpected local-addr: %v`, dia.LocalAddr)
    }
    if dia.DualStack {
        t.Errorf(`expected dual-stack to be false`)
    }
    if dia.FallbackDelay != time.Duration(0) {
        t.Errorf(`unexpected fallback-delay: %v`, dia.FallbackDelay)
    }
    if dia.KeepAlive != time.Duration(0) {
        t.Errorf(`unexpected keep-alive: %v`, dia.KeepAlive)
    }
    if !reflect.DeepEqual(dia.KeepAliveConfig, KeepAliveConfig{}) {
        t.Errorf(`unexpected keep-alive: %+v`, dia.KeepAliveConfig)
    }
    if dia.Resolver != nil {
        t.Errorf(`unexpected resolver: %+v`, dia.Resolver)
    }
    if dia.Cancel != nil {
        t.Errorf(`unexpected cancel: %+v`, dia.Cancel)
    }
    if dia.Control != nil {
        t.Errorf(`unexpected control: %p`, dia.Control)
    }
    if dia.ControlContext != nil {
        t.Errorf(`unexpected control-context: %p`, dia.ControlContext)
    }
}

func TestNewDialer_WithDeadline(t *testing.T) {
    var net  = NewNet()
    var deadline = time.Now().Add(time.Minute)
    var intf = net.NewDialer(WithDeadline(deadline))

    impl, ok := intf.(dialerFacade)
    if !ok {
        t.Fatalf(`unexpected type: %T`, intf)
    }
    
    if impl.dialer == nil {
        t.Fatalf(`expected non-nil dialer`)
    }

    var dia = impl.dialer

    if dia.Timeout != time.Duration(0) {
        t.Errorf(`unexpected timeout: %v`, dia.Timeout)
    }
    if !dia.Deadline.Equal(deadline) {
        t.Errorf(`unexpected deadline: %v`, dia.Deadline)
    }
    if dia.LocalAddr != nil {
        t.Errorf(`unexpected local-addr: %v`, dia.LocalAddr)
    }
    if dia.DualStack {
        t.Errorf(`expected dual-stack to be false`)
    }
    if dia.FallbackDelay != time.Duration(0) {
        t.Errorf(`unexpected fallback-delay: %v`, dia.FallbackDelay)
    }
    if dia.KeepAlive != time.Duration(0) {
        t.Errorf(`unexpected keep-alive: %v`, dia.KeepAlive)
    }
    if !reflect.DeepEqual(dia.KeepAliveConfig, KeepAliveConfig{}) {
        t.Errorf(`unexpected keep-alive: %+v`, dia.KeepAliveConfig)
    }
    if dia.Resolver != nil {
        t.Errorf(`unexpected resolver: %+v`, dia.Resolver)
    }
    if dia.Cancel != nil {
        t.Errorf(`unexpected cancel: %+v`, dia.Cancel)
    }
    if dia.Control != nil {
        t.Errorf(`unexpected control: %p`, dia.Control)
    }
    if dia.ControlContext != nil {
        t.Errorf(`unexpected control-context: %p`, dia.ControlContext)
    }
}

func TestNewDialer_WithAll(t *testing.T) {
    var net  = NewNet()
    var timeout  = time.Minute
    var deadline = time.Now().Add(time.Minute)
    var localAddr = IPAddr{ IP: []byte{ 10, 0, 0, 3 } }
    var fallbackDelay = 2 * time.Minute
    var keepAlive = 30 * time.Second
    var keepAliveConfig = KeepAliveConfig{
        Idle: 40 * time.Second,
        Interval: time.Minute,
        Count: 4,
    }
    var resolver = net.NewResolver()
    var cancel = make(chan struct{})
    var control = func(string, string, syscall.RawConn) error { return nil }
    var controlContext = func(context.Context, string, string, syscall.RawConn) error { return nil }
    var intf = net.NewDialer(
        WithTimeout(timeout),
        WithDeadline(deadline),
        WithLocalAddr(&localAddr),
        WithDualStack(),
        WithFallbackDelay(fallbackDelay),
        WithKeepAlive(keepAlive),
        WithKeepAliveConfig(keepAliveConfig),
        WithResolver(resolver),
        WithCancel(cancel),
        WithControl(control),
        WithControlContext(controlContext),
        WithSetMultipathTCP(true))

    impl, ok := intf.(dialerFacade)
    if !ok {
        t.Fatalf(`unexpected type: %T`, intf)
    }
    
    if impl.dialer == nil {
        t.Fatalf(`expected non-nil dialer`)
    }

    var dia = impl.dialer

    if dia.Timeout != timeout {
        t.Errorf(`unexpected timeout: %v`, dia.Timeout)
    }
    if !dia.Deadline.Equal(deadline) {
        t.Errorf(`unexpected deadline: %v`, dia.Deadline)
    }
    if dia.LocalAddr != &localAddr {
        t.Errorf(`unexpected local-addr: %v`, dia.LocalAddr)
    }
    if !dia.DualStack {
        t.Errorf(`expected dual-stack to be true`)
    }
    if dia.FallbackDelay != fallbackDelay {
        t.Errorf(`unexpected fallback-delay: %v`, dia.FallbackDelay)
    }
    if dia.KeepAlive != keepAlive {
        t.Errorf(`unexpected keep-alive: %v`, dia.KeepAlive)
    }
    if !reflect.DeepEqual(dia.KeepAliveConfig, keepAliveConfig) {
        t.Errorf(`unexpected keep-alive: %+v`, dia.KeepAliveConfig)
    }
    if dia.Resolver != resolver.getUnderlyingResolver() {
        t.Errorf(`unexpected resolver: %+v`, dia.Resolver)
    }
    if dia.Cancel != cancel {
        t.Errorf(`unexpected cancel: %+v`, dia.Cancel)
    }
    if dia.Control == nil {
        t.Errorf(`expected non-nil control`)
    }
    if dia.ControlContext == nil {
        t.Errorf(`expected non-nil control-context`)
    }
}
