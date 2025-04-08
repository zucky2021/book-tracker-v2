package domain

type Book struct {
	ID         string     `json:"id"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	VolumeInfo VolumeInfo `json:"volumeInfo"` // 書棚に含まれる書籍の数
}

type SaleInfo struct {
	BuyLink     string `json:"buyLink"`     // 購入リンク
	Saleability string `json:"saleability"` // 販売状態 ("FOR_SALE", "FREE", "NOT_FOR_SALE", "FOR_PREORDER")
}

type VolumeInfo struct {
	Title       string      `json:"title"`
	Authors     []string    `json:"authors,omitempty"`     // 著者（オプション）
	Description string      `json:"description,omitempty"` // 説明（オプション）
	Categories  []string    `json:"categories,omitempty"`  // カテゴリ（オプション）
	ImageLinks  *ImageLinks `json:"imageLinks,omitempty"`  // サムネイル画像（オプション）
	InfoLink    string      `json:"infoLink"`              // 詳細リンク
}

type ImageLinks struct {
	Thumbnail string `json:"thumbnail"` // サムネイル画像リンク
}
type BookRepository interface {
	// google books apiを使用して本を取得する
	FindAll(userId string, shelfId int, startIndex int, maxResults int) ([]Book, error)
}
