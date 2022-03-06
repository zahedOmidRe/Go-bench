package main

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Config struct {
	Age  int
	Name string
}

type Product struct {
	DBId                 int          `db:"id"`
	DBPrice              int          `db:"price"`
	DBVerticalPosition   int          `db:"vertical_position"`
	DBHorizontalPosition int          `db:"horizontal_position"`
	DBDescriptionType    int          `db:"description_type"`
	DBLimitation         int          `db:"limitation"`
	DBPID                uuid.UUID    `db:"pid"`
	DBVentureID          uuid.UUID    `db:"venture_id"`
	DBCreatedAt          time.Time    `db:"created_at"`
	DBUpdatedAt          time.Time    `db:"updated_at"`
	DBDeletedAt          time.Time    `db:"deleted_at"`
	DBWhitelisted        sql.NullBool `db:"whitelisted"`
	DBEnabled            bool         `db:"enabled"`
	DBConfig             Config       `db:"config"`
	DBImageUrl           string       `db:"image_url"`
	DBIconUrl            string       `db:"icon_url"`
	DBContent            string       `db:"content"`
	DBName               string       `db:"name"`
	DBVentureUrl         string       `db:"venture_url"`
}

func (p Product) ID() int {
	return p.DBId
}

func (p Product) Name() string {
	return p.DBName
}

func (p Product) Price() int {
	return p.DBPrice
}

func (p Product) ImageUrl() string {
	return p.DBImageUrl
}

func (p Product) IconUrl() string {
	return p.DBIconUrl
}

func (p Product) DescriptionType() int {
	return p.DBDescriptionType
}

func (p Product) Content() string {
	return p.DBContent
}

func (p Product) VerticalPosition() int {
	return p.DBVerticalPosition
}

func (p Product) HorizontalPosition() int {
	return p.DBHorizontalPosition
}

func (p Product) Enabled() bool {
	return p.DBEnabled
}

func (p Product) CreatedAt() time.Time {
	return p.DBCreatedAt
}

func (p Product) UpdatedAt() time.Time {
	return p.DBUpdatedAt
}

func (p Product) DeletedAt() time.Time {
	return p.DBDeletedAt
}

func (p Product) Whitelisted() sql.NullBool {
	return p.DBWhitelisted
}

func (p Product) VentureUrl() string {
	return p.DBVentureUrl
}

func (p Product) Config() Config {
	return p.DBConfig
}

func (p Product) PID() uuid.UUID {
	return p.DBPID
}

func (p Product) Limitation() int {
	return p.DBLimitation
}

func (p Product) VentureID() uuid.UUID {
	return p.DBVentureID
}

func (p *Product) SetId(id int) {
	p.DBId = id
}

type ProductPointer struct {
	DBId                 int          `db:"id"`
	DBPrice              int          `db:"price"`
	DBVerticalPosition   int          `db:"vertical_position"`
	DBHorizontalPosition int          `db:"horizontal_position"`
	DBDescriptionType    int          `db:"description_type"`
	DBLimitation         int          `db:"limitation"`
	DBPID                uuid.UUID    `db:"pid"`
	DBVentureID          uuid.UUID    `db:"venture_id"`
	DBCreatedAt          time.Time    `db:"created_at"`
	DBUpdatedAt          time.Time    `db:"updated_at"`
	DBDeletedAt          time.Time    `db:"deleted_at"`
	DBWhitelisted        sql.NullBool `db:"whitelisted"`
	DBEnabled            bool         `db:"enabled"`
	DBConfig             Config       `db:"config"`
	DBImageUrl           string       `db:"image_url"`
	DBIconUrl            string       `db:"icon_url"`
	DBContent            string       `db:"content"`
	DBName               string       `db:"name"`
	DBVentureUrl         string       `db:"venture_url"`
}

func (p *ProductPointer) ID() int {
	return p.DBId
}

func (p *ProductPointer) Name() string {
	return p.DBName
}

func (p *ProductPointer) Price() int {
	return p.DBPrice
}

func (p *ProductPointer) ImageUrl() string {
	return p.DBImageUrl
}

