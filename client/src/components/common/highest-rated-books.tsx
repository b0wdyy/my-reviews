import { getBooks } from "@/lib/server";

export default async function HighestRatedBooks() {
    const books = await getBooks();
    console.log(books);

    return (
        <div>
            <h1>Highest Rated Books</h1>
        </div>
    );
}
