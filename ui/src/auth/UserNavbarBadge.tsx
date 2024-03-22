import SvgIconButton from "../core/SvgIconButton";
import {useCurrentUser} from "../api/hooks/useUser";
import {Nav, Spinner} from "react-bootstrap";

const userIcon = <SvgIconButton size={18} classNames={"fill-dark dark:fill-white"}>
    <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0"/>
    <path
        d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1"/>
</SvgIconButton>

function UserNavbarBadge() {
    const {data, isLoading, error} = useCurrentUser()

    let content = <></>
    if (isLoading) {
        content = <div className={"flex align-content-center"}>
            {userIcon}
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (error) {
        content = <div className={"flex align-content-center"}>
            {userIcon}
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get user!</div>
        </div>
    } else if (data) {
        content = <button className={"flex justify-content-center"}>
            {userIcon}
            <span className={"ml-2"}>
                {data.display_name}
            </span>
        </button>
    }

    return <>
        <Nav.Link className="justify-content-end">
            {content}
        </Nav.Link>
    </>
}

export default UserNavbarBadge