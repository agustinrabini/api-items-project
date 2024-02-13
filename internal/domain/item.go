package domain

type ItemsOutput struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID          string     `json:"id" bson:"_id,omitempty"`
	Name        string     `json:"name" bson:"name,$set,omitempty"`
	ShopID      string     `json:"shop_id" bson:"shop_id,$set,omitempty"`
	UserID      string     `json:"user_id" bson:"user_id,$set,omitempty" binding:"required"`
	Category    Category   `json:"category" bson:"category,$set,omitempty"`
	Price       Price      `json:"price" bson:"-"`
	Description string     `json:"description" bson:"description,$set,omitempty"`
	Status      string     `json:"status" bson:"status,omitempty" default:"active"`
	Images      []Image    `json:"images" bson:"images,$set,omitempty"`
	Attributes  Attributes `json:"attributes" bson:"attributes,$set,omitempty"`
	Eligible    []Eligible `json:"eligible,omitempty" bson:"eligible,$set,omitempty"`
}

type Category struct {
	ID   string `json:"id" bson:"_id,$set,omitempty"`
	Name string `json:"name" bson:"name,$set,omitempty"`
}

type Eligible struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Type       string   `json:"type"`
	IsRequired bool     `json:"is_required"`
	Options    []Option `json:"options"`
}

type Attributes map[string]string
type Option string
type Image string

type ItemsIds struct {
	Items []string `json:"items"`
}

func (i *Item) Validate() {
	if len(i.Eligible) == 0 {
		i.Eligible = []Eligible{}
	}

	/*
		 	if len(i.Attributes) == 0 {
				i.Attributes = Attributes{}
			}
	*/
}
