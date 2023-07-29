import HighestRatedBooks from '@/components/common/highest-rated-books'

export default function Home() {
  return (
    <>
      <div className="grid grid-cols-2">
        <div>
          <h2 className="text-3xl text-text tracking-wider font-bold">
            Your highest rated books
          </h2>

          <p>
            Here are your highest rated books. If you want to add more notes,
            make sure to do so.
          </p>

          <button>Watch full list</button>
        </div>

        <HighestRatedBooks />
      </div>

      <div className="grid grid-cols-2">
        <div>
          <h2 className="text-3xl text-text tracking-wider font-bold">
            Still need to finish
          </h2>

          <p>
            Here are your books that haven&lsquo;t been finished yet. Make sure
            to finish these books and take notes if needed!
          </p>

          <button>Watch full list</button>
        </div>

        <HighestRatedBooks />
      </div>
    </>
  )
}
