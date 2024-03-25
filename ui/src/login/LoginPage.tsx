import React from "react";
import {LoginVariables, useLoginMutation} from "../api/hooks/useLogin";
import {useNavigate} from "react-router";
import {useAuth} from "../auth/AuthProvider";

function LoginPage() {
    const navigate = useNavigate();
    const {setToken} = useAuth();

    const onLoginError = (error: Error) => {
        setLoginError(error.message)
    }
    const onLoginSuccess = (token: string) => {
        setLoginError("")
        setToken(token)
        navigate("/")
    }
    const loginMutation = useLoginMutation({onSuccess: onLoginSuccess, onError: onLoginError})

    const [loginError, setLoginError] = React.useState("");
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");

    const onEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value);
    }
    const onPasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    }

    const handleLogin = (e: React.FormEvent<HTMLFormElement>) => {
        // 👇️ prevent page refresh
        e.preventDefault();

        loginMutation.mutate({email: email, password: password} as LoginVariables);
    };

    return <div className={"fixed inset-0 flex items-center justify-center -z-10"}>
        <div className="position-absolute grid place-items-center">
            <div className="w-full max-w-sm px-4 py-6 space-y-6 bg-white rounded-md dark:bg-darker">
                <div className={"select-none"}>
                    <div className={"flex justify-center items-center"}>
                        <img
                            alt=""
                            src="/logo.png"
                            width="96"
                            height="96"
                            className="d-inline-block align-top"
                        />{' '}
                        <div className={"pl-2 font-logo text- text-5xl align-self-center"}>
                            YGODraft
                        </div>
                    </div>
                </div>
                <h1 className="text-xl font-semibold text-center">Login</h1>
                <form onSubmit={handleLogin} className="space-y-6">
                    <input
                        className="w-full px-4 py-2 border rounded-md dark:bg-darker dark:border-gray-700 focus:outline-none focus:ring focus:ring-primary-100 dark:focus:ring-primary-darker"
                        type="email"
                        name="email"
                        placeholder="Email address"
                        value={email}
                        onChange={onEmailChange}
                        required
                    />
                    <input
                        className="w-full px-4 py-2 border rounded-md dark:bg-darker dark:border-gray-700 focus:outline-none focus:ring focus:ring-primary-100 dark:focus:ring-primary-darker"
                        type="password"
                        name="password"
                        placeholder="Password"
                        value={password}
                        onChange={onPasswordChange}
                        required
                    />
                    <div className="flex items-center justify-between">
                        <label className="flex items-center">
                        </label>

                        <a href="forgot-password.html" className="text-sm text-blue-600 hover:underline">Forgot
                            Password?</a>
                    </div>
                    {loginError ? <div className={"bg-danger rounded-1 text-white pl-2 pr-2 pt-1 pb-1"}>
                        Cannot log in. Check your credentials and try again!
                    </div> : <></>}
                    <div>
                        <button
                            type="submit"
                            className="flex w-full px-4 py-2 font-medium text-center text-white transition-colors duration-200 rounded-md bg-primary hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-1 dark:focus:ring-offset-darker"
                        >
                            Login
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
}

export default LoginPage