"use client"
import Link from 'next/link';
import { dataStyle, labelStyle, inputStyle, linkStyle } from '../style/style';
import { useState } from 'react';
import { LoadingSpinner } from '../component/LoadingSpinner';
import { useRouter } from 'next/navigation';

export function RegisterForm() {

    const [errorState, setErrorState] = useState<{message: string, success: boolean}>({message: "", success: false});
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter();    

    async function onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        
        setIsLoading(true);
        const formData = new FormData(e.currentTarget);

        const res = await fetch(process.env.NEXT_PUBLIC_API_URL + '/api/auth/register/', {
            method: 'POST',
            body: formData
        });

        if(!res.ok){
            setIsLoading(false);
            const data = await res.json();
            setErrorState({
                message: data.message,
                success: false
            });
            return;
        }
        setIsLoading(false);
        setErrorState({
            message: "Register success. Please login.",
            success: true
        });
        // router.push('/login');
    }

    function getErrorMessage(error: string|null|undefined): string{
        if(!error) return "Something went wrong. Please try again.";
        if(error === "CredentialsSignin") return "Invalid email or password.";
        return "Something went wrong. Please try again.";
    }

    return <>
<form onSubmit={onSubmit} className="space-y-2 md:space-y-4" action="#">
    <div className="grid grid-cols-2 gap-2">
        <div className="grow">
            <label htmlFor="username" className={labelStyle}>Username</label>
            <input type="text" name="username" id="username" className={inputStyle} placeholder="Chicken Mountain" required/>
        </div>
        <div>
            <label htmlFor="nim" className={labelStyle}>NIM</label>
            <input type="text" name="nim" id="nim" className={inputStyle} placeholder="13522180" required/>
        </div>
    </div>


    <div className="grow">
        <label htmlFor="email" className={labelStyle}>Your email</label>
        <input type="email" name="email" id="email" className={inputStyle} placeholder="name@gmail.com" required/>
    </div>
    <div>
        <label htmlFor="password" className={labelStyle}>Password</label>
        <input type="password" name="password" id="password" placeholder="••••••••" className={inputStyle} required/>
    </div>
    <div></div>

    {isLoading ? 
        <div className="w-full text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 flex justify-center">
            <LoadingSpinner></LoadingSpinner>
            <div className='ml-1'>Registering...</div>
        </div>
    :
        <button type="submit" className="w-full text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Register</button>
    }
    
    
    <p className={`${errorState.success ? "text-green-500" : "text-red-500"}`}>{errorState.message}</p>
    <p className='my-0'>Already have a registered device ? <Link className={linkStyle} href="/login">Login here.</Link></p>
</form>
    </>
}