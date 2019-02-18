package DataStruct

import "errors"
/**
 *link element
 */
type Element struct {
	No   int
	Val  int
	Next *Element
}

type SingleLinkIf interface {
	IsEmpty() (bool)
	GetElement(index int) (*Element)
	GetHeadElement() (*Element)
	GetTailElement() (*Element)
	InsertElement(element *Element, index int) (error)
	InsertHeadElement(element *Element)
	InsertTailElement(element *Element)
	DeleteElement(element *Element, index int) (error)
	DeleteHeadElement() (error)
	DeleteTailElement() (error)
	Destroy()
	Size() (int)
}

type SingleLink struct {
	head *Element
	size int
	next *Element
}

/**
 *create single link
 */
func NewSingleLink(size int) (*SingleLink) {
	return &SingleLink{nil, size, nil}
}

/**
 *is empty
 */
func (s *SingleLink) IsEmpty() (bool) {

	if s.size == 0 {
		return true
	}

	return false
}

/**
 *create a element
 */
func (s *SingleLink) CreateElement(val int) (*Element) {
	return &Element{0, val, nil}
}

/**
 * get element by index
 */
func (s *SingleLink) GetElement(index int) (*Element) {
	node := s.head

	for node != nil {
		if node.No == index {
			return node
		}
		node = node.Next
	}
	return nil
}

/**
 *get single link of element
 */

func (s *SingleLink) GetHeadElement() (*Element) {
	return s.head
}

/**
 *get single link of tail
 */

func (s *SingleLink) GetTailElement() (*Element) {
	node := s.head

	for node != nil {
		node = node.Next
	}

	return node
}

/**
 *add element
 */
func (s *SingleLink) InsertElement(element *Element, index int) (error) {

	if s.head == nil {
		s.head = element
		s.size++
		return nil
	}

	if s.size < index {
		return errors.New("over index")
	}

	node := s.head
	count := 1
	for node != nil {
		count++
		if count == (index - 1) {
			element.No = index
			element.Next = node.Next
			node.Next = element
			s.size++
		}
		node = node.Next

	}

	if index > s.size {
		node.Next = element
		s.size++
	}

	return nil

}

/**
 *add element to single link of head
 */
func (s *SingleLink) InsertHeadElement(element *Element) {
	node := s.head
	element.Next = node
	s.head = element
	s.size++
}

/**
 *add element to single link of tail
 */
func (s *SingleLink) InsertTailElement(element *Element) {
	node := s.head

	for node != nil {
		node = node.Next
	}
	node.Next = element
	s.size++
}

/**
 *delete element
 */
func (s *SingleLink) DeleteElement(index int) (error) {
	node := s.head
	count := 1
	if index > s.size {
		return errors.New("over index")
	}
	var prev *Element
	for node != nil {
		prev = node
		if count == index {
			prev.Next = node.Next
			s.size--
			return nil
		}

		node = node.Next
		count++
	}

	return nil

}

/**
 *delete link of head element
 */
func (s *SingleLink) DeleteHeadElement() (error) {
	s.head = s.head.Next
	s.size--
	return nil
}

/**
 *delete link of tail element
 */

func (s *SingleLink) DeleteTailElement() (error) {
	node := s.head

	var prev *Element

	for node != nil {
		prev = node
		node = node.Next
	}
	prev.Next = nil
	s.size--
	return nil
}

/**
 *destroy single link
 */

func (s *SingleLink) Destroy() {
	s.head = nil
}

/**
 *get single link size
 */

func (s *SingleLink) Size() (int) {
	return s.size
}
