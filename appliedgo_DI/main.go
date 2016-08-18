package main

import "fmt"

type Poem struct {
	content []byte
	storage PoemStorage
}

type PoemStorage interface {
	Type() string
	Load(string) []byte
	Save(string, []byte)
}

func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		content: []byte("I am a poem from a " + ps.Type() + "."),
		storage: ps,
	}
}

func (p *Poem) Save(name string) {
	p.storage.Save(name, p.content)
}

func (p *Poem) Load(name string) {
	p.content = p.storage.Load(name)
}

func (p *Poem) String() string {
	return string(p.content)
}

type Notebook struct {
	poems map[string][]byte
}

func NewNotebook() *Notebook {
	return &Notebook{
		poems: map[string][]byte{},
	}
}

func (n *Notebook) Save(name string, contents []byte) {
	n.poems[name] = contents
}

func (n *Notebook) Load(name string) []byte {
	return n.poems[name]
}

func (n *Notebook) Type() string {
	return "Notebook"
}

type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.poem = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Type() string {
	return "Napkin"
}

func main() {
	notebook := NewNotebook()
	napkin := NewNapkin()

	poem := NewPoem(notebook)
	poem.Save("My first poem")

	poem = NewPoem(notebook)
	poem.Load("My first poem")
	fmt.Println(poem)

	poem = NewPoem(napkin)
	poem.Save("My second poem")
	poem = NewPoem(napkin)
	poem.Load("My second poem")
	fmt.Println(poem)

}
