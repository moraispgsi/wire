package fbp

type Initiable interface {
	Init()
}

//TODO: Component Builder
type Component interface {
	Initiable
	GetName() string
	GetInputs() map[string]InputPort
	GetOutputs() map[string]OutputPort
}

type component struct {
	name     string
	inPorts  map[string]InputPort
	outPorts map[string]OutputPort
}

//Overridable
func (c *component) Init() {

}

func NewComponent(name string, inPorts map[string]InputPort, outPorts map[string]OutputPort) Component {

	c := component{
		name:     name,
		inPorts:  inPorts,
		outPorts: outPorts,
	}

	return &c

}

func (c *component) GetInputs() map[string]InputPort {
	return c.inPorts
}

func (c *component) GetOutputs() map[string]OutputPort {
	return c.outPorts
}

func (c *component) GetName() string {
	return c.name
}

type ComponentBuilder struct {
	name     string
	inPorts  map[string]InputPort
	outPorts map[string]OutputPort
}

func NewComponentBuilder(name string) ComponentBuilder {

	return ComponentBuilder{
		name:     name,
		inPorts:  make(map[string]InputPort),
		outPorts: make(map[string]OutputPort),
	}

}

func (cb *ComponentBuilder) AddInPort(portName string, InPort InputPort) {
	cb.inPorts[portName] = InPort
}
func (cb *ComponentBuilder) AddSimpleInPort(portName string) {
	cb.inPorts[portName] = NewSimpleInputPort()
}

func (cb *ComponentBuilder) AddOutPort(portName string, outPort OutputPort) {
	cb.outPorts[portName] = outPort
}

func (cb *ComponentBuilder) AddSimpleOutPort(portName string) {
	cb.outPorts[portName] = NewSimpleOutputPort()
}

func (cb *ComponentBuilder) Build() Component {
	return NewComponent(cb.name, cb.inPorts, cb.outPorts)
}
