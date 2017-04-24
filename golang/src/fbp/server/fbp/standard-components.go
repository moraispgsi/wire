package fbp

import (
	"errors"
	"reflect"
)

type openedBracketIP struct{}

func (IP openedBracketIP) String() string {
	return "{"
}

type closedBracketIP struct{}

func (IP closedBracketIP) String() string {
	return "}"
}

func IsBracketIP(IP interface{}) bool {
	_, ok1 := IP.(openedBracketIP)
	_, ok2 := IP.(closedBracketIP)
	return ok1 || ok2
}

func IsOpenedBracketIP(IP interface{}) bool {
	_, ok := IP.(openedBracketIP)
	return ok
}

func IsClosedBracketIP(IP interface{}) bool {
	_, ok := IP.(closedBracketIP)
	return ok
}

func MakeOpenedBracketIP() interface{} {
	return openedBracketIP{}
}
func MakeClosedBracketIP() interface{} {
	return closedBracketIP{}
}

func SendSubstream(substream []interface{}, sender Sender) error {

	sender.Send(MakeOpenedBracketIP())
	for _, IP := range substream {
		err := sender.Send(IP)
		if err != nil {
			return err
		}
	}
	sender.Send(MakeClosedBracketIP())
	return nil
}

func RecvSubstream(receiver Receiver) ([]interface{}, error) {

	var substream = make([]interface{}, 0)
	c := 0
	for {
		IP, err := receiver.Recv()
		if err != nil {
			return nil, errors.New("Substream is incomplete.")
		}

		switch {
		case IsOpenedBracketIP(IP):
			c++
		case IsClosedBracketIP(IP):
			c--
		default:
			substream = append(substream, IP)
		}

		if c <= 0 {
			return substream, nil
		}

	}
}

type compareComponent struct {
	Component
}

func (c *compareComponent) Init() {

	inMap := c.GetInputs()
	outMap := c.GetOutputs()
	IN := inMap["IN"]
	FORWARD := outMap["FORWARD"]
	COMPARISON := outMap["COMPARISON"]

	defer FORWARD.Close()
	defer COMPARISON.Close()
	for {

		substream, err := RecvSubstream(IN)
		if err != nil {
			break
		}

		if len(substream) < 2 {
			COMPARISON.Send(true)
			SendSubstream(substream, FORWARD)
			continue
		}

		var latest interface{}
		var compare = true
		for i, IP := range substream {

			if i != 0 {
				if !compare {
					continue
				}

				compare = reflect.DeepEqual(latest, IP)

			}

			latest = IP

		}

		COMPARISON.Send(compare)
		SendSubstream(substream, FORWARD)

	}
}

func MakeCompareComponent() Component {

	cb := NewComponentBuilder("Compare")
	cb.AddSimpleInPort("IN")
	cb.AddSimpleOutPort("FORWARD")
	cb.AddSimpleOutPort("COMPARISON")
	component := cb.Build()

	return &compareComponent{
		Component: component,
	}

}

//Pairs 2 streams in substreams wrapped by brackets,
type pairComponent struct {
	Component
}

func (c *pairComponent) Init() {
	inMap := c.GetInputs()
	outMap := c.GetOutputs()
	IN1 := inMap["IN1"]
	IN2 := inMap["IN2"]
	OUT := outMap["OUT"]

	defer OUT.Close()
	for {
		substream1, err := RecvSubstream(IN1)
		if err != nil {
			break
		}

		substream2, err := RecvSubstream(IN2)
		if err != nil {
			SendSubstream(substream1, OUT)
			break
		}

		finalSubstream := substream1
		for _, IP := range substream2 {
			finalSubstream = append(finalSubstream, IP)
		}

		SendSubstream(finalSubstream, OUT)
	}
}

func MakePairComponent() Component {

	cb := NewComponentBuilder("Pair")
	cb.AddSimpleInPort("IN1")
	cb.AddSimpleInPort("IN2")
	cb.AddSimpleOutPort("OUT")
	component := cb.Build()

	return &pairComponent{
		Component: component,
	}

}

type splitComponent struct {
	Component
}

func (c *splitComponent) Init() {
	inMap := c.GetInputs()
	outMap := c.GetOutputs()
	IN := inMap["IN"]
	OUT1 := outMap["OUT1"]
	OUT2 := outMap["OUT2"]

	defer OUT1.Close()
	defer OUT2.Close()
	var current = false
	for {
		substream, err := RecvSubstream(IN)
		if err != nil {
			break
		}
		if current {
			SendSubstream(substream, OUT1)
		} else {
			SendSubstream(substream, OUT2)
		}
		current = !current
	}
}

func MakeSplitComponent() Component {

	cb := NewComponentBuilder("Split")
	cb.AddSimpleInPort("IN")
	cb.AddSimpleOutPort("OUT1")
	cb.AddSimpleOutPort("OUT2")
	component := cb.Build()

	return &splitComponent{
		Component: component,
	}

}

func NewSTDComponent(name string) Component {

	switch {
	case name == "Compare":
		return MakeCompareComponent()
	case name == "Pair":
		return MakePairComponent()
	case name == "Split":
		return MakeSplitComponent()
	default:
		return nil
	}

}
