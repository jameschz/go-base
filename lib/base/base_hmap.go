package base

// Hmap :
type Hmap struct {
	size int
	data map[string]interface{}
}

// NewHmap :
func NewHmap() *Hmap {
	hash := make(map[string]interface{}, 0)
	return &Hmap{
		size: 0,
		data: hash,
	}
}

// Len :
func (h *Hmap) Len() int {
	return h.size
}

// Set :
func (h *Hmap) Set(s string, v interface{}) {
	h.data[s] = v
	h.size++
}

// Get :
func (h *Hmap) Get(s string) interface{} {
	return h.data[s]
}

// Delete :
func (h *Hmap) Delete(s string) {
	delete(h.data, s)
	h.size--
}
