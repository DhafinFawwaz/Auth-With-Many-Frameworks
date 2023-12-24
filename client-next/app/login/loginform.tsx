"use client"

import { signIn } from "next-auth/react";
import Link from "next/link";
import { labelStyle, inputStyle, linkStyle } from "../style/style";
import { useState } from "react";
import { LoadingSpinner } from "../component/LoadingSpinner";
import { useRouter } from "next/navigation";
export default function LoginForm(){
    const [errorMessage, setErrorMessage] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter();
    function onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        setIsLoading(true);
        const form = e.currentTarget;
        const email = form.email.value;
        const password = form.password.value;
        signIn('credentials',
            {
                email,
                password,
                callbackUrl: `${window.location.origin}/profile`,
                redirect: false
            }
        ).then(res => {
            if (res && res.ok) {
                router.push('/profile');
            } else {
                setIsLoading(false);
                setErrorMessage(getErrorMessage(res?.error));
                window.history.replaceState(null, "New Page Title", "/login") // reset url
            }
        })
    }

    function getErrorMessage(error: string|null|undefined): string{
        if(!error) return "Something went wrong. Please try again.";
        if(error === "CredentialsSignin") return "Invalid email or password.";
        return "Something went wrong. Please try again.";
    }

    return <>
<form onSubmit={onSubmit} className="space-y-2 md:space-y-4" action="#">
    <div className="grow">
        <label htmlFor="email" className={labelStyle}>Your email</label>
        <input type="email" name="email" id="email" className={inputStyle} placeholder="name@gmail.com" required/>
    </div>
    <div>
        <label htmlFor="password" className={labelStyle}>Password</label>
        <input type="password" name="password" id="password" placeholder="••••••••" className={inputStyle} required/>
    </div>

    {isLoading ? 
        <div className="w-full text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 flex justify-center">
            <LoadingSpinner></LoadingSpinner>
            <div className='ml-1'>Loging in...</div>
        </div>
    :
        <button type="submit" className="w-full text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Login</button>
    }
    
    <p className="text-red-500">{errorMessage}</p>
    <p className='my-0'>{`Don't have a registered device ? `}<Link className={linkStyle} href="/register">Register here.</Link></p>
</form>
    </>
}