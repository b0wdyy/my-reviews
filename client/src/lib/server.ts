export async function getBooks() {
    const res = await fetch('http://127.0.0.1:8080/api/books');
    const books = await res.json();
    return books;
}
