import Input from '../form/input'
import { MagnifyingGlassIcon } from '@heroicons/react/24/outline'
import { AcademicCapIcon } from '@heroicons/react/24/solid'

export default function Nav() {
  return (
    <nav className="flex items-center justify-between p-8" id="main-nav">
      <div className="flex items-center gap-3">
        <AcademicCapIcon className="h-8 w-8 text-primary" />
        <h1 className="font-bold text-text text-2xl">MY_REVIEWS</h1>
      </div>

      <Input
        prefixIcon={<MagnifyingGlassIcon className="h-4 w-4 text-text" />}
        className="w-96"
        placeholder="Search for a title or an author"
      />

      <div>this will be the logged in user?</div>
    </nav>
  )
}
