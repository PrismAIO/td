package message

import "github.com/PrismAIO/td/tg"

// Contact adds contact attachment.
func Contact(contact tg.InputMediaContact, caption ...StyledTextOption) MediaOption {
	return Media(&contact, caption...)
}
