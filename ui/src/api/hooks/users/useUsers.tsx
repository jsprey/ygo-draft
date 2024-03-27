import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";

export type User = {
    id: number
    email: string
    display_name: string
    is_admin: boolean
}

export type GetUsersResponse = {
    numberOfUsers: number
    numberOfPages: string
    users: User[]
}

const GetUsersPageParameter = "page"
const GetUsersPageSizeParameter = "page_size"

export function useUsers(page: number, pageSize: number, queryOptions: any = {}): UseQueryResult<GetUsersResponse> {
    const queryParameters = new Map<string, string>();
    queryParameters.set(GetUsersPageParameter, page.toString());
    queryParameters.set(GetUsersPageSizeParameter, pageSize.toString());
    return useMagicMethodAxios<GetUsersResponse>(["users", page], `users`, queryParameters , queryOptions)
}