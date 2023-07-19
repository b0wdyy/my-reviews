import { cn } from "@/lib/util";
import { InputHTMLAttributes } from "react";

type InputProps = InputHTMLAttributes<HTMLInputElement> & {
    prefixIcon?: React.ReactNode;
    suffixIcon?: React.ReactNode;
}

export default function Input({ prefixIcon, suffixIcon, ...rest }: InputProps) {
    return (
        <div className="bg-off-white rounded-xl focus-within:ring focus-within:ring-text border border-text p-1 px-2 inline-flex gap-2 items-center focus-within:ring-opacity-25">
            {prefixIcon}
            <input {...rest} className={cn("bg-transparent focus:outline-none", rest.className)} />
            {suffixIcon}
        </div>
    )
}
