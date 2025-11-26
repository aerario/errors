package errors

const (
	LabelUserFriendly Label = "user-friendly"
)

type Label string

type LabelList []Label

func (self LabelList) Add(labels ...Label) LabelList {
	var newTags = make(LabelList, 0, len(self)+len(labels))

	newTags = append(newTags, self...)
	newTags = append(newTags, labels...)

	return newTags
}

func (self LabelList) Has(label Label) bool {
	for _, l := range self {
		if l == label {
			return true
		}
	}

	return false
}
