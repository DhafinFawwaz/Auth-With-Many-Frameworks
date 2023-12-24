import { ensureNotAuthenticated } from '../api/auth/[...nextauth]/route';
import LoginForm from './loginform';

export default async function Login() {
    await ensureNotAuthenticated();

    return (
<section className="bg-gray-50 dark:bg-gray-900 h-screen justify-center">
    <div className="flex flex-col items-center justify-center p-5 mx-auto h-screen lg:py-0">
        <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
            <div className="p-4 space-y-4 md:space-y-6 sm:p-8">
                <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                    Login
                </h1>
                <LoginForm></LoginForm>
            </div>
        </div>
    </div>
</section>
    );
}