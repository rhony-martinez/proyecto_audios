package entity

type MetadataAudiolibro struct {
	tituloLibro string
	autor       string
	narrador    string
	editorial   string
	isbn        string
	capitulo    string
	archivo     string
}

func NewMetadataAudiolibro(titulo, autor, narrador, editorial, isbn, capitulo, archivo string) MetadataAudiolibro {
	return MetadataAudiolibro{
		tituloLibro: titulo, autor: autor, narrador: narrador,
		editorial: editorial, isbn: isbn, capitulo: capitulo, archivo: archivo,
	}
}

func (a *MetadataAudiolibro) GetTituloLibro() string { return a.tituloLibro }
func (a *MetadataAudiolibro) GetAutor() string       { return a.autor }
func (a *MetadataAudiolibro) GetNarrador() string    { return a.narrador }
func (a *MetadataAudiolibro) GetEditorial() string   { return a.editorial }
func (a *MetadataAudiolibro) GetIsbn() string        { return a.isbn }
func (a *MetadataAudiolibro) GetCapitulo() string    { return a.capitulo }
func (a *MetadataAudiolibro) GetArchivo() string     { return a.archivo }