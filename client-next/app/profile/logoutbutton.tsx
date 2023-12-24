"use client"
import { signOut } from "next-auth/react";

export function LogoutButton(){
    function onLogout(e: React.MouseEvent<HTMLElement>) {
        e.preventDefault();
        signOut({
            callbackUrl: `${window.location.origin}/login`
        });
    }

    return <>
<button onClick={onLogout} type="submit" className="w-full text-white font-bold bg-gray-600 hover:bg-gray-700 focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-lg text-sm px-5 py-2.5 text-center dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-900">
    Logout
</button>
    </>
}