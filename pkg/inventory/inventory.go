package inventory

import "github.com/cosrnic/b173-server/pkg/util"

const PLAYER_INVENTORY_SIZE = 45

type Item struct {
	TypeId int16
	Count  byte
	Uses   int16
}

func (item *Item) Serialize() []byte {
	w := util.NewPacketWriter()

	w.WriteShort(uint16(item.TypeId))
	w.WriteByte(item.Count)
	w.WriteShort(uint16(item.Uses))

	return w.Bytes()
}

func NewItem(typeId int16, count byte) Item {
	return Item{
		TypeId: typeId,
		Count:  count,
		Uses:   0,
	}
}

type Inventory struct {
	Size  uint16
	Items []Item
}

func (inv *Inventory) Serialise() []byte {
	itemIds := util.NewPacketWriter()
	counts := util.NewPacketWriter()

	for i := range inv.Items {
		item := inv.Items[i]
		itemIds.WriteShort(uint16(item.TypeId))
		counts.WriteByte(item.Count)
		counts.WriteShort(uint16(item.Uses))
	}

	w := util.NewPacketWriter()
	w.Write(itemIds.Bytes())
	w.Write(counts.Bytes())

	return w.Bytes()
}

func NewInventory(size uint16) Inventory {
	inv := Inventory{
		Size:  size,
		Items: make([]Item, size),
	}

	for i := range inv.Items {
		inv.Items[i] = NewItem(-1, 1)
	}

	return inv
}
