import {Navbar, Spinner} from "react-bootstrap";
import {useCurrentUser} from "../api/hooks/useUser";

function UserNavbarBadge() {
    const {data, isLoading, error} = useCurrentUser()

    let content = <></>
    if (isLoading) {
        content = <div className={"flex align-content-center"}>
            <div className={"mr-2 align-self-center"}>Signed in as:
            </div>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (error) {
        content = <div className={"flex align-content-center"}>
            <div className={"mr-2 align-self-center"}>Signed in as:</div>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get user!</div>
        </div>
    } else if (data) {
        content = <span className={"pl-2"}>Signed in as: <b>{data.display_name}</b></span>
    }

    return <>
        <Navbar.Collapse className="justify-content-end">
            <Navbar.Text>
                {content}
            </Navbar.Text>
        </Navbar.Collapse>
    </>
}

export default UserNavbarBadge