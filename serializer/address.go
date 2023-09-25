package serializer

import "E-commerce_system/model"

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	CreateAt int64  `json:"createAt"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:       item.ID,
		UserId:   item.UserId,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		CreateAt: item.CreatedAt.Unix(),
	}
}
func BuildAddresses(items []*model.Address) []Address {
	var addresses []Address
	for _, v := range items {
		address := BuildAddress(v)
		addresses = append(addresses, address)
	}
	return addresses
}
