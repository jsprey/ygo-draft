import {useAuth} from "../auth/AuthProvider";
import {Navigate, Outlet} from "react-router";

export type ProtectedRouteProps = {
    children?: JSX.Element[] | JSX.Element;
}

export const ProtectedRoute = (_: ProtectedRouteProps): JSX.Element => {
    const { token } = useAuth();

    // Check if the user is authenticated
    if (!token) {
        // If not authenticated, redirect to the login page
        return <Navigate to="/login" />;
    }

    // If authenticated, render the child routes
    return <Outlet />;
};