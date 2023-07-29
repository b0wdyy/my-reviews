import { Nav, SideBar } from '@/components/common'
import './globals.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <Nav />
        <SideBar />

        <main id="content" className="p-4 mr-8 bg-off-white rounded-xl">
          {children}
        </main>
      </body>
    </html>
  )
}
