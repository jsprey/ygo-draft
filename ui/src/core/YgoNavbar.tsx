import {Link} from "react-router-dom";
import {useAuth} from "../auth/AuthProvider";
import {useLocation, useNavigate} from "react-router";
import UserNavbarBadge from "../users/UserNavbarBadge";
import React from "react";
import ThemeSwitcher from "./ThemeSwitcher";
import {useCurrentUser} from "../api/hooks/users/useUser";
import classNames from "classnames";
import {AdminPath, DraftDeckPath, HomePath, LoginPath, RandomDeckPath, UserPath} from "../routes/AppRouter";
import YgoIcon from "./YgoIcon";

function YgoNavbar() {
    const {pathname} = useLocation();
    const navigation = useNavigate();
    const {token, setToken} = useAuth();
    const user = useCurrentUser()

    const logout = () => {
        setToken(null)
        navigation(LoginPath)
    }

    const unfocusedLinkCN = classNames("m-1 p-2 rounded", "bg-light dark:bg-dark", "hover:bg-light-1 dark:hover:bg-dark-1", "active:bg-light-2 dark:active:bg-dark-2")
    const focusedLinkCN = classNames("m-1 p-2 rounded", "bg-primary hover:bg-primary-hover active:bg-primary-active", "text-light")

    let userInformation = <></>
    if (token) {
        userInformation = <>
            <UserNavbarBadge
                className={classNames("justify-end", pathname === UserPath ? focusedLinkCN : unfocusedLinkCN)}
                focused={pathname === UserPath}/>
            <span className={classNames("justify-end", unfocusedLinkCN)} onClick={logout}>Logout</span>
        </>
    }


    // for class navBarContainer see index.scss
    // noinspection TypeScriptValidateTypes
    return <>
        <div className={classNames("navBarContainer", "bg-light dark:bg-dark", "text-dark dark:text-light")}>
            <div className={classNames("container mx-auto flex items-center")}>
                <Link to={HomePath} className={classNames(unfocusedLinkCN, "flex items-center")}>
                    <img
                        alt=""
                        src="/logo.png"
                        width="30"
                        height="30"
                        className="d-inline-block"
                    />{' '}
                    <span className={"text-xl font-bold ml-1"}>YgoDraft</span>
                </Link>
                <div className="me-auto flex">
                    <Link to={HomePath}
                          className={classNames(pathname === HomePath ? focusedLinkCN : unfocusedLinkCN)}>Home</Link>
                    {token ? <Link to={RandomDeckPath}
                                   className={pathname === RandomDeckPath ? focusedLinkCN : unfocusedLinkCN}>Mode:
                        Random</Link> : <></>}
                    {token ? <Link to={DraftDeckPath}
                                   className={pathname === DraftDeckPath ? focusedLinkCN : unfocusedLinkCN}>Mode:
                        Draft</Link> : <></>}
                </div>
                <div className={classNames("flex")}>
                    <ThemeSwitcher className={unfocusedLinkCN}/>
                    {user && user.data?.is_admin ?
                        <Link to={AdminPath} className={classNames("flex items-center", pathname === AdminPath ? focusedLinkCN : unfocusedLinkCN)}>
                            <YgoIcon icon={"cog"} size={18} classNames={classNames(pathname === AdminPath ? "fill-light" : "fill-dark dark:fill-light")}/>
                        </Link> : <></>}
                    {!token ? <Link to={LoginPath}
                                    className={pathname === LoginPath ? focusedLinkCN : unfocusedLinkCN}>Login</Link> : <></>}
                    {token ? userInformation : <></>}
                </div>
            </div>
        </div>
    </>
}

export default YgoNavbar