func (p *ProductPointer) IconUrl() string {
	return p.DBIconUrl
}

func (p *ProductPointer) DescriptionType() int {
	return p.DBDescriptionType
}

func (p *ProductPointer) Content() string {
	return p.DBContent
}

func (p *ProductPointer) VerticalPosition() int {
	return p.DBVerticalPosition
}

func (p *ProductPointer) HorizontalPosition() int {
	return p.DBHorizontalPosition
}

func (p *ProductPointer) Enabled() bool {
	return p.DBEnabled
}

func (p *ProductPointer) CreatedAt() time.Time {
	return p.DBCreatedAt
}

func (p *ProductPointer) UpdatedAt() time.Time {
	return p.DBUpdatedAt
}

func (p *ProductPointer) DeletedAt() time.Time {
	return p.DBDeletedAt
}

func (p *ProductPointer) Whitelisted() sql.NullBool {
	return p.DBWhitelisted
}

func (p *ProductPointer) VentureUrl() string {
	return p.DBVentureUrl
}

func (p *ProductPointer) Config() Config {
	return p.DBConfig
}

func (p *ProductPointer) PID() uuid.UUID {
	return p.DBPID
}

func (p *ProductPointer) Limitation() int {
	return p.DBLimitation
}

func (p *ProductPointer) SetId(id int) {
	p.DBId = id
}

func (p *ProductPointer) VentureID() uuid.UUID {
	return p.DBVentureID
}

type ProducMedium struct {
	DBId                 int          `db:"id"`
	DBPrice              int          `db:"price"`
	DBVerticalPosition   int          `db:"vertical_position"`
	DBHorizontalPosition int          `db:"horizontal_position"`
	DBDescriptionType    int          `db:"description_type"`
	DBLimitation         int          `db:"limitation"`
	DBPID                uuid.UUID    `db:"pid"`
	DBVentureID          uuid.UUID    `db:"venture_id"`
	DBCreatedAt          time.Time    `db:"created_at"`
	DBUpdatedAt          time.Time    `db:"updated_at"`
	DBDeletedAt          time.Time    `db:"deleted_at"`
	DBWhitelisted        sql.NullBool `db:"whitelisted"`
	DBEnabled            bool         `db:"enabled"`
	DBConfig             Config       `db:"config"`
	DBImageUrl           string       `db:"image_url"`
	DBIconUrl            string       `db:"icon_url"`
	DBContent            string       `db:"content"`
	DBName               string       `db:"name"`
	DBVentureUrl         string       `db:"venture_url"`
}

type ProductInterface interface {
	ID() int
	Name() string
	Price() int
	ImageUrl() string
	IconUrl() string
	DescriptionType() int
	Content() string
	VerticalPosition() int
	HorizontalPosition() int
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() time.Time
	Whitelisted() sql.NullBool
	VentureUrl() string
	Config() Config
	PID() uuid.UUID
	Limitation() int
	VentureID() uuid.UUID
	SetId(int)
}

func makeProductInterfaceOnPointer() ProductInterface {
	return &ProductPointer{}
}

func makeProductInterfaceOnStruct() ProductInterface {
	return &Product{}
}

func CastToMedium(p ProductInterface) ProducMedium {
	return ProducMedium{
		DBId:                 p.ID(),
		DBName:               p.Name(),
		DBPrice:              p.Price(),
		DBImageUrl:           p.ImageUrl(),
		DBIconUrl:            p.IconUrl(),
		DBDescriptionType:    p.DescriptionType(),
		DBContent:            p.Content(),
		DBVerticalPosition:   p.VerticalPosition(),
		DBHorizontalPosition: p.HorizontalPosition(),
		DBEnabled:            p.Enabled(),
		DBCreatedAt:          p.CreatedAt(),
		DBUpdatedAt:          p.UpdatedAt(),
		DBDeletedAt:          p.DeletedAt(),
		DBWhitelisted:        p.Whitelisted(),
		DBVentureUrl:         p.VentureUrl(),
		DBConfig:             p.Config(),
		DBPID:                p.PID(),
		DBLimitation:         p.Limitation(),
		DBVentureID:          p.VentureID(),
	}
}
