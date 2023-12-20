package main

type Pulse bool

const (
	Low  Pulse = false
	High Pulse = true
)

type Module interface {
	GetName() string
	Send(origin Module, pulse Pulse, queue *Queue)
	SendTargets(pulse Pulse, queue *Queue)
	SetTargets(targets []Module)
	GetTargets() []Module
	SetOrigin(orig Module)
	GetOrigins() []string
}

// FlipFlop Module

type FlipFlop struct {
	name    string
	targets []Module
	state   bool
}

func (f *FlipFlop) Send(_ Module, pulse Pulse, queue *Queue) {
	if !pulse {
		if f.state {
			f.state = false
			f.SendTargets(Low, queue)
		} else {
			f.state = true
			f.SendTargets(High, queue)
		}
	}
}

func (f *FlipFlop) SendTargets(pulse Pulse, queue *Queue) {
	for _, t := range f.targets {
		queue.push(NewQueueItem(f, t, pulse))
	}
}

func (f *FlipFlop) SetTargets(targets []Module) {
	f.targets = targets
}

func (f *FlipFlop) GetTargets() []Module {
	return f.targets
}

func (f *FlipFlop) GetName() string {
	return f.name
}

func (f *FlipFlop) SetOrigin(orig Module) {
}

func (f *FlipFlop) GetOrigins() []string {
	return nil
}

// Conjunction Module

type Conjunction struct {
	name    string
	targets []Module
	memory  map[string]Pulse
}

func (c *Conjunction) Send(origin Module, pulse Pulse, queue *Queue) {
	c.memory[origin.GetName()] = pulse
	if c.allHigh() {
		c.SendTargets(Low, queue)
	} else {
		c.SendTargets(High, queue)
	}
}

func (c *Conjunction) SendTargets(pulse Pulse, queue *Queue) {
	for _, t := range c.targets {
		queue.push(NewQueueItem(c, t, pulse))
	}
}

func (c *Conjunction) SetTargets(targets []Module) {
	c.targets = targets
}

func (c *Conjunction) GetTargets() []Module {
	return c.targets
}

func (c *Conjunction) GetName() string {
	return c.name
}

func (c *Conjunction) SetOrigin(orig Module) {
	if c.memory == nil {
		c.memory = make(map[string]Pulse)
	}
	c.memory[orig.GetName()] = Low
}

func (c *Conjunction) GetOrigins() []string {
	var origins []string
	for k, _ := range c.memory {
		origins = append(origins, k)
	}
	return origins
}

func (c *Conjunction) allHigh() bool {
	for _, m := range c.memory {
		if !m {
			return false
		}
	}
	return true
}

// Broadcaster Module

type BroadcasterModule struct {
	targets []Module
}

func (b *BroadcasterModule) Send(_ Module, pulse Pulse, queue *Queue) {
	b.SendTargets(pulse, queue)
}

func (b *BroadcasterModule) SendTargets(pulse Pulse, queue *Queue) {
	for _, t := range b.targets {
		queue.push(NewQueueItem(b, t, pulse))
	}
}

func (b *BroadcasterModule) SetTargets(targets []Module) {
	b.targets = targets
}

func (b *BroadcasterModule) GetTargets() []Module {
	return b.targets
}

func (b *BroadcasterModule) GetName() string {
	return "broadcaster"
}

func (b *BroadcasterModule) SetOrigin(orig Module) {
}

func (b *BroadcasterModule) GetOrigins() []string {
	return nil
}

// NoOp Module

type NoopModule struct {
	name string
}

func (n NoopModule) Send(origin Module, pulse Pulse, queue *Queue) {
}

func (n NoopModule) SendTargets(pulse Pulse, queue *Queue) {
}

func (n NoopModule) SetTargets(targets []Module) {
}

func (n NoopModule) GetTargets() []Module {
	return nil
}

func (n NoopModule) GetName() string {
	return n.name
}

func (n NoopModule) SetOrigin(orig Module) {
}

func (n NoopModule) GetOrigins() []string {
	return nil
}
