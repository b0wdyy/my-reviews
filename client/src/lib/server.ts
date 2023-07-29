export async function getBooks() {
  const res = await fetch('http://127.0.0.1:8080/api/books')
  const books = await res.json()

  return books
}

export async function getBook(id: string) {
  const res = await fetch(`http://127.0.0.1:8080/api/books/${id}`)
  const book = await res.json()

  return book
}

export async function createBook(book: FormData) {
  const res = await fetch(`http://127.0.0.1:8080/api/books`, {
    method: 'POST',
    body: book,
  })
  const newBook = await res.json()

  return newBook
}
