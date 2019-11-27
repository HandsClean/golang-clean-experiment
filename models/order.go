package models

// Order represent the author model
type Order struct {
	ID        int64  `json:"id"`
	Number1c  string `json:"number1c"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// // Author represent the author model
// type Acceptance struct {
// 	ID        int64  `json:"id"`
// 	number1c  string `json:"number1c"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// 	Children []Acceptance `json:"children,omitempty"`
// }
