'use client'

import { cn } from "@/lib/util";
import { BookmarkIcon, HomeIcon, StarIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import { usePathname } from 'next/navigation'

export default function SideBar() {
    const pathName = usePathname()

    const items = [
        {
            name: "Home",
            to: "/",
            icon: <HomeIcon className="h-6 w-6 text-text" />,
        },
        {
            name: "Favorites",
            to: "/favorites",
            icon: <StarIcon className="h-6 w-6 text-text" />,
        },
        {
            name: "To Read",
            to: "/to-read",
            icon: <BookmarkIcon className="h-6 w-6 text-text" />,
        },
    ]

    return (
        <nav className="flex flex-col gap-4 px-8 pb-4" id="side-bar">
            {items.map((item) => {
                const isActive = pathName.startsWith(item.to)
                const className = isActive ? "bg-white" : ""

                return (
                    <Link href={item.to} key={item.name} className={cn("flex items-center gap-2 p-4 hover:bg-white transition-colors rounded-xl", className)}>
                        {item.icon}
                        <p className="text-text">{item.name}</p>
                    </Link>
                )
            }
            )}
        </nav>
    )
}
