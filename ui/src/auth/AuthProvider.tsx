import {Context, createContext, useContext, useEffect, useMemo, useState} from "react";

export type AuthContextType = {
    token: string | null,
    setToken: (newToken: any) => void
}

const AuthContext: Context<AuthContextType> = createContext({} as AuthContextType);

export type AuthProviderProps = {
    children: JSX.Element[] | JSX.Element;
}

const AuthProvider = (props: AuthProviderProps) => {
    // State to hold the authentication token
    const [token, setToken_] = useState(localStorage.getItem("token"));

    // Function to set the authentication token
    const setToken = (newToken:string): void => {
        setToken_(newToken);
    };

    useEffect(() => {
        // if (token) {
        //     axios.defaults.headers.common["Authorization"] = "Bearer " + token;
        //     localStorage.setItem('token',token);
        // } else {
        //     delete axios.defaults.headers.common["Authorization"];
        //     localStorage.removeItem('token')
        // }
    }, [token]);

    // Memoized value of the authentication context
    const contextValue = useMemo(
        () => ({
            token,
            setToken,
        }),
        [token]
    );

    // Provide the authentication context to the children components
    return (
        <AuthContext.Provider value={contextValue}>{props.children}</AuthContext.Provider>
    );
};

export const useAuth = () => {
    return useContext(AuthContext);
};

export default AuthProvider